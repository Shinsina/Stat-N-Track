import { defineConfig } from 'astro/config';
import tailwind from "@astrojs/tailwind";
import prefetch from "@astrojs/prefetch";

// https://astro.build/config
export default defineConfig({
  site: "https://shinsina.github.io",
  base: "/Stat-N-Track",
  integrations: [tailwind(), prefetch()]
});
