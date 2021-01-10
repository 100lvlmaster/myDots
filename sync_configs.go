package main

import (
	"fmt"
	"log"
	"math/rand"
	"os/exec"
)

//CopyDir will copy files to this dir
func CopyDir(src, dst string) error {
	cmd := exec.Command("cp", "-R", src, dst)
	log.Printf("Running cp -a")
	fmt.Println(src)
	return cmd.Run()
}

func pushToGit() {
	val := exec.Command("git", "init")
	val.Run()
	nol := exec.Command("git", "add", ".")
	nol.Run()
	commit := exec.Command("git", "commit", "-m", fmt.Sprintf("%d", rand.Intn(100000)))
	commit.Run()
	all := exec.Command("git", "push", "origin", "master")
	all.Run()
}
func main() {
	/// Home Directory
	var homeDir = "/home/navin"
	var configDir = homeDir + "/.config"
	/// Map of config files
	configs := map[string]string{
		"alacritty": configDir + "/alacritty",
		"dunst":     configDir + "/dunst",
		"polybar":   configDir + "/polybar",
		"i3":        configDir + "/i3",
	}
	/// Map of misc files
	miscFiles := map[string]string{
		"zshrc":  homeDir + "/.zshrc",
		"bashrc": homeDir + "/.bashrc",
		"walls":  homeDir + "/wallpapers",
	}
	for k := range miscFiles {
		CopyDir(miscFiles[k], "./")
	}
	for k := range configs {
		CopyDir(configs[k], "./.config/")
	}
	pushToGit()
}
