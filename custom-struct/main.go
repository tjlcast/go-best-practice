package main

import (
	"fmt"
)

type Notifier interface {
	notify()
}

// 定义 User 结构体
type User struct {
	Name string
	Email string
}

type Admin struct {

}

func (admin Admin) notify() {

}

func SendNotification(n Notifier) {

}

// 值接收者 Value Receiver: 调用时使用这个值的副本来执行
// 指针接收者 Pointer Receiver: 共享调用方法时接受者所指向的值
func (self User) notify() {
	self.Name = self.Name + " notify"
}

func (self *User) notifyPtr() {
	self.Name = self.Name + " notify"
}

func main() {
	user := User {
		Name: "hello",
		Email: "hello@qq.com",
	}

	user.notify()
	fmt.Printf(user.Name)
}