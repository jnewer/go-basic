package main

import (
	"fmt"
	"sort"
)

// type NewInts []uint

// func (n NewInts) Len() int {
// 	return len(n)
// }
// func (n NewInts) Less(i, j int) bool {
// 	fmt.Println(i, j, n[i] < n[j], n)
// 	return n[i] < n[j]
// }

// func (n NewInts) Swap(i, j int) {
// 	n[i], n[j] = n[j], n[i]
// }

// type testSlice [][]int

// func (l testSlice) Len() int {
// 	return len(l)
// }

// func (l testSlice) Swap(i, j int) {
// 	l[i], l[j] = l[j], l[i]
// }

// func (l testSlice) Less(i, j int) bool {
// 	return l[i][1] < l[j][1]
// }

// type testSlice []map[string]float64

// func (l testSlice) Len() int {
// 	return len(l)
// }

// func (l testSlice) Swap(i, j int) {
// 	l[i], l[j] = l[j], l[i]
// }

// func (l testSlice) Less(i, j int) bool {
// 	return l[i]["a"] < l[j]["a"] //按照"a"对应的值排序
// }
type People struct {
	Name string
	Age  int
}
type testSlice []People

func (l testSlice) Len() int {
	return len(l)
}

func (l testSlice) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l testSlice) Less(i, j int) bool {
	return l[i].Age < l[j].Age
}

func main() {
	// n := []uint{1, 3, 2}
	// sort.Sort(NewInts(n))
	// fmt.Println(n)
	// 1 0 false [1 3 2]
	// 2 1 true [1 3 2]
	// 1 0 false [1 2 3]
	// [1 2 3]

	// []float64
	// f := []float64{1.1, 4.4, 5.5, 3.3, 2.2}
	// sort.Float64s(f)
	// fmt.Printf("f: %v\n", f) //f: [1.1 2.2 3.3 4.4 5.5]

	// // []int
	// i := []int{3, 5, 1, 2, 4}
	// sort.Ints(i)
	// fmt.Printf("i: %v\n", i) //i: [1 2 3 4 5]

	// // string
	// //字符串排序，先比较高位，相同的再比较低位
	// ls := sort.StringSlice{
	// 	"100",
	// 	"42",
	// 	"41",
	// 	"3",
	// 	"2",
	// }

	// fmt.Println(ls) //[100 42 41 3 2]
	// sort.Strings(ls)
	// fmt.Println(ls) //[100 2 3 41 42]

	// //字符串排序，先比较高位，相同的再比较低位
	// ls = sort.StringSlice{
	// 	"d",
	// 	"ac",
	// 	"c",
	// 	"ab",
	// 	"e",
	// }

	// fmt.Println(ls) //[d ac c ab e]
	// sort.Strings(ls)
	// fmt.Println(ls) //[ab ac c d e]

	// ////汉字排序，依次比较byte大小
	// ls = sort.StringSlice{
	// 	"啊",
	// 	"博",
	// 	"次",
	// 	"得",
	// 	"饿",
	// 	"周",
	// }
	// fmt.Println(ls) //[啊 博 次 得 饿 周]
	// sort.Strings(ls)
	// fmt.Println(ls) //[博 周 啊 得 次 饿]

	// for _, v := range ls {
	// 	fmt.Println(v, []byte(v))
	// }

	// 博 [229 141 154]
	// 周 [229 145 168]
	// 啊 [229 149 138]
	// 得 [229 190 151]
	// 次 [230 172 161]
	// 饿 [233 165 191]

	// 复杂结构：[][]int

	// ls := testSlice{
	// 	{1, 4},
	// 	{9, 3},
	// 	{7, 5},
	// }

	// fmt.Println(ls) //[[1 4] [9 3] [7 5]]
	// sort.Sort(ls)
	// fmt.Println(ls) //[[9 3] [1 4] [7 5]]

	// 复杂结构体：[]map[string]int [{"k":0},{"k1":1},{"k2":2]
	// ls := testSlice{
	// 	{"a": 4, "b": 12},
	// 	{"a": 3, "b": 11},
	// 	{"a": 5, "b": 13},
	// }
	// fmt.Println(ls) //[map[a:4 b:12] map[a:3 b:11] map[a:5 b:13]]
	// sort.Sort(ls)
	// fmt.Println(ls) //[map[a:3 b:11] map[a:4 b:12] map[a:5 b:13]]

	// 复杂结构体：[]struct
	ls := testSlice{
		{Name: "n1", Age: 12},
		{Name: "n2", Age: 11},
		{Name: "n3", Age: 10},
	}
	fmt.Println(ls) //[{n1 12} {n2 11} {n3 10}]
	sort.Sort(ls)
	fmt.Println(ls) //[{n3 10} {n2 11} {n1 12}]
}
