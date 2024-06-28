package parser

// takes in raw input from the user. then sends to sub functions to
// parse WHERE, SELECT, LIMIT, etc.

func parseInput(rawInput string) {

	// {rawWhere, rawSelect, rawFrom, rawLimit} = parseArgs(rawInput)

	// whereExp = parseWhere(rawWhere)
	// parseSelect(rawSelect)

}

// returns rawSelect, rawFrom, rawWhere, rawLimit

func parseArgs(rawInput string) {

	//

}

// takes in raw where string. returns a predicate of some kind, that can be plugged into
// the function that actually iterates over the array and runs the test
func parseWhere(rawWhere string) {

}

// given the raw select string returns a list of columns that need to be returned
// (ensure it could be ["*"])
func parseSelect(rawSelect string) []string {

}

func findMatching(whereExp string, table []string) {
	for index, jsonElem := range table {

	}
}

// first step is WHERE function. WHERE query parser

// 1. parseInput(inputString)
// 2. whereParser creates a boolean expression. this might be recursive.
// 3. filter method runs filter on each json element, using the query array. it outputs to
// matchArr. MatchArr basically is passed to
// 4. selectCols (matchArr, cols) which returns only the information in the columns.
// cols is likely going to be a string[], served by selectParser
// if cols is "*" returns the whole {object}[] array. if not it will
// add only the key-value pairs in cols. for each over matchArr (this might be a for loop limited by LIMIT) and then a
// nested loop within that's like cols.foreach add jsonElem[key] to the return. need to think about
// what to do - need to figure out if this gets added
//
// and then return that
//5. might add something to selectCols so that it picks from a table
