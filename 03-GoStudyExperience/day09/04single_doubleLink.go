package main

import "fmt"

type heroNode struct {
	no   int
	name string
	pre  *heroNode
	next *heroNode
}

func insertNode(head *heroNode, newNode *heroNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newNode
}
func insertNode2(head *heroNode, newNode *heroNode) {
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newNode.no {
			break
		} else if temp.next.no == newNode.no {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		fmt.Println("已存在该排名", newNode.no)
	} else {
		newNode.next = temp.next
		temp.next = newNode
	}
}
func singleListLink(head *heroNode) {
	temp := head
	if temp == nil {
		fmt.Println("link is empty")
		return
	}
	for {
		fmt.Printf("[%v %v]->", temp.next.no, temp.next.name)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}
func delNode(head *heroNode, id int) {
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
	} else {
		fmt.Println("不存在改变号的节点")
	}
}
func main() {
	head := &heroNode{}
	node1 := &heroNode{
		no:   1,
		name: "宋江",
		next: nil,
	}
	node2 := &heroNode{
		no:   2,
		name: "卢俊义",
		next: nil,
	}
	node3 := &heroNode{
		no:   3,
		name: "林冲",
		next: nil,
	}
	node4 := &heroNode{
		no:   4,
		name: "吴用",
		next: nil,
	}
	//insertNode2(head, node3)
	//insertNode2(head, node1)
	//insertNode2(head, node2)
	//insertNode2(head, node4)
	//singleListLink(head)
	//delNode(head, 2)
	//fmt.Println()

	dInsertNode(head, node1)
	dInsertNode(head, node2)
	//dInsertNode(head, node3)
	dInsertNode(head, node4)
	singleListLink(head)
	fmt.Println()
	dListLink(head)
	doubleInsertNode1(head, node3)
	fmt.Println()
	dListLink(head)
	fmt.Println()
	dDelNode(head, 2)
	dListLink(head)

}

func dInsertNode(head *heroNode, newNode *heroNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newNode
	newNode.pre = temp
}
func dListLink(head *heroNode) {
	temp := head
	if temp == nil {
		fmt.Println("link is empty")
		return
	}
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	for {
		fmt.Printf("[%v %v]->", temp.no, temp.name)
		temp = temp.pre
		if temp.pre == nil {
			break
		}
	}
}
func doubleInsertNode1(head *heroNode, newNode *heroNode) {
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newNode.no {
			break
		} else if temp.next.no == newNode.no {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		fmt.Println("已存在该排名", newNode.no)
	} else {
		newNode.next = temp.next
		newNode.pre = temp
		if temp.next != nil {
			temp.next.pre = newNode
		}
		temp.next = newNode

	}
}
func dDelNode(head *heroNode, id int) {
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Println("不存在改变号的节点")
	}
}
