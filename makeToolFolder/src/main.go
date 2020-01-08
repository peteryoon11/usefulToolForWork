package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {

	/**
	폴더명 받기
	*/
	fmt.Println(os.Args[1:])
	if len(os.Args[1:]) == 0 {
		fmt.Println("폴더 명이 없습니다. / need folder name")
		return
	}
	folderName := os.Args[1:][0]

	fmt.Println(folderName)
	fmt.Println()
	// os.CreateDirIfNotExist("./test")
	// os.Mkdir("."+string(filepath.Separator)+"TestDir", 0777)
	mkfolderList := []string{"src", "pkg"}
	for i := 0; i < len(mkfolderList); i++ {
		// fmt.Println(mkfolderList[i])
		// fmt.Println(filepath.Separator)
		// fmt.Println(filepath.Separator)
		// fmt.Println(string(filepath.Separator))
		filePath := "." + string(filepath.Separator) + folderName + string(filepath.Separator) + mkfolderList[i]
		gitkeepPath := "." + string(filepath.Separator) + folderName + string(filepath.Separator) + mkfolderList[i] + string(filepath.Separator) + ".gitkeep"
		mainFile := "." + string(filepath.Separator) + folderName + string(filepath.Separator) + mkfolderList[i] + string(filepath.Separator) + "main.go"
		fmt.Println("filePath " + filePath)
		os.MkdirAll(filePath, 0777)
		cmd := exec.Command("touch", gitkeepPath)
		cmd.Run()
		if strings.EqualFold(mkfolderList[i], "src") {
			cmd = exec.Command("touch", mainFile)
			cmd.Run()
			d1 := []byte(`package main
						  import "fmt"
			
						  func main() {
						  	fmt.Println("Hello World")
						  }
			`)
			err := ioutil.WriteFile(mainFile, d1, 0644)
			if err != nil {

			}
		}

		// cmd := exec.Command("ls", "-lah")
		// if runtime.GOOS == "windows" {
		// 	cmd = exec.Command("tasklist")
		// }

		// fmt.Println(cmd)
	}

}

func makeIninMain() {
	//test
}
