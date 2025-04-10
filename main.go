package main

import (
	"bufio"
	"errors"
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

	//split input to seperate command and arguements
	args:= strings.Split(input, " ")
	
	// check for built-in commands
	switch args[0]{
	case "cd":
		// 'cd' to home dir with empty path not yet supported
		if len(args)>2{
			return errors.New("path required")
		}
		
		//change directory and return possible errors
		return os.Chdir(args [1])
		
	case "exit":
		os.Exit(0)
	}
	
	//prepare the cmd to execute
	cmd:= exec.Command(args[0], args[1:]...)

	//set correct output device
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	
	//execute command and return possible errors
	return cmd.Run()
}

