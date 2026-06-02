const isDev = import.meta.env.DEV;

function error(...args: unknown[]) {
  if (isDev) {
    console.error(...args);
  }
}

function warn(...args: unknown[]) {
  if (isDev) {
    console.warn(...args);
  }
}

function debug(...args: unknown[]) {
  if (isDev) {
    console.debug(...args);
  }
}

function info(...args: unknown[]) {
  if (isDev) {
    console.info(...args);
  }
}

export const logger = { debug, error, info, warn };