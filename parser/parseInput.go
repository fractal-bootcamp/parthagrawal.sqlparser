package parser

// takes in raw input from the user. then sends to sub functions to
// parse WHERE, SELECT, LIMIT, etc.

func parseInput(rawInput string) {

	// {rawWhere, rawSelect, from, limit} = parseArgs(rawInput)

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

// given an input predicate expression, perhaps defined recursively somewhere
// else, and the data table, return matchArr to show all rows in the database
// that match the requirements of where(). this will be processed by
// selectCols

func findMatching(whereExp string, table []string) []string {
	// for index, jsonElem := range table {

	// }
}

// given the number of cols to return and the number of rows to return
func returnSelected(matchArr []string, selectedCols []string, limit int) []string {

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
