package main

import (
	"fmt"
)

type Intrinsic struct {
	color string
}
type Extrinsic struct {
	x int
	y int
}
type Tree struct {
	intrinsic Intrinsic
	extrinsic Extrinsic
}
type IntrinsicFactory struct {
	colorMap map[string]*Intrinsic
}

func NewFactory() *IntrinsicFactory {
	return &IntrinsicFactory{
		colorMap: make(map[string]*Intrinsic),
	}
}
func (factory *IntrinsicFactory) getIntrinsic(color string) *Intrinsic {
	if _, ok := factory.colorMap[color]; !ok {
		factory.colorMap[color] = &Intrinsic{color: color}
	}
	return factory.colorMap[color]
}
func main() {
	tree := Tree{}
	factory := NewFactory()
	for i := 0; i < 100000000; i++ {
		tree = Tree{intrinsic: *factory.getIntrinsic("red"), extrinsic: Extrinsic{x: 0, y: i}}
	}
	fmt.Println(tree)
}
