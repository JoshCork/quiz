package main

import "fmt"
import "os"
import "io/ioutil"
import "bufio"

func check(e error){
	if e != nil {
		panic(e)
	}
}

func main() {

	// Read and print the entire file
    dat, err := ioutil.ReadFile("problems.csv")
    check(err)
    fmt.Print(string(dat))

	fmt.Println("Starting the read")
	f, err := os.Open("problems.csv")
	check(err)
	defer f.Close()
	

	theBytes := make([]byte, 5)
	newBytes, err := f.Read(theBytes)
	check(err)
	fmt.Printf("%d bytes: %s\n", newBytes, string(theBytes))

	_, err = f.Seek(0, 0)
    check(err)

	newRead := bufio.NewReader(f)
	buffRead, err := newRead.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(buffRead))

	_, err = f.Seek(0, 0)
	check(err)
	
	// reads the entire file and splits it by line.
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	  }


}