package main

import "fmt"

type Counter struct {
	value int
}

func (c *Counter) Value() int {
	return c.value
}

func (c *Counter) Inc() {
	c.value += 1
}

func main() {
	fmt.Println("Hello World!")
}
