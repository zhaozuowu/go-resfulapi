package main

import "fmt"

const (
	WHITE  = iota
	BLACK
	RED
	BLUE
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box

func (box Box) Volume() float64 {

	return box.width * box.height * box.depth
}

type Human struct {
	name   string
	age    int
	height float64
}

type Person struct {
	Human
	speical string
}

func (human *Human) SayHi() {

	fmt.Printf("name is :%s,age:%d,height:%f\n", human.name, human.age, human.height)
}

func (person *Person) SayHi()  {

	fmt.Printf("sname is :%s,sage:%d,sheight:%f\n", person.name, person.age, person  .height)
}
func (boxlist BoxList) BiggestColor() Color {

	result := 0.00

	color := Color(WHITE)

	for _, box := range boxlist {

		if volume := box.Volume(); volume > result {

			result = volume
			color = box.color
		}
	}

	return color
}

func (color Color) String() string {

	colorList := []string{"WHITE", "BLACK", "RED", "BLUE", "YELLOW"}

	return colorList[color]
}

func (box *Box) SetColor(color Color) {

	box.color = color
}

func (boxlist BoxList) ChangeBoxListColor() {

	for index := range boxlist {

		boxlist[index].SetColor(BLACK)
	}
}

func main() {

	boxlist := BoxList{
		Box{width: 6, height: 7, depth: 8, color: Color(WHITE)},
		Box{width: 7, height: 7, depth: 8, color: Color(BLACK)},
		Box{width: 8, height: 7, depth: 8, color: Color(RED)},
		Box{width: 9, height: 7, depth: 8, color: Color(BLUE)},
		Box{width: 10, height: 7, depth: 8, color: Color(YELLOW)},
	}

	fmt.Printf("thre first volume is:%f	\n", boxlist[0].Volume())

	fmt.Printf("the biggestVolume color is :%d\n", boxlist.BiggestColor())
	fmt.Printf("the biggestVolume color is :%s\n", boxlist[len(boxlist)-1].color.String())
	fmt.Printf("the biggestVolume color is :%s\n", boxlist.BiggestColor().String())

	boxlist.ChangeBoxListColor()
	fmt.Printf("the biggestVolume color is :%s\n", boxlist[0].color.String())


	human := Person{Human{"zhangsan",20,1.80},"TESHU"}
	human.SayHi()

}
