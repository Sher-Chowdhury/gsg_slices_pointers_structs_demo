package main

import (
	"fmt"
	"reflect"

	"github.com/Sher-Chowdhury/gsg_slices_pointers_structs_demo/avengers"
)

func main() {

	avengers.AddAvenger("IronMan", "Tony Stark")
	avengers.AddAvenger("Spiderman", "Peter Parker")

	memberslist := avengers.GetAvengers()

	fmt.Println(reflect.TypeOf(memberslist)) // prints []*avengers.Avenger
	fmt.Println(memberslist)                 // [0xc000094000 0xc000094030]
	fmt.Println(*memberslist[0])             // {1 IronMan Tony Stark}
	fmt.Println(*memberslist[1])             //{2 Spiderman Peter Parker}
}
