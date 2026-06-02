/**
 * Pusher channel names & event names for the admin system.
 * Keep in sync with backend modules/system/pkg/websocket/ constants.
 */

// Channel name patterns
export const Channels = {
  // Private user channel — receives personal notifications
  user(userId: number | string) {
    return `private-user-${userId}`;
  },

  // System-wide public channel for broadcast messages
  systemBroadcast: 'system-broadcast',

  // Presence channel for online admin users
  adminPresence: 'presence-admin',
} as const;

// Event names
export const Events = {
  // New notification/message pushed to a user
  NOTIFICATION_NEW: 'notification:new',

  // Notification read status updated
  NOTIFICATION_READ: 'notification:read',

  // System broadcast message
  SYSTEM_NOTICE: 'system:notice',

  // User came online / went offline (presence)
  MEMBER_ADDED: 'pusher:member_added',
  MEMBER_REMOVED: 'pusher:member_removed',

  // Subscription succeeded (presence)
  SUBSCRIPTION_SUCCESS: 'pusher:subscription_succeeded',
} as const;

// Event data types
export interface NotificationNewData {
  id: number;
  title: string;
  content: string;
  content_type: string;
  send_by: number;
  send_user?: { id: number; nickname: string; username: string };
  created_at: string;
}

export interface NotificationReadData {
  ids: number[];
}

export interface SystemNoticeData {
  title: string;
  content: string;
  type: string;
  created_at: string;
}

export interface PresenceMemberData {
  id: number;
  nickname: string;
  username: string;
}
