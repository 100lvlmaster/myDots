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
	var userName string = "navin"
	fileNames := []string{"wallpapers", ".bashrc", ".zshrc"}
	var srcdir = "/home/" + userName + "/"
	var destdir = "./"
	for i := 0; i < len(fileNames); i++ {
		CopyDir(srcdir+fileNames[i], destdir)

	}
	var confiLocation = srcdir + ".config/"
	configFiles := []string{"alacritty", "i3", "dunst"}
	for j := 0; j < len(configFiles); j++ {
		CopyDir(confiLocation+configFiles[j], "./.config/")

	}

	pushToGit()
}
