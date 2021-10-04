package FileReader

import (
	"ReadFileProj/ice_cream"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//Use the appropriate functions from the os package, JSON package, ioutils package etc. To read the input
// from a file, unmarshal it into an IceCream struct and return the struct
// Return an error if any of your functions that you call return an error
func ReadFile(fileName string) (ice_cream.IceCream, error) {
	file, err := ioutil.ReadFile(fileName)

	data := ice_cream.IceCream{}

	err = json.Unmarshal([]byte(file), &data)

	fmt.Println(data)
	return data, err
}
