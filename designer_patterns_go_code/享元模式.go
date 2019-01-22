package main

import "fmt"

type user struct {
	name string
}

type website interface {
	getName(u user)
}

type Awebsite struct {
	web string
}

type Bwebsite struct {
	web string
}

type webFactory struct {
	data map[string]website
}

func(me *webFactory)getWebsite(name string)website{
	if _,ok:=me.data[name];ok{
		return me.data[name]
	}else{
		switch name {
		case "tian":
			return new(Awebsite)
		case "di":
			return new(Bwebsite)
		}
	}
	return new(Bwebsite)
}



func (me *Awebsite) getName(u user) {
	fmt.Println("A website ", me.web, ":", u.name)
}

func (me *Bwebsite) getName(u user) {
	fmt.Println("B website ", me.web, ":", u.name)
}



func main() {
	use:=user{name:"llla"}
	use1:=user{name:"111"}
	webfac:=new(webFactory)
	web:=webfac.getWebsite("tian")
	web.getName(use)
	web.getName(use1)
	return
}
