package main

import (
	"fmt"
	"os"
)

var continue_shell bool = true

func makeshell() {
	for {
		fmt.Printf("%s", "\n"+prompt)

		var user_input string = input()

		if len(user_input) < 1 {
			continue
		}

		var parsed_user_input []string = CmdParse(user_input)

		var command_name = parsed_user_input[0]

		switch command_name {
			case("sessions"):
				display_sessions(sessions)
			case("exec"):
				exec_command(parsed_user_input)
			case("shell"):
				shell_command(parsed_user_input)
			case("generate"):
				generate_command(parsed_user_input)
			case("alias"):
				alias_command(parsed_user_input)
			case("reset"):
				reset_command()
			case("server"):
				serverinfo_command()
			case("clear"):
				clear_command()
			case("help"):
				help_command(parsed_user_input)
			case("exit"):
				os.Exit(0)
			default:
				Log.Error("Command not found.")
		}
	}
}