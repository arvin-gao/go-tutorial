package packages

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// refer to: https://gobyexample.com/reading-files

// TODO: https://www.devdungeon.com/content/working-files-go
func TestFile(t *testing.T) {
	check := func(e error) {
		if e != nil {
			panic(e)
		}
	}

	dat, err := os.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	f, err := os.Open("/tmp/dat")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()
}

// refer to: https://gobyexample.com/writing-files
func TestWriteFile(t *testing.T) {
	check := func(e error) {
		if e != nil {
			panic(e)
		}
	}

	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	f, err := os.Create("/tmp/dat2")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}

// refer to: https://gobyexample.com/file-paths

func TestFilePath(t *testing.T) {
	p := filepath.Join("dir1", "dir2", "filename")
	pln("p:", p)

	pln(filepath.Join("dir1//", "filename"))
	pln(filepath.Join("dir1/../dir1", "filename"))

	pln("Dir(p):", filepath.Dir(p))
	pln("Base(p):", filepath.Base(p))

	pln(filepath.IsAbs("dir/file"))
	pln(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	ext := filepath.Ext(filename)
	pln(ext)

	pln(strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	pln(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	pln(rel)
}

// refer to: https://gobyexample.com/directories
func TestDirectories(t *testing.T) {
	check := func(e error) {
		if e != nil {
			panic(e)
		}
	}

	visit := func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		pln(" ", p, info.IsDir())
		return nil
	}

	err := os.Mkdir("subdir", 0755)
	check(err)

	defer os.RemoveAll("subdir")

	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	c, err := os.ReadDir("subdir/parent")
	check(err)

	pln("Listing subdir/parent")
	for _, entry := range c {
		pln(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("subdir/parent/child")
	check(err)

	c, err = os.ReadDir(".")
	check(err)

	pln("Listing subdir/parent/child")
	for _, entry := range c {
		pln(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	check(err)

	pln("Visiting subdir")

	err = filepath.Walk("subdir", visit)
	pln(err)
}

func TestTempFileAndDir(t *testing.T) {
	check := func(e error) {
		if e != nil {
			panic(e)
		}
	}

	f, err := os.CreateTemp("", "sample")
	check(err)

	pln("Temp file name:", f.Name())

	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	pln("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}

// TODO: .
// refer to: https://gobyexample.com/embed-directive
// func TestEmbedDirective(t *testing.T) {

// 	//go:embed folder/single_file.txt
// 	var fileString string

// 	//go:embed folder/single_file.txt
// 	var fileByte []byte

// 	//go:embed folder/single_file.txt
// 	//go:embed folder/*.hash
// 	var folder embed.FS

// 	print(fileString)
// 	print(string(fileByte))

// 	content1, _ := folder.ReadFile("folder/file1.hash")
// 	print(string(content1))

// 	content2, _ := folder.ReadFile("folder/file2.hash")
// 	print(string(content2))
// }
