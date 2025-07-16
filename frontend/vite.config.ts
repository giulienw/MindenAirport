import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";
import tsconfigPaths from "vite-tsconfig-paths";
import tailwindcss from "@tailwindcss/vite";
import electron from 'vite-plugin-electron';

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    react(), 
    tsconfigPaths(), 
    tailwindcss(),
    electron([
      {
        entry: 'electron/main.ts',
        onstart(options) {
          // Notify the Renderer-Process to reload the page when the Main-Process restart
          options.reload();
        },
        vite: {
          build: {
            outDir: 'dist-electron',
            rollupOptions: {
              external: ['electron']
            }
          }
        }
      }
    ])
  ],
  server: {
    port: 3000,
    proxy: {
      "/api": {
        target: "http://localhost:8080",
        changeOrigin: true,
      }
    }
  },
  build: {
    outDir: 'dist'
  }
});
