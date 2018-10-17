package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func splitCommand(cmd string, cmdChan chan []string) {
	cmdArr := strings.Split(cmd, " ")
	args := strings.Join(cmdArr[1:], " ")
	outCmd := []string{cmdArr[0], args}
	cmdChan <- outCmd
}

func executeCMD(cmd []string, outputChan chan []byte) {
	if len(cmd[1]) != 0 {
		cmdO := exec.Command(cmd[0], cmd[1])
		output, err := cmdO.Output()
		if err != nil {
			fmt.Printf("%s failed to execute.\n", cmd)
		}
		outputChan <- output
	} else {
		cmdO := exec.Command(cmd[0])
		output, err := cmdO.Output()
		if err != nil {
			fmt.Printf("%s failed to execute.\n", cmd)
		}
		outputChan <- output
	}
}

func main() {
	cmdChan := make(chan []string)
	outputChan := make(chan []byte)
	if len(os.Args) < 2 {
		fmt.Println("No commands supplied to consh.")
	} else {
		for _, cmd := range os.Args[1:] {
			go splitCommand(cmd, cmdChan)
			go executeCMD(<-cmdChan, outputChan)
			fmt.Println(string(<-outputChan))
		}
	}
}
