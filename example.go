package main

import (
	"fmt"
	"os"
)

func main() {

	//create file pointer
	var infile *streamdata

	//initialize the pointer
	infile = new(streamdata)

	//take the filename from argument
	filename := os.Args[1]

	//URI parameter name to be read
	uriparameter := os.Args[2]

	//pass the file name to the file object created
	infile.openstream(filename)

	//loop parsestream, which returns false on EOF
	for infile.parsestream() {

		//print the value of uri parameter that was passed as second command line argument
		fmt.Println(infile.urlparsefunc(uriparameter))

	}
}
