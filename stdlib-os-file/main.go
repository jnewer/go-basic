package main

import (
	"fmt"
	"os"
)

func createFile() {
	f, err := os.Create("./test.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f: %v\n", f)
	}
}

func createDir() {
	// err := os.Mkdir("test", os.ModePerm)
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// }

	err := os.MkdirAll("test/a/b", os.ModePerm)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
func removeDir() {
	err := os.RemoveAll("test")

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func getWd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("dir: %v\n", dir)
	}
}

func chWd() {
	err := os.Chdir("d:/")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Println(os.Getwd())

}

func renameFile() {
	err := os.Rename("test.txt", "test2.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func readFile() {
	b, err := os.ReadFile("test2.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("b: %v\n", string(b[:]))
	}
}

func writeFile() {
	s := "hello world"
	os.WriteFile("test2.txt", []byte(s), os.ModePerm)
}
func main() {
	// createFile()
	// createDir()
	// removeDir()
	// getWd()
	// chWd()
	// renameFile()
	// readFile()
	writeFile()

}
