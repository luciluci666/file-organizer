package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func moveFile(source string, destination string) error {
	// Check if the source file exists
	_, err := os.Stat(source)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("source file does not exist: %s", source)
		}
		return err
	}

	// Check if the destination directory exists, create if not
	destDir := filepath.Dir(destination)
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		// Create the destination directory and any necessary parent directories
		err := os.MkdirAll(destDir, 0755)
		if err != nil {
			return fmt.Errorf("error creating destination directory: %v", err)
		}
	}

	// Rename (move) the file
	err = os.Rename(source, destination)
	if err != nil {
		return fmt.Errorf("error moving file: %v", err)
	}

	return nil
}

func getDestination(ext string) string {
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".tiff", ".webp":
		return "images"
	case ".ts", ".py", ".go", ".js", ".html", ".css", ".cpp", ".c", ".java", ".php", ".rb", ".swift", ".kt", ".rs", ".pl", ".sh", ".bat", ".ps1", ".vbs", ".lua", ".r", ".json", ".xml", ".yaml", ".yml":
		return "code"
	case ".mp3", ".wav":
		return "audio"
	case ".mp4", ".avi", ".mkv", ".mov", ".wmv", ".flv", ".webm", ".mpg", ".mpeg":
		return "video"
	case ".pdf", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".txt", ".rtf", ".csv", ".md":
		return "documents"
	case ".exe", ".msi", ".deb", "dmg", "apk", ".jar", ".rpm", ".appimage":
		return "executable"
	case ".zip", ".rar", ".7z", ".tar", ".gz", ".xz", ".bz2":
		return "compressed"
	default:
		return "unrecognized"
	}
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	// If err is nil, the path exists
	return !os.IsNotExist(err)
}

func main() {
	var dirPath string

	// Get the folder path from the user
	fmt.Print("Enter the folder path: ")
	fmt.Scanln(&dirPath)

	if !pathExists(dirPath) {
		fmt.Println("The folder path does not exist")
		return
	}

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	var successCount int = 0

	// Loop through and print file names
	for _, entry := range entries {
		if !entry.IsDir() { // Check if it's a file
			path := filepath.Join(dirPath, entry.Name())
			ext := filepath.Ext(path)
			destination := filepath.Join(dirPath, getDestination(ext), entry.Name())
			err := moveFile(path, destination)
			if err != nil {
				fmt.Println("Error moving file:", err)
			}
			successCount++
		}
	}

	if successCount != 0 {
		fmt.Printf("%d files moved successfully\n", successCount)
	} else {
		fmt.Println("No files found to move")
	}
}
