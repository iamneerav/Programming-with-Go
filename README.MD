# Programming-with-Go
This repository contains my hands-on practice while learning Go (Golang)
# Go Practice (Udemy)

This repository contains my hands-on practice while learning Go (Golang), primarily following:

- **Course:** Learn How To Code: Google's Go (golang) Programming Language
- **Instructor:** Todd McLeod

The goal of this repo is to track my progress, practice core Go concepts, and build a strong public learning log on GitHub.

## Repository Structure

The repo is organized into multiple `hands-on-*` folders.
Each folder contains focused exercises for specific topics from the course and revision practice.

Examples:

- `hands-on-109/`
- `hands-on-110/`
- `hands-on-121(revision_and_practice)/`
- `hands-on-understanding-Functions- Part 1/`

## How to Run

From the repository root, run any exercise folder using:

```bash
go run ./hands-on-109
```

If a folder name has spaces, wrap it in quotes:

```bash
go run "./hands-on-understanding-Functions- Part 1"
```

## Why This Repo Exists

- Practice Go fundamentals consistently
- Maintain a visible learning journey
- Prepare for production-grade Go projects
- Build a portfolio that shows discipline and growth

## Learning Focus

This practice includes topics such as:

- Variables, types, constants
- Control flow and loops
- Functions and methods
- Structs and interfaces
- Working with standard library packages

## Development Status

This repository is currently in active development. I regularly add new exercises and revisions.

I am also building a fresh **Windows + WSL** development environment and automating tool setup using **Ansible**.

## CI Pipeline (GitHub Actions)

The repository uses GitHub Actions workflow:

- Workflow file: `.github/workflows/go.yml`
- Script executed by workflow: `source/workflow.sh`

The script checks each `hands-on-*` (and `hands-on-understanding-*`) folder and runs the appropriate command based on folder contents:

- `go test ./...` + `go build ./...` for module-based folders
- `go run main.go` for main-only folders

## Dev Environment (Windows)

- Setup guide: [docs/dev-environment-windows-wsl.md](docs/dev-environment-windows-wsl.md)
- Ansible bootstrap playbook: [ansible/dev-tools.yml](ansible/dev-tools.yml)
- `direnv` template: [.envrc.example](.envrc.example)

## Roadmap

- [ ] Add a short description to each `hands-on-*` folder
- [ ] Group folders by topic (basics, functions, structs, interfaces, etc.)
- [ ] Add example outputs/screenshots for selected exercises
- [ ] Add mini-projects after core fundamentals
- [ ] Add tests for selected practice modules
- [ ] Complete WSL setup and use WSL as primary Go environment
- [ ] Expand Ansible playbook with more dev tools (Docker, Node, Python, Terraform)
- [ ] Add repository-level `.envrc` from template and standardize local environment variables

## Author

**Neerav Bhatia**

If you are also learning Go, feel free to fork and practice along.