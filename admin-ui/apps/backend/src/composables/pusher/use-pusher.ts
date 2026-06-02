import Pusher from 'pusher-js';
import type { Channel } from 'pusher-js';
import { readonly, ref } from 'vue';

import { useAccessStore } from '@vben/stores';

import { logger } from '#/utils/logger';

export interface PusherConfig {
  appKey: string;
  wsHost?: string;
  wsPort?: number;
  wssPort?: number;
  forceTLS?: boolean;
  authEndpoint?: string;
  cluster?: string;
}

// Default config — reads from env, falls back to sensible defaults
const defaultConfig: PusherConfig = {
  appKey: (import.meta.env.VITE_PUSHER_APP_KEY as string) || 'devinggo-app-key',
  wsHost: (import.meta.env.VITE_PUSHER_WS_HOST as string) || window.location.hostname,
  wsPort: (import.meta.env.VITE_PUSHER_WS_PORT as number) || 8070,
  wssPort: (import.meta.env.VITE_PUSHER_WSS_PORT as number) || 8070,
  forceTLS: window.location.protocol === 'https:',
  authEndpoint: `${import.meta.env.VITE_GLOB_API_URL}/system/pusher/auth`,
  cluster: (import.meta.env.VITE_PUSHER_CLUSTER as string) || 'local',
};

let pusherInstance: Pusher | null = null;
const connectionState = ref<string>('disconnected');

/**
 * Initialize or return the singleton Pusher client.
 */
export function usePusher(config?: Partial<PusherConfig>) {
  const merged = { ...defaultConfig, ...config };

  function getInstance(): Pusher {
    if (pusherInstance) return pusherInstance;

    pusherInstance = new Pusher(merged.appKey, {
      cluster: merged.cluster || 'local',
      wsHost: merged.wsHost,
      wsPort: merged.wsPort,
      wssPort: merged.wssPort,
      forceTLS: merged.forceTLS,
      enabledTransports: ['ws', 'wss'],
      authEndpoint: merged.authEndpoint,
      auth: {
        headers: {
          // Read token dynamically so expired tokens are refreshed
          get Authorization() {
            const accessStore = useAccessStore();
            return accessStore.accessToken ?? '';
          },
        },
      },
    });

    // Connection state monitoring
    pusherInstance.connection.bind(
      'state_change',
      (states: { current: string; previous: string }) => {
        connectionState.value = states.current;
        logger.debug('[Pusher] state:', states.previous, '→', states.current);
      },
    );

    pusherInstance.connection.bind('error', (err: Error) => {
      logger.error('[Pusher] connection error:', err);
    });

    return pusherInstance;
  }

  /** Subscribe to a public channel */
  function subscribe(channelName: string): Channel {
    return getInstance().subscribe(channelName);
  }

  /** Subscribe to a private channel (requires auth) */
  function subscribePrivate(channelName: string): Channel {
    const name = channelName.startsWith('private-')
      ? channelName
      : `private-${channelName}`;
    return getInstance().subscribe(name);
  }

  /** Subscribe to a presence channel (requires auth, tracks users) */
  function subscribePresence(channelName: string): Channel {
    const name = channelName.startsWith('presence-')
      ? channelName
      : `presence-${channelName}`;
    return getInstance().subscribe(name);
  }

  /** Unsubscribe from a channel */
  function unsubscribe(channelName: string): void {
    pusherInstance?.unsubscribe(channelName);
  }

  /** Bind a callback to a channel event, returns unbind fn */
  function bind<T = unknown>(
    channel: Channel,
    event: string,
    callback: (data: T) => void,
  ) {
    channel.bind(event, callback);
    return () => channel.unbind(event, callback);
  }

  /** Disconnect and clean up */
  function disconnect(): void {
    if (pusherInstance) {
      pusherInstance.disconnect();
      pusherInstance = null;
      connectionState.value = 'disconnected';
    }
  }

  return {
    getInstance,
    subscribe,
    subscribePrivate,
    subscribePresence,
    unsubscribe,
    bind,
    disconnect,
    state: readonly(connectionState),
  };
}
