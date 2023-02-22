package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	animate("Starting the BUll-Framework")
	init_help()

	rand.Seed(time.Now().UnixNano())

	if len(os.Args) > 1 {
		Server_Info["engine_server"] = os.Args[1]
	}

	fmt.Printf("%s", logo+"\n")

	Log.Info("BUll is on version "+Fore.YELLOW+Server_Info["version"]+Fore.RESET+"\n")
	Log.Info("Starting BULL-ENGINE on "+Fore.YELLOW+Server_Info["engine_server"]+Fore.RESET+"\n\n")

	Log.Warning("Reverse-shells are not fully interactive.\n")

	http.HandleFunc("/", Handle)

    go func() {
        if err := http.ListenAndServe(":8080", nil); err != nil {
            Log.Error(err.Error())
        }
    }()

	makeshell()
}