<p align="center">
  <a href="https://opencode.ai">
    <picture>
      <source srcset="packages/web/src/assets/logo-ornate-dark.svg" media="(prefers-color-scheme: dark)">
      <source srcset="packages/web/src/assets/logo-ornate-light.svg" media="(prefers-color-scheme: light)">
      <img src="packages/web/src/assets/logo-ornate-light.svg" alt="ollamacode logo">
    </picture>
  </a>
</p>
<p align="center">AI coding agent with local model support, built for the terminal.</p>
<p align="center">
  <a href="https://www.npmjs.com/package/ollamacode-ai"><img alt="npm" src="https://img.shields.io/npm/v/ollamacode-ai?style=flat-square" /></a>
  <a href="https://github.com/88plug/ollamacode/actions/workflows/publish.yml"><img alt="Build status" src="https://img.shields.io/github/actions/workflow/status/88plug/ollamacode/publish.yml?style=flat-square&branch=dev" /></a>
</p>

[![ollamacode Terminal UI](packages/web/src/assets/lander/screenshot.png)](https://opencode.ai)

---

### Installation

```bash
# YOLO
curl -fsSL https://opencode.ai/install | bash

# Package managers
npm i -g ollamacode-ai@latest      # or bun/pnpm/yarn
brew install 88plug/tap/ollamacode # macOS
paru -S ollamacode-bin             # Arch Linux
```

> [!TIP]
> Remove versions older than 0.1.x before installing.

### Local Model Support

ollamacode supports local models via:
- **Ollama**: Run models locally with `ollama serve` (auto-detected on localhost:11434)
- **vLLM**: High-performance inference with `vllm serve` (auto-detected on localhost:8000)

All existing cloud providers (Anthropic, OpenAI, etc.) are also supported.

### Documentation

For more info on configuration, see the upstream [opencode docs](https://opencode.ai/docs).

### Contributing

opencode is an opinionated tool so any fundamental feature needs to go through a
design process with the core team.

> [!IMPORTANT]
> We do not accept PRs for core features.

However we still merge a ton of PRs - you can contribute:

- Bug fixes
- Improvements to LLM performance
- Support for new providers
- Fixes for env specific quirks
- Missing standard behavior
- Documentation

Take a look at the git history to see what kind of PRs we end up merging.

> [!NOTE]
> If you do not follow the above guidelines we might close your PR.

To run opencode locally you need.

- Bun
- Golang 1.24.x

And run.

```bash
$ bun install
$ bun run packages/opencode/src/index.ts
```

#### Development Notes

**API Client**: After making changes to the TypeScript API endpoints in `packages/opencode/src/server/server.ts`, you will need the opencode team to generate a new stainless sdk for the clients.

### FAQ

#### How is this different than Claude Code?

It's very similar to Claude Code in terms of capability. Here are the key differences:

- 100% open source
- Not coupled to any provider. Although Anthropic is recommended, opencode can be used with OpenAI, Google or even local models. As models evolve the gaps between them will close and pricing will drop so being provider agnostic is important.
- A focus on TUI. opencode is built by neovim users and the creators of [terminal.shop](https://terminal.shop); we are going to push the limits of what's possible in the terminal.
- A client/server architecture. This for example can allow opencode to run on your computer, while you can drive it remotely from a mobile app. Meaning that the TUI frontend is just one of the possible clients.

#### What's the other repo?

The other confusingly named repo has no relation to this one. You can [read the story behind it here](https://x.com/thdxr/status/1933561254481666466).

---

**Join our community** [Discord](https://discord.gg/opencode) | [YouTube](https://www.youtube.com/c/sst-dev) | [X.com](https://x.com/SST_dev)
