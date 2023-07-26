package my_other_module

import (
	double_linked_list "github.com/Mexicoder/my-golang-monorepo/deeper-module/double-linked-list"
	linked_list "github.com/Mexicoder/my-golang-monorepo/linked-list"
)

func main() {
	list := linked_list.SingleLinkedList{}
	_ = list
	otherlist := double_linked_list.DoublyLinkedList{}
	_ = otherlist
}

//git remote add origin https://github.com/Mexicoder/my-golang-monorepo.git
//git push -u origin main
