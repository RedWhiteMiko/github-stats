# Github stats

Github stats output to CSV text format to stdout

## Instruction
- Requirements: `Docker or Golang (with libraries)`
  - Golang libraries:
    - "github.com/google/go-github/github"
	- "golang.org/x/oauth2"
- Execute Build: `make build` (Golang required)
- Execute Build & Run: `make run` (Golang required)
- Docker Build: `make build-docker` (Docker required)
- Docker Build & Run: `make run-docker` (Docker required)

### Notes
- Executable binary parameter: `./gh_scrap [input_file]`
- If input files not specified, it will read from stdin
- In case of Docker build, it will set to read from `input.txt` in root directory, I'm not sure how to set `docker run` to use interactive input.

## Structure
- `src/def`: contain classes definition
  - `src/def/repo`: Git repository object
- `src/scanner`: Scan string to parser from file or stdin
    - `src/scanner/file`: File scanner
    - `src/scanner/shell`: Shell stdin scanner
- `src/parser`: Process scanned string to output

## Assumption
- Latest commit is the latest commit in default branch, not latest commit from all branches
- Repository name is `owner/repo_name`, otherwise it will output `Not valid Github Repository`
- `GITHUB_TOKEN` is set to empty by default, if your API quota run out, please update the Token in `Dockerfile` to your token