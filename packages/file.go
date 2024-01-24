package packages

import (
	"bufio"
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
}

func ReadDirectories(path string) {
	dirs, _ := os.ReadDir(path)
	for _, dir := range dirs {
		dir.IsDir()
		dir.Name()
		dir.Type()
	}
}

// TODO
func OtherDirectories(path string) {
	// Changes the current working directory to the named directory.
	os.Chdir(path)

	// ?
	visit := func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		ptr(" ", p, info.IsDir())
		return nil
	}
	filepath.Walk("subdir", visit)
}

func CreateFile() {
	// Create a temporary file, which named with random numbers.
	tempF, _ := os.CreateTemp("", "filename_*.txt")
	defer os.Remove(tempF.Name())
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
func WriteFile(path string, data []byte) {
	os.WriteFile(path, data, 0755)

	var f os.File
	data = []byte{115, 111, 109, 101, 10}
	// Write by bytes.
	f.Write(data)
	// Write by string.
	f.WriteString(string(data))

	f.Sync()

	writer := bufio.NewWriter(&f)
	writer.WriteString(string(data))
	writer.Flush()
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
