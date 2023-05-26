package main

import (
	"fmt"
)

type check struct {
	n []int
}

func (c *check) add(a int) {
	c.n = append(c.n, a)
}

func (c *check) remove(b int) {
	for i, num := range c.n {
		if num == b {
			c.n = append(c.n[:i], c.n[i+1:]...)
			break
		}
	}
}

func (c *check) dis() []int {
	return c.n
}

func main() {
	obj := check{}
	choice := 1
	for choice != 0 {
		fmt.Println("0. Exit")
		fmt.Println("1. Add")
		fmt.Println("2. Delete")
		fmt.Println("3. Display")
		fmt.Print("Enter choice: ")
		fmt.Scanln(&choice)

		if choice == 1 {
			var n int
			fmt.Print("Enter number to append: ")
			fmt.Scanln(&n)
			obj.add(n)
			fmt.Println("List:", obj.dis())

		} else if choice == 2 {
			var n int
			fmt.Print("Enter number to remove: ")
			fmt.Scanln(&n)
			obj.remove(n)
			fmt.Println("List:", obj.dis())

		} else if choice == 3 {
			fmt.Println("List:", obj.dis())

		} else if choice == 0 {
			fmt.Println("Exiting!")
		} else {
			fmt.Println("Invalid choice!!")
		}
	}
}
