package main

import (
	"fmt"
	"time"
)

/*--------------------------*/

type TMP_BACK struct {
	RESET         string
	BLACK         string
	BLUE          string
	CYAN          string
	GREEN         string
	MAGENTA       string
	RED           string
	WHITE         string
	YELLOW        string
	LIGHT_BLACK   string
	LIGHT_BLUE    string
	LIGHT_CYAN    string
	LIGHT_GREEN   string
	LIGHT_MAGENTA string
	LIGHT_RED     string
	LIGHT_WHITE   string
	LIGHT_YELLOW  string
}

var Back TMP_FORE = TMP_FORE{
	RESET        : "\x1b[49m",
	BLACK        : "\x1b[40m",
	BLUE         : "\x1b[44m",
	CYAN         : "\x1b[46m",
	GREEN        : "\x1b[42m",
	MAGENTA      : "\x1b[45m",
	RED          : "\x1b[41m",
	WHITE        : "\x1b[47m",
	YELLOW       : "\x1b[43m",
	LIGHT_BLACK  : "\x1b[100m",
	LIGHT_BLUE   : "\x1b[104m",
	LIGHT_CYAN   : "\x1b[106m",
	LIGHT_GREEN  : "\x1b[102m",
	LIGHT_MAGENTA: "\x1b[105m",
	LIGHT_RED    : "\x1b[101m",
	LIGHT_WHITE  : "\x1b[107m",
	LIGHT_YELLOW : "\x1b[103m",
}

type TMP_FORE struct {
	RESET         string
	BLACK         string
	BLUE          string
	CYAN          string
	GREEN         string
	MAGENTA       string
	RED           string
	WHITE         string
	YELLOW        string
	LIGHT_BLACK   string
	LIGHT_BLUE    string
	LIGHT_CYAN    string
	LIGHT_GREEN   string
	LIGHT_MAGENTA string
	LIGHT_RED     string
	LIGHT_WHITE   string
	LIGHT_YELLOW  string
}

var Fore TMP_FORE = TMP_FORE{
	RESET        : "\x1b[39m",
	BLACK        : "\x1b[30m",
	BLUE         : "\x1b[34m",
	CYAN         : "\x1b[36m",
	GREEN        : "\x1b[32m",
	MAGENTA      : "\x1b[35m",
	RED          : "\x1b[31m",
	WHITE        : "\x1b[37m",
	YELLOW       : "\x1b[33m",
	LIGHT_BLACK  : "\x1b[90m",
	LIGHT_BLUE   : "\x1b[94m",
	LIGHT_CYAN   : "\x1b[96m",
	LIGHT_GREEN  : "\x1b[92m",
	LIGHT_MAGENTA: "\x1b[95m",
	LIGHT_RED    : "\x1b[91m",
	LIGHT_WHITE  : "\x1b[97m",
	LIGHT_YELLOW : "\x1b[93m",
}

/*--------------------------*/

type ZERO_STRUCT struct {}

var Log ZERO_STRUCT = ZERO_STRUCT{}

func (mc *ZERO_STRUCT) Error(str string) {
    fmt.Printf("["+Fore.RED+"Err"+Fore.RESET+"] ["+Fore.GREEN+string(time.Now().Format("15:04:05"))+Fore.RESET+"] %s", str)
}

func (mc *ZERO_STRUCT) Info(str string) {
    fmt.Printf("["+Fore.GREEN+"Inf"+Fore.RESET+"] ["+Fore.GREEN+string(time.Now().Format("15:04:05"))+Fore.RESET+"] %s", str)
}

func (mc *ZERO_STRUCT) Warning(str string) {
    fmt.Printf("["+Fore.YELLOW+"Wrn"+Fore.RESET+"] ["+Fore.GREEN+string(time.Now().Format("15:04:05"))+Fore.RESET+"] %s", str)
}

func (mc *ZERO_STRUCT) Debug(str string) {
    fmt.Printf("["+Fore.BLUE+"Dbg"+Fore.RESET+"] ["+Fore.GREEN+string(time.Now().Format("15:04:05"))+Fore.RESET+"] %s", str)
}