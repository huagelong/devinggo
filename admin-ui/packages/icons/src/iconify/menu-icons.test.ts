import { describe, expect, it } from 'vitest';

import { localIconEntries, menuIconNames } from './menu-icons';

describe('menu icon presets', () => {
  it('exports unique bare lucide names for the menu icon picker', () => {
    expect(menuIconNames).toHaveLength(40);
    expect(new Set(menuIconNames).size).toBe(menuIconNames.length);
    expect(menuIconNames.every((name) => !name.includes(':'))).toBe(true);
  });

  it('has local icon data for each selectable menu icon', () => {
    for (const name of menuIconNames) {
      expect(
        localIconEntries.some(([icon]) => icon === `lucide:${name}`),
      ).toBe(true);
    }
  });
});
