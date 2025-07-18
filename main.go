package main

import (
	"encoding/base64"
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"time"
) 

func main() {
	runtime.GOMAXPROCS(2) // Set jumlah goroutine yang berjalan bersamaan
	execution()
}

func array() {
	var fruits = [4]string{"Apple", "Banana", "Cherry", "Date"}

	for _, fruit := range fruits {
		fmt.Printf("Fruit %s\n", fruit)
	}
}

func backup() {
	var firstName string = "Jane"
	lastName := "Smith"
	middleName := new(string)

	fmt.Printf("Hello, %s %v %s!\n", firstName, middleName, lastName)
}

func slice() {
	var fruits = []string{"Apple", "Banana", "Cherry", "Date"}
	var newFruits = fruits[1:3]

	fmt.Printf("Fruits is %s\n", newFruits)
}

func uinput() {
	var input string
	fmt.Println("Input some number: ")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Printf("You input number %d\n", number)
	} else {
		fmt.Println("Input is not a number")
	}
}

func curTime() {
	_ = time.Now()
	var time2 = time.Date(2025, time.July, 18, 21, 19, 11, 21, time.UTC)
	fmt.Printf("Current time now is %v\n", time2)
}

func timeSpace() {
	var beginAtime time.Time = time.Now()

	time.Sleep(5 * time.Second)

	var beginBtime time.Time = time.Now()

	space := beginBtime.Sub(beginAtime)

	fmt.Println("Time elapsed per second is ", space.Seconds())
	fmt.Println("Time elapsed per minute is ", space.Minutes())
	fmt.Println("Time elapsed per hour is ", space.Hours())
}

func strconversion() {
	var str = "1234"
	var num, err = strconv.Atoi(string(str))
	var num2 = 1234
	var orinum = strconv.Itoa(num2)

	if err == nil {
		fmt.Printf("This is string ! with value: %d \n", num)
		fmt.Println(orinum)
	} else {
		fmt.Printf("This is not a string ! with value: %v", num)
	}
}

func regex() {
	var placeholder string = "This is a placeholder for regex example"
	var regex, err = regexp.Compile(`[a-zA-Z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	var rgxPattern1 = regex.FindAllString(placeholder, 2)
	fmt.Printf("Regex pattern 1: %v\n", rgxPattern1)
	var rgxPattern2 = regex.FindAllString(placeholder, 1)
	fmt.Printf("Regex pattern 2: %v\n", rgxPattern2)
}

func encode() {
	var secretsause string = "lorem ipsum dolor sit amet, consectetur adipiscing elit"

	var encodedSauce = base64.StdEncoding.EncodeToString([]byte(secretsause))
	fmt.Printf("Encoded sauce: %s\n", encodedSauce)

	var decodedSauce,_ = base64.StdEncoding.DecodeString(encodedSauce)
	fmt.Printf("Decoded sauce: %s\n", string(decodedSauce))
}

func execution() {
	if runtime.GOOS == "windows" {
		var output1, err = exec.Command("cmd", "/C", "systeminfo").Output()

		if err == nil {
			fmt.Println("This is the output: ", string(output1))
		} else {
			fmt.Println("Error executing command: ", err.Error())
		}
	}
}