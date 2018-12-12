package tree

import (
	"fmt"
	"io/ioutil"
	"log"
)

func Tree(path string, needFiles bool) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}
