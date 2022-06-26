package main

import (
	"fmt"
	"time"
)

type Questions struct {
	questions      string
	answers        []string
	correctAnswers string
}

func addQuestion(question string, answer []string, correctAnswer string) Questions {
	askedQuestion := Questions{questions: question, answers: answer, correctAnswers: correctAnswer}
	return askedQuestion
}

func (q Questions) askQuestion() {
	fmt.Println(q.questions)
	fmt.Println(q.answers)
}

func main() {
	j := 0
	questions1 := addQuestion("Highest mountain?", []string{"nepal", "india", "china"}, "nepal")
	questions2 := addQuestion("Tallest animal?", []string{"lion", "tiger", "giraffe"}, "giraffe")
	questions3 := addQuestion("Fastest animal?", []string{"tiger", "cheetah", "rhino"}, "cheetah")
	questions4 := addQuestion("Largest animal?", []string{"elephant", "snake", "whale"}, "whale")

	questionList := []Questions{questions1, questions2, questions3, questions4}
	timer := time.NewTicker(20 * time.Second)
	correctAnswers := 0
	incorrectAnswers := 0

	for j < 1{
		for i := range questionList{
			fmt.Println("makes new channel")
			var userAnswer string
			questionChannel := make(chan string)
			questionList[i].askQuestion()
			
			go func ()  {
				fmt.Scan(&userAnswer)
				fmt.Println(userAnswer)
				questionChannel <- userAnswer
			}()
			select {
			case ans := <-questionChannel:
				fmt.Println(ans)
				if ans == questionList[i].correctAnswers{
					fmt.Println("Correct Answer")
					correctAnswers += 1
					continue
				} else{
					fmt.Println("Incorrect Answer")
					incorrectAnswers += 1
					continue
				}
			case <- timer.C:
				fmt.Println("timeout")
				break
			}
			break
		}
		break
	}
	fmt.Println("Correct: ", correctAnswers)
	fmt.Println("Incorrect: ", incorrectAnswers)
}
