package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	//编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值
	var param01 int = 10
	addTen(&param01)
	fmt.Printf("通过指针加十后的指针值为%v \n", param01)
	// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
	param02 := []int{1, 2, 3, 4}
	var param02Num = multiplyByTwo(&param02)
	fmt.Printf("切片相乘后的值为%v \n", param02Num)
	//定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
	// 然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，
	// 创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
	var rect Rectangle = Rectangle{width: 10, height: 20}
	var cir Circle = Circle{3.65}
	fmt.Printf("Rectangle 面积为 Area=%v Rectangle 圆形为 Perimeter=%v\n", rect.Area(), rect.Perimeter())
	fmt.Printf("Circle 面积为 Area=%v Circle 圆形为 Perimeter=%v\n", cir.Area(), cir.Perimeter())
	// 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
	// 再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
	employee := Employee{Person: Person{Name: "张三", Age: "20"}, EmployeeID: "001"}
	employee.PrintInfo()
	//编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
	ch := make(chan int)
	wg.Add(1)
	go inData(ch)
	wg.Add(1)
	go outData(ch)
	wg.Wait()
	//实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
	ch1 := make(chan int, 20)
	wg.Add(1)
	go producer(ch1)
	wg.Add(1)
	go consumer(ch1)
	wg.Wait()
	//编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go countSum()
	}
	wg.Wait()
	fmt.Printf("计数器值为%v\n", count)
	//使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go atomicCountSum()
	}
	wg.Wait()
	fmt.Printf("原子计数器值为%v\n", atomic.LoadInt64(&atomicCount))
}

var atomicCount int64 = 0

func atomicCountSum() {
	for i := 1; i <= 1000; i++ {
		atomic.AddInt64(&atomicCount, 1)
	}
	defer wg.Done()
}

var count int = 0

func countSum() {
	defer wg.Done()
	mutex.Lock()
	for i := 1; i <= 1000; i++ {
		count++
	}
	mutex.Unlock()
}

func producer(ch chan int) {
	for i := 1; i <= 100; i++ {
		ch <- i
		fmt.Printf("带缓冲写入数据为%v\n", i)
	}
	close(ch)
	wg.Done()
}

func consumer(ch chan int) {
	for v := range ch {
		fmt.Printf("带缓冲读取的值为%v\n", v)
	}
	wg.Done()
}

func inData(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("写入数据为%v\n", i)
	}
	close(ch)
	wg.Done()
}

func outData(ch chan int) {
	for v := range ch {
		fmt.Printf("读取的值为%v\n", v)
	}
	wg.Done()
}

// 员工结构体
type Employee struct {
	Person
	EmployeeID string
}

// 人员结构体
type Person struct {
	Name string
	Age  string
}

func (e Employee) PrintInfo() {
	fmt.Printf("员工姓名：%v 员工年龄：%v 员工ID：%v\n", e.Name, e.Age, e.EmployeeID)
}

// 定义一个函数，用来增加10
func addTen(num *int) {
	*num += 10
}

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func multiplyByTwo(nums *[]int) []int {
	result := make([]int, len(*nums))
	for k, v := range *nums {
		result[k] = v * 2
	}
	return result
}

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
	width  int
	height int
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() int {
	return r.width * r.height
}

func (r Rectangle) Perimeter() int {
	return r.width * r.height * 2
}

func (c Circle) Area() float64 {
	return c.radius
}

func (c Circle) Perimeter() float64 {
	return c.radius * 2
}
