package main

import "bufio"
import "fmt"
import "os"
import "os/exec"
import "path/filepath"

import "github.com/spf13/viper"

func main() {
	//usage() and exit when run without any arguments
	if len(os.Args) == 1 {
		usage()
	}

	//check needed dependencies else this wont do much good
	if !(doesBinaryExistInPath("git")) {
		fmt.Println("git is not installed")
		os.Exit(1)
	}
	if !(doesBinaryExistInPath("direnv")) {
		fmt.Println("direnv is not installed")
		os.Exit(1)
	}

	//use viper for config management
	v := viper.New()
	v.SetConfigType("toml")
	configfile := filepath.Join(os.Getenv("HOME"), ".config", "gosca.toml")
	v.SetConfigFile(configfile)

	//read config settings or prompt for settings and create it if no config is present yet
	if err := v.ReadInConfig(); err != nil {
		fmt.Println(err)
		fmt.Println()
		fmt.Println("no config present yet, please provide desired default settings:")
		fmt.Println("full path of desired workspace?")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		ws := scanner.Text()
		//if given workspace path doesnt exist just exit
		if _, err := os.Stat(ws); err != nil {
			fmt.Println("workspace path must already exist!")
			os.Exit(1)
		}
		fmt.Println("github account name?")
		scanner.Scan()
		githubName := scanner.Text()
		gh := "github.com/" + githubName
		v.Set("workspace", ws)
		v.Set("github", gh)
		if err := v.WriteConfigAs(configfile); err != nil {
			fmt.Println(err)
		}
		fmt.Println("config written to " + configfile)
	}
	workspace := v.GetString("workspace")
	github := v.GetString("github")

	if len(os.Args) != 2 {
		fmt.Println()
		fmt.Println("wrong number of arguments.")
		fmt.Println()
		usage()
	}

	projectname := os.Args[1]
	projectpath := filepath.Join(workspace, projectname)
	path := filepath.Join(workspace, projectname, "src", github, projectname)
	//cmdpath := filepath.Join(path, "cmd", projectname)

	checkIfWorkspaceFolderExistsOrQuit(workspace)
	createFolderStructure(projectpath, path)
	createEnvRC(projectpath)
	//os.MkdirAll(cmdpath, os.ModePerm)
	//createMain(cmdpath)
	createMain(path)
	gitInit(path)
	fmt.Println("cd " + path + "  ## dev here")
}

func usage() {
	fmt.Println()
	fmt.Println("usage:")
	fmt.Println()
	fmt.Println(os.Args[0] + "  PROJECTNAME")
	fmt.Println()
	fmt.Println("config file is here: " + filepath.Join(os.Getenv("HOME"), ".config", "gosca.toml"))
	os.Exit(0)
}

func doesBinaryExistInPath(commandname string) bool {
	cmd := exec.Command("/bin/sh", "-c", "command -v "+commandname)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func checkIfWorkspaceFolderExistsOrQuit(workspace string) {
	if _, err := os.Stat(workspace); os.IsNotExist(err) {
		fmt.Printf("Error: "+workspace+" folder is missing; %s\n", err)
		os.Exit(1)
	}
}

func createFolderStructure(projectpath, path string) {
	srcpath := filepath.Join(projectpath, "src")
	binpath := filepath.Join(projectpath, "bin")
	pkgpath := filepath.Join(projectpath, "pkg")

	os.MkdirAll(srcpath, os.ModePerm)
	os.MkdirAll(binpath, os.ModePerm)
	os.MkdirAll(pkgpath, os.ModePerm)
	os.MkdirAll(path, os.ModePerm)

	fmt.Println("folders created.")
}

func gitInit(path string) {
	cmd := "git init -q " + path
	exec.Command("sh", "-c", cmd).Output()

	fmt.Println("git init.")
}

func createEnvRC(projectpath string) {
	content := `#export GOPATH=$(pwd):$GOPATH
#export PATH=$(pwd)/bin:$PATH
layout go
`

	dest := filepath.Join(projectpath, ".envrc")
	if f, err := os.Create(dest); err == nil {
		f.WriteString(content)
		f.Sync()
		defer f.Close()
	}

	fmt.Println("wrote .envrc.")

}

func createMain(path string) {
	content := `package main

import "fmt"

func main() {

	fmt.Println("asdf")
}
`

	dest := filepath.Join(path, "main.go")
	if f, err := os.Create(dest); err == nil {
		defer f.Close()
		f.WriteString(content)
		f.Chmod(0755)
		f.Sync()
	}

	fmt.Println("wrote main.go.")

}

func direnvAllow(projectpath string) {
	cmd := "direnv allow " + projectpath
	exec.Command("sh", "-c", cmd).Output()

	fmt.Println("direnv allow.")

}
