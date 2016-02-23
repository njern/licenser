package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ryanuber/go-license"
)

func main() {
	searchDir := flag.String("d", "", "directory to scan")
	flag.Parse()

	if *searchDir == "" {
		fmt.Println("please specify a directory with -d")
		os.Exit(0)
	}

	var subFolders []string
	err := filepath.Walk(*searchDir, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			subFolders = append(subFolders, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	for _, folder := range subFolders {
		l, err := license.NewFromDir(folder)
		if err != nil {
			if err.Error() != license.ErrNoLicenseFile {
				fmt.Println(folder, err.Error())
				os.Exit(-1)
			}
		}

		if l != nil {
			fmt.Printf("%v => %v\n", folder, l.Type)
		}
	}
}
