#!/bin/bash
# This script is executed after the creation of a new project.

echo 'deb [trusted=yes] https://repo.goreleaser.com/apt/ /' | sudo tee /etc/apt/sources.list.d/goreleaser.list
sudo apt update
sudo apt install -y goreleaser direnv

cp scripts/pre-commit.sh .git/hooks/pre-commit
cp scripts/pre-push.sh .git/hooks/pre-push
chmod 755 .git/hooks/pre-commit
chmod 755 .git/hooks/pre-push