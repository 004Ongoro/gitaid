# Gitaid ✅

**AI-powered git commit generator** that analyzes your staged diff and suggests a Conventional Commits–style message using Google Gemini (via the `genai` client).

---

## 🔍 What is Gitaid?

Gitaid stands for **Git AI Daemon** it reads your staged Git changes (what you'd commit with `git commit`), generates a clear, conventional commit message using a language model, and offers an interactive flow to accept, regenerate, or discard the suggestion.

- Purpose: reduce commit message friction and improve commit quality.
- Behavior: reads `git diff --cached` and prompts Gemini(Any model, everyone uses their own API Key model which are stored locally in json file, eg i create my json file at C:/Users/george/.config/gitaid/config.json) to generate a Conventional Commit message.

contents of C:/Users/your_name/.config/gitaid/config.json

```json
{
  "gemini_key": "YOUR_API_KEY",
  "model": "gemini-flash-latest"
}
```

---

## ⚙️ Key features

- Generates Conventional Commits-format messages (types like `feat`, `fix`, `docs`, etc.)
- Interactive CLI: accept, regenerate, or quit
- Configurable model (defaults to `gemini-1.5-flash`)
- Designed for simple pre-commit or manual usage

---

## 💾 Install

There are several ways to install Gitaid.

### 1) From GitHub Releases (recommended)

Download the appropriate binary from the Releases page:

```bash
https://github.com/004Ongoro/gitaid/releases
```

Unpack and place the binary in your `PATH`.

### 2) Homebrew (macOS / Linux)

Goreleaser is configured to publish a Homebrew formula to `004Ongoro/homebrew-tap`.

```bash
brew tap 004Ongoro/homebrew-tap
brew install gitaid
```

### 3) Scoop (Windows)

Install Scoop package manager if you havent already, it handles downloads of CLI tools and adds them to PATH

A Scoop bucket `004Ongoro/scoop-bucket` is configured in `goreleaser`.

```powershell
scoop bucket add 004Ongoro https://github.com/004Ongoro/scoop-bucket
scoop install gitaid
```

### 4) From source (Go)

You can build or install directly using Go (Go 1.20+ recommended):

```bash
# Build locally
git clone https://github.com/004Ongoro/gitaid.git
cd gitaid
go build -o gitaid ./...

# Or install via `go install`
go install github.com/004Ongoro/gitaid@latest
```

---

## 🔧 Configuration

Gitaid expects a JSON config file at:

- Linux/macOS/WSL: `$HOME/.config/gitaid/config.json`
- Windows: `%USERPROFILE%\.config\gitaid\config.json`

Example `config.json`:

```json
{
  "gemini_key": "YOUR_GEMINI_API_KEY",
  "model": "gemini-1.5-flash"
}
```

- `gemini_key` (required): API key used by the Gemini client.
- `model` (optional): model name to use; defaults to `gemini-1.5-flash`.

Important: do not commit your API key to source control.

---

## ▶️ Usage

1. Stage your changes:

```bash
git add <files>
```

2. Run Gitaid in the repository root:

```bash
gitaid
```

Gitaid will:
- Validate it's inside a git repo
- Get the staged diff (`git diff --cached`)
- Send the diff + system prompt to Gemini
- Print the suggested commit message and prompt you:
  - `(a) accept & commit` — runs `git commit -m "..."`
  - `(r) regenerate` — re-generate with the same diff
  - `(q) quit` — discard message and exit

Example interactive flow:

> Analyzing changes and generating commit message...
>
> --- Suggested Commit Message ---
> feat(auth): add session timeout enforcement
>
> The commit adds a 30-minute session timeout which expires tokens and
> redirects users to the login page. This prevents long-lived sessions and
> reduces risk in case of compromised tokens.
> --------------------------------
>
> Options: (a)ccept & commit, (r)egenerate, (q)uit:

---



## ⚠️ Troubleshooting

- "config file not found" — create the config file at `$HOME/.config/gitaid/config.json` with `gemini_key` as shown above.
- "not a git repository" — run `gitaid` from inside a git repository.
- "no staged changes found" — stage changes with `git add` before running.
- "AI generation failed" — check your API key, quota, and network connectivity.

Privacy note: The diff content is sent to the Gemini API — avoid staging secrets or PII.

---

## Contributing ✨

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a branch: `git checkout -b feat/my-feature`
3. Make changes and add tests where appropriate
4. Run `go vet`, `gofmt`, and `go test` (or your preferred linters)
5. Open a Pull Request and describe the change

Guidelines:

- Keep changes small and focused
- Follow Go idioms and existing code style
- Write tests for new behavior when possible

Maintainers will review, request changes if needed, and merge once ready.

---

## Releasing

This project uses `goreleaser` for cross-platform builds and publishing to GitHub Releases, Homebrew, and Scoop.

Typical release flow (maintainers):

```bash
export GORELEASER_TOKEN="<token-with-repo-scope>"
goreleaser release --rm-dist
```

See `.goreleaser.yaml` for configured taps/buckets and targets.

---

## License 📄

Gitaid is released under the **MIT License**. See `LICENSE` for details.

---

## Authors & Acknowledgements

- [George Ongoro](https://ongoro.top) — original author
- See `go.mod` for dependencies (including `google.golang.org/genai`)

---

If anything is missing or you'd like docs expanded (examples, pre-commit scripts, or CI samples), please open an issue or a PR — contributions are very welcome! ✅
