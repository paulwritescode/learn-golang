pnackage main

import "fmt"

type USER struct {
	Name  string
	Email string
	Age   int
}

func main() {
	ron := USER{"Rob", "rob@gmail.com", 67}
	fmt.Printf("%+v\n", ron)

	// the new keyword : it only allocates but it doesnt initialize memory
	/**
	. In golang NEW keyword is used to allocate memory for a value of specified type and returns a pointer to it
	. User creates a pointer to an instance of USER. It allocates memory for the struct but does not initialize it
	**/
	var SAM = new(USER)
	SAM.Name = "Samwel"
	SAM.Email = "mwaniki@tajiri.com"
	SAM.Age = 25
	fmt.Printf("%+v\n", SAM)

	var TOBY = &USER{"Toby", "toby@tajiri.com", 45}
	fmt.Printf("%+v\n", TOBY)

	fmt.Println("Structs baby!")
}

/**
. In structs we use other data-types to create our own data type
**/
