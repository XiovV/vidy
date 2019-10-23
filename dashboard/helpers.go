package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadDir(u string) []string {
	dirs := []string{}

	withSpaceChar := strings.Replace(u, "%20", " ", -1)

	files, err := ioutil.ReadDir("./" + withSpaceChar)
	if err != nil {
		fmt.Println(err)
	}

	for _, f := range files {
		dirs = append(dirs, f.Name())
	}

	return dirs
}

// This function replaces "video" with "library" and strips off a .mp4 string
func ConvertURLToReadablePath(u string) string {
	filepathSplit := strings.Split(strings.Replace(u, "video", "library", -1), "/")
	filepath := strings.Replace(strings.Join(filepathSplit[:len(filepathSplit)-1], "/"), "%20", " ", -1)

	return filepath
}
