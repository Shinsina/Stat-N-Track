import { defineConfig } from "astro/config";
import tailwind from "@astrojs/tailwind";
import prefetch from "@astrojs/prefetch";
import alpinejs from "@astrojs/alpinejs";
import sitemap from "@astrojs/sitemap";
import db from "@astrojs/db";

// https://astro.build/config
export default defineConfig({
  site: "https://shinsina.github.io",
  base: "/Stat-N-Track",
  integrations: [
    tailwind(),
    prefetch(),
    alpinejs({ entrypoint: "/src/lib/alpine" }),
    sitemap(),
    db(),
  ],
});
