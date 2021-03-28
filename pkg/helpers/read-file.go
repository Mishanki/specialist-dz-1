package helpers

import "io/ioutil"

func MyReadFile() []byte {
	b, err := ioutil.ReadFile("db.json")
	if err != nil {
		panic(err)
	}

	return b
}
