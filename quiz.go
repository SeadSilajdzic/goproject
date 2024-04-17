package main

import "fmt"

func main() {
	fmt.Println("Idiot detector starting...")

	var playerName string
	fmt.Print("Enter your name: ")
	fmt.Scan(&playerName)

	var age uint
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	if age <= 10 {
		fmt.Println("You are not old enough!")
		return
	}

	fmt.Printf("Hello %v, welcome to quiz game. \n", playerName)
	fmt.Println("--------------Game is starting-------------")
	fmt.Println("Question no.1")

	var score int = 0

	var answer1 string
	var answer2 string
	fmt.Print("What is better, the RTX 3080 or RTX 3090? ")
	fmt.Scan(&answer1, &answer2)

	var answer = answer1 + " " + answer2

	if answer == "RTX 3090" || answer == "rtx 3090" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Println("Question no.2")

	var cores int
	fmt.Print("How many cores does the Ryzen 9 3900x have? ")
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	var num_questions = 2
	var percent = (float32(score) / float32(num_questions)) * 100
	fmt.Printf("Your score is %v/%v, which is %v%%. \n", score, num_questions, percent)

	if percent < 100 {
		fmt.Println("Congrats!")
	} else {
		fmt.Println("Good job anyways!")
	}

}
