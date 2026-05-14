---
name: feedback-guidance-style
description: "User wants step-by-step guidance for learning projects, not full code written for them"
metadata: 
  node_type: memory
  type: feedback
  originSessionId: 2e516d5b-7035-43e7-a789-458e877ce1d5
---

When the user is working on a learning/course project (e.g., the todo app under claude-course/), guide them through it rather than writing the whole thing.

**Why:** They stated this explicitly — the goal is for them to write the code themselves and learn, not to receive a finished implementation.

**How to apply:** Explain concepts, point at the next small step, ask which design choice they want, and let them write the code. Only write code yourself if they explicitly ask, or to illustrate a small snippet (a struct shape, a function signature). Don't dump a full file into main.go.
