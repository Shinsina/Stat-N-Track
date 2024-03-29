import { defineConfig } from "astro/config";
import tailwind from "@astrojs/tailwind";
import prefetch from "@astrojs/prefetch";
import sitemap from "@astrojs/sitemap";
import db from "@astrojs/db";

// https://astro.build/config
export default defineConfig({
  site: "https://shinsina.github.io",
  base: "/Stat-N-Track",
  integrations: [tailwind(), prefetch(), sitemap(), db()],
});
