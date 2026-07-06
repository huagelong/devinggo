import { describe, expect, it } from 'vitest';

import {
  buildCrontabParameter,
  parseCrontabParameter,
} from '../utils/crontab-parameter';

describe('crontab parameter helpers', () => {
  it('builds common task parameters from key-value rows', () => {
    const result = buildCrontabParameter({
      commonParams: [
        { key: 'name', value: 'name1' },
        { key: 'value', value: 'shuju1' },
        { key: '', value: 'ignored' },
      ],
      mode: 'common',
      rawJson: '',
      url: {
        headers: [],
        method: 'get',
        params: [],
        url: '',
      },
      useRawJson: false,
    });

    expect(result).toBe(
      JSON.stringify({ name: 'name1', value: 'shuju1' }, null, 2),
    );
  });

  it('builds url task parameters from url form fields', () => {
    const result = buildCrontabParameter({
      commonParams: [],
      mode: 'url',
      rawJson: '',
      url: {
        headers: [{ key: 'Authorization', value: 'Bearer token' }],
        method: 'post',
        params: [{ key: 'page', value: '1' }],
        url: 'https://example.com/api',
      },
      useRawJson: false,
    });

    expect(result).toBe(
      JSON.stringify(
        {
          url: 'https://example.com/api',
          method: 'post',
          header: { Authorization: 'Bearer token' },
          params: { page: '1' },
        },
        null,
        2,
      ),
    );
  });

  it('parses existing url json into form fields', () => {
    const result = parseCrontabParameter(
      JSON.stringify({
        header: { Authorization: 'Bearer token' },
        method: 'get',
        params: { page: '1' },
        url: 'https://example.com/api',
      }),
      'url',
    );

    expect(result).toEqual({
      commonParams: [{ key: '', value: '' }],
      mode: 'url',
      rawJson: '',
      url: {
        headers: [{ key: 'Authorization', value: 'Bearer token' }],
        method: 'get',
        params: [{ key: 'page', value: '1' }],
        url: 'https://example.com/api',
      },
      useRawJson: false,
    });
  });

  it('keeps invalid json in advanced mode', () => {
    const result = parseCrontabParameter('{bad json', 'common');

    expect(result).toMatchObject({
      rawJson: '{bad json',
      useRawJson: true,
    });
  });
});
