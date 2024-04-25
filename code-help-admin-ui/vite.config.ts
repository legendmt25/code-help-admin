import { defineConfig, loadEnv } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
  process.env = { ...process.env, ...loadEnv(mode, process.cwd()) };

  return {
    plugins: [svelte()],
    optimizeDeps: { exclude: ["svelte-navigator"] },
    base: process.env.VITE_BASE_ROUTE
  };
});
