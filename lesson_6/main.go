package main

import "lesson_6/LinkedList"

func main() {
	var list = []int{1, 2, 3, 4, 5}
	var ll = LinkedList.CreatList(list)
	ll.Info()
	var dupl = []int{0, 1, 2, 2, 3, 4, 4, 5}
	var lll = LinkedList.CreatList(dupl)
	LinkedList.DeleteDuplicates(lll)
	lll.Info()

}
