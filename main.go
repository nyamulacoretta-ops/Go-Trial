package main

import "fmt"

const pMoney int = 32

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

	fmt.Println("starting Textio server")
	fmt.Println(herName, "is ", herAgeInt)
	fmt.Println(msg)
	fmt.Println(statement)
}

//for conditionals, we use if, else if, else....
