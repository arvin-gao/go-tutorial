package packages

import (
	"os"
	"path"
	"path/filepath"
	"testing"
)

func TestPath(t *testing.T) {
	//* Path process.
	// aa.jpg
	ptr(path.Base("http://domain/file/aa.jpg"))
	// c:\file/abc/aa.jpg
	ptr(path.Clean("c:\\file//abc///aa.jpg"))
	// Current path, <nil>
	ptr(os.Getwd())
	// http:/domain/aa
	ptr(path.Dir("http://domain/aa/aaa.jpg"))
	// c:/a/b/c
	ptr(path.Dir("c:/a/b/c/d.txt"))
	// c:\a
	ptr(path.Dir("c:\\a/b.txt"))
	// .txt
	ptr(path.Ext("c:\\a/b.txt"))
	// false
	ptr(path.IsAbs("c:/wind/aa/bb/b.txt"))
	// c:/aa/bb/cc.txt
	ptr(path.Join("c:", "aa", "bb", "cc.txt"))
	// true <nil>
	ptr(path.Match("c:/windows/*/", "c:/windows/system/"))
	// c:/windows/system/ aaa.jpg
	ptr(path.Split("c:/windows/system/aaa.jpg"))

	//* FilePath process.
	ptr(filepath.IsAbs("c:\\wind\\aa\\bb\\b.txt")) // true
	// Check absolute representation of path
	// "D:\Projects\path <nil>"
	ptr(filepath.Abs("."))
	// Get filename with extension.
	// "baa.exe"
	ptr(filepath.Base("c:\\aa\\baa.exe"))
	ptr(filepath.Clean("c:\\\\aa/c\\baa.exe")) // c:\aa\c\baa.exe
	ptr(filepath.Clean("aa/c\\baa.exe"))       // aa\c\baa.exe
	// "aa\c"
	ptr(filepath.Dir("aa/c\\baa.exe"))
	ptr(filepath.EvalSymlinks("./path.exe")) // path.exe <nil>. That can be used for file exists.
	// Get extension name with dot.
	// ".exe"
	ptr(filepath.Ext("./path.exe"))
	ptr(filepath.FromSlash("c:\\windows\\aa//bb/cc//path.exe")) // c:\windows\aa\\bb\cc\\path.exe. '/' instead of '\\'
	ptr(filepath.ToSlash("c:\\windows\\aa/bb/cc/path.exe"))     // c:/windows/aa/bb/cc/path.exe. '\\' instead of '/'
	// Get volume name.
	// "c:"
	ptr(filepath.VolumeName("c:\\windows\\"))
	ptr(filepath.Glob("c:\\windows\\*.exe"))                       // Get all of files with the 'exe' execution name
	ptr(filepath.IsAbs("http://www.baidu.com/aa.jpg"))             // false
	ptr(filepath.Join("a", "\\bb\\", "cc", "/d", "e\\", "ff.txt")) // a\bb\cc\d\e\ff.txt
	ptr(filepath.Match("c:/windows/*/", "c:/windows/system/"))     // true <nil>
	ptr(filepath.Rel("c:/windows", "c:/windows/system/"))          // 取得第二参的路径中，相对于前面的路径的相对路径。  //system <nil>
	ptr(string(filepath.Separator))                                // windows下返回\\
	ptr(filepath.Split("c:/windows/system/abc.exe"))               //c:/windows/system/ abc.exe
	ptr(filepath.SplitList("c:/windows/system/abc.exe"))           //[c:/windows/system/abc.exe]
	filepath.Walk("../../syntax", WalkFunc)
	/*
	   File: ../../syntax IsDir: true size: 0
	   File: ..\..\syntax\painc IsDir: true size: 0
	   File: ..\..\syntax\painc\main.go IsDir: false size: 813
	   File: ..\..\syntax\painc\painc.exe IsDir: false size: 2498048
	   File: ..\..\syntax\path IsDir: true size: 0
	   File: ..\..\syntax\path\path.exe IsDir: false size: 2851328
	   File: ..\..\syntax\path\path.go IsDir: false size: 3419
	*/

}
func WalkFunc(path string, info os.FileInfo, err error) error {
	ptr("File:", path, "IsDir:", info.IsDir(), "size:", info.Size())
	return nil
}
