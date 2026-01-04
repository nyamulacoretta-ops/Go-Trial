package main

import "fmt"

const pMoney int = 32

type messageToSend struct {
	phoneNumber int
	message     string
}

type Corie struct {
	phoneNumber int
	address     string
	homeTown    home
	gender      bool
}
type home struct {
	district string
	ta       string
	village  string
	country  string
}

var myDetails Corie

func test(n messageToSend) {
	fmt.Printf("Sending message from: '%v'\n message: %s\n", n.phoneNumber, n.message)
	fmt.Printf("----------------\n")
}
func calc(pMoney int, rMoney float64) int {
	remainingMoney := int(rMoney)
	remainingMoney = 1200 - pMoney
	return remainingMoney + pMoney
}

func main() {

	herAge := 28.0
	herAgeInt := int(herAge)
	var herName = "Coretta Nyamula"
	const name = "Corie"
	const openrate = 14.8
	//formats tis into a string
	msg := fmt.Sprintf("Hi %s, your open rate is %1f percent", name, openrate)
	var rMoney float64
	myMoney := calc(pMoney, rMoney)
	statement := fmt.Sprintf("My whole inheritance is %v", myMoney)

	test(messageToSend{
		phoneNumber: 265887539584,
		message:     "I really love you"})

	test(messageToSend{
		phoneNumber: 265981831528,
		message:     "I really hate you"})

	fmt.Println("starting Textio server")
	fmt.Println(herName, "is ", herAgeInt)
	fmt.Println(msg)
	fmt.Println(statement)
	fmt.Println(myDetails.homeTown.district)
}

//for conditionals, we use if, else if, else....
