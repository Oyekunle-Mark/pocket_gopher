package main

import (
	"log"
	"os"
	"os/exec"
)

var cmdChain = []*exec.Cmd{
	exec.Command("lib/synonymsapp"),
	exec.Command("lib/sprinkleapp"),
	exec.Command("lib/coolifyapp"),
	exec.Command("lib/domainifyapp"),
	exec.Command("lib/availableapp"),
}

func main() {
	cmdChain[0].Stdin = os.Stdin
	cmdChain[len(cmdChain)-1].Stdout = os.Stdout

	for i := 0; i < len(cmdChain)-1; i++ {
		thisCmd := cmdChain[i]
		nextCmd := cmdChain[i+1]
		stdout, err := thisCmd.StdoutPipe()

		if err != nil {
			log.Fatalln(err)
		}

		nextCmd.Stdin = stdout
	}

	for _, cmd := range cmdChain {
		if err := cmd.Start(); err != nil {
			log.Fatalln(err)
		} else {
			defer cmd.Process.Kill()
		}
	}

	for _, cmd := range cmdChain {
		if err := cmd.Wait(); err != nil {
			log.Fatalln(err)
		}
	}
}
