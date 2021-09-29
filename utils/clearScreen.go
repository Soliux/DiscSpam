package utils

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	/*
		Here we are creating a cross platform clear screen function. This function has the ability to detect the operating system it is being ran on and then clear the screen accordingly.
	*/
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if runtime.GOOS == "linux" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

}
