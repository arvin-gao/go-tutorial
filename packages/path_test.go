package packages

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
)

func TestPath(t *testing.T) {
	// Current path, <nil>
	// A rooted path name corresponding to the current directory.
	os.Getwd()

	// Get filename with extension.
	// - http://domain/file/aa.jpg -> aa.jpg
	// - c:\\aa\\baa.exe           -> baa.exe
	_ = path.Base("http://domain/file/aa.jpg")
	_ = filepath.Base("c:\\aa\\baa.exe")

	// Get extension of the filename of path.
	// c:\\a/b.txt -> .txt
	// ./path.exe  -> .exe
	_ = path.Ext("c:\\a/b.txt")
	_ = filepath.Ext("./path.exe")

	// Get filename.
	{
		// c:\\aa\\baa.exe -> baa
		filename := filepath.Base("c:\\aa\\baa.exe")
		_ = strings.TrimSuffix(filename, filepath.Ext(filename))
	}

	// Get the shortest path name equivalent to path by purely lexical processing.
	// - c:\\file//abc///aa.jpg -> c:\file/abc/aa.jpg
	// - c:\\\\aa/c\\baa.exe"   -> c:\aa\c\baa.exe
	// - aa/c\\baa.exe"         -> aa\c\baa.exe
	path.Clean("c:\\file//abc///aa.jpg")
	filepath.Clean("c:\\\\aa/c\\baa.exe")

	// Get all but the last element of path, typically the path's directory.
	// 移除路徑的最後元素。
	// - http://domain/aa/aaa.jpg -> http:/domain/aa
	// - c:/a/b/c/d.txt           -> c:/a/b/c
	// - c:\\a/b.txt              -> c:\a
	// - aa/c\\baa.exe            -> aa\c
	path.Dir("c:\\a/b.txt")
	filepath.Dir("aa/c\\baa.exe")

	// Joins
	// Joins any number of path elems into a single path,
	// separating them with slashes
	// - "c:", "aa", "bb", "cc.txt" -> c:/aa/bb/cc.txt
	_ = path.Join("c:", "aa", "bb", "cc.txt")
	// Joins any number of path elems into a single path,
	// separating them with an OS specific Separator.
	// - "a", "\\bb\\", "cc", "/d", "e\\", "ff.txt" -> a\bb\cc\d\e\ff.txt
	_ = filepath.Join("a", "\\bb\\", "cc", "/d", "e\\", "ff.txt")

	// Reports whether name matches the shell pattern.
	// true <nil>
	_, _ = path.Match("c:/windows/*/", "c:/windows/system/")
	// Splits path immediately following the final slash
	// :/windows/system/, aaa.jpg
	_, _ = path.Split("c:/windows/system/aaa.jpg")

	// Check the path is absolute.
	// - c:\\wind\\aa\\bb\\b.txt   -> true
	// - http://www.xxx.com/aa.jpg -> false
	_ = path.IsAbs("c:/wind/aa/bb/b.txt")
	_ = filepath.IsAbs("c:\\wind\\aa\\bb\\b.txt")

	// Check absolute representation of path
	// "D:\Projects\path <nil>"
	_, _ = filepath.Abs(".")

	// Get the path name after the evaluation of any symbolic
	// #That can be used for file exists.
	// - path.exe <nil>.
	filepath.EvalSymlinks("./path.exe")

	_ = filepath.FromSlash("c:\\windows\\aa//bb/cc//path.exe") // c:\windows\aa\\bb\cc\\path.exe. '/' instead of '\\'
	_ = filepath.ToSlash("c:\\windows\\aa/bb/cc/path.exe")     // c:/windows/aa/bb/cc/path.exe. '\\' instead of '/'
	// Get volume name.
	// "c:"
	_ = filepath.VolumeName("c:\\windows\\")
	_, _ = filepath.Glob("c:\\windows\\*.exe")                   // Get all of files with the 'exe' execution name
	_, _ = filepath.Match("c:/windows/*/", "c:/windows/system/") // true <nil>
	_, _ = filepath.Rel("c:/windows", "c:/windows/system/")      // 取得第二参的路径中，相对于前面的路径的相对路径。  //system <nil>
	_ = string(filepath.Separator)                               // windows下返回\\
	_, _ = filepath.Split("c:/windows/system/abc.exe")           //c:/windows/system/ abc.exe
	_ = filepath.SplitList("c:/windows/system/abc.exe")          //[c:/windows/system/abc.exe]

	/*
	   File: ../../syntax IsDir: true size: 0
	   File: ..\..\syntax\painc IsDir: true size: 0
	   File: ..\..\syntax\painc\main.go IsDir: false size: 813
	   File: ..\..\syntax\painc\painc.exe IsDir: false size: 2498048
	   File: ..\..\syntax\path IsDir: true size: 0
	   File: ..\..\syntax\path\path.exe IsDir: false size: 2851328
	   File: ..\..\syntax\path\path.go IsDir: false size: 3419
	*/
	filepath.Walk("../../syntax", WalkFunc)
}

func WalkFunc(path string, info os.FileInfo, err error) error {
	ptr("File:", path, "IsDir:", info.IsDir(), "size:", info.Size())
	return nil
}
