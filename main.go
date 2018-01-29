package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// take a project name as an arg
	n := flag.String("notepath", notePath(), "Main notebook directory set by env var NOTEPATH")
	s := flag.String("section", "new_section", "Name for new section of this notebook")
	flag.Parse()

	// create a directory with the name flag
	r := strings.Join([]string{*n, *s}, "/")
	fmt.Printf("Creating new section: %v\n", r)
	err := os.MkdirAll(r, 0755)
	if err != nil {
		fmt.Println(err)
	}

	// create a README.md file with h1 of notebook name
	fmt.Println("Creating README")
	p := strings.Join([]string{r, "README.md"}, "/")
	f, err := os.Create(p)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	h := strings.Join([]string{"#", r}, " ")
	f.WriteString(h)

	// create subdirectories for data, scripts, notes
	subDirs := []string{"data", "scripts", "notes"}
	for _, d := range subDirs {
		dirPath := strings.Join([]string{r, d}, "/")
		err := os.Mkdir(dirPath, 0755)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Creating %v section\n", d)
	}
}

func notePath() string {
	if os.Getenv("NOTEPATH") != "" {
		return os.Getenv("NOTEPATH")
	}
	return "notebook"
}
