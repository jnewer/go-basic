package main

import (
	"fmt"
	"time"
)

func test1() {
	now := time.Now()
	fmt.Printf("current time: %v\n", now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	fmt.Printf("%T,%T,%T,%T,%T,%T,%T\n", now, year, month, day, hour, minute, second)
	// time.Time,int,time.Month,int,int,int,int
}

func test2() {
	now := time.Now()
	fmt.Printf("Timestamp type:%T, Timestamp:%v", now.Unix(), now.Unix())
}

func test3() {
	now := time.Now()
	fmt.Printf("Timestamp type:%T, Timestamp:%v", now.UnixNano(), now.UnixNano())

}

func timestampToTime() {
	timestamp := time.Now().Unix()
	timeObj := time.Unix(timestamp, 0)
	fmt.Println(timeObj)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

}

func add(h, m, s, mls, msc, ns time.Duration) {
	now := time.Now()
	fmt.Println(now.Add(time.Hour*h + time.Minute*m + time.Second*s + time.Millisecond*mls + time.Microsecond*msc + time.Nanosecond*ns))
}

func sub() {
	now := time.Now()
	targetTime := now.Add(time.Hour)
	fmt.Println(targetTime.Sub(now))

}

func tick() {
	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Println(i)
	}
}

func fomart() {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}
func main() {
	// test1()
	// test2()
	// test3()
	// timestampToTime()
	// add(3, 4, 5, 6, 7, 8)
	// sub()
	// tick()
	// fomart()

	now := time.Now()
	fmt.Println(now)

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}

	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))
}
