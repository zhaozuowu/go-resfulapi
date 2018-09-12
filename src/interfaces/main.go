package main

import "fmt"

type Action interface {
	Poping()
}
type Human interface {
	Action
	Say()
	Eat(str string)
}

type Args interface {
}
type Person struct {
	Name string
	Age  int
	Sex  string
}

func (person Person) Say() {

	fmt.Printf("say:%s\n", person.Name)
}

func (person Person) callback(args ...Args) {

	for _, item := range args {

		//if value, ok := item.(int); ok {
		//
		//	fmt.Printf("type is:int,value:%d\n", value)
		//} else if value,ok := item.(string);ok {
		//	 	fmt.Printf("type is:string,value:%s\n", value)
		//}else if value,ok := item.(Person); ok {
		//	 	fmt.Printf("type is:struct,value:%#v\n", value)
		//}

		switch value := item.(type) {
			case int:
				fmt.Printf("type is:int,value:%d\n", value)
			break
			case string:
				fmt.Printf("type is:int,value:%s\n", value)
			break
			case  Person:
				fmt.Printf("type is:struct,value:%#v\n", value)
			break
		}

	}

}
func (person Person) Eat(str string) {

	fmt.Printf("eat:%s\n", str)

}
func (person Person) Poping() {

	fmt.Printf("poping\n")
}
func main() {

	var human Human
	person := Person{Name: "赵作武1", Age: 28, Sex: "男"}
	person1 := Person{Name: "赵作武2", Age: 28, Sex: "男"}
	person2 := Person{Name: "赵作武3", Age: 28, Sex: "男"}

	human = person

	human.Say()

	list := make([]Human, 3)
	list[0], list[1], list[2] = person, person1, person2

	for _, item := range list {

		item.Say()
	}

	x, y, z := 1, "hello", human

	person.callback(x, y, z)
	person.Poping()

}
