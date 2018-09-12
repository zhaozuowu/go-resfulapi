package main

import "fmt"

type Notify interface {

	Send(str string)
}

type EmailNotify struct {
	str string
}

type Element interface {

}
type PhoneNoTity struct {
	str string
}

func (email EmailNotify)  Send(str string) {

	fmt.Printf("send message:%s by email",email.str)
	fmt.Println()

}

func (phone PhoneNoTity) Send()  {

	fmt.Printf("send message:%s by phone",phone.str)
}

type Human interface {
	SayHi(str string)
	GetListById(id int) []Element
}

type Person interface {
	Human
	Work()
}

func (email EmailNotify) SayHi(str string)  {

}

func (email EmailNotify) GetListById(id int) []Element  {

	 list := make([]Element,3)

	 list[0],list[1] = 10,"test"

	 return list
}
func (email EmailNotify) Work()  {

}
func main()  {

	emailNotify :=  EmailNotify{"gaofandi1236677@12366.com"}
	var notify Notify
	notify = emailNotify

	var person Person

	notify.Send("s")


	person = emailNotify

	person.GetListById(10)

}