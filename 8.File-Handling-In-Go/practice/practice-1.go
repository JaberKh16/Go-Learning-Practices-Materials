package main

import (
	"fmt"
	"os"
	"time"
)

// =========================
// 1. Create a file
// =========================
func createFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Println("File created:", filename)
}


// =========================
// 2. Write to file
// =========================
func writeFile(filename string, content string) {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("Written to file:", filename)
}


// =========================
// 3. Read file
// =========================
func readFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// direct print content
	contentDirectPrint(string(data))	
}



func contentDirectPrint(data string){
	fmt.Println("File content:")
	fmt.Println(data)
}


// =========================
// 4. Delete file
// =========================
func deleteFile(filename string) {
	err := os.Remove(filename)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
	fmt.Println("File deleted:", filename)
}


// =========================
// 5. Create directory
// =========================
func createDir(dir string) {
	err := os.Mkdir(dir, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	fmt.Println("Directory created:", dir)
}


// =========================
// 6. List directory files
// =========================
func listDir(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	fmt.Println("Directory contents:")
	for _, file := range files {
		fmt.Println("-", file.Name())
	}
}


// =========================
// 7. Get file info
// =========================
func fileInfo(filename string) {
	info, err := os.Stat(filename)
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}

	fmt.Println("File Info:")
	fmt.Println("Name:", info.Name())
	fmt.Println("Size:", info.Size(), "bytes")
	fmt.Println("Permissions:", info.Mode())
	fmt.Println("Modified:", info.ModTime())
	fmt.Println("Is Directory:", info.IsDir())
}


// =========================
// 8. Environment variables
// =========================
func envExample() {
	os.Setenv("APP_NAME", "GoOSApp")

	fmt.Println("APP_NAME:", os.Getenv("APP_NAME"))

	// List all env (partial example)
	fmt.Println("Some environment variables:")
	for i, env := range os.Environ() {
		if i > 5 {
			break
		}
		fmt.Println(env)
	}
}


// =========================
// 9. Working directory
// =========================
func workingDir() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}

	fmt.Println("Current working directory:", dir)

	err = os.Chdir("..") // move one level up
	if err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}

	newDir, _ := os.Getwd()
	fmt.Println("After change directory:", newDir)
}


// =========================
// MAIN FUNCTION
// =========================
func main() {

	filename := "test.txt"
	dir := "files"

	// File operations
	createFile(filename)
	writeFile(filename, "Hello Go OS Package!\nWorking with files is easy.")
	readFile(filename)
	fileInfo(filename)

	// Directory operations
	createDir(dir)
	listDir(".")

	// Environment variables
	envExample()

	// Working directory
	workingDir()

	// Cleanup
	time.Sleep(1 * time.Second)
	deleteFile(filename)
}