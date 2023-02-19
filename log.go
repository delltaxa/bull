package main

import (
	"fmt"
)

func Log_Error(m string) {
	fmt.Printf("["+Fore.RED+"Error"+Fore.RESET+"] "+m)
}

func Log_Info(m string) {
	fmt.Printf("["+Fore.GREEN+"Info"+Fore.RESET+"] "+m)
}

func Log_Warning(m string) {
	fmt.Printf("["+Fore.YELLOW+"Warning"+Fore.RESET+"] "+m)
}