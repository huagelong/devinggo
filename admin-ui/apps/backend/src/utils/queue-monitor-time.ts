import dayjs from 'dayjs';

export function formatQueueMonitorTime(value: Date | string | undefined) {
  if (!value) {
    return '';
  }

  if (typeof value === 'string') {
    const match = value.match(
      /^(\d{4})-(\d{2})-(\d{2})[T\s](\d{2}):(\d{2}):(\d{2})/,
    );
    if (match) {
      const [, year, month, day, hour, minute, second] = match;
      return `${year}-${month}-${day} ${hour}:${minute}:${second}`;
    }
  }

  const instance = dayjs(value);
  if (!instance.isValid()) {
    return String(value);
  }

  return instance.format('YYYY-MM-DD H:mm:ss');
}
