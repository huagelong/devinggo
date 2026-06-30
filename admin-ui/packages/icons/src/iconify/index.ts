import { addCollection, addIcon, createIconifyIcon } from '@vben-core/icons';

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

addIcon('mdi:keyboard-esc', {
  body: '<path fill="currentColor" d="M1 7h6v2H3v2h4v2H3v2h4v2H1zm10 0h4v2h-4v2h2a2 2 0 0 1 2 2v2c0 1.11-.89 2-2 2H9v-2h4v-2h-2a2 2 0 0 1-2-2V9c0-1.1.9-2 2-2m8 0h2a2 2 0 0 1 2 2v1h-2V9h-2v6h2v-1h2v1c0 1.11-.89 2-2 2h-2a2 2 0 0 1-2-2V9c0-1.1.9-2 2-2"/>',
  height: 24,
  width: 24,
});

export * from '@vben-core/icons';

export const MdiKeyboardEsc = createIconifyIcon('mdi:keyboard-esc');

export { loadIconifyCollection };
