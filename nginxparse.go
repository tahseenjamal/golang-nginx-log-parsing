package main

import (
	"bufio"
	"net/url"
	"os"
	"strings"
)

//nginx data fields as per the Split Function created by me

type datagram struct {
	ipaddress string
	datestamp string

	hour     string
	minute   string
	second   string
	timezone string

	reqeuststype string
	requesturi   string
	httpversion  string
	response     string
}

//struct for the object

type streamdata struct {
	datagram

	filehandle *os.File
	fileerr    error
	scanner    *bufio.Scanner
}

//Reading the file
func (instream *streamdata) openstream(filename string) {

	instream.filehandle, instream.fileerr = os.Open(filename)

	instream.scanner = bufio.NewScanner(instream.filehandle)

}

//This function is used by strings->FieldFunc to identify multiple delimiters, to split the line basis that
func (instream *streamdata) Split(r rune) bool {

	return r == '[' || r == ']' || r == ' ' || r == ':'

}

//Here lines are read and mapped to datagram struct
func (instream *streamdata) parsestream() bool {

	if instream.scanner.Scan() {

		data := strings.FieldsFunc(instream.scanner.Text(), instream.Split)

		instream.datagram = datagram{data[0], data[3], data[4], data[5], data[6], data[7], data[8], data[9], data[10], data[11]}

		//information read so return true
		return true

	} else {

		//return false if end of file and no more lines
		return false

	}

}

//this function is used to extract GET Parameters from the URL. parameter is the name of the Get parameter in the UR, whose value we want to extract
func (instream *streamdata) urlparsefunc(parameter string) string {

	//current line in buffer is parsed
	uriparsed, err := url.Parse(instream.requesturi)

	//if no error then go inside
	if err == nil {

		//extract map of parameters and values in keyvalue from the request URI extracted from the nginx log line currently in buffer
		keyvalue, err := url.ParseQuery(uriparsed.RawQuery)

		//if no error then go inside
		if err == nil {

			// data would hold value as string of the name parameter. Note ok is true is the map has parameter key
			data, ok := keyvalue[parameter]

			if ok == true {

				//This is the data value of the parameter of URI
				return data[0]

			} else {

				// return empty string in case of missing paramer. Although, I will change this function to (string, bool) and return false in case of faiure
				// instead of empty string as empty string can also mean that the parameter has no value. Empty string doesn't tell that the parameter is not present
				return ""

			}

		}

	}

	return ""

}

//close the file handle
func (instream *streamdata) closestream() {

	instream.filehandle.Close()

}
