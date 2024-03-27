/// <reference path="../.astro/db-types.d.ts" />
/// <reference path="../.astro/types.d.ts" />
/// <reference types="astro/client" />
interface ImportMetaEnv {
  readonly MONGODB_URI: string;
  readonly MONGODB_DB: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
