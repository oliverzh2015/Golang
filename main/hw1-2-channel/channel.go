// 基于 Channel 编写一个简单的单线程生产者消费者模型：

//     队列：
//     队列长度 10，队列元素类型为 int
//     生产者：
//     每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
//     消费者：
//     每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞

package main

import (
	"fmt"
	"time"
)

// func main() {
// 	ch := make(chan int, 10)
// 	//timer := time.NewTimer(time.Second)
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			println("this is child thread")
// 			ch <- i
// 		}
// 	}()
// 	fmt.Println("print channel values")
// 	//j := <-ch
// 	//println(i)
// 	for j := range ch {
// 		fmt.Println("receiving: ", j)
// 	}
// }

func main() {
	messages := make(chan int, 10)
	done := make(chan bool)

	defer close(messages)
	// consumer
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-done:
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Printf("send message: %d\n", <-messages)
			}
		}
	}()

	// producer
	for i := 0; i < 10; i++ {
		messages <- i
	}
	time.Sleep(5 * time.Second)
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("main process exit!")
}
