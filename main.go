package main

import "parth/sqlparser/parse"

// accepts the input commands. calls parseInput

func main() {
	parse.ParseInput("SELECT state FROM table WHERE pop > 100000 AND state != 'California' LIMIT 5;")
}
