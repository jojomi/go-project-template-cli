# go-project-template

## How to Use

1. Clone this repo: `git clone https://github.com/jojomi/go-project-template.git [project name]`
2. Change to project directory: `cd [project name]`
3. Set the project name in `build_info.go`[build_info.go] (variable `ToolName`)
4. Execute `go mod init github.com/[your name]/[project name] && go mod tidy` to initialize a Go module (change URL as you see fit)
5. Remove `go.mod` and `go.sum` lines from `.gitignore` so that they become part of your repo
6. If you did not use this as a template in Github: Change git remote using `git remote set-url origin https://github.com/[your name]/[project name]` to abandon the mothership
7. Replace this file (`README.md`) with more useful content for your project
8. Sail away!

Good luck with your project and never forget where you came from ;).