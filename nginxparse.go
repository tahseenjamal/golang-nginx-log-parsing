package main

import (
	"bufio"
	"net/url"
	"os"
	"regexp"
	"time"
)

//nginx data fields as per the Split Function created by me

type datagram struct {
	ipaddress string

	datetime_stamp time.Time

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

//Here lines are read and mapped to datagram struct
func (instream *streamdata) parsestream() bool {

	if instream.scanner.Scan() {

		logLine := instream.scanner.Text()

		ipaddress_reg := regexp.MustCompile(`([0-9]{1,3}\.){3}[0-9]{1,3}`)
		ipaddress_data := ipaddress_reg.FindAllString(logLine, -1)[0]

		datetime_reg := regexp.MustCompile(`\[\d{1,2}\/\w{3}\/\d{1,4}(:[0-9]{1,2}){3} \+([0-9]){1,4}\]`)
		datetime_data := datetime_reg.FindAllString(logLine, -1)[0]
		parsed_time, _ := time.Parse("[02/Jan/2006:15:04:05 -0700]", datetime_data)

		url_reg := regexp.MustCompile(`"\w+\s[^\s]+`)
		url_data := url_reg.FindAllString(logLine, -1)[0]

		request_type_re := regexp.MustCompile(`"\w+`)
		request_data := request_type_re.FindAllString(logLine, -1)[0]

		http_version_reg := regexp.MustCompile(`HTTP\/\d.\d"`)
		http_version_data := http_version_reg.FindAllString(logLine, -1)[0]

		response_and_byte_reg := regexp.MustCompile(`([0-9]{1,3}) \d+`)
		response_and_byte_data := response_and_byte_reg.FindAllString(logLine, -1)[0]

		instream.datagram = datagram{ipaddress_data, parsed_time, request_data, url_data, http_version_data, response_and_byte_data}

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
