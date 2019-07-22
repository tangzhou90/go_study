package main 

import (
	"fmt"
	"time"
)

func task(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d: %d\n", id, i)
		time.Sleep(time.Second)
	}
}

//使用channel和goroutine实现

//消费者
func consumer(data chan int, done chan bool) {
	for v := range data {
		fmt.Println("recv: ", v)
	}
	done <- true
}

//生产者
func producer(data chan int) {
	for i := 0; i < 5; i++ {
		data <- i //往通道里面写入数据
	}
	
	close(data) //关闭通道
}

func main () {
	go task(1)
	go task(2)
	
	time.Sleep(time.Second * 6)
	
	done := make(chan bool) //用于接收结束信号
	data := make(chan int) //数据管道
	
	go consumer(data, done) //启动消费者
	go producer(data) //启动生产者
	<-done //阻塞;直到消费者发回结束信号, 类似于上面的Sleep;
}
