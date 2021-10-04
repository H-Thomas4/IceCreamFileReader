package main

import (
	//"ReadFileProj/input"
	//"ReadFileProj/FileReader"
	"ReadFileProj/Repo"
	"ReadFileProj/Validation"
	filereader "ReadFileProj/input"
	"fmt"
	"os"
)

// Import the necessary packages and call the appropriate methods to read a files input and convert it to an IceCream
// You should then validate IceCream if it fails validation print that out.
// If the icecream passes validation, store the IceCream in the db

func main() {
	files := []string{"ic1", "ic2", "ic3", "ic4"}

	for _, v := range files {
		fn := "icjsons/" + v + ".json"

		ic, err := filereader.ReadInput(fn)
		if err != nil {
			fmt.Println(err)

			fmt.Printf("%+v", ic)

			passed, err := Validation.Validate(ic)
			if err != nil {
				fmt.Println(err)
				return
			}

			if passed == false {
				fmt.Printf("Unable to save iceCream, invalid topping.")
				os.Exit(1)
			}

			fmt.Println(passed)

			err = Repo.AddIceCream(ic)
			if err != nil {
				fmt.Println(err)
				return
			}

		}

		_, err = Repo.GetIceCreamByName("Mint Chocolate Chips")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}