---
name: project-todo-cli
description: "Single-file Go CLI for learning Go — `todo` with add/list/complete/delete subcommands, persisted to tasks.json"
metadata: 
  node_type: memory
  type: project
  originSessionId: 2e516d5b-7035-43e7-a789-458e877ce1d5
---

`~/golang/claude-course/todo/` — a small Go CLI built as a learning exercise.

**Why:** Training project, not production code. The goal is for the user to write Go and learn the language, so simplicity beats robustness when there's a tradeoff.

**How to apply:**
- When suggesting an approach, default to the simpler option unless the user asks about alternatives.
- Don't push for premature abstractions, extra packages, DI, or tests unless the user brings them up.
- Guide step-by-step per [[feedback-guidance-style]] — don't write the whole feature for them.

**Agreed design decisions (so I don't re-litigate them):**
- Single `main.go`, no extra packages.
- Subcommand CLI: `add <title>`, `list`, `complete <id>`, `delete <id>`.
- Storage file: `./tasks.json` (cwd-relative).
- ID strategy: max existing ID + 1, recomputed on each add (rejected: a `next_id` counter stored in the JSON — overkill for this scale).
- Handler shape: each `cmd*` function does load → mutate → save end-to-end (option A from our discussion); `main()` dispatch stays a dumb switch.

**Output spec for `list`** (agreed but not yet implemented):
```
[ ] 1: Buy milk (created just now)
[x] 2: Take out trash (created 1 minute ago)
```
Relative time buckets: `just now` / `N minute(s) ago` / `N hour(s) ago` / `N day(s) ago`. Singular vs plural matters (`1 minute ago`, not `1 minutes ago`).

**Repo:** `github.com/igorfazlyev/todo` — origin already set, `main` tracking. SSH setup quirks: see [[reference-github-ssh]].
