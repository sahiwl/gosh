package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		//reading keyboard input
		input, err:= reader.ReadString('\n')
		if err != nil{
			fmt.Fprintln(os.Stderr, err)
		}

		//handle execution of input
		if err = executeInput(input); err !=nil{
			fmt.Fprintln(os.Stderr, err)
		}
	}	
}


func executeInput(input string) error {
	//remove the newline char
	input = strings.TrimSuffix(input, "\n")

	//command will execute
	cmd:= exec.Command(input)

	//set correct output device
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	//execute command and return possible errors
	return cmd.Run()
}

