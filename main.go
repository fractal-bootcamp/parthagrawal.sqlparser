package main

import "parth/sqlparser/parse"

// accepts the input commands. calls parseInput

func main() {
	parse.ParseInput("SELECT gang, ayo FROM hello WHERE etc LIMIT 5;")
}
