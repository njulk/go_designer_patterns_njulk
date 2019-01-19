package main

import "fmt"

type Visitor interface {
	getYan() int
	getYong() int
	getSheng() int
}

type Age struct {
}

type Height struct {
}

func (me *Age) getYan() int {
	return 16
}
func (me *Age) getYong() int {
	return 20
}
func (me *Age) getSheng() int {
	return 22
}

func (me *Height) getYan() int {
	return 130
}
func (me *Height) getYong() int {
	return 140
}
func (me *Height) getSheng() int {
	return 150
}

type Element interface {
	Accept(visitor Visitor)
}

type Yan struct {
}
type Yong struct {
}
type Sheng struct {
}

func (me *Yan) Accept(visitor Visitor) {
	fmt.Println(visitor.getYan())
}

func (me *Yong) Accept(visitor Visitor) {
	fmt.Println(visitor.getYong())
}

func (me *Sheng) Accept(visitor Visitor) {
	fmt.Println(visitor.getSheng())
}



func main() {
	height:=new(Height)
	element:=new(Yan)
	element.Accept(height)
	return
}
