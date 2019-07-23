package main 

import(
	"fmt"
)

//初始化
type data struct {
	x int
	s string
}

func main () {
	const v = 20
	var a byte = 10
	c := a + v //v 自动转换成byte
	fmt.Printf("%v, %T\n", c, c) //uint8
	
	var b float32 = 1.2
	d := b + v //v 自动转换成float32
	fmt.Printf("%v, %T\n", d, d) //float32

	m := 1
	//++a 没有这种操作
	//if (a++) > 1 {}  语句不能作为表达式使用
	
	p := &m
	*p++
	fmt.Println(m); //2
	
	//指针 专门存放地址的类型；其本身也是有一个地址
	x := 10
	var px *int = &x
	*px += 20
	fmt.Println(px, *px) //0xc00004a0a8 30
	
	//初始化
	y := data{ //括号不能换行
		1,
		"str", //需要有逗号或者括号不换行
	}
	fmt.Println(y)
	
	//switch
	r, u, i, z := 1, 2, 3, 6
	switch z {
		case r, u: //这个是 or; 只需要匹配一个上就行
			fmt.Println("r | u")
		case i:
			fmt.Println("i")
		case 4:
			fmt.Println("d")
		default:  //default顺序无关,不会因为在上面就优先执行
			fmt.Println("z")
		case 6:
			fmt.Println(6)
			fallthrough //贯通 继续执行下一case 不再匹配条件 break可以终止
		case 7:
			fmt.Println(7) //如果6使用了fallthrough 将要被执行
		case 8:
			fmt.Println(8)			
	}
	
	//break,continue
	/*
		0:0 0:1 0:2
		1:0 1:1 1:2
		2:0 2:1 2:2
	*/
	outer:
		for p := 0; p < 5; p++ {
			for t :=0; t < 10; t++ {
				if t > 2 {
					fmt.Println()
					continue outer
				}
				
				if p > 2 {
					break outer
				}
				
				fmt.Print(p, ":", t, " ")
			}
		}
}