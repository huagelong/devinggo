import {
    WebsocketBuilder,
    LinearBackoff,
} from 'websocket-ts';

import {v4 as uuidv4} from 'uuid';

class WsOptions {
    constructor() {
        this.sessionId = '';
        this.token = '';
        this.heartTime = 3000;
        this.retryLockTime = 120000;
        this.bufferNumber = 1000;
        this.onProgress = (detail, ws) => {};
        this.onConnected = (data, ws) => {};
    }
}

class WsMessage {
    constructor() {
        this.be = ""; // bind event
        this.e = ''; //Event
        this.d = null; //Data
        this.r = ''; //requestId
        this.t = 0; //Time
        this.c = 0; //Code
        this.m = ''; //Message
        this.cb = 0; //CallBack
    }
}

class WsMessageReq {
    constructor() {
        this.e = ''; //Event
        this.d = null; //Data
        this.t = 0; //Time
        this.r = ''; //requestId
    }
}

class messageData {
    constructor() {
        this.message = new WsMessageReq();
        this.callback = (data) => {};
    }
}

const RequestEvents = {
    Error: 'error',
    Connected: 'connected',
    Close: 'close',
    Subscribe: 'subscribe',
    Unsubscribe: 'unsubscribe',
    Publish: 'publish',
    Ping: 'ping',
    PingAll: 'pingAll',
    Pong: 'pong',
    IdMsg: 'idMsg',
    BroadcastMsg: 'broadcastMsg',
};

let _hTimer;
class Ws {
    constructor(url, option) {
        this._isConnected = false;
        this._onEvent = new Map();
        this._messageQueue = new Map();
        this._messageCallbackQueue = new Map();
        const defaultOption = new WsOptions();
        this._option = { ...defaultOption, ...option };
        url = this._formatUrl(url, this._option.sessionId, this._option.token);
        const ws = new WebsocketBuilder(url)
            .withBackoff(
                new LinearBackoff(3000, 3000, this._option.retryLockTime),
            )
            .onOpen((i, ev) => {
                this._isConnected = true;
                this._sendAll_messageQueue(i);
                this._heartCheck(i);
            })
            .onClose((i, ev) => {
                console.log('closed');
                this._isConnected = false;
            })
            .onError((i, ev) => {
                console.log('error');
                console.log(ev);
                this._isConnected = false;
            })
            .onMessage((i, ev) => {
                const message = JSON.parse(ev.data);
                console.log(`received message:${ev.data}`);
                const isContain = this._dispath_messageCallbackQueue(
                    message,
                    i,
                );
                if (!isContain) {
                    if (
                        message.cb === 0 &&
                        message.e !== RequestEvents.PingAll
                    ) {
                        console.log(`_onEvent:${this._onEvent}`);
                        if(this._onEvent.has(message.be)){
                            const callback = this._onEvent.get(message.be);
                            callback(message, i);
                        }else{
                            console.warn(`WebSocket 消息事件[${message.be}]未绑定...`);
                        }
                    }
                }

                if (
                    this._option.onConnected &&
                    message.e === RequestEvents.Connected
                ) {
                    console.log(`onConnected`);
                    this._option.onConnected(message, i);
                }
                if (message.e === RequestEvents.PingAll) {
                    this.pong();
                }
            })
            .onRetry((i, ev) => {
                if (this._option.onProgress) {
                    this._option.onProgress(ev.detail, i);
                }
            })
            .build();
        this._ws = ws;
    }

    _formatUrl(url, sid, token) {
        if (!token) return '';
        const myURL = new URL(url);
        const params = myURL.search;
        if (params) {
            const searchParams = new URLSearchParams(params);
            if (searchParams.get('sessionId')) {
                url = this._replaceParamVal(url, 'sessionId', sid);
            }
            if (searchParams.get('token')) {
                url = this._replaceParamVal(url, 'token', token);
            }
            return url;
        }
        return `${url}?sessionId=${sid}&token=${token}`;
    }

    _replaceParamVal(url, name, value) {
        const re = new RegExp(`${name}=[^&]*`, 'gi');
        return url.replace(re, `${name}=${value}`);
    }

    _formatRequestName(message) {
        if (!message.r) {
            message.r = '0';
        }
        return `${message.e}_${message.r}`;
    }

