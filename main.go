package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	touchAll(currentDir())
}

func touchAll(baseDir string) {
	files := listFiles(baseDir)
	for _, file := range files {
		absPath := filepath.Join(baseDir, file.Name())
		log.Printf("abs path: %s", absPath)
		if file.IsDir() {
			touchAll(absPath)
		}
		touch(absPath)
	}
}

func currentDir() (dir string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Printf("get current dir error: %v\n", err)
		return
	}
	return
}

func listFiles(dir string) (files []os.FileInfo) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("list files error: %v\n", err)
		return
	}
	return
}

func touch(absPath string) (err error) {
	file, err := os.Stat(absPath)
	if err != nil {
		log.Printf("stat file error: %v\n", err)
		return
	}

	err = os.Chtimes(absPath, file.ModTime(), time.Now())
	if err != nil {
		log.Printf("change time error: %v\n", err)
		return
	}
	return
}
