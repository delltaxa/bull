package main

import (
	"regexp"
	"strings"
)

func CmdParse(cmdLine string) []string {
	re := regexp.MustCompile(`"([^"\\]*(?:\\.[^"\\]*)*)"|'([^'\\]*(?:\\.[^'\\]*)*)'|(\S+)`)

	args := re.FindAllString(cmdLine, -1)

	for i:=0;i<len(args);i++ {
		args[i] = strings.Trim(regexp.MustCompile(`\\(.)`).ReplaceAllString(args[i], "$1"), " ") 

		if strings.HasPrefix(args[i], "\"") && strings.HasSuffix(args[i], "\"") ||
			strings.HasPrefix(args[i], "'") && strings.HasSuffix(args[i], "'") {
			args[i] = args[i][1:len(args[i]) - 1]
		}
	}

	return args
}