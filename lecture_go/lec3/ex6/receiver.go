package main

import "fmt"

type Persons struct {
	Name string
	Age  int
	Pnum string
}

func (p Persons) Alice() {
	p.Name = "alice"
	p.Age = 7
	p.Pnum = "01012345678"
}

func (p *Persons) Bob() {
	p.Name = "bob"
	p.Age = 5
	p.Pnum = "01098765432"
}

func main() {
	p := &Persons{"codz", 8, "0327894561"}
	fmt.Println(p)
	p.Alice()
	fmt.Println(p)
	p.Bob()
	fmt.Println(p)

}
