package main

var prompt string = "bull > "

type session struct {
	SessionId string
	IPAddress string
	Username string
	Hostname string
	OperatingSystem string
	Job string
	Alias string ///////
	Seen int64
}

var selected_session string
var sessions []session = []session{}
var generated_session_ids []string


var help_text string = Fore["GREEN"]+`Command                     Function`+Fore["RESET"]+`
-----------------------     ------------------
help                        shows this message
exec `+Fore["RED"]+`<session> <command>`+Fore["RESET"]+`    executes a command against a target
shell `+Fore["RED"]+`<session>`+Fore["RESET"]+`             get's and shell for an sesssion
generate `+Fore["RED"]+`?os ?lhost ?lport`+Fore["RESET"]+`  generate a payload
sessions                    display all available sessions
alias `+Fore["RED"]+`<session> <alias>`+Fore["RESET"]+`     alias a session
clear                       clear the terminal
reset                       reset the session list [warning: this will reset all aliases]
server                      display server info`+"\n"