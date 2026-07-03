import { describe, expect, it } from 'vitest';

import { formatQueueMonitorTime } from '../queue-monitor-time';

describe('formatQueueMonitorTime', () => {
  it('formats RFC3339 strings to yyyy-mm-dd HH:mm:ss', () => {
    expect(formatQueueMonitorTime('2026-07-03T03:44:00Z')).toBe('2026-07-03 03:44:00');
  });

  it('formats Date objects to yyyy-mm-dd HH:mm:ss', () => {
    expect(formatQueueMonitorTime(new Date('2026-07-03T11:08:09+08:00'))).toBe('2026-07-03 11:08:09');
  });

  it('returns empty string for empty values', () => {
    expect(formatQueueMonitorTime('')).toBe('');
  });
});
