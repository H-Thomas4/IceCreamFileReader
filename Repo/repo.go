package Repo

import (
	"ReadFileProj/ice_cream"
	json "encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

						//returned value
func ConvertDataToGO() (GolangDb, error) {
	////create variable
	data := GolangDb{}
	////first read the file							//absolute path
	jsonBytes, err := ioutil.ReadFile("/home/heather/Downloads/ReadFileProj/icDB.json")
	if err != nil {
		return data, err
	}
	//second
	err = json.Unmarshal(jsonBytes, &data)//all the data and memory use the Unmarshal to grab all as one piece
	if err != nil {
		return data, err
	}
	return data, nil


	//filePtr, err := os.Open("/home/heather/Downloads/ReadFileProj/icDB.json")
	////filePtr implements the reader interface
	//
	//json.NewDecoder(filePtr).Decode(&data)
	//if err != nil {
	//	return data, err
	//}
	//return data, nil
	//

}

func ConvertBackToJson(db GolangDb) error {
	jsonBytes, err := json.Marshal(db)
	if err != nil {
		return err
	}
	//this function over writes the file by default
	err = ioutil.WriteFile("/home/heather/Downloads/ReadFileProj/icDB.json",jsonBytes, 0644)
	return nil
}

// Create a function that accepts an IceCream, converts it to JSON and stores it on the next available line in iceCreamDb.txt
// Return an error if any of the functions you call return an error
func AddIceCream(ice ice_cream.IceCream) error {
	GolangDb, err := ConvertDataToGO()
	if err != nil {
		return err
	}
	//have to use fields to append to fields
	GolangDb.IceCreams = append(GolangDb.IceCreams, ice)
	for _, v := range GolangDb.IceCreams {
		if v.Name == ice.Name {
			//returns string
			statement := fmt.Sprintf("Ice cream already db: %s", ice.Name)
			return errors.New(statement)
		}
	}
	return ConvertBackToJson(GolangDb)
}

// Create function that can retrieve all the ice creams from iceCreamDb.txt
// The function should unmarshal each one into an ice cream struct and return a slice of IceCream --> []IceCream
// Return an error if any of the functions you call return an error
func GetAllIceCreams() ([]ice_cream.IceCream, error) {

	GolangDb, err := ConvertDataToGO()
	if err != nil {
		return nil, err
	}
	return GolangDb.IceCreams, nil
}
//	file, err := os.Open("/home/heather/Downloads/ReadFileProj/icDB.json")
//
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	scanner := bufio.NewScanner(file)
//	scanner.Split(bufio.ScanLines)
//
//	var iceSlice []string
//	var iceCream []ice_cream.IceCream
//
//	ic := ice_cream.IceCream{}
//
//	for scanner.Scan() {
//		iceSlice = append(iceSlice, scanner.Text())
//	}
//
//	for _, val := range iceSlice {
//		_ = json.Unmarshal([]byte(val), &ic)
//		iceCream = append(iceCream, ic)
//
//	}
//	return iceCream, err
//}

func GetIceCreamByName(name string) (ice_cream.IceCream, error) {
	ic := ice_cream.IceCream{}

	GolangDb, err := ConvertDataToGO()
	if err != nil {
		return ic, err
	}

	for _, v := range GolangDb.IceCreams{
		if v.Name == name {
			return v, nil
		}
	}

	return ic, errors.New(fmt.Sprintf("Ice cream not found in db: %s", name))

}

func UpdateIceCreamByName(ic ice_cream.IceCream) error {
	GolangDb, err := ConvertDataToGO()
	if err != nil{
		return err
	}

				// |_| when loops it creates a container and every time you loop around it plugs in something else in container
	for i, v := range GolangDb.IceCreams {
		if v.Name == ic.Name {
			GolangDb.IceCreams[i] = ic
		}
	}
	return ConvertBackToJson(GolangDb)
}



func DeleteIceCreamByName(name string) error {
	GolangDb, err := ConvertDataToGO()
	if err != nil {
		return  err
	}

	 newSlice := []ice_cream.IceCream{}
	for i, v := range GolangDb.IceCreams {
		if v.Name == name {
			itemsBefore := GolangDb.IceCreams[:i]
			itemsAfter := GolangDb.IceCreams[i+1:]
			newSlice = append(itemsBefore, itemsAfter...)
		}
	}

	GolangDb.IceCreams = newSlice
	return ConvertBackToJson(GolangDb)
	}




