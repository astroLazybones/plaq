package main

import (
	//"io/ioutil"
	"log"
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
	if len(os.Args) != 3 {
        	log.Fatal("Specify type (e.g -S for download) and the download link, nothing more/less!")
    	} else {
		return
	}
}

func main() {
	check_arg_valid()
	var style = os.Args[1]
	var link = os.Args[2]
	fmt.Println(style, link)
	fmt.Println(askconfirm())
}


