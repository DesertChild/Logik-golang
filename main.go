package main

//imports

import (
	"fmt"
	"logic/repl"
	"os"
	"os/user"
)

func main(){
	user,err := user.Current()
	if err != nil{
		panic(err)
	}
	fmt.Printf("Hello %s! welcome to the new world of LOGIK!\n",user.Username)
	fmt.Printf("Feel free to type in the commands\n")
	repl.Start(os.Stdin,os.Stdout)
}