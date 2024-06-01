/// <reference types="vite/client" />

import { ConfigEnv, defineConfig, loadEnv, UserConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";

export default defineConfig(({ mode }: ConfigEnv): UserConfig => {
  process.env = { ...process.env, ...loadEnv(mode, process.cwd()) };

  return {
    server: {
      port: 30000
    },
    plugins: [svelte()],
    optimizeDeps: { exclude: ["svelte-navigator"] },
    base: process.env.VITE_BASE_ROUTE
  };
});
