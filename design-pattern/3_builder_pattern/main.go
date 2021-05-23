package main

import "fmt"

// https://refactoringguru.cn/design-patterns/builder/go/example

/*
分为多个步骤进行构建的潜在问题是， 构建不完整的和不稳定的产品可能会被暴露给客户端。
生成器模式能够在产品完成构建之前使其处于私密状态。
*/
type iBuilder interface {
	setWindowType()
	setDoorType()
	SetNumFloor()
	getHouse() house
}

func getBuilder(builderType string) iBuilder {
	switch builderType {
	case "normal":
		return &normalBuilder{}
	// case "igloo":
	// 	return &iglooBuilder{}
	default:
		return nil
	}
}

type normalBuilder struct {
	house
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (n *normalBuilder) setWindowType() { n.windowType = "normal" }
func (n *normalBuilder) setDoorType()   { n.doorType = "normal" }
func (n *normalBuilder) SetNumFloor()   { n.floor = 2 }
func (n *normalBuilder) getHouse() house {
	// 最后从这里返回的总是产品
	return house{
		windowType: n.windowType,
		doorType:   n.doorType,
		floor:      n.floor,
	}
}

type house struct {
	windowType string
	doorType   string
	floor      int
}

// Director ·这层封装所需要的就是交付的完整性
type Director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) buildHouse() house {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.SetNumFloor()

	return d.builder.getHouse()
}

func main() {
	normal_builder := getBuilder("normal")

	director := newDirector(normal_builder)

	normal_house := director.buildHouse()

	fmt.Println(normal_house)
}
