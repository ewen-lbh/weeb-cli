package main

import (
	"fmt"
	"os"
	"io/ioutil"
	
	"github.com/docopt/docopt-go"
)

func readFile(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	return string(b)
}

func main() {
	usage := readFile("./README")
	args, _ := docopt.ParseDoc(usage)
	fmt.Println(args)
}
