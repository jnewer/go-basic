package main

import (
	"fmt"
	"os"
)

func openCloseFile() {
	f, _ := os.Open("a.txt")
	fmt.Printf("f.Name(): %v\n", f.Name()) //f.Name(): a.txt

	f2, _ := os.OpenFile("a1.txt", os.O_RDWR|os.O_CREATE, 0755)
	fmt.Printf("f2.Name(): %v\n", f2.Name()) //f2.Name(): a1.txt
	err := f.Close()
	fmt.Printf("err: %v\n", err) //err: <nil>

	err2 := f2.Close()
	fmt.Printf("err2: %v\n", err2) //err2: <nil>
}

func createFile() {
	f, _ := os.Create("a2.txt")
	fmt.Printf("f.Name(): %v\n", f.Name()) //f.Name(): a2.txt

	f2, _ := os.CreateTemp("", "temp")
	fmt.Printf("f2.Name(): %v\n", f2.Name()) //f2.Name(): C:\Users\kang\AppData\Local\Temp\temp1225017327
}

func readOps() {
	/* f, _ := os.Open("a.txt")
	for {
		buf := make([]byte, 6)
		n, err := f.Read(buf)
		fmt.Println(string(buf))
		fmt.Printf("n: %v\n", n)
		if err == io.EOF {
			break
		}
	}
	f.Close() */

	/* buf := make([]byte, 10)
	f2, _ := os.Open("a.txt")
	n, _ := f2.ReadAt(buf, 5)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("string(buf): %v\n", string(buf))
	f2.Close() */

	/* f, _ := os.Open("a")
	de, _ := f.ReadDir(-1)
	for _, v := range de {
		fmt.Printf("v.IsDir(): %v\n", v.IsDir())
		fmt.Printf("v.Name(): %v\n", v.Name())
	} */

	f, _ := os.Open("a.txt")
	f.Seek(3, 0)
	buf := make([]byte, 10)
	n, _ := f.Read(buf)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("string(buf): %v\n", string(buf))
	f.Close()

}
func main() {
	// openCloseFile()
	// createFile()
	readOps()

}
