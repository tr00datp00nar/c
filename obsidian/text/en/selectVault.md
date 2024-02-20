Tool for selecting and opening Obsidian vaults.

Looks through the directories `$HOME/projects/obsidian` and `$HOME/Repos` to find any directories that contain a `.obsidian` directory. Then builds a fzf picker with the results. Once the user selects one of the directories, a URL is built and used to open Obsidian in the specified vault.
