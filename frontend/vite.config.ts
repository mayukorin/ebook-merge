// eslint-disable-next-line @typescript-eslint/triple-slash-reference
/// <reference path="./env/ProcessEnv.d.ts" />
import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import { resolve } from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    react({
      jsxImportSource: "@emotion/react",
    }),
  ],
  esbuild: {
    logOverride: { "this-is-undefined-in-esm": "silent" },
  },
  build: {
    rollupOptions: {
      output: {
        dir: resolve(__dirname, "dist"),
      },
    },
  },
  server: {
    host: true,
  },
  define: {
    "process.env.API_HOST": JSON.stringify(process.env.API_HOST),
    "process.env.FIREBASE_API_KEY": JSON.stringify(process.env.FIREBASE_API_KEY),
    "process.env.FIREBASE_AUTH_DOMAIN": JSON.stringify(process.env.FIREBASE_AUTH_DOMAIN),
  },
});