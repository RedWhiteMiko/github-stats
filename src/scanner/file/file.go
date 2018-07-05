package file

import (
	"bufio"
	"fmt"
	"os"
	"parser"
)

type File struct {
	parser *parser.Parser
}

// Parse - Parse file to command
func Parse(fileName string) error {
	file := new(File)
	file.parser = parser.Init()

	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	fmt.Print("Name,Clone URL,Date Commit,Author\n")
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Print(file.parser.Map(line))
	}
	return nil
}
