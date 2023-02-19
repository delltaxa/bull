package main

var logo string = `
BUll Framework`+Fore.BLUE+`     ^     ^
 ____  _   _ _ _  ( \___/ )
| __ )| | | | | |  [ `+Fore.RED+`0 0`+Fore.BLUE+` ]
|  _ \| | | | | |   \   /
| |_) | |_| | | |   ( 0 )
|____/ \___/|_|_|    ---`+Fore.RESET+"\n"

func init_help() {
	_HELP_COMMANDS = append(_HELP_COMMANDS, _tmp_help{arguments: true, name: "exec", functionality: "Execute a command against a target.", doc: "Usage: exec <session> <command>\nThis Command will Execute a specified command on a Target machine."})
	_HELP_COMMANDS = append(_HELP_COMMANDS, _tmp_help{arguments: true, name: "shell", functionality: "Get a shell to a target.", doc: "Usage: shell <session>\nThis Command will spawn a interactive shell to a session."})
	_HELP_COMMANDS = append(_HELP_COMMANDS, _tmp_help{arguments: true, name: "generate", functionality: "Generate backdoor/payload.", doc: "Usage: generate arg:os,lhost,lport>\nThis Command will generate a backdoor payload.\n\nWindows Example: generate os=windows lhost=127.0.0.1 lport=8080\nLinux Example: generate os=linux lhost=127.0.0.1 lport=8080"})
	_HELP_COMMANDS = append(_HELP_COMMANDS, _tmp_help{arguments: false, name: "sessions", functionality: "Display all sessions.", doc: "Usage: sessions\nThis Command will display all sessions."})
	_HELP_COMMANDS = append(_HELP_COMMANDS, _tmp_help{arguments: true, name: "alias", functionality: "Alias session.", doc: "Usage: alias <session_id> <alias>\nThis Command will alias a session with a name."})
	_HELP_COMMANDS = append(_HELP_COMMANDS, _tmp_help{arguments: false, name: "reset", functionality: "Reset session list.", doc: "Usage: reset\nThis Command will Reset the session list."})
	_HELP_COMMANDS = append(_HELP_COMMANDS, _tmp_help{arguments: false, name: "server", functionality: "Shows info about server.", doc: "Usage: server\nThis Command will Show basic info about the server."})
	_HELP_COMMANDS = append(_HELP_COMMANDS, _tmp_help{arguments: false, name: "clear", functionality: "Clears the Terminal.", doc: "Usage: clear\nThis Command will simply clear the Terminal."})
	_HELP_COMMANDS = append(_HELP_COMMANDS, _tmp_help{arguments: false, name: "exit", functionality: "Exits bull.", doc: "Usage: exit\nThis Command will simply exit the bull-framework."})
}