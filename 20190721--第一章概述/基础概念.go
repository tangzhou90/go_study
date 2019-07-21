package main

import(
	"fmt"
	"errors"
)

type user struct {
	name string
	age byte
}

func (u user) ToSting() (string) {
	return fmt.Sprintf("%+v", u) //Sprintf是返回一个字符串
}	

type manager struct {
	user  //类似于继承 user 结构体
	title string 
}

//实现接口
func (u user) Print() {
	fmt.Printf("%+v\n", u); //Printf 直接打印
}
type Printer interface { //接口类型
	Print()
}

func main () {
	fmt.Println("hello, world")
	
	//变量定义
	var x int32 //零值
	var s = "hello,world" //可以省略类型
	y := 456 //类型推导
	fmt.Println(x, s, y);
	
	//If，else条件表达式
	if x > 0 {
		fmt.Println("大于0");
	} else if x < 0 {
		fmt.Println("小于0");
	} else {
		fmt.Println("等于0");
	}
	
	//swtich语句
	switch {
		case y > 0:
			fmt.Println("大于0")
		case y < 0:
			fmt.Println("小于0")
		default:
			fmt.Println("等于0")
	}
	
	//for语句
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	
	//for range 遍历
	z := []int{1,2,3}
	for i, n := range z {
		fmt.Println(i, ":", n) //i 为key,n value
	}
	
	//函数
	c, err := div(1, 0)
	fmt.Println(c, err)
	d, err := div(10, 2)
	fmt.Println(d, err)
	
	//匿名函数
	f := test(100);
	f();
	
	//defer延迟调用,常用来释放资源、解除锁定,或执行一些清理操作
	test_defer(10, 1)
	
	//切片slice 类似动态的数组
	sl := make([]int, 0, 5) //创建一个容量为5的切片
	for i := 0; i < 8; i++ {
		sl = append(sl, i) //自动扩容
	}
	fmt.Println(sl)
	
	//map
	m := make(map[string]int)
	m["a"] = 1 //赋值
	v,ok := m["b"] //可以用来判断key是否存在map中(ok-idiom)
	fmt.Println(v, ok) 
	delete(m, "a"); //删除key

	//结构体 struct
	var mg manager
	mg.name = "tang"
	mg.age = 18
	mg.title = "CTO"
	fmt.Println(mg);//打印这个 {{tang 18} CTO}
	
	//方法(理解成结构体中的函数)
	fmt.Println(mg.ToSting()) //{name:tang age:18} user转换成字符串
	
	//接口
	var p Printer = mg
	p.Print()
}

func div(a, b int) (int,error) {
	if b == 0 {
		return 0, errors.New("分母不能为0")
	}
	
	return a / b , nil
}

func test(x int) func() {
	return func(){
		fmt.Println(x)
	}
}


func test_defer(a, b int) {
	defer fmt.Println("defer被调用")
	
	fmt.Println( a / b)
}