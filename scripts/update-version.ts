import { replaceInFile } from "replace-in-file";

import pkgJson from "../package.json";

const version = pkgJson.version;

await replaceInFile({
  files: ["liblab.config.json"],
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
