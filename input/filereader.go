package filereader

import (
	"ReadFileProj/ice_cream"
	"encoding/json"
	"io/ioutil"
)

func ReadInput(fn string) (ice_cream.IceCream, error) {
	ic := ice_cream.IceCream{}

	jsonBytes, err :=ioutil.ReadFile(fn)
	if err != nil {
		return ic, err
	}

	err = json.Unmarshal(jsonBytes,&ic)
	if err != nil {
		return ic, err
	}
	return ic, nil
}