package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

var (
	DIRNAMES = []string{"src", "pkg", "bin"}
	noSrc    = flag.Bool("no-src", false, "If src folder has to be skipped")
	args     []string
	path     string
	executableName
)

// Returns absolute path or relative path of a directory and err if not found
// Searches for an executable binary named file in the directories
// Named by the PATH environment variable
func lookPath(path string) (string, error) {
	return exec.LookPath(path)

}

// Get the current working directory (directory where the command was invoked)
func getwd() (string, error) {
	return os.Getwd()
}

func setGoPath(path string) {
	log.Printf("Setting GOPATH: %s", path)
	os.Setenv("GOPATH", path)
}

func createDirIfNotExists(path string) {
	for _, dirName := range DIRNAMES[1:] {
		if _, errDir := os.Stat(path + dirName); os.IsNotExist(errDir) {
			log.Printf("Creating directory %v", dirName)
			os.Mkdir(path+dirName, os.ModePerm)
		}
	}
}

func createGoPathLayout(dir string) {
	setGoPath(dir)
	createDirIfNotExists(dir + "/")
}

// This will find the path before the "first" occurance of src folder
// Then set the path as GOPATH
// eg. $HOME/workspace/client1/src/myproject/src then,
// GOPATH will be $HOME/workspace/client1/
func goPathLayout(dir string) {
	index := -1

	if index = strings.Index(dir, DIRNAMES[0]); index == -1 {
		if _, errDir := os.Stat(dir + "/" + DIRNAMES[0]); os.IsNotExist(errDir) {
			log.Fatal("Please make sure your directory has atleast three directories at it's root - %v", DIRNAMES)
		} else {
			createGoPathLayout(dir)
		}

	} else {
		createGoPathLayout(dir[:index-1])
	}
}

func parseCmdArgs(command []string) (executableName string, args []string) {
	return command[0], command
}

func main() {
	flag.Parse()

	if *noSrc == false {
		executableName, args := parseCmdArgs(os.Args[1:])
		dir, getwdErr := getwd()
		if getwdErr != nil {
			log.Fatal(getwdErr)
		}
		goPathLayout(dir)
	} else {
		executableName, args := parseCmdArgs(flag.Args())
	}

	binary, lookErr := lookPath(executableName)
	if lookErr != nil {
		log.Fatal(lookErr)
	}

	execErr := syscall.Exec(binary, args, os.Environ())
	if execErr != nil {
		log.Fatal(execErr)
	}
}
