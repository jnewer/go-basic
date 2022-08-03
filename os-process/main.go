package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//当前进程id
	fmt.Printf("os.Getpid(): %v\n", os.Getpid())
	//父id
	fmt.Printf("os.Getppid(): %v\n", os.Getppid())

	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Env:   os.Environ(),
	}
	p, err := os.StartProcess("C:\\Windows\\System32\\notepad.exe", []string{"C:\\Windows\\System32\\notepad.exe", "D:\\a.txt"}, attr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
	fmt.Println("进程ID", p.Pid)

	p2, _ := os.FindProcess(p.Pid)
	fmt.Println(p2)

	time.AfterFunc(time.Second*10, func() {
		p.Signal(os.Kill)
	})

	ps, _ := p.Wait()
	fmt.Println(ps.String())
}
