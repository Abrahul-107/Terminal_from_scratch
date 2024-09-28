package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		in, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		clearIn := strings.TrimSpace(in) // Use TrimSpace to remove all whitespace including newline
		cmds := strings.Fields(clearIn)   // Split by whitespace and ignore empty strings

		if len(cmds) == 0 { // Handle empty input
			continue
		}

		switch cmds[0] {
		case "exit":
			os.Exit(0)
		case "echo":
			fmt.Println(strings.Join(cmds[1:], " "))
		case "type":
			if len(cmds) < 2 {
				fmt.Println("type: missing argument")
				continue
			}
			switch cmds[1] {
			case "exit", "echo", "type":
				fmt.Printf("%s is a shell builtin\n", cmds[1])
			default:
				paths := strings.Split(os.Getenv("PATH"), string(os.PathListSeparator))
				isFound := false
				for _, path := range paths {
					fullPath := filepath.Join(path, cmds[1])
					if _, err := os.Stat(fullPath); !os.IsNotExist(err) {
						fmt.Printf("%s is %s\n", cmds[1], fullPath)
						isFound = true
						break
					}
				}
				if !isFound {
					fmt.Printf("%s: not found\n", cmds[1])
				}
			}
		default:
			// Check if the command is your shell script
			if cmds[0] == "program_1234" { // Replace with the name of your script
				if len(cmds) < 3 {
					fmt.Println("Usage: program_1234 <name> <code>")
					continue
				}
				command := exec.Command("./shell.sh", cmds[1], cmds[2])
				command.Stderr = os.Stderr
				command.Stdout = os.Stdout
				err := command.Run()
				if err != nil {
					fmt.Printf("Error executing script: %s\n", err)
				}
			} else {
				// Execute other commands
				command := exec.Command(cmds[0], cmds[1:]...)
				command.Stderr = os.Stderr
				command.Stdout = os.Stdout
				err := command.Run()
				if err != nil {
					fmt.Printf("%s: command not found\n", cmds[0])
				}
			}
		}
	}
}
