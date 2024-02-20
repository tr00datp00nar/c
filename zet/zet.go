package zet

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

var ZetDir, _ = filepath.Abs("/home/micah/Repos/github.com/tr00datp00nar/zet/")

type Note struct {
	Title string `yaml:"title"`
	Body  string `yaml:"body"`
}

const zetTmpl = `---
creation_date: <% tp.file.creation_date() %>
---
# <% tp.file.title %>

{{ .Body }}
`

func NewZet() {
	var title string

	fmt.Print("Enter the title of the note: ")
	reader := bufio.NewReader(os.Stdin)
	title, _ = reader.ReadString('\n')
	title = strings.TrimSpace(title)

	// Create a temporary file for the body
	tmpFile, err := os.CreateTemp("", "note-*.md")
	if err != nil {
		fmt.Println("Error creating temp file:", err)
		return
	}
	tmpFileName := tmpFile.Name()
	tmpFile.Close() // Close the file so vim can open it

	// Open Vim to edit the body
	cmd := exec.Command("nvim", tmpFileName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error launching Neovim:", err)
		return
	}

	// Read the content from the temporary file
	bodyBytes, err := os.ReadFile(tmpFileName)
	if err != nil {
		fmt.Println("Error reading temp file:", err)
		return
	}
	body := string(bodyBytes)

	// Clean up the temporary file
	os.Remove(tmpFileName)

	tmpl := template.Must(template.New("zet").Parse(zetTmpl))

	zetNote := Note{
		Title: title,
		Body:  body,
	}

	file, err := os.Create(filepath.Join(ZetDir, title+".md"))
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	err = tmpl.Execute(file, zetNote)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	fmt.Println("Note saved as", title+".md")
}

func ZetObsidian() {
	files, err := os.ReadDir(ZetDir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	fileNames := make([]string, len(files))
	for i, file := range files {
		fileNames[i] = file.Name()
	}

	fzfInput := strings.Join(fileNames, "\n")
	fzfCmd := exec.Command("fzf", "--multi", "--delimiter=\\n", "--tac")
	fzfCmd.Stdin = strings.NewReader(fzfInput)
	fzfCmd.Stderr = os.Stderr
	fzfOutput, err := fzfCmd.Output()
	if err != nil {
		fmt.Println("Error running fzf:", err)
		return
	}

	selectedNotes := strings.Split(strings.TrimSpace(string(fzfOutput)), "\n")
	for _, selectedNote := range selectedNotes {
		openCmd := exec.Command("obsidian", "obsidian://open?vault=zet&file="+selectedNote)
		err := openCmd.Run()
		if err != nil {
			fmt.Println("Error opening note:", err)
		}
		applyTemplater := exec.Command("obsidian", "obsidian://advanced-uri?zet&filepath="+selectedNote+"&commandid=replace-in-file-templater")
		err = applyTemplater.Run()
		if err != nil {
			fmt.Println("Error applying template:", err)
		}
	}
}
