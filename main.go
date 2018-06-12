package main

import "fmt"
import "os"
import "os/exec"
import "path/filepath"

func main() {
	workspace := filepath.Join(os.Getenv("HOME"), "wrk")
	github := "github.com/sjas"

	if len(os.Args) != 2 {
		fmt.Println()
		fmt.Println("wrong number of arguments.")
		usage()
		os.Exit(1)
	}

	projectname := os.Args[1]
	projectpath := filepath.Join(workspace, projectname)
	path := filepath.Join(workspace, projectname, "src", github, projectname)

	checkIfWorkspaceFolderExistsOrQuit(workspace)
	createFolderStructure(projectpath, path)
	createEnvRC(projectpath)
	createMain(path)
	direnvAllow(projectpath)
	gitInit(path)
	fmt.Println("cd " + path + "  ## dev here")
}

func usage() {
	fmt.Println()
	fmt.Println("usage:")
	fmt.Println()
	fmt.Println(os.Args[0] + "  PROJECTNAME")
	fmt.Println()
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
	content := `export GOPATH=$(pwd):$GOPATH
export PATH=$(pwd)/bin:$PATH
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
		f.WriteString(content)
		f.Chmod(0755)
		f.Sync()
		defer f.Close()
	}

	fmt.Println("wrote main.go.")

}

func direnvAllow(projectpath string) {
	cmd := "direnv allow " + projectpath
	exec.Command("sh", "-c", cmd).Output()

	fmt.Println("direnv allow.")

}
