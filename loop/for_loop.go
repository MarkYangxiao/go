package main

import "fmt"

type Item struct {
	value int
}

func demo1() {
	// 基本数据类型 情况1：arr 里本来都是地址
	var basicItemsP []*int = []*int{NewIntP(1), NewIntP(2), NewIntP(3)}
	var all []*int
	for _, item := range basicItemsP {
		all = append(all, item)
	}
	fmt.Printf("all item=%+v\n", all) //all item=[0x14000020098 0x140000200a0 0x140000200a8]
	fmt.Println("-----------------------------")
	var all1 []*int
	var basicItems = []int{1, 2, 5}
	for _, item := range basicItems {
		all1 = append(all1, &item)
	}
	fmt.Printf("all1 item =%+v\n", all1) //all1 item =[0x1400018c028 0x1400018c028 0x1400018c028]
	// 可以发现如果是取地址，全部都是最后一个slice的最后一个元素的地址
	fmt.Println(*all1[0], *all1[1], *all1[2]) // 5,5,5
	fmt.Println("===========================")
	item1 := Item{value: 1}
	item2 := Item{value: 2}
	item3 := Item{value: 3}
	var allItems = []Item{item1, item2, item3}
	var allItemsP []*Item
	for _, item := range allItems {
		allItemsP = append(allItemsP, &item)
	}
	fmt.Printf("allItemsP item=%+v\n", allItemsP)
	fmt.Println(*allItemsP[0], *allItemsP[1], *allItemsP[2]) // {3}, {3}, {3}
	// 同上

	//得出结论： 如果是for循环slice里的每个元素的地址，得到的全部都是slice最后一个元素的地址
	// 如何解决： 构造临时变量
	var allItemsPP []*Item
	for _, item := range allItems {
		item := item // 构造一个临时变量
		allItemsPP = append(allItemsPP, &item)
	}
	fmt.Println(*allItemsPP[0], *allItemsPP[1], *allItemsPP[2]) // {1}, {2}, {3}
	allItemsPP = allItemsPP[0:0]
	for i := range allItems {
		item := allItems[i] // 构造一个临时变量
		allItemsPP = append(allItemsPP, &item)
	}
	fmt.Println(*allItemsPP[0], *allItemsPP[1], *allItemsPP[2]) // {1}, {2}, {3}

	//以上两种构造临时变量的方法
}

func demo2() {
	var prints []func()
	for _, v := range []int{1, 2, 3} {
		prints = append(prints, func() {
			fmt.Println(v)
		})
	}

	for _, print := range prints {
		print() // 3, 3, 3
	}
	prints = prints[0:0] //清空prints
	fmt.Println("===========================")
	//for 循环结束，闭包函数的结果都变成了3
	for _, v := range []int{1, 2, 3} {
		v := v
		prints = append(prints, func() {
			fmt.Println(v)
		})
	}
	for _, print := range prints {
		print() // 1,2,3
	}
}

func main() {
	demo1()
	demo2()
	// demo1,demo2举例了两种情况下for循环会导致的问题，如果记不住，可以统一都再次赋值给临时变量(两种方式)
}

func NewIntP(i int) *int {
	return &i
}
