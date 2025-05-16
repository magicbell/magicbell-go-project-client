#! /usr/bin/env -S node --experimental-strip-types
import { replaceInFile } from "replace-in-file";
import fs from 'node:fs/promises';

const pkgJson = JSON.parse(await fs.readFile('./package.json', 'utf-8'));
const version = pkgJson.version;

await replaceInFile({
  files: ["liblab.config.*.json"],
  from: /"sdkVersion": "\d.\d.\d"/g,
  to: `"sdkVersion": "${version}"`,
});

await replaceInFile({
  files: ["README.md"],
  from: [
    /SDK \d.\d.\d/g,
    /SDK version: `\d.\d.\d`/g,
  ],
  to: [
    `SDK ${version}`,
    `SDK version: \`${version}\``,
  ]
});
