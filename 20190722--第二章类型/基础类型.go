package main

import(
	"fmt"
)

var x int = 100 //全局变量
type flags byte //自定义类型

func main() {
	fmt.Println(&x, x)
	//变量的定义 
	var x int //零值 和上面的全局不变量不是一个
	fmt.Println(&x, x)
	var y = false //不用指定类型
	z := 123 //类型推导
	var a,s = 100, "abc" //定义多个变量  a,s := 1,'abc' 这样也行
	fmt.Println(a,s,y,z)
	
	{
		x,y := 100,200 //不同作用域,全局变量定义
		fmt.Println(&x, x, y)
	}
	
	//常量 常量没有使用不会引起报错
	const m,n int = 1,2
	const t = "hello,world"
	const c = '我' //rune类型
	const (
		i, f = 1, 0.123 //int,float64(小数默认是float64)
		b = false
	)
	
	const (
		a1 uint16 =120
		a2     				//与上一行a1的类型相同,右值相同
		a3 = "abc"
		a4					//和a3一样
	)
	
	fmt.Printf("%T, %v\n", a2, a2) //uint16, 120
	fmt.Printf("%T, %v\n", a4, a4) //string, abc
	
	//枚举
	const (
		_, _ = iota, iota * 10  //0,0*10
		sa, sb 					//1, 1*10
		sc, sd					//2, 2*10
	)
	const (
		xa = iota  //0
		xb 		  //1
		xc = 100   //100
		xd         //100
		xe = iota //4
		xf        //5   
	)
	
	//byte 别名 uint8 rune 别名int32
	var t1 byte = 0x11
	var t2 uint8 = t1
	var t3 uint8 = t1 + t2
	fmt.Println(t3)
	
	//var t4 int = 100
	//var t5 int64 = t4 //这是错误的不能用来计算 t4 + t5;
	
	//引用类型 slice,map,channel 他们的长度,key,val类型不一样就视为不同类型
	z1 := mkSlice();
	fmt.Println(z1[0])
	z2 := mkmap();
	fmt.Println(z2["a"])
	
	//自定义类型
	const (
		read flags = 1 << iota
		write
		exec
	)
	//组
	type(
		test struct {
			name string
			age uint8
		}
		
		event func(string) bool //函数类型
	)
	s1 := test{"tang", 18}
	fmt.Println(s1)
	
	var s2 event = func (str string) bool {
		fmt.Println(str)
		return s != ""
	}
	s2("abc")
}

func mkSlice() []int {
	s := make([]int, 0 , 10)
	s = append(s, 10)
	return s
}

func mkmap() map[string]int {
	m := make(map[string]int)
	m["a"] = 1
	return m 
}
