/* 
CHALLENGE 01 - Write your own Go program titled, ~/goBasicsChallenge01.go. No need for anything fancy; 

using the fmt.Println() function, print your name to the screen, and a second with your City and State.
Be sure your code executes when you run it with go run, and it builds a binary when you execute go build. 

Note: Be sure to read any build errors you get. They should indicate the line number on which your error is occurring.
*/

package main

import "fmt"

func main(){
	var name = "John Doe"
	var city = "Minneapolis"
	var state = "Minnesota"

	fmt.Println("Name:",name)
	fmt.Println("City: " + city + ", State: " + state)
}

