package main

import "fmt"

type Request struct {
	requestType string
	num         int
}

type Manager interface {
	SetSuperior(manager Manager)
	RequestApplication(request *Request)
}

type CommonManager struct {
	superior Manager
	name     string
}

type Majordomo struct {
	superior Manager
	name     string
}

type GeneralManager struct {
	superior Manager
	name     string
}

func (me *CommonManager) SetSuperior(manager Manager) {
	me.superior = manager
}

func (me *CommonManager) RequestApplication(request *Request) {
	if request.requestType == "请假" && request.num < 3 {
		fmt.Println(me.name, ":", request.requestType, ",", request.num)
	} else {
		if me.superior != nil {
			me.superior.RequestApplication(request)
		} else {
			fmt.Println("权力之外的申请，禁止申请")
		}
	}
}

func (me *Majordomo) SetSuperior(manager Manager) {
	me.superior = manager
}
func (me *Majordomo) RequestApplication(request *Request) {
	if request.requestType == "请假" {
		fmt.Println(me.name, ":", request.requestType, ",", request.num)
	} else if request.requestType == "涨薪" {
		if request.num < 1000 {
			fmt.Println(me.name, ":", request.requestType, ",", request.num)
		} else {
			if me.superior != nil {
				me.superior.RequestApplication(request)
			} else {
				fmt.Println("权力之外的申请，禁止申请")
			}
		}
	}
}

func (me *GeneralManager) SetSuperior(manager Manager) {
	me.superior = manager
}
func (me *GeneralManager) RequestApplication(request *Request) {
	if request.requestType == "请假" && request.num < 10 {
		fmt.Println(me.name, ":", request.requestType, ",", request.num)
	} else if request.requestType == "涨薪" && request.num < 5000 {
		fmt.Println(me.name, ":", request.requestType, ",", request.num)
	} else {
		fmt.Println("权力之外的申请，禁止申请")
	}
}

func main() {
	commonManager := &CommonManager{
		name: "manager",
	}
	major := &Majordomo{
		name: "major",
	}
	general := &GeneralManager{
		name: "general",
	}
	commonManager.SetSuperior(major)
	major.SetSuperior(general)

	req:=&Request{
		requestType:"涨薪",
		num:3000,
	}

	commonManager.RequestApplication(req)

}
