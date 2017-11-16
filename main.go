package main

import "fmt"
import "os"
// import "io/ioutil"
import "bufio"
import "strings"

func check(e error){
	if e != nil {
		panic(e)
	}
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	type quizItem struct {
		question 		string
		expectedAnswer 	string
		actualAnswer 	string
	}
	
	quiz := make([]quizItem,0)
	score := 0
	
	
	
	fmt.Println("Starting the read")
	f, err := os.Open("problems.csv")
	check(err)
	defer f.Close()
	

	// Move back to the start of file to run method two
	_, err = f.Seek(0, 0)
    check(err)
	
	// reads the entire file and splits it by line.
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	
	for scanner.Scan() {

		
		splitQ := strings.Split(scanner.Text(),",")
		theQuestion := splitQ[0]
		fmt.Println(theQuestion)
		theExpectedAnswer := splitQ[1]
		theAnswer, err := reader.ReadString('\n')
		theAnswer = strings.Replace(theAnswer, "\n", "", -1)

		if err != nil {
			println(err)
		}

		if theExpectedAnswer == theAnswer {
			score++
		}
		
		quiz = append(quiz, quizItem{
			question: theQuestion,
			expectedAnswer: theExpectedAnswer,
			actualAnswer: theAnswer,
		})	
		
		


		// fmt.Println(scanner.Text())

	  }
	  fmt.Printf("You scored %d out of a possible %d \n", score, len(quiz))

}