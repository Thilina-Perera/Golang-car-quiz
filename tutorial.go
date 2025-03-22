package main

import "fmt"

func main() {
	fmt.Println("Welcome to the car quiz!")

	fmt.Printf("What is your name?\n")
	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello %s, let's start the quiz!\n", name)

	fmt.Printf("Enter your age:")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("You are old enough to play the quiz!")
		fmt.Println("Lets Start!")
	} else {
		fmt.Println("You are too young to play the quiz!")
		return
	}

	score := 0
	num_questions := 2

	fmt.Printf("What Italian car company is associated closely with the color red? ")
	var answer string
	fmt.Scan(&answer)

	if answer == "Ferrari" || answer == "ferrari" || answer == "FERRARI" {
		fmt.Println("Correct!")
		score += 1
	} else {
		fmt.Println("Wrong!")
	}

	fmt.Printf("In what year was the first car invented?\n A) 1885\n B) 1886\n C) 1867\n D) 1898\n")
	var q2answer string
	fmt.Scan(&q2answer)

	if q2answer == "B" || q2answer == "b" {
		fmt.Println("Correct!")
		score += 1
	} else {
		fmt.Println("Wrong!")
	}

	fmt.Printf("Your score is %v out of %v.\n", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100
	fmt.Printf("You got %v%% correct!", percent)
}
