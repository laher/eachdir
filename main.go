package main

import (
	"context"
	"log"
	"os"
	"os/exec"
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
		log.Printf("cd %s\n", dir)
		if err := os.Chdir(dir); err != nil {
			panic(err)
		}
		// run command
		c := exec.CommandContext(ctx, args[0], args[1:]...)
		c.Stdout = os.Stdout
		if err := c.Run(); err != nil {
			panic(err)
		}
		log.Printf("cd %s\n", wd)
		if err := os.Chdir(wd); err != nil {
			panic(err)
		}
	}
}
