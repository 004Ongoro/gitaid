# Roadmap: Project Gitaid

## Phase 1: Core Functionality (Current)
- [ ] Initialize Go module and project structure.
- [ ] Implement local configuration management (JSON/YAML).
- [ ] Integrate Google Gemini API client.
- [ ] Logic for `git diff --cached` extraction.
- [ ] Prompt engineering for Detailed Conventional Commits.

## Phase 2: User Experience & CLI
- [ ] Interactive mode: Prompt user to accept, edit, or regenerate the message.
- [ ] "Breaking Change" detection (adding `!` to the commit type).
- [ ] Support for custom system prompts via config.

## Phase 3: Infrastructure & CI/CD
- [ ] Setup GitHub Actions for testing.
- [ ] Configure `GoReleaser` for cross-platform binaries.
- [ ] Automate releases to **Homebrew** (macOS/Linux) and **Scoop** (Windows).

## Phase 4: Community & Documentation
- [ ] Create `CONTRIBUTING.md`.
- [ ] Comprehensive README with installation via `brew` and `scoop`.
- [ ] Example gallery of generated commits.