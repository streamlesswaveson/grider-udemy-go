package main

import "fmt"

type person struct {
	firstName string
	lastName string
	contact contactinfo
}

type contactinfo struct {
	email string
	zip int
}

func main() {

	jim := person{
		firstName:"Jim",
		lastName:"Shandy",
		contact: contactinfo{
			email:"a@b.com",
			zip: 12345,
		},
	}

	jim.print()

	// legal way to do it
	jimPointer := &jim
	fmt.Println(jimPointer)

	(*jimPointer).updateFirstname("Howard")

	jim.print()

	// another way - less code
	jim.updateFirstname("Steve-o")
	jim.print()

}

func (p person) print()  {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateFirstname(name string){
	p.firstName = name
}

func displayProperties(){
	var aidan person
	aidan.firstName = "Aidan"
	aidan.lastName = "Blair"

	fmt.Println(aidan)
	fmt.Printf("%+v\n", aidan)


}

func explicitDeclaration(){
	iain := person{firstName:"Iain", lastName:"Blair"}
	fmt.Println(iain)
}

func implicitDeclaration() {
	//iain := person{"Iain", "Blair", }
	//fmt.Println(iain)

}