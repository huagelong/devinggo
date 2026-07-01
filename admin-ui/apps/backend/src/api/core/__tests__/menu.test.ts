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

  it('normalizes bare menu icon names to lucide icon keys', async () => {
    vi.mocked(getSystemInfoApi).mockResolvedValue({
      codes: [],
      roles: [],
      routers: [
        {
          children: [],
          component: 'system/menu/index',
          id: 1,
          meta: {
            hidden: false,
            hiddenBreadcrumb: false,
            icon: 'user',
            title: 'Menu',
            type: 'M',
          },
          name: 'SystemMenu',
          parent_id: 0,
          path: '/system/menu',
          redirect: '',
        },
      ],
      user: {
        avatar: '',
        id: 1,
        nickname: 'Admin',
        username: 'admin',
      },
    } as any);

    await expect(getAllMenusApi()).resolves.toMatchObject([
      {
        meta: {
          icon: 'lucide:user',
        },
      },
    ]);
  });
});
