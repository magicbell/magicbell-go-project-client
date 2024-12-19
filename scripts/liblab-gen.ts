import { execSync } from "node:child_process";
import * as fs from "node:fs/promises";
import * as path from "node:path";

async function move(oldPath: string, newPath: string) {
  await fs.rm(newPath, { recursive: true, force: true });
  await fs.mkdir(path.dirname(newPath), { recursive: true });
  await fs.rename(oldPath, newPath);
}

async function readJSON(path: string) {
  return fs
    .readFile(path, { encoding: "utf-8" })
    .then((data) => JSON.parse(data));
}

function hasChangesInPath(path: string) {
  try {
    execSync(`git diff --quiet HEAD ${path}`, { stdio: "ignore" });
    return false;
  } catch (error) {
    return true;
  }
}

async function build() {
  let pkgJson = await readJSON("./package.json");
  const liblabConfig = await readJSON("./liblab.config.json");
  console.log(
    `Building '${pkgJson.name}' using spec '${liblabConfig.specFilePath}'`
  );

  execSync("npx -y liblab@latest build -y --skip-validation", {
    stdio: "inherit",
  });

  await move('output/go/pkg', './pkg');
  await move('output/go/internal', './internal');
  await move('output/go/go.mod', './go.mod');

  if (!hasChangesInPath("./pkg") && !hasChangesInPath("./internal") && !hasChangesInPath("./go.mod")) {
    console.log("No changes detected in output.");
    return;
  }

  await move('output/go/cmd', './cmd');
  await move('output/go/documentation', './documentation');
  await move('output/go/README.md', './README.md');
}

void build().finally(() => fs.rm("output", { recursive: true, force: true }));
