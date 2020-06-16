package main

import ("fmt")

// Return the full chain of a number (including number)
// Input consists of a single integer
// Output is a slice consisting of the inputs total population chain (ascending order)
func totalChain(num int) []int{
	var inc = incChain(num)
	var dec = append(decChain(num), num)
	return append(dec, inc...)
}

// Return the increasing population chain of a number (excluding number)
// Input consists of a single integer
// Output is a slice of the increasing population chain of the input number (ascending order)
func incChain(num int) []int{
	if ((num % 2) == 0) {
		var inc = (num * 3) / 2
		var chain = incChain(inc)
		// prepend to start of chain
		return append([]int{inc}, chain...)
	} else {
		return []int{}
	}
}


// Return decreasing population chain of number (excluding number)
// Input consists of a single integer
// Output is a slice of the decreasing chain of the input number (ascending order)
func decChain(num int) []int{
	if ((num % 3) == 0) {
		var dec = (num * 2) / 3
		var chain = decChain(dec)
		// Normal append to end of chain
		return append(chain, dec)
	} else {
		return []int{}
	}
}


func main() {
	fmt.Print("Enter a number to check it's population sequence: ")
	
	var target int
	fmt.Scan(&target)

	fmt.Printf("%d : %d\n", target, totalChain(target))
}