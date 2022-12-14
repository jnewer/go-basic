package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/* func show(msg string) {
	for i := 1; i < 5; i++ {
		fmt.Printf("msg: %v\n", msg)
		time.Sleep(time.Millisecond * 100)
	}
} */

func responseSize(url string) {
	fmt.Println("Step1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step2: ", url)
	defer response.Body.Close()

	fmt.Println("Step3: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step4: ", len(body))
}

func main() {
	// go show("java")
	// show("golang") // 在main协程中执行，如果它前面也添加go，程序没有输出
	// fmt.Println("end...")

	go responseSize("https://www.duoke360.com")
	go responseSize("https://baidu.com")
	go responseSize("https://jd.com")
	time.Sleep(10 * time.Second)

}