    _dispath_messageCallbackQueue(message, i) {
        const requestName = this._formatRequestName(message);
        if (this._messageCallbackQueue.has(requestName)) {
            const callback = this._messageCallbackQueue.get(requestName);
            if (callback) {
                callback({
                    message,
                    ws: i,
                });
                this._messageCallbackQueue.delete(requestName);
            }
            return true;
        } else {
            return false;
        }
    }

    _sendAll_messageQueue(i) {
        if (this._messageQueue.size > 0) {
            for (const [requestName, messageData] of this._messageQueue) {
                this._sendCallback(
                    i,
                    messageData.message,
                    messageData.callback,
                );
            }
            this._messageQueue.clear();
        }
    }

    _sendCallback(i, message, callback) {
        message.t = new Date().getTime();
        message.r = this._requestId();
        const requestName = this._formatRequestName(message);
        if (this._isConnected) {
            this._messageCallbackQueue.set(requestName, callback);
            this._sendMsg(i, message);
        } else {
            this._messageQueue.set(requestName, {
                message,
                callback,
            });
        }
    }

    _requestId() {
        let uuid = uuidv4();
        return uuid;
    }

    _sendMsg(i, message) {
        if (!message.t) {
            message.t = new Date().getTime();
        }
        if (!message.r) {
            message.r = this._requestId();
        }
        i.send(JSON.stringify(message));
    }

    _heartCheck(i) {
        clearTimeout(_hTimer);
        _hTimer = setTimeout(() => {
            this._heartCheck(i);
        }, this._option.heartTime);

        this._sendMsg(i, {
            e: RequestEvents.Ping,
            d: {},
        });

        if (i.readyState !== 1) {
            i.close();
        }
    }

    // 公开方法

    getWs() {
        return this._ws;
    }

    pong() {
        const that = this;
        return new Promise((resolve, reject) => {
            that._sendCallback(
                that.getWs(),
                {
                    e: RequestEvents.Pong,
                    d: {},
                },
                (data) => {
                    resolve(data);
                },
            );
        });
    }

    close() {
        const that = this;
        return new Promise((resolve, reject) => {
            that._sendCallback(
                that.getWs(),
                {
                    e: RequestEvents.Close,
                    d: {},
                },
                (data) => {
                    resolve(data);
                },
            );
            that._ws.close();
        });
    }

    subscribe(topic) {
        const that = this;
        return new Promise((resolve, reject) => {
            that._sendCallback(
                that.getWs(),
                {
                    e: RequestEvents.Subscribe,
                    d: {
                        topic,
                    },
                },
                (data) => {
                    resolve(data);
                },
            );
        });
    }

    unsubscribe(topic) {
        const that = this;
        return new Promise((resolve, reject) => {
            that._sendCallback(
                that.getWs(),
                {
                    e: RequestEvents.Unsubscribe,
                    d: {
                        topic,
                    },
                },
                (data) => {
                    resolve(data);
                },
            );
        });
    }

    on(be, callback) {
        console.log(`WebSocket 绑定事件[${be}]...`);
        if(be){
            this._onEvent.set(be, callback);
        }
    }

    sendBroadcastMsg(message,bindEvent) {
        const that = this;
        return new Promise((resolve, reject) => {
            that._sendCallback(
                that.getWs(),
                {
                    be: bindEvent,
                    e: RequestEvents.BroadcastMsg,
                    d: {
                        message,
                    },
                },
                (data) => {
                    resolve(data);
                },
            );
        });
    }

    sendMsgBySid(toSessionId,bindEvent, message) {
        const that = this;
        return new Promise((resolve, reject) => {
            that._sendCallback(
                that.getWs(),
                {
                    be: bindEvent,
                    e: RequestEvents.IdMsg,
                    d: {
                        toSessionId,
                        message,
                    },
                },
                (data) => {
                    resolve(data);
                },
            );
        });
    }

    publish(topic, message) {
        const that = this;
        return new Promise((resolve, reject) => {
            that._sendCallback(
                that.getWs(),
                {
                    e: RequestEvents.Publish,
                    d: {
                        topic,
                        message,
                    },
                },
                (data) => {
                    resolve(data);
                },
            );
        });
    }
}

export default Ws;