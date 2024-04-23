package main

import (
	"fmt"
)

type IChair interface {
	chairToString() string
}

type PlasticChair struct {
}

type WoodChair struct {
}

func (chair *PlasticChair) chairToString() string {
	return "plastic chair"
}
func (chair *WoodChair) chairToString() string {
	return "wood chair"
}

type ITable interface {
	tableToString() string
}

type PlasticTable struct {
}

type WoodTable struct {
}

func (table *PlasticTable) tableToString() string {
	return "plastic table"
}
func (table *WoodTable) tableToString() string {
	return "wood table"
}

type FurnitureAbstractFactory interface {
	createTable() ITable
	createChair() IChair
}
type PlasticFactory struct{}
type WoodFactory struct{}

func (*PlasticFactory) createTable() ITable {
	return &PlasticTable{}
}
func (*PlasticFactory) createChair() IChair {
	return &PlasticChair{}
}
func (*WoodFactory) createTable() ITable {
	return &WoodTable{}
}
func (*WoodFactory) createChair() IChair {
	return &WoodChair{}
}

type FurnitureFactory struct {
}
type MaterialType int

const (
	PLASTIC = iota
	WOOD
)

func (*FurnitureFactory) getFactory(materialType MaterialType) FurnitureAbstractFactory {
	if materialType == PLASTIC {
		return &PlasticFactory{}
	} else if materialType == WOOD {
		return &WoodFactory{}
	}
	fmt.Println("This furniture is unsupported")
	return nil
}

func main() {
	factory := FurnitureFactory{}
	abstractFactory := factory.getFactory(WOOD)
	fmt.Println(abstractFactory.createTable().tableToString())
	fmt.Println(abstractFactory.createChair().chairToString())

	abstractFactory = factory.getFactory(PLASTIC)
	fmt.Println(abstractFactory.createTable().tableToString())
	fmt.Println(abstractFactory.createChair().chairToString())
}
