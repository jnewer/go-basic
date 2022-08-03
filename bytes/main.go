package main

import (
	"bytes"
	"fmt"
)

func main() {
	// var i int = 1
	// var j byte = 2
	// j = byte(i)
	// fmt.Printf("j: %v\n", j) //j: 1

	// //Contains
	// b := []byte("duoke360.com")
	// sublice1 := []byte("duoke360")
	// sublice2 := []byte("Duoke360")
	// fmt.Println(bytes.Contains(b, sublice1)) //true
	// fmt.Println(bytes.Contains(b, sublice2)) //false

	// //Count
	// s := []byte("hellooooooooo")
	// sep1 := []byte("h")
	// sep2 := []byte("l")
	// sep3 := []byte("o")
	// fmt.Println(bytes.Count(s, sep1)) //1
	// fmt.Println(bytes.Count(s, sep2)) //2
	// fmt.Println(bytes.Count(s, sep3)) //9

	// //Repeat
	// b = []byte("hi")
	// fmt.Println(string(bytes.Repeat(b, 1))) //hi
	// fmt.Println(string(bytes.Repeat(b, 3))) //hihihi

	// //Replace
	// s = []byte("hello,world")
	// old := []byte("o")
	// news := []byte("ee")
	// fmt.Println(string(bytes.Replace(s, old, news, 0)))  //hello,world
	// fmt.Println(string(bytes.Replace(s, old, news, 1)))  //hellee,world
	// fmt.Println(string(bytes.Replace(s, old, news, 2)))  //hellee,weerld
	// fmt.Println(string(bytes.Replace(s, old, news, -1))) //hellee,weerld

	// //Runes
	// s = []byte("你好世界")
	// r := bytes.Runes(s)
	// fmt.Println("转换前字符串的长度：", len(s)) //转换前字符串的长度： 12
	// fmt.Println("转换后字符串的长度：", len(r)) //转换后字符串的长度： 4

	// // Join
	// s2 := [][]byte{[]byte("你好"), []byte("世界")}
	// sep4 := []byte(",")
	// fmt.Println(string(bytes.Join(s2, sep4))) //你好,世界
	// sep5 := []byte("#")
	// fmt.Println(string(bytes.Join(s2, sep5))) //你好#世界

	data := "123456789"
	re := bytes.NewReader([]byte(data))
	fmt.Printf("re.Len(): %v\n", re.Len()) //re.Len(): 9

	fmt.Printf("re.Size(): %v\n", re.Size()) //re.Size(): 9
	fmt.Println("----------")

	buf := make([]byte, 2)
	for {
		n, err := re.Read(buf)
		if err != nil {
			break
		}

		fmt.Println(string(buf[:n]))
	}
	fmt.Println("----------")

	re.Seek(0, 0)
	for {
		b, err := re.ReadByte()
		if err != nil {
			break
		}

		fmt.Println(string(b))
	}

	fmt.Println("----------")

	re.Seek(0, 0)
	off := int64(0)
	for {
		n, err := re.ReadAt(buf, off)
		if err != nil {
			break
		}

		off += int64(n)
		fmt.Println(off, string(buf[:n]))
	}
}
