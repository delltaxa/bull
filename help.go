package main

import "fmt"

type _tmp_help struct {
	name string
	functionality string
	arguments bool
	doc string
}

var _HELP_COMMANDS []_tmp_help = []_tmp_help {
	_tmp_help{name: "help", functionality: "display this message.", doc: "help {comand}=display all commands+func, help", arguments: true},
}

type _ZERO_STRUCT struct {}
var help _ZERO_STRUCT = _ZERO_STRUCT{}

func (mc *_ZERO_STRUCT) display_all() {
	var longest_name int = 7
	var longest_func int = 13
	for i:=0;i<len(_HELP_COMMANDS);i++ {
		if longest_name < len(_HELP_COMMANDS[i].name) {
			longest_name = len(_HELP_COMMANDS[i].name)
		}

		if longest_func < len(_HELP_COMMANDS[i].functionality) {
			longest_func = len(_HELP_COMMANDS[i].functionality)
		}
	}

	fmt.Println("Command     Functionality")
	fmt.Println("-------     "+get_char("-", longest_func))
	
	for i:=0;i<len(_HELP_COMMANDS);i++ {
		fmt.Printf("%s", _HELP_COMMANDS[i].name+get_char(" ", longest_name - len(_HELP_COMMANDS[i].name)))

		if _HELP_COMMANDS[i].arguments {
			fmt.Printf("\x1b[32m%s\x1b[39m", " [+] ")
		} else {
			fmt.Printf("%s", "     ")
		}

		fmt.Printf("%s", _HELP_COMMANDS[i].functionality+"\n")
	}
	

	fmt.Printf("\n%s\n", "To get more information and usage about a command, use \x1b[33m<help command_name>\x1b[39m")
}

func (mc *_ZERO_STRUCT) help(str string) {
	for i:=0;i<len(_HELP_COMMANDS);i++ {
		if _HELP_COMMANDS[i].name == str {
			fmt.Println(_HELP_COMMANDS[i].doc)
			
			return
		}
	}

	Log.Error("Help: Command not found.")
}


func get_char(c string, l int) string {
	var result string

	for i:=0;i<l;i++ {
			result+=c
	}

	return result
}         