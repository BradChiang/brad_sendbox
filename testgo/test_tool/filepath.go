package test_tool

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Filepath() {
	pln := fmt.Println
	pln("file path test start")
	p := filepath.Join("dir1", "dir2", "filename")
	pln("path : ", p)
	pln(filepath.Join("dir1//", "filename"))
	pln(filepath.Join("dir1/../dir1", "filename"))
	pln("Dir(p) : ", filepath.Dir(p))
	pln("Base(p) : ", filepath.Base(p))

	pln(filepath.IsAbs("dir/file"))
	pln(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	ext := filepath.Ext(filename)
	pln(ext)

	pln(strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	check(err)
	pln(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	check(err)
	pln(rel)

	pln("file path test end")
}
func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println("-", p, info.IsDir())
	return nil
}

func Director() {
	p := fmt.Println
	p("director test start")

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
	p("List subdir/parent :")

	for _, entry := range c {
		p("-", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("subdir/parent/child")
	check(err)

	c, err = os.ReadDir(".")
	check(err)
	p("List subdir/parent/child :")
	for _, entry := range c {
		p("-", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	check(err)
	fmt.Println("Visit subdir : ")
	_ = filepath.Walk("subdir", visit)
	p("director test end")

}

func TempfileDirector() {
	fmt.Println("temporary file & director test start")
	f, err := os.CreateTemp("", "sample")
	check(err)

	fmt.Println("Temp file name : ", f.Name())

	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	defer os.Remove(dname)

	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
	fmt.Println("temporary file & director test end")
}
