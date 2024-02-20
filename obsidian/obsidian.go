package obsidian

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func selectVault() (vaultName string) {
	var selectedVault string
	var dirsToPick []string

	userHomeDir, _ := os.UserHomeDir()
	searchDirs := []string{filepath.Join(userHomeDir, "Repos")}

	for _, searchDir := range searchDirs {
		err := filepath.Walk(searchDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() && strings.HasSuffix(path, ".obsidian") {
				dirsToPick = append(dirsToPick, filepath.Dir(path))
			}
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	fzfInput := strings.Join(dirsToPick, "\n")
	cmd := exec.Command("fzf", "--multi", "--delimiter=\\n", "--tac")
	cmd.Stdin = strings.NewReader(fzfInput)
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	selectedVault = strings.TrimSpace(string(out))
	return selectedVault
}

func selectFile(vaultPath string) (file string) {
	var choices []string
	var selectedFile string

	err := filepath.Walk(vaultPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".md") {
			choices = append(choices, filepath.Base(path))
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error reading files:", err)
	}

	fzfInput := strings.Join(choices, "\n")
	fzfCmd := exec.Command("fzf", "--multi", "--delimiter=\\n", "--tac")
	fzfCmd.Stdin = strings.NewReader(fzfInput)
	// fzfCmd.Stdout = os.Stdout
	fzfCmd.Stderr = os.Stderr
	out, err := fzfCmd.Output()
	if err != nil {
		fmt.Println("Error running fzf:", err)
	}
	selectedFile = strings.TrimSpace(string(out))
	return selectedFile
}

func openVault(vaultName string) error {
	if vaultName == "" {
		return nil
	}
	vaultName = filepath.Base(vaultName)
	uri := fmt.Sprintf("obsidian://open?vault=%s", vaultName)
	return exec.Command("obsidian", uri).Run()
}

func openFile(vaultName string, fileName string) {
	selectedVault := vaultName
	selectedFile := fileName
	uri := fmt.Sprintf("obsidian://open?vault=%s&file=%s", selectedVault, selectedFile)
	openCmd := exec.Command("obsidian", uri)
	err := openCmd.Run()
	if err != nil {
		fmt.Println("Error opening note:", err)
	}
}

// files, err := os.ReadDir(selectedVault)
// if err != nil {
// 	fmt.Println("Error reading directory:", err)
// }

// fileNames := make([]string, len(files))
// for i, file := range files {
// 	fileNames[i] = file.Name()
// }

// fzfInput := strings.Join(fileNames, "\n")
// fzfCmd := exec.Command("fzf", "--multi", "--delimiter=\\n", "--tac")
// fzfCmd.Stdin = strings.NewReader(fzfInput)
// fzfCmd.Stderr = os.Stderr
// fzfOutput, err := fzfCmd.Output()
// if err != nil {
// 	fmt.Println("Error running fzf:", err)
// }

// Choose a file

// selectedVault = filepath.Base(selectedVault)
// selectedNotes := strings.Split(strings.TrimSpace(string(fzfOutput)), "\n")
// for _, selectedNote := range selectedNotes {
// 	if selectedNote == "" {
// 		return
// 	}
