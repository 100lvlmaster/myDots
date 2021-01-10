package main

import (
	"fmt"
	"log"
	"os/exec"
)

//CopyDir will copy files to this dir
func CopyDir(src, dst string) error {
	cmd := exec.Command("cp", "-R", src, dst)
	log.Printf("Running cp -a")
	fmt.Println(src)
	return cmd.Run()
}

// func pushToGit() {
// 	val := exec.Command("git", "init")
// 	val.Run()
// 	nol := exec.Command("git", "add", ".")
// 	nol.Run()
// 	commit := exec.Command("git", "commit", "-m", fmt.Sprintf("%d", rand.Intn(100000)))
// 	commit.Run()
// 	all := exec.Command("git", "push", "origin", "master")
// 	all.Run()

// }
func main() {
	var mainDir = "~/.config"
	configs := map[string]string{
		"alacritty": mainDir + "/alacritty",
		"dunst":     mainDir + "/dunst",
		"polybar":   mainDir + "/polybar",
		"i3":        mainDir + "/i3",
	}
	miscFiles := map[string]string{
		"zshrc":  "~/.zshrc",
		"bashrc": "~/.bashrc",
		"walls":  "~/wallpapers",
	}

	// for j := 0; j < len(configFiles); j++ {

	// }

	for k := range miscFiles {
		CopyDir(miscFiles[k], "./")
	}
	for k := range configs {
		CopyDir(configs[k], "./.config/")
		fmt.Printf("key[%s] value[%s]\n", k, configs[k])
	}
	// pushToGit()
}
