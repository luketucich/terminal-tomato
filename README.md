# 🍅 Terminal Tomato

A simple Pomodoro timer in your terminal using Go and Cobra.

## 📦 Install

```bash
go install github.com/luketucich/terminal-tomato@latest
```

## ⌛ Starting a Timer
```bash
terminal-tomato start
```

## ⚠️ Post-install Check

If running `terminal-tomato start` gives a **command not found** error, your terminal can't find the app.

Run this to check:
```bash
which terminal-tomato
```

If it returns nothing, add Go's bin directory to your PATH:

```bash
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc
source ~/.zshrc
```

*Note: Replace `~/.zshrc` with `~/.bashrc` if you use Bash instead of Zsh.*

## 💡 Optional: Create a shorter alias

For convenience, you can create an alias to use `tomato` instead of `terminal-tomato`:

```bash
echo 'alias tomato="terminal-tomato"' >> ~/.zshrc
source ~/.zshrc
```

Then you can use:
```bash
tomato start
```