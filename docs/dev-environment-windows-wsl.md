# Windows Dev Environment Setup (WSL + Ansible + direnv)

This guide is for a **Windows-first** setup where development runs in WSL.

## 1) Enable WSL on Windows

Run in PowerShell (Admin):

```powershell
wsl --install
```

Then reboot and verify:

```powershell
wsl --status
```

Recommended distro: Ubuntu (LTS).

## 2) Inside WSL: base packages

```bash
sudo apt update
sudo apt install -y curl git unzip build-essential software-properties-common
```

## 3) Install Ansible in WSL

```bash
sudo apt update
sudo apt install -y ansible
ansible --version
```

## 4) Bootstrap tools with Ansible

From this repository root in WSL:

```bash
ansible-playbook ansible/dev-tools.yml --ask-become-pass
```

This installs common CLI tools and sets up `direnv` shell hook.

## 5) Enable direnv

After playbook completion, reload shell:

```bash
exec "$SHELL"
```

Allow this repo env once:

```bash
direnv allow
```

## 6) Verify core tools

```bash
go version
git --version
gh --version
direnv version
```

## Notes

- Keep code inside WSL filesystem for best performance (`~/...`), or use VS Code Remote - WSL.
- If `gh auth login` is needed, run it once after setup.
- You can extend the playbook for Docker, Node, Python, Terraform, etc.
