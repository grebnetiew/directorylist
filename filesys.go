package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type FileList []os.FileInfo

const slash = string(os.PathSeparator)

func main() {
	// The only means of input is using command line arguments.
	// Produce an error if no suitable command line is given.
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <directory to recurse>\n", os.Args[0])
	}

	// List the given directory
	rootDirName := strings.Join(os.Args[1:], " ")
	ListDirectory(rootDirName)
}

func ListDirectory(path string) {
	// Attempt to open the given directory
	rootDir, err := os.Stat(path)
	if err != nil || !rootDir.IsDir() {
		log.Fatalf("<%s> is not a directory.\n", path)
	}
	// Start recursive listing
	recurseDirectory(path+slash, rootDir, 0)
}

func recurseDirectory(path string, info os.FileInfo, depth int) {
	// Print the current directory name
	fmt.Println(strings.Repeat("\t", depth) + "[" + info.Name() + "]")

	// Attempt to list the files in this directory
	dirFile, err := os.Open(path)
	if err != nil {
		log.Printf("Couldn't open directory <%s>, %v\n", path, err)
		return
	}
	var list FileList
	list, err = dirFile.Readdir(0)
	if err != nil {
		log.Printf("Couldn't list directory <%s>, %v\n", path, err)
		return
	}
	// Sort the list
	sort.Sort(list)

	// Print or recurse each file resp. directory
	for _, file := range list {
		if file.IsDir() {
			recurseDirectory(path+slash+file.Name(), file, depth+1)
		} else {
			fmt.Println(strings.Repeat("\t", depth+1) + file.Name())
		}
	}
}

// The sort library requires that FileList implement sort.Interface.
func (fl FileList) Len() int {
	return len(fl)
}
func (fl FileList) Swap(i, j int) {
	fl[i], fl[j] = fl[j], fl[i]
}
func (fl FileList) Less(i, j int) bool {
	// If one is a dir but the other is not, list the dir first
	if isDir := fl[i].IsDir(); isDir != fl[j].IsDir() {
		return isDir
	}
	// Otherwise, just sort lexicographically
	return fl[i].Name() < fl[j].Name()
} 
