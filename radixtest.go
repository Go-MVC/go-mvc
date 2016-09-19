package main

import (
		"fmt"
		"github.com/bs_radix"		
)


func main() {
	fmt.Println("Starting Main")
	t := bs_radix.New()
	fmt.Println("t is ", t)
	t.Insert(":445L/PG", "445LPG")
	t.Insert(":445L/PG/RegDoc", "445LPG")
	t.Insert(":445L/PG/SampleDoc", "bob")
	t.Insert(":445L/PG/SampleCode", "bobby")
	val, found := t.Get(":445L/PG/SampleCode")
	if found {
		fmt.Println("445L/PG/SampleCode was found. Value: ", val)
	}
	val, found = t.Get(":445L/PG/SampleDoc")
	if found {
		fmt.Println("445L/PG/SampleDoc was found. Value: ", val)
	}
	//look for /Sample. Should return not found
	val, found = t.Get(":445L/PG/Sample")
	if found {
		fmt.Println("445L/PG/Sample was found. Value: ", val)
	}
	//now insert /Sample and check again
	t.Insert(":445L/PG/Sample", "i'm here")
	val, found = t.Get(":445L/PG/Sample")
	if found {
		fmt.Println("445L/PG/Sample was found. Value: ", val)
	}


}