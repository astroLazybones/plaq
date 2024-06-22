package main

import (
	//"io/ioutil"
	//"log"
	"os"
	"fmt"
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

var style = ""
var link = ""

func askconfirm() bool {
	for {
		fmt.Print("Would you like to download this file? [" + Green + yes + Reset + "/" + Red + no + Reset + "]: ")
		var confi string
		fmt.Scanln(&confi)
		if len(confi) != 0 {
			if string(confi[0]) == string(yes[0]) {
				return true
			} else if string(confi[0]) == string(no[0]) {
				return false
			} else{
				fmt.Println("Invalid option\n")
			}
		} else {
			fmt.Println("No option specified\n")
		}
	}
}

func check_arg_valid() {
	if len(os.Args) <= 1 {
        	fmt.Println("Invalid usage. Read -h page.")
		os.Exit(1)
    	} else {
		if os.Args[1] == "-h" {
			return
		}
	}
}

func main() {
	check_arg_valid()
	//var style = os.Args[1]
	//var link = os.Args[2]
	if style != "install" {
		fmt.Println("Mode: " + style)
		fmt.Println("Link: " + link)
		fmt.Println(askconfirm())
	} else {
		fmt.Println("Invalid usage. Read -h page.")
		os.Exit(1)
	}
}
