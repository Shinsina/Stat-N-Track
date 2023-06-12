import { defineConfig } from 'astro/config';
import image from '@astrojs/image';
import tailwind from "@astrojs/tailwind";
import prefetch from "@astrojs/prefetch";
import netlify from "@astrojs/netlify/functions";

// https://astro.build/config
export default defineConfig({
  output: "server",
  adapter: netlify(),
  integrations: [tailwind(), prefetch(), image({ serviceEntryPoint: '@astrojs/image/sharp' })]
});
