package main

import "fmt"
import "os"
// import "io/ioutil"
import "bufio"
import "strings"
import "flag"
import "time"

func check(e error){
	if e != nil {
		panic(e)
	}
}

type QuizItem struct {
	question 		string
	expectedAnswer 	string
	actualAnswer 	string
	questionScore 	int
}

func main() {	
	
	
	fnPtr := flag.String("fileName", "problems.csv", "name of the file in root that contains the quiz")
	timeLimitPtr := flag.Int("timeLimit", 10, "the amount of time in seconds that you want to set as a limit for the quiz")	
	quizBegin := bufio.NewScanner(os.Stdin)
	quiz := make([]QuizItem,0)	
	
	// Once all flags are declared, call `flag.Parse()`
    // to execute the command-line parsing.
    flag.Parse()
	
	fmt.Println("fileName:",*fnPtr)	
	
	f, err := os.Open(*fnPtr)
	check(err)
	defer f.Close()

	// parse file into slice of quiz objects		
	// reads the entire file and splits it by line.
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		splitQ := strings.Split(scanner.Text(),",")
		theQuestion := splitQ[0]
		theExpectedAnswer := splitQ[1]
		quiz = append(quiz, QuizItem{
			question: theQuestion,
			expectedAnswer: theExpectedAnswer,
			actualAnswer: "",
			questionScore: 0,
		})
	}
	
	fmt.Println("Quiz file has been read in, are you ready to start?  Press enter to continue!")

	for quizBegin.Scan() {
		inputReceived := quizBegin.Text()
		if inputReceived == "" {			
			// Move back to the start of file to run method two
			_, err = f.Seek(0, 0)
			check(err)		
			giveQuiz(&quiz, *timeLimitPtr)			
			break
		}
	}

	

}

func giveQuiz(quizFile *[]QuizItem, timeLimit int) {

	quizTimer := time.NewTimer(time.Second * time.Duration(timeLimit))		
	totalScore := 0

	for i, item := range *quizFile {
		fmt.Printf("Question #%d: %s = \n", i+1, item.question)
		answerCh := make(chan string)

		go func() {
			var actualAnswer string
			fmt.Scanf("%s\n", &actualAnswer)
			answerCh <- actualAnswer
		}()

		select {
		case <- quizTimer.C:
			fmt.Printf("\nThe quiz timer has expired.\n")
			fmt.Printf("You scored %d out of a possible %d \n", totalScore, len(*quizFile))
			return		
		case actualAnswer := <-answerCh:
			fmt.Printf("\nactual answer: %s | expected answer: %s\n", actualAnswer, item.expectedAnswer)
			if actualAnswer == item.expectedAnswer {
				totalScore++
			}
			
		}

	}

	fmt.Printf("You scored %d out of a possible %d \n", totalScore, len(*quizFile))

}

