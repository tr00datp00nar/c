Create a new entry in the Obsidian Zettelkasten.

{{ aka }} will prompt the user for a note title and then creates a temporary file in the `/tmp` directory and opens the file in neovim. Once neovim is closed, {{ aka }} copies the contents of the temporary file to the vault, and cleans up `/tmp`.

The template below will be applied to each note.

    ---
    creation_date: <% tp.file.creation_date() %>
    ---
    # <% tp.file.title %>
    Body of Note
