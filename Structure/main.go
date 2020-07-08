package main

import "fmt"

//DS- Collection of properties that are related together
type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	// shweta := person{ "Shweta" , "Shrinivas"}
	var shweta person
	fmt.Println(shweta) //Output: {  { 0}}  //Nil values

	shweta.firstName = "Shweta"
	shweta.lastName = "Shrinivas"
	fmt.Println(shweta) //Output:  {Shweta Shrinivas { 0}}

	smita := person{
		firstName: "Smita",
		lastName:  "Shrinivas",
		contactInfo: contactInfo{
			email:   "xyz@gmail.com",
			zipCode: 99999,
		},
	}

	smita.updatedName("Smitu") //Doesn't change to smitu . Hence use pointer
	smita.print()

	//& gives the memory address of the value this variable is pointing at
	//* gives the value this memory address is pointing at

	/*shwetaPointer := &shweta  //Go shortcut fot not having to write & is adding Pointer to the variablename
	shwetaPointer.updateName("Shwetu")
	shwetaPointer.print()*/

	shweta.updateName("Shweti")
	shweta.print()
}

//Function to update name without pointer - doesnt work
func (p person) updatedName(newFirstName string) {
	p.firstName = newFirstName
}

//Go is a pass by value language

//*person - a pointer that points at the person (type description)
//*pointerToPerson - an operator. it means we want to maniplulate the value the point is referencing

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
