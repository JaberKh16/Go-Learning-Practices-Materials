// Mini CLI File Manager
package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// =========================
// HELP
// =========================
func help() {
	fmt.Println(`
Mini File Manager Commands:

ls                -> list files
pwd               -> print working directory
cd <dir>          -> change directory
mkdir <name>      -> create folder
touch <file>      -> create file
cat <file>        -> read file
rm <file/dir>     -> delete file or folder
mv <src> <dest>   -> move file
cp <src> <dest>   -> copy file
exit              -> quit
`)
}

// =========================
// LS
// =========================
func ls() {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, f := range files {
		if f.IsDir() {
			fmt.Println("[DIR] ", f.Name())
		} else {
			fmt.Println("      ", f.Name())
		}
	}
}

// =========================
// PWD
// =========================
func pwd() {
	dir, _ := os.Getwd()
	fmt.Println(dir)
}

// =========================
// CD
// =========================
func cd(path string) {
	err := os.Chdir(path)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// =========================
// MKDIR
// =========================
func mkdir(name string) {
	err := os.Mkdir(name, 0755)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// =========================
// TOUCH
// =========================
func touch(name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	file.Close()
}

// =========================
// CAT
// =========================
func cat(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(data))
}

// =========================
// RM (file or folder)
// =========================
func rm(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// =========================
// MV
// =========================
func mv(src, dest string) {
	err := os.Rename(src, dest)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// =========================
// CP
// =========================
func cp(src, dest string) {
	input, err := os.Open(src)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer input.Close()

	output, err := os.Create(dest)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer output.Close()

	_, err = io.Copy(output, input)
	if err != nil {
		fmt.Println("Error copying:", err)
	}
}

// =========================
// COMMAND PARSER
// =========================
func handleCommand(input string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return
	}

	cmd := parts[0]

	switch cmd {

	case "ls":
		ls()

	case "pwd":
		pwd()

	case "cd":
		if len(parts) < 2 {
			fmt.Println("Usage: cd <dir>")
			return
		}
		cd(parts[1])

	case "mkdir":
		if len(parts) < 2 {
			fmt.Println("Usage: mkdir <name>")
			return
		}
		mkdir(parts[1])

	case "touch":
		if len(parts) < 2 {
			fmt.Println("Usage: touch <file>")
			return
		}
		touch(parts[1])

	case "cat":
		if len(parts) < 2 {
			fmt.Println("Usage: cat <file>")
			return
		}
		cat(parts[1])

	case "rm":
		if len(parts) < 2 {
			fmt.Println("Usage: rm <file/dir>")
			return
		}
		rm(parts[1])

	case "mv":
		if len(parts) < 3 {
			fmt.Println("Usage: mv <src> <dest>")
			return
		}
		mv(parts[1], parts[2])

	case "cp":
		if len(parts) < 3 {
			fmt.Println("Usage: cp <src> <dest>")
			return
		}
		cp(parts[1], parts[2])

	case "help":
		help()

	case "exit":
		os.Exit(0)

	default:
		fmt.Println("Unknown command:", cmd)
	}
}

// =========================
// MAIN LOOP
// =========================
func main() {
	fmt.Println("Mini CLI File Manager (type 'help')")

	for {
		dir, _ := os.Getwd()
		fmt.Print(filepath.Base(dir), " $ ")

		var input string
		fmt.Scanln(&input)

		handleCommand(input)
	}
}