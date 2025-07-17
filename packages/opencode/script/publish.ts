#!/usr/bin/env bun

import { $ } from "bun"

import pkg from "../package.json"

const dry = process.argv.includes("--dry")
const snapshot = process.argv.includes("--snapshot")

const version = snapshot
  ? `0.0.0-${new Date().toISOString().slice(0, 16).replace(/[-:T]/g, "")}`
  : await $`git describe --tags --abbrev=0`
      .text()
      .then((x) => x.substring(1).trim())
      .catch(() => {
        console.error("tag not found")
        process.exit(1)
      })

console.log(`publishing ${version}`)

const GOARCH: Record<string, string> = {
  arm64: "arm64",
  x64: "amd64",
}

const targets = [
  ["linux", "arm64"],
  ["linux", "x64"],
  ["darwin", "x64"],
  ["darwin", "arm64"],
  ["windows", "x64"],
]

await $`rm -rf dist`

const optionalDependencies: Record<string, string> = {}
const npmTag = snapshot ? "snapshot" : "latest"
for (const [os, arch] of targets) {
  console.log(`building ${os}-${arch}`)
  const name = `${pkg.name}-${os}-${arch}`
  await $`mkdir -p dist/${name}/bin`
  await $`CGO_ENABLED=0 GOOS=${os} GOARCH=${GOARCH[arch]} go build -ldflags="-s -w -X main.Version=${version}" -o ../opencode/dist/${name}/bin/tui ../tui/cmd/opencode/main.go`.cwd(
    "../tui",
  )
  await $`bun build --define OPENCODE_VERSION="'${version}'" --compile --minify --target=bun-${os}-${arch} --outfile=dist/${name}/bin/ollamacode ./src/index.ts ./dist/${name}/bin/tui`
  await $`rm -rf ./dist/${name}/bin/tui`
  await Bun.file(`dist/${name}/package.json`).write(
    JSON.stringify(
      {
        name,
        version,
        os: [os === "windows" ? "win32" : os],
        cpu: [arch],
      },
      null,
      2,
    ),
  )
  if (!dry) await $`cd dist/${name} && bun publish --access public --tag ${npmTag}`
  optionalDependencies[name] = version
}

await $`mkdir -p ./dist/${pkg.name}`
await $`cp -r ./bin ./dist/${pkg.name}/bin`
await $`cp ./script/postinstall.mjs ./dist/${pkg.name}/postinstall.mjs`
await Bun.file(`./dist/${pkg.name}/package.json`).write(
  JSON.stringify(
    {
      name: pkg.name + "-ai",
      bin: {
        [pkg.name]: `./bin/${pkg.name}`,
      },
      scripts: {
        postinstall: "node ./postinstall.mjs",
      },
      version,
      optionalDependencies,
    },
    null,
    2,
  ),
)
if (!dry) await $`cd ./dist/${pkg.name} && bun publish --access public --tag ${npmTag}`

if (!snapshot) {
  // Github Release
  for (const key of Object.keys(optionalDependencies)) {
    await $`cd dist/${key}/bin && zip -r ../../${key}.zip *`
  }

  const previous = await fetch("https://api.github.com/repos/88plug/ollamacode/releases/latest")
    .then((res) => res.json())
    .then((data) => data.tag_name)

  const commits = await fetch(`https://api.github.com/repos/88plug/ollamacode/compare/${previous}...HEAD`)
    .then((res) => res.json())
    .then((data) => data.commits || [])

  const notes = commits
    .map((commit: any) => `- ${commit.commit.message.split("\n")[0]}`)
    .filter((x: string) => {
      const lower = x.toLowerCase()
      return (
        !lower.includes("ignore:") &&
        !lower.includes("chore:") &&
        !lower.includes("ci:") &&
        !lower.includes("wip:") &&
        !lower.includes("docs:") &&
        !lower.includes("doc:")
      )
    })
    .join("\n")

  if (!dry) await $`gh release create v${version} --title "v${version}" --notes ${notes} ./dist/*.zip`

  // Calculate SHA values
  const arm64Sha = await $`sha256sum ./dist/${pkg.name}-linux-arm64.zip | cut -d' ' -f1`.text().then((x) => x.trim())
  const x64Sha = await $`sha256sum ./dist/${pkg.name}-linux-x64.zip | cut -d' ' -f1`.text().then((x) => x.trim())
  const macX64Sha = await $`sha256sum ./dist/${pkg.name}-darwin-x64.zip | cut -d' ' -f1`.text().then((x) => x.trim())
  const macArm64Sha = await $`sha256sum ./dist/${pkg.name}-darwin-arm64.zip | cut -d' ' -f1`.text().then((x) => x.trim())

  // AUR package
  const pkgbuild = [
    "# Maintainer: dax",
    "# Maintainer: adam",
    "",
    "pkgname='${pkg}'",
    `pkgver=${version.split("-")[0]}`,
    "options=('!debug' '!strip')",
    "pkgrel=1",
    "pkgdesc='The AI coding agent built for the terminal.'",
    "url='https://github.com/88plug/ollamacode'",
    "arch=('aarch64' 'x86_64')",
    "license=('MIT')",
    "provides=('opencode')",
    "conflicts=('opencode')",
    "depends=('fzf' 'ripgrep')",
    "",
    `source_aarch64=("\${pkgname}_\${pkgver}_aarch64.zip::https://github.com/88plug/ollamacode/releases/download/v${version}/ollamacode-linux-arm64.zip")`,
    `sha256sums_aarch64=('${arm64Sha}')`,
    "",
    `source_x86_64=("\${pkgname}_\${pkgver}_x86_64.zip::https://github.com/88plug/ollamacode/releases/download/v${version}/ollamacode-linux-x64.zip")`,
    `sha256sums_x86_64=('${x64Sha}')`,
    "",
    "package() {",
    `  install -Dm755 ./${pkg.name} "\${pkgdir}/usr/bin/${pkg.name}"`,
    "}",
    "",
  ].join("\n")

  for (const aurPkg of ["ollamacode", "ollamacode-bin"]) {
    await $`rm -rf ./dist/aur-${aurPkg}`
    await $`git clone ssh://aur@aur.archlinux.org/${aurPkg}.git ./dist/aur-${aurPkg}`
    await Bun.file(`./dist/aur-${aurPkg}/PKGBUILD`).write(pkgbuild.replace("${pkg}", aurPkg))
    await $`cd ./dist/aur-${aurPkg} && makepkg --printsrcinfo > .SRCINFO`
    await $`cd ./dist/aur-${aurPkg} && git add PKGBUILD .SRCINFO`
    await $`cd ./dist/aur-${aurPkg} && git commit -m "Update to v${version}"`
    if (!dry) await $`cd ./dist/aur-${aurPkg} && git push`
  }

  // Homebrew formula
  const homebrewFormula = [
    "# typed: false",
    "# frozen_string_literal: true",
    "",
    "# This file was generated by GoReleaser. DO NOT EDIT.",
    "class Ollamacode < Formula",
    `  desc "The AI coding agent built for the terminal."`,
    `  homepage "https://github.com/88plug/ollamacode"`,
    `  version "${version.split("-")[0]}"`,
    "",
    "  on_macos do",
    "    if Hardware::CPU.intel?",
    `      url "https://github.com/88plug/ollamacode/releases/download/v${version}/ollamacode-darwin-x64.zip"`,
    `      sha256 "${macX64Sha}"`,
    "",
    "      def install",
    '        bin.install "${pkg.name}"',
    "      end",
    "    end",
    "    if Hardware::CPU.arm?",
    `      url "https://github.com/88plug/ollamacode/releases/download/v${version}/ollamacode-darwin-arm64.zip"`,
    `      sha256 "${macArm64Sha}"`,
    "",
    "      def install",
    '        bin.install "${pkg.name}"',
    "      end",
    "    end",
    "  end",
    "",
    "  on_linux do",
    "    if Hardware::CPU.intel? and Hardware::CPU.is_64_bit?",
    `      url "https://github.com/88plug/ollamacode/releases/download/v${version}/ollamacode-linux-x64.zip"`,
    `      sha256 "${x64Sha}"`,
    "      def install",
    '        bin.install "${pkg.name}"',
    "      end",
    "    end",
    "    if Hardware::CPU.arm? and Hardware::CPU.is_64_bit?",
    `      url "https://github.com/88plug/ollamacode/releases/download/v${version}/ollamacode-linux-arm64.zip"`,
    `      sha256 "${arm64Sha}"`,
    "      def install",
    '        bin.install "${pkg.name}"',
    "      end",
    "    end",
    "  end",
    "end",
    "",
    "",
  ].join("\n")

  await $`rm -rf ./dist/homebrew-tap`
  await $`git clone https://${process.env["GITHUB_TOKEN"]}@github.com/88plug/homebrew-tap.git ./dist/homebrew-tap`
  await Bun.file("./dist/homebrew-tap/ollamacode.rb").write(homebrewFormula)
  await $`cd ./dist/homebrew-tap && git add ollamacode.rb`
  await $`cd ./dist/homebrew-tap && git commit -m "Update to v${version}"`
  if (!dry) await $`cd ./dist/homebrew-tap && git push`
}
