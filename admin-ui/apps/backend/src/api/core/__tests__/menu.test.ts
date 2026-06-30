import { beforeEach, describe, expect, it, vi } from 'vitest';

import { getSystemInfoApi } from '../user';
import { getAllMenusApi } from '../menu';

vi.mock('../user', () => ({
  getSystemInfoApi: vi.fn(),
}));

describe('getAllMenusApi', () => {
  beforeEach(() => {
    vi.mocked(getSystemInfoApi).mockReset();
  });

  it('returns an empty menu list when backend routers is not an array', async () => {
    vi.mocked(getSystemInfoApi).mockResolvedValue({
      codes: [],
      roles: [],
      routers: null,
      user: {
        avatar: '',
        id: 1,
        nickname: 'Admin',
        username: 'admin',
      },
    } as any);

    await expect(getAllMenusApi()).resolves.toEqual([]);
  });
});
