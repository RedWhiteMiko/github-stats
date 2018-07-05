package parser

import (
	"def/repo"
	"fmt"
	"strings"
)

type Parser struct {
	git *repo.Github
}

func Init() *Parser {
	parser := new(Parser)
	return parser
}

// Map - Dumb command mapper.
// TODO: Probably change it to use reflection (not sure how to do it in Go)
func (p *Parser) Map(input string) string {
	input = strings.Trim(input, " \n\t")

	name := strings.Split(input, "/")

	if len(name) < 2 {
		return "Not valid Github Repository\n"
	}
	p.git = repo.Init(name[0], name[1])

	lastCommit, err := p.git.GetLastCommit()
	if err != nil {
		return fmt.Sprintf("Last commit not found, %v\n", err)
	}
	cloneURL, err := p.git.GetCloneURL()
	if err != nil {
		return fmt.Sprintf("Clone URL not found, %v\n", err)
	}

	return fmt.Sprintf("%s,%s,%s,%s\n",
		input, cloneURL, lastCommit.Commit.Author.Date, *lastCommit.Commit.Author.Name)
}
