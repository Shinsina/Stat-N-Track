import { defineConfig } from "astro/config";
import tailwind from "@astrojs/tailwind";
import alpinejs from "@astrojs/alpinejs";
import sitemap from "@astrojs/sitemap";
import db from "@astrojs/db";

// https://astro.build/config
export default defineConfig({
  site: "https://shinsina.github.io",
  base: "/Stat-N-Track",
  integrations: [
    tailwind(),
    alpinejs({ entrypoint: "/src/lib/alpine" }),
    sitemap(),
    db(),
  ],
  prefetch: {
    defaultStrategy: 'viewport',
    prefetchAll: true
  },
  image: {
    domains: ['images-static.iracing.com']
  }
});
