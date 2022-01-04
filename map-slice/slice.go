package main

import (
	"fmt"
	"encoding/json"
)

// 数组: 数组长度已被定义
// 数组赋值，整体复制
var numArr = [5]int{1, 2, 3, 4, 5}
var wordArr = [...]string{"hello", "world"}

// 切片: （没有长度）
// {起始的指针, 长度, 容量}
var slice = []string{"hello", "world"}
// make -> slice, map, chan(均为特殊指针, 传引用)
// 生成长度为2的slice, '零'值填充
var newSlice = make([]int, 2)
// 长度可变
// newSlice = append(newSlice, 1)

func main() {
	fmt.Println("Len numArr", len(numArr))
	fmt.Println("Cap numArr", cap(numArr))

	newSlice = append(newSlice, 1)
	fmt.Println(len(newSlice), cap(newSlice))

	subSlice := newSlice[0:2]
	bs, _ := json.Marshal(subSlice)
	fmt.Println(string(bs))

	testArr := [...]string{}
	bs, _ = json.Marshal(testArr)
	fmt.Println(string(bs))
	fmt.Println("Len: ", len(testArr))
}