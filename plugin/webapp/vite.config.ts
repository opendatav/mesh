import {defineConfig} from 'vite';
import react from '@vitejs/plugin-react';
import svg from '@svgr/rollup';
import {resolve} from "path";

// https://vitejs.dev/config/
export default defineConfig({
    base: 'mesh',
    plugins: [react(), svg()],
    resolve: {
        alias: {
            '@': resolve(__dirname, './src'),
        },
    },
    build: {
        outDir: resolve(__dirname, './static'),
    },
    assetsInclude: ['node_modules/vscode-oniguruma/**/*.wasm'],
    server: {
        proxy: {
            '/mpc/invoke': {
                secure: false,
                target: 'https://127.0.0.1',
            },
            '/mesh/lsp': {
                rewrite: path => path.substring(5),
                secure: false,
                target: 'ws://127.0.0.1:8889',
            }
        },
        fs: {
            allow: [
                resolve(__dirname, '../../client/javascript/dist'),
                resolve(__dirname, './'),
            ]
        }
    }
})
