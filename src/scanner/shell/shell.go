package shell

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"parser"
)

type Shell struct {
	parser *parser.Parser
}

// Parse - Parse shell to command
func Parse() error {
	shell := new(Shell)
	shell.parser = parser.Init()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input (Put empty to finish):")
	var outBuffer bytes.Buffer
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		if line == "\n" {
			fmt.Print("Name,Clone URL,Date Commit,Author\n")
			fmt.Print(outBuffer.String())
			return nil
		}
		outBuffer.WriteString(shell.parser.Map(line))
	}
	return nil
}
