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
	num_questions := 10

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

	fmt.Printf("How many wheels does a sadan car have?")
	var q3answer int
	fmt.Scan(&q3answer)

	if q3answer == 4 {
		fmt.Println("Correct!")
		score += 1
	} else {
		fmt.Println("Wrong!")
	}

	fmt.Printf("Which company created the Model T?\n A) Ford\n B) Chevrolet\n C) Mercedes\n D) BMW\n")
	var q4answer string
	fmt.Scan(&q4answer)

	if q4answer == "A" || q4answer == "a" {
		fmt.Println("Correct!")
		score += 1
	} else {
		fmt.Println("Wrong!")
	}

	fmt.Printf("What is the top speed of a Bugatti Chiron (in mph)?\n A) 261\n B) 250\n C) 273\n D) 280\n")
	var q5answer string
	fmt.Scan(&q5answer)

	if q5answer == "A" || q5answer == "a" {
		fmt.Println("Correct!")
		score += 1
	} else {
		fmt.Println("Wrong!")
	}

	fmt.Printf("Which country is famous for manufacturing Lamborghini cars?\n A) Germany\n B) France\n C) Italy\n D) USA\n")
	var q6answer string
	fmt.Scan(&q6answer)

	if q6answer == "C" || q6answer == "c" {
		fmt.Println("Correct!")
		score += 1
	} else {
		fmt.Println("Wrong!")
	}

	fmt.Printf("How many cylinders does a typical V8 engine have?")
	var q7answer int
	fmt.Scan(&q7answer)

	if q7answer == 8 {
		fmt.Println("Correct!")
		score += 1
	} else {
		fmt.Println("Wrong!")
	}

	fmt.Printf("What does ABS stand for in a car's braking system?\n A) Automatic Braking System\n B) Anti-lock Braking System\n C) Advanced Braking System\n D) Active Braking System\n")
	var q8answer string
	fmt.Scan(&q8answer)

	if q8answer == "B" || q8answer == "b" {
		fmt.Println("Correct!")
		score += 1
	} else {
		fmt.Println("Wrong!")
	}

	fmt.Printf("Which of these fuels is NOT commonly used in cars?\n A) Diesel\n B) Petrol\n C) Kerosene\n D) Electricity\n")
	var q9answer string
	fmt.Scan(&q9answer)

	if q9answer == "C" || q9answer == "c" {
		fmt.Println("Correct!")
		score += 1
	} else {
		fmt.Println("Wrong!")
	}

	fmt.Printf("What is the unit used to measure engine power?\n A) Newton\n B) Horsepower\n C) Joule\n D) Watt\n")
	var q10answer string
	fmt.Scan(&q10answer)

	if q10answer == "B" || q10answer == "b" {
		fmt.Println("Correct!")
		score += 1
	} else {
		fmt.Println("Wrong!")
	}

	fmt.Printf("Your score is %v out of %v.\n", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100
	fmt.Printf("You got %v%% correct!", percent)
}
