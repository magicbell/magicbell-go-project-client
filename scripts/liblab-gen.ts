#! /usr/bin/env node --experimental-strip-types
/* eslint-disable no-console */
import { execSync } from 'node:child_process';
import * as fs from 'node:fs/promises';
import * as path from 'node:path';
import { parseArgs } from 'node:util';
import { replaceInFile } from 'replace-in-file';
import * as md from '@magicbell/codegen/markdown';

async function move(oldPath: string, newPath: string) {
  await fs.rm(newPath, { recursive: true, force: true });
  await fs.mkdir(path.dirname(newPath), { recursive: true });
  await fs.rename(oldPath, newPath);
}

async function readJSON(path: string) {
  return JSON.parse(await readFile(path));
}

async function writeJSON(path: string, data: any) {
  data = typeof data === 'string' ? data : JSON.stringify(data, null, 2) + '\n';
  return fs.writeFile(path, data, { encoding: 'utf-8' });
}

async function readFile(path: string): Promise<string> {
  if (/^https?:\/\//.test(path)) {
    return fetch(path, { headers: { 'content-type': 'application/json' } }).then((x) => x.text());
  }

  return fs.readFile(path, { encoding: 'utf-8' });
}

const { values: args } = parseArgs({
  options: {
    spec: { type: 'string' },
    target: { type: 'string' },
  },
});

function hasChangesInPath(path: string) {
  try {
    execSync(`git add .`);
    execSync(`git diff --quiet HEAD ${path}`, { stdio: 'ignore' });
    return false;
  } catch (error) {
    return true;
  }
}


async function writeReadme(source: string, dest: string, title: string) {
  // process readme
  const readmeAst = await md.read(source);
  md.removeAllBeforeHeading(readmeAst, 'Setup & Configuration');
  md.reIndentHeadings(readmeAst, 1);
  md.mapLinks(readmeAst, (url) => {
    if (!url.startsWith('documentation/')) return;
    return url.replace('documentation/', '');
  });
  md.insertFrontMatter(readmeAst, {
    title: title,
  });
  await md.write(readmeAst, dest);
}

async function mergeReadme(source: string, dest: string, block: string) {
  const sourceAst = await md.read(source);
  md.removeAllBeforeHeading(sourceAst, 'Authentication');
  md.mapLinks(sourceAst, (url) => {
    if (url.startsWith('http://') || url.startsWith('https://')) return url;
    return path.join(path.dirname(source), url);
  });

  const mainAst = await md.read(dest);
  md.replaceBlock(mainAst, block, sourceAst);
  await md.write(mainAst, dest);
}

async function build(options: typeof args) {
  if (!options.spec) {
    console.log(`error: missing required argument '--spec <file>'`);
    process.exit(1);
  }

  if (!options.target) {
    console.log(`error: missing required argument '--target <dir>'`);
    process.exit(1);
  }

  const dest = `./pkg/${args.target}`;
  const liblabCfgFile = `./liblab.config.${args.target}.json`;
  const liblabConfig = await readJSON(liblabCfgFile);

  console.log(`Building '${args.target}' using spec '${options.spec}'`);
  const openapiJSON = await readFile(options.spec);
  await writeJSON(liblabConfig.specFilePath, openapiJSON);

  execSync(`npx -y liblab@latest build --skip-validation --approve-docs --liblab-config=${liblabCfgFile}`, {
    stdio: "inherit",
  });
  await fs.rm(liblabConfig.specFilePath);


  await move('output/go/pkg', `${dest}`);
  await move('output/go/internal', `${dest}/internal`);
  await move('output/go/go.mod', `./go.mod`);

  if (!hasChangesInPath(dest)) {
    console.log("No changes detected in output.");
    return;
  }

  await move('output/go/cmd', `./cmd/${args.target}`);
  await move('output/go/documentation', `./docs/${args.target}`);
  await writeReadme('output/go/README.md', `./docs/${args.target}/README.md`, args.target);

  const modName = liblabConfig.languageOptions.go.goModuleName;
  const files = [
    `./cmd/${args.target}/**/*.go`,
    `./pkg/${args.target}/**/*.go`,
    `./docs/${args.target}/**/*.{md,json}`,
  ];

  await replaceInFile({
    files,
    from: new RegExp(`${modName}/pkg/`, 'g'),
    to: `${modName}/pkg/${args.target}/`,
  })

  await replaceInFile({
    files,
    from: new RegExp(`${modName}/internal/`, 'g'),
    to: `${modName}/pkg/${args.target}/internal/`,
  })

  await mergeReadme(`./docs/${args.target}/README.md`, './README.md', args.target);
}

void build(args).finally(() => fs.rm("output", { recursive: true, force: true }));
