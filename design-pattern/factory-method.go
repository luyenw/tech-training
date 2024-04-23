package main

import (
	"errors"
	"fmt"
)

type IBank interface {
	toString() string
}
type VPBank struct {
	name string
}
type VietcomBank struct {
	name string
}

func (bank *VPBank) toString() string {
	return bank.name
}
func (bank *VietcomBank) toString() string {
	return bank.name
}

type BankFactory struct {
}

func (*BankFactory) getBank(name string) (error, IBank) {
	switch name {
	case "vpbank":
		return nil, &VPBank{name: "VPBank"}
	case "vcbank":
		return nil, &VietcomBank{name: "VietcomBank"}
	}
	return errors.New("This bank type is unsupported"), nil
}

func main() {
	factory := BankFactory{}
	fmt.Println(factory.getBank("vpbank"))
	fmt.Println(factory.getBank("vcbank"))
	fmt.Println(factory.getBank("vtbank"))
}
