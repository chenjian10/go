package main

import "fmt"

type USB interface {
    Name() string
    Connecter
}

type Connecter interface {
	Connect()	
}
type PhoneConnecter struct {
	name string	
}

func (pc PhoneConnecter) Name() string {
	 return pc.name
}

func (pc PhoneConnecter) Connect()  {
	fmt.Println("Connect:", pc.name)	
}

func main() {
	a := PhoneConnecter{"PhoneConnecter"}
	a.Connect()
	Disconnet(a)
	fmt.Println(a)
	
	pc := PhoneConnecter{"PhoneConnecter"}
	var b Connecter
	b = Connecter(pc)
	b.Connect()
	
	pc.name = "pc"
	b.Connect()
}

func Disconnet(usb USB)  {
	/*if pc,ok := usb.(PhoneConnecter); ok {
		fmt.Println("Disconnet:", pc.name)
		return
	}*/

	switch v:=usb.(type) {
		case PhoneConnecter:
			fmt.Println("Disconnet:", v.name)
		default:
			fmt.Println("unKnown Disconnet:")
	}

}

