package packages

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

/*
Ref:
	1. https://gobyexample.com/reading-files
	2. https://gobyexample.com/directories
	3. https://www.devdungeon.com/content/working-files-go
	4. https://gobyexample.com/file-paths
	5. https://gobyexample.com/embed-directive
*/

// Create directories.
func CreateDirectory(folderName, path string) {
	// Create a single directory.
	os.Mkdir(folderName, 0755)

	// Create a hierarchy of directories(nested directories).
	os.MkdirAll(path, os.ModePerm)

	// Create a temporary directory, which named with random numbers.
	os.MkdirTemp("./temp", "directory_*")
	defer os.RemoveAll("./temp")

	// Create an empty directory
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, 0755); err != nil {
			fmt.Println("Error making directory", err)
		}
	}
}

func ReadDirectories(path string) {
	dirs, _ := os.ReadDir(path)
	for _, dir := range dirs {
		dir.IsDir()
		dir.Name()
		dir.Type()
	}

	// Read directory recursively
	filepath.Walk(path, func(fn string, fi os.FileInfo, err error) error {
		if err != nil {
			ptr(err)
			return err
		}

		if fi.IsDir() {
			ptr("Directory: ", fn)
			if fi.Name() == "skipme" {
				return filepath.SkipDir
			}
		} else {
			ptr("File: ", fn)
		}
		return nil
	})
}

// Delete directory
func DeleteDirect(src string) {
	os.RemoveAll(src)
}

func CreateFile() {
	// Create a temporary file, which named with random numbers.
	tempFile, _ := os.CreateTemp("", "filename_*.txt")
	defer os.Remove(tempFile.Name())

	// Create file in the directory
	emptyFile, _ := os.Create("emptyFile.txt")
	defer emptyFile.Close()
}

// Rename the file.
func RenameFile(oldPath, newPath string) {
	os.Rename(oldPath, newPath)
}

func FileInformation(filePath string) {
	// Get file information
	fileStat, _ := os.Stat(filePath)
	ptr("File Name:", fileStat.Name())        // Name of the file
	ptr("Size:", fileStat.Size())             // Length in bytes
	ptr("Last Modified:", fileStat.ModTime()) // Last modification time and date

}

// Move directory to new path.
func MoveFile(oldPath, newPath string) error {
	return os.Rename(oldPath, newPath)
}

func RemoveFile(filename, folder string) {
	// Remove the file.
	os.Remove(filename)

	// Remove the folder.
	os.RemoveAll(folder)
}

func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	return io.Copy(destination, source)
}

func ControlFile(path string) {
	// Open file.
	os.Open(path)
	// Read file.
	os.ReadFile(path)
	// Read bytes.
	var f os.File
	var fBytes []byte
	f.Read(fBytes)

	//?? Set offset?
	f.Seek(6, 0)
	io.ReadAtLeast(&f, fBytes, 2)
	f.Seek(0, 0)

	reader := bufio.NewReader(&f)
	reader.Peek(5)
	f.Close()
}

// refer to: https://gobyexample.com/writing-files
func WriteFile(filePath, content string) error {
	return os.WriteFile(filePath, []byte(content), 0755)
}

func WriteFileByBufio(filePath, content string) error {
	var f os.File
	// Write by bytes.
	f.Write([]byte(content))
	// Write by string.
	f.WriteString(string(content))

	f.Sync()

	writer := bufio.NewWriter(&f)
	if _, err := writer.WriteString(string(content)); err != nil {
		return err
	}
	return writer.Flush()
}

func WriteAppendFile(filePath string) {
	// Append to an existing file
	af, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error appending to file:", err)
	}
	defer af.Close()

	if _, err = af.WriteString("\nAppending this text"); err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func ReduceFileSize(filePath string) error {
	return os.Truncate(filePath, 100)
}

func FilePath(dir1, dir2, filename string) {
	newPath := filepath.Join(dir1, dir2, filename)

	ptr(filepath.Join("dir1//", "filename"))
	ptr(filepath.Join("dir1/../dir1", "filename"))

	ptr("Dir(p):", filepath.Dir(newPath))
	ptr("Base(p):", filepath.Base(newPath))

	ptr(filepath.IsAbs("dir/file"))
	ptr(filepath.IsAbs("/dir/file"))

	// Get the extension of the filename.
	filepath.Ext(filename)

	// Get the filename without extension.
	strings.TrimSuffix(filename, filepath.Ext(filename))

	// ?
	filepath.Rel("a/b", "a/b/t/file")
	filepath.Rel("a/b", "a/c/t/file")
}
