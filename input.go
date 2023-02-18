package main

import (
	"bufio"
	"os"
)

func input() string {
    reader := bufio.NewReader(os.Stdin)
    
    userinput, err := reader.ReadString('\n')

    if err != nil {
            return ""
    }

    return userinput[0:len(userinput) - 1]
}       