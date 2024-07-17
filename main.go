package main

import (
	//"io/ioutil"
	//"log"
	"os"
	"fmt"
	//"net/http"
	//"container/list"
)

var helper = `
plaq helper
help / -h / --help	displays plaq help
install / -i [link]	installs content from mediafire link

			-plaq has magic cat powers :3
`

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

var style = check_arg_valid()

func str_i_slc(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

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
			} else {
				fmt.Println("Invalid option\n")
			}
		} else {
			fmt.Println("No option specified\n")
		}
	}
}

func check_arg_valid() string {
	if len(os.Args) <= 1 {
        	badusage()
    	} else {
		if os.Args[1] == "help" || os.Args[1] == "--help" || os.Args[1] == "-h" {
			return "help"
		} else if os.Args[1] == "install" || os.Args[1] == "-i" || os.Args[1] == "--install"{
			if len(os.Args) == 3 {
				return "install"
			} else {
				badusage()
			}
		} else if os.Args[1] == "meow" || os.Args[1] == "--meow" {
			return "meow"
		} else {
			badusage()
		}
	}
	badusage()
	return "fatal"
}

func check_package_valid(x string) {
	if x == "install" {
		fmt.Println("Mode: " + style)
		fmt.Println("Link: " + os.Args[2])
		askconfirm()
	} else if x == "help" {
		fmt.Println(helper)
	} else if x == "meow"{
		fmt.Println(`
meow!!!
   _____
 _|___ /
(_) |_ \
 _ ___) |
(_)____/
`)
	} else {
		badusage()
	}
}

func badusage() {
        fmt.Println("Bad usage. Read the -h page.")
	os.Exit(1)
}

func main() {
	if str_i_slc("--help", os.Args) || str_i_slc("help", os.Args) || str_i_slc("-h", os.Args) {
		check_package_valid("help")
	} else {
		check_arg_valid()
		/*var style = os.Args[1]
		var link = os.Args[2]*/
		check_package_valid(style)
	}
}
