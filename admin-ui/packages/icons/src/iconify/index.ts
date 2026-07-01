import { addCollection, addIcon, createIconifyIcon } from '@vben-core/icons';

import { localIconEntries } from './menu-icons';

const localCollectionLoaders = {
  carbon: () => import('@iconify-json/carbon/icons.json'),
  lucide: () => import('@iconify-json/lucide/icons.json'),
  mdi: () => import('@iconify-json/mdi/icons.json'),
};

const loadedLocalCollections = new Set<string>();

async function loadIconifyCollection(prefix: string) {
  const loader =
    localCollectionLoaders[prefix as keyof typeof localCollectionLoaders];
  if (!loader || loadedLocalCollections.has(prefix)) {
    return;
  }

  const collection = await loader();
  addCollection(collection.default);
  loadedLocalCollections.add(prefix);
}

localIconEntries.forEach(([name, data]) => addIcon(name, data));

export * from '@vben-core/icons';
export * from './menu-icons';

export const MdiKeyboardEsc = createIconifyIcon('mdi:keyboard-esc');

export { loadIconifyCollection };
