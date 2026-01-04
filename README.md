# godsa2026


Daily DSA practice using **GitHub Gists** as a public thinking log.

Goal:
- One problem per day
- Clean solution
- Short explanation
- Minimal friction

---

## Folder Structure (local)

Each day lives in its own folder before publishing:

---

## Daily Workflow (5–10 minutes)

1. Create a new daily folder (e.g. `dsa_03_01_2026`)
2. Solve one DSA problem
3. Fill in:
   - `solution.go` → clean, readable code
   - `problem.md` → short summary (no copy-paste)
   - `notes.md` → insights, trade-offs, complexity
4. Publish as a **public gist**
5. Move on
---

## Requirements

- GitHub account
- GitHub CLI installed
- Authenticated with GitHub

```bash
gh auth login
gh auth status

```

## Deploying Gists via CLI (Primary Flow)
```bash
gh gist create solution.go problem.md notes.md \
  -d "DSA Daily — Contains Duplicate (LeetCode 217)" \
  -p
```
