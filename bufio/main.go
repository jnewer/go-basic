package main

import (
	"bufio"
	// "bytes"
	"fmt"
	"strings"
)

// "io"
// "strings"

func main() {
	// s := strings.NewReader("ABCDEFG")
	// str := strings.NewReader("123456")
	// br := bufio.NewReader(s)
	// b, _ := br.ReadString('\n')
	// fmt.Println(b)
	// br.Reset(str)
	// b, _ = br.ReadString('\n')
	// fmt.Println(b)

	// s := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	// br := bufio.NewReader(s)
	// p := make([]byte, 10)
	// for {
	// 	n, err := br.Read(p)
	// 	if err == io.EOF {
	// 		break
	// 	} else {
	// 		fmt.Printf("string(p[0:n]): %v\n", string(p[0:n]))
	// 	}
	// }

	// s := strings.NewReader("ABCDEFG")
	// br := bufio.NewReader(s)

	// c, _ := br.ReadByte()
	// fmt.Printf("c: %v\n", c)

	// c, _ = br.ReadByte()
	// fmt.Printf("c: %v\n", c)

	// br.UnreadByte()
	// c, _ = br.ReadByte()
	// fmt.Printf("c: %v\n", c)

	// s := strings.NewReader("你好，世界！")
	// br := bufio.NewReader(s)

	// c, size, _ := br.ReadRune()
	// fmt.Printf("%c %v\n", c, size)

	// c, size, _ = br.ReadRune()
	// fmt.Printf("%c %v\n", c, size)

	// br.UnreadRune()
	// c, size, _ = br.ReadRune()
	// fmt.Printf("%c %v\n", c, size)

	// s := strings.NewReader("ABC\nDEF\r\nGHI\r\nJKL")
	// br := bufio.NewReader(s)

	// w, isPrefix, _ := br.ReadLine()
	// fmt.Printf("%q %v\n", w, isPrefix)

	// w, isPrefix, _ = br.ReadLine()
	// fmt.Printf("%q %v\n", w, isPrefix)

	// w, isPrefix, _ = br.ReadLine()
	// fmt.Printf("%q %v\n", w, isPrefix)

	// w, isPrefix, _ = br.ReadLine()
	// fmt.Printf("%q %v\n", w, isPrefix)

	// s := strings.NewReader("ABC,DEF,GHI,JKL")
	// br := bufio.NewReader(s)

	// w, _ := br.ReadSlice(',')
	// fmt.Printf("w: %q\n", w)

	// w, _ = br.ReadSlice(',')
	// fmt.Printf("w: %q\n", w)

	// w, _ = br.ReadSlice(',')
	// fmt.Printf("w: %q\n", w)

	// s := strings.NewReader("ABC DEF GHI JKL")
	// br := bufio.NewReader(s)

	// w, _ := br.ReadBytes(' ')
	// fmt.Printf("%q\n", w)

	// w, _ = br.ReadBytes(' ')
	// fmt.Printf("%q\n", w)

	// w, _ = br.ReadBytes(' ')
	// fmt.Printf("%q\n", w)

	// s := strings.NewReader("ABC DEF GHI JKL")
	// br := bufio.NewReader(s)

	// w, _ := br.ReadString(' ')
	// fmt.Printf("%q\n", w)

	// w, _ = br.ReadString(' ')
	// fmt.Printf("%q\n", w)

	// w, _ = br.ReadString(' ')
	// fmt.Printf("%q\n", w)

	// s := strings.NewReader("ABCDEFGHIJKLMN")
	// br := bufio.NewReader(s)
	// b := bytes.NewBuffer(make([]byte, 0))
	// br.WriteTo(b)
	// fmt.Printf("b: %s\n", b)

	// b := bytes.NewBuffer(make([]byte, 0))
	// bw := bufio.NewWriter(b)
	// bw.WriteString("123456789")
	// c := bytes.NewBuffer(make([]byte, 0))
	// bw.Reset(c)
	// bw.WriteString("456")
	// bw.Flush()
	// fmt.Println(b)
	// fmt.Println(c)

	// b := bytes.NewBuffer(make([]byte, 0))
	// bw := bufio.NewWriter(b)
	// fmt.Println(bw.Available()) //4096
	// fmt.Println(bw.Buffered())  //0

	// bw.WriteString("ABCDEFGHIJKLMN")
	// fmt.Println(bw.Available()) //4082
	// fmt.Println(bw.Buffered())  //14
	// fmt.Printf("b: %q\n", b)    //b: ""

	// bw.Flush()
	// fmt.Println(bw.Available()) //4096
	// fmt.Println(bw.Buffered())  //0
	// fmt.Printf("b: %q\n", b)    //b: "ABCDEFGHIJKLMN"

	// b := bytes.NewBuffer(make([]byte, 0))
	// bw := bufio.NewWriter(b)

	// bw.WriteByte('H')
	// bw.WriteByte('e')
	// bw.WriteByte('l')
	// bw.WriteByte('l')
	// bw.WriteByte('o')
	// bw.WriteByte(' ')
	// bw.WriteRune('世')
	// bw.WriteRune('界')
	// bw.WriteRune('！')

	// bw.Flush()
	// fmt.Println(b)

	// b := bytes.NewBuffer(make([]byte, 0))
	// s := strings.NewReader("Hello 世界！")
	// bw := bufio.NewWriter(b)
	// bw.ReadFrom(s)
	// fmt.Println(b)

	// b := bytes.NewBuffer(make([]byte, 0))
	// bw := bufio.NewWriter(b)
	// s := strings.NewReader("123")
	// br := bufio.NewReader(s)
	// rw := bufio.NewReadWriter(br, bw)
	// p, _ := rw.ReadString('\n')
	// fmt.Println(string(p))
	// rw.WriteString("asdf")
	// rw.Flush()
	// fmt.Println(b)

	// s := strings.NewReader("ABC DEF GHI JKL")
	// bs := bufio.NewScanner(s)
	// bs.Split(bufio.ScanWords)
	// for bs.Scan() {
	// 	fmt.Println(bs.Text())
	// }

	s := strings.NewReader("Hello 世界！")
	bs := bufio.NewScanner(s)
	// bs.Split(bufio.ScanBytes)//乱码
	bs.Split(bufio.ScanRunes)
	for bs.Scan() {
		fmt.Printf("%s ", bs.Text())
	}

}
