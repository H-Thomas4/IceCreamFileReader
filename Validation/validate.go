package Validation

import (
	"ReadFileProj/ice_cream"
	"encoding/json"
	"io/ioutil"
)

// Write a function that takes in an ice cream then reads the file toppings.json.
// Iterate through the toppings returned from toppings.json and iterate throught the toppings in the IceCream
// to verify that the IceCream ony contains toppings listed in toppings.json.
// If the ice cream fails validation return false, otherwise return true.
func Validate(ic ice_cream.IceCream) (bool, error) {
	data, err := ioutil.ReadFile("/home/heather/Downloads/ReadFileProj/input/toppings.json")

	top := struct {
		Toppings []string
	}{}

	err = json.Unmarshal([]byte(data), &top)
	if err != nil {
		return false, err
	}

	for _, v := range ic.Toppings {
		if !sliceContains(top.Toppings, v) {
			return false, nil
		}
	}
	return true, nil

}

func sliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}

	}
	return false
}
