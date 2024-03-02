package zet

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
	"time"
)

var ZetDir, _ = filepath.Abs("/home/micah/Repos/github.com/tr00datp00nar/zet/")

type Note struct {
	Title        string `yaml:"title"`
	Body         string `yaml:"body"`
	CreationDate string `yaml:"creation_date"`
}

type FileData struct {
	FileName     string
	LastModified string
}

const zetTmpl = `---
creation_date: {{ .CreationDate }}
---
# {{ .Title }}

{{ .Body }}
`

func newZet() {
	var title string

	fmt.Print("Enter the title of the note: ")
	reader := bufio.NewReader(os.Stdin)
	title, _ = reader.ReadString('\n')
	title = strings.TrimSpace(title)

	// Create a temporary file for the body
	tmpFile, err := os.CreateTemp("", "note-*.md")
	if err != nil {
		log.Println("Error creating temp file:", err)
		return
	}
	tmpFileName := tmpFile.Name()
	tmpFile.Close() // Close the file so vim can open it

	// Open the editor to edit the body
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi" // Use vim as a default editor if $EDITOR is not set
	}

	// Open Vim to edit the body
	cmd := exec.Command(editor, tmpFileName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Println("Error launching editor:", err)
		return
	}

	// Read the content from the temporary file
	bodyBytes, err := os.ReadFile(tmpFileName)
	if err != nil {
		log.Println("Error reading temp file:", err)
		return
	}
	body := string(bodyBytes)

	// Clean up the temporary file
	os.Remove(tmpFileName)

	tmpl := template.Must(template.New("zet").Parse(zetTmpl))

	zetNote := Note{
		Title:        title,
		Body:         body,
		CreationDate: getCreationDate(),
	}

	file, err := os.Create(filepath.Join(ZetDir, title+".md"))
	if err != nil {
		log.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	err = tmpl.Execute(file, zetNote)
	if err != nil {
		log.Println("Error executing template:", err)
		return
	}

	fmt.Println("Note saved as", title+".md")
	updateIndexList(ZetDir)
}

func getCreationDate() string {
	creationDate := time.Now().UTC().Format("200601021504")
	return creationDate
}

func updateIndexList(directory string) error {
	indexFilePath := filepath.Join(directory, "index.md")

	// Read existing content of the index file
	existingContent, err := os.ReadFile(indexFilePath)
	if err != nil {
		return err
	}

	// Parse existing content to extract table and file data list
	tableStart := strings.Index(string(existingContent), "| File | Last Modified |")
	if tableStart == -1 {
		_, err = fmt.Println("Table not found in index file")
		return err
	}
	tableEnd := strings.LastIndex(string(existingContent), "|")
	if tableEnd == -1 {
		_, err = fmt.Println("Invalid table format in index file")
		return err
	}
	tableContent := string(existingContent)[tableStart:tableEnd]

	var fileDataList []FileData
	rows := strings.Split(tableContent, "\n")[2:] // Skip header rows
	for _, row := range rows {
		if len(row) > 0 {
			cells := strings.Split(row, "|")
			fileName := strings.TrimSpace(cells[1])
			lastModified := strings.TrimSpace(cells[2])
			fileDataList = append(fileDataList, FileData{FileName: fileName, LastModified: lastModified})
		}
	}

	// Search for files in the given directory and sort them by LastModified
	var files []FileData
	err = filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() != "index.md" {
			files = append(files, FileData{FileName: info.Name(), LastModified: info.ModTime().Format("200601021422")})
		}
		return nil
	})
	if err != nil {
		return err
	}
	sort.Slice(files, func(i, j int) bool {
		return files[i].LastModified > files[j].LastModified
	})

	// Construct updated table content
	var updatedTable strings.Builder
	updatedTable.WriteString("| File | Last Modified |\n")
	updatedTable.WriteString("|----|-------------|\n")
	for _, fd := range files {
		row := fmt.Sprintf("|[[%s]]|%s|\n", fd.FileName, fd.LastModified)
		updatedTable.WriteString(row)
	}

	// Replace existing table content with updated table content
	newContent := strings.Replace(string(existingContent), tableContent, updatedTable.String(), 1)

	// Write the entire content back to the index file
	err = os.WriteFile(indexFilePath, []byte(newContent), 0644)
	if err != nil {
		return err
	}

	return nil
}
