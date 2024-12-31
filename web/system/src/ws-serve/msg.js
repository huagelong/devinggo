import Ws from '@/utils/ws';
import tool from "@/utils/tool.js";

class Msg {
    ws = null;
    constructor() {
        var wsUrl = import.meta.env.VITE_APP_WS_URL;
        var token = tool.local.get(import.meta.env.VITE_APP_TOKEN_PREFIX);
        this.ws = new Ws(wsUrl, {
            token: token,
        });
    }
}
export default Msg;