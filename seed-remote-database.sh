#!/usr/bin/env bash

set -a
source .env
set +a

npx astro db execute ./db/seed.ts --remote
