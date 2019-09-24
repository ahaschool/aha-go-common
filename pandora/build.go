package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/urfave/cli"
)

type build struct {
	Name  string
}

func buildAction(c *cli.Context) error {
	base, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if &b.Name == nil {
		panic("need name option")
	}
	dir := buildDir(base, "cmd", 5)
	for _, v := range obj {
		fmt.Printf("%s build ...\n", v)
		dirName := fmt.Sprintf("../bin/%s_%s", b.Name, v)
		cmdName := fmt.Sprintf("%s.go", v)
		args := append([]string{"build", "-o", dirName, cmdName}, c.Args()...)
		cmd := exec.Command("go", args...)
		cmd.Dir = dir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			panic(err)
		}
	}
	fmt.Printf("pandora: %s\n", Version)
	fmt.Println("build success.")
	return nil
}

func buildDir(base string, cmd string, n int) string {
	dirs, err := ioutil.ReadDir(base)
	if err != nil {
		panic(err)
	}
	for _, d := range dirs {
		if d.IsDir() && d.Name() == cmd {
			return path.Join(base, cmd)
		}
	}
	if n <= 1 {
		return base
	}
	return buildDir(filepath.Dir(base), cmd, n-1)
}
