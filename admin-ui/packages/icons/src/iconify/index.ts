import { addCollection, createIconifyIcon } from '@vben-core/icons';

// 导入本地图标集
import lucideIcons from '@iconify-json/lucide/icons.json';
import mdiIcons from '@iconify-json/mdi/icons.json';
import carbonIcons from '@iconify-json/carbon/icons.json';

// 注册本地图标集
addCollection(lucideIcons);
addCollection(mdiIcons);
addCollection(carbonIcons);

export * from '@vben-core/icons';

export const MdiKeyboardEsc = createIconifyIcon('mdi:keyboard-esc');
