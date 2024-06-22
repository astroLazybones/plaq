package main

import "fmt"

import (
   //"io/ioutil"
   //"log"
   //"net/http"
   //"container/list"
)

var Reset = "\033[0m" 
var Red = "\033[31m" 
var Green = "\033[32m" 
var Yellow = "\033[33m" 
var Blue = "\033[34m" 
var Magenta = "\033[35m" 
var Cyan = "\033[36m" 
var Gray = "\033[37m" 
var White = "\033[97m"

var yes = "y"
var no = "n"

var options = []string{yes, no}

var dw_dir = "./.plaqs"

func main() {
	askconfirm()
}

func askconfirm() {
	for {
		fmt.Println("[" + Green + yes + Reset + "/" + Red + no + Reset + "]")
		var confi string
		fmt.Scanln(&confi)
		if len(confi) != 0 {
			if string(confi[0]) == string(yes[0]) {
				return //true
			} else if string(confi[0]) == string(no[0]) {
				return //false
			} else{
				fmt.Println("Invalid option")
			}
		} else {
			fmt.Println("No option specified")
		}
	}
}
