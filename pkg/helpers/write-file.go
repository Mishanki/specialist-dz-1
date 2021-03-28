package helpers

import "io/ioutil"

func MyWrite(data []byte) {
	err := ioutil.WriteFile("db.json", data, 0644)
	if err != nil {
		panic(err)
	}
}
