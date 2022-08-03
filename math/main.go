package main

import (
	"fmt"
	// "math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	// fmt.Printf("float64的最大值是:%.f\n", math.MaxFloat64)
	// fmt.Printf("float64的最小值是:%.f\n", math.SmallestNonzeroFloat64)
	// fmt.Printf("float32的最大值是:%.f\n", math.MaxFloat32)
	// fmt.Printf("float32的最小值是:%.f\n", math.SmallestNonzeroFloat32)
	// fmt.Printf("Int8的最大值是:%d\n", math.MaxInt8)
	// fmt.Printf("Int8的最小值是:%d\n", math.MinInt8)
	// fmt.Printf("Uint8的最大值是:%d\n", math.MaxUint8)
	// fmt.Printf("Int16的最大值是:%d\n", math.MaxInt16)
	// fmt.Printf("Int16的最小值是:%d\n", math.MinInt16)
	// fmt.Printf("Uint16的最大值是:%d\n", math.MaxUint16)
	// fmt.Printf("Int32的最大值是:%d\n", math.MaxInt32)
	// fmt.Printf("Int32的最小值是:%d\n", math.MinInt32)
	// fmt.Printf("Uint32的最大值是:%d\n", math.MaxUint32)
	// fmt.Printf("Int64的最大值是:%d\n", math.MaxInt64)
	// fmt.Printf("Int64的最小值是:%d\n", math.MinInt64)
	// fmt.Printf("圆周率默认为:%.200f\n", math.Pi)

	// fmt.Printf("[-3.14]的绝对值为:[%.2f]\n", math.Abs(-3.14))
	// fmt.Printf("[2]的16次方为:[%.f]\n", math.Pow(2, 16))
	// fmt.Printf("10的[3]次方为:[%.f]\n", math.Pow10(3))
	// fmt.Printf("[64]的开平方为:[%.f]\n", math.Sqrt(64))
	// fmt.Printf("[27]的开立方为:[%.f]\n", math.Cbrt(27))
	// fmt.Printf("[3.14]向上取整为:[%.f]\n", math.Ceil(3.14))
	// fmt.Printf("[8.75]向下取整为:[%.f]\n", math.Floor(8.75))
	// fmt.Printf("[8.15]四舍五入取整为:[%.f]\n", math.Round(8.15))
	// fmt.Printf("[8.75]四舍五入取整为:[%.f]\n", math.Round(8.75))
	// fmt.Printf("[10/3]的余数为:[%.f]\n", math.Mod(10, 3))

	// Integer, Decimal := math.Modf(3.14159265358979)
	// fmt.Printf("[3.14159265358979]的整数部分为:[%.f],小数部分为:[%.14f]\n", Integer, Decimal)

	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Println(a)
	}

	fmt.Println("------------")

	for i := 0; i < 10; i++ {
		a := rand.Intn(100)
		fmt.Println(a)
	}
	fmt.Println("------------")

	for i := 0; i < 10; i++ {
		a := rand.Float32()
		fmt.Println(a)
	}

}
