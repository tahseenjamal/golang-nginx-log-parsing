## Program to use nginx log file

There are two program. One is the package file and the other is example program, which uses the package program

package program: nginxlog.go

test program: sample.go


## Purpose

nginxlog.go can be included as a package and be easily used to process logs easily. The struct and methods manage it all

Package has the struct and method created

Please check the example program on how to use the package


## How to build

Simplest way is to 

go build example.go nginxlog.go


## Sample file

Sample file is just for test purpose - datasample
This file has one line of read nginx log. number and another are two uri parameters

## How to use

./example datasample number

Where ./example is the executable
datasample is the sample file
number is the uri parameter, whose value would be printed
