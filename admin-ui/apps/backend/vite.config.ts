import { defineConfig } from '@vben/vite-config';
import { fileURLToPath } from 'node:url';
import path from 'node:path';

const __dirname = path.dirname(fileURLToPath(import.meta.url));

export default defineConfig(async () => {
  return {
    application: {},
    vite: {
      server: {
        fs: {
          // 允许访问 monorepo 的 packages 目录
          allow: [
            // 当前工作区
            path.resolve(__dirname, '../..'),
            // packages 目录（用于 @vben/icons 等包）
            path.resolve(__dirname, '../../packages'),
          ],
        },
        proxy: {
          '/api': {
            changeOrigin: true,
            rewrite: (path) => path.replace(/^\/api/, ''),
            // 后端 Go 服务地址
            target: 'http://localhost:8070',
            ws: true,
          },
        },
      },
    },
  };
});
