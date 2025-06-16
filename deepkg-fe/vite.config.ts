import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'

// https://vite.dev/config/
export default defineConfig({
    build: {
      outDir: 'build',
    },
    base: './',
    plugins: [
        react()
    ],
    resolve: {
    },
    css: {
      preprocessorOptions: {
        less: {
          modifyVars: {
            //引入less基础变量
            hack: `true;`,
          },
          charset: false,
          javascriptEnabled: true,
        }
      }
    },


})