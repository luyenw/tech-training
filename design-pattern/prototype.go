package main

import "fmt"

type Cloneable interface {
	clone() Cloneable
}

type Computer struct {
	os      string
	browser string
}

func (c *Computer) clone() Computer {
	return Computer{
		os:      c.os,
		browser: c.browser,
	}
}

func main() {
	c := Computer{"windows", "firefox"}
	clone := c.clone()
	fmt.Printf("%+v, %p\n", c, &c)
	fmt.Printf("%+v, %p\n", clone, &clone)
}
