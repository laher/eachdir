package main

import (
	"context"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	eachdir := os.Getenv("EACHDIR")

	if eachdir == "" {
		log.Fatal("no EACHDIR specified")
	}
	args := os.Args[1:]
	if len(args) < 1 {
		args = []string{"ls"}
	}
	dirs := strings.Split(eachdir, " ")
	// fmt.Println("running ", args, " on ", dirs)
	wd, _ := os.Getwd()
	ctx := context.Background()
	for _, dir := range dirs {
		matches, err := filepath.Glob(dir)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		for _, match := range matches {
			s, err := os.Stat(match)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			if s.IsDir() {
				log.Printf("cd %s\n", match)
				if err := os.Chdir(match); err != nil {
					log.Println(err)
					os.Exit(1)
				}
				// log.Printf("run %s\n", strings.Join(args, ","))
				// run command
				c := exec.CommandContext(ctx, args[0], args[1:]...)
				c.Stdout = os.Stdout
				if err := c.Run(); err != nil {
					log.Println(err)
					os.Exit(1)
				}
				// log.Printf("cd %s\n", wd)
				if err := os.Chdir(wd); err != nil {
					log.Println(err)
					os.Exit(1)
				}
			}
		}
	}
}
