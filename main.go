package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"time"

	"github.com/jackc/pgx"
) 

var baseSvr = "xxxx"
var secretsauce = "xxxx"
var databaseUrl = "xxxx"

var data = []student{
	{"E001", "Alice", 90},
	{"E002", "Bob", 85},
	{"E003", "Charlie", 92},
	{"F001", "David", 88},
	{"F002", "Eva", 95},
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

type student struct {
	ID string
	Name string
	Grade int
}

type changelog struct {
	ReleaseDate string   `json:"release_date"`
	Version     string   `json:"version"`
	Changelog   []string `json:"changelog"`
}

type changelogResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	App     interface{} `json:"app"`
	Data    []changelog `json:"data"`
}

func main() {
	runtime.GOMAXPROCS(2) // Set jumlah goroutine yang berjalan bersamaan
	dbTest()
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

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, I'm learning Golang!")
}

func handleIndex() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Heloo from Golang!")
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Name":	"John Doe",
			"Message": "Welcome to the home page!",
		}

		var homeTemp,err = template.ParseFiles("template.html")
		if err != nil {
			fmt.Println("Error parsing template:", err)
			return
		}
		homeTemp.Execute(w, data)
	})

	http.HandleFunc("/api/user", WebAPILearnUser)
	http.HandleFunc("/api/users", WebAPILearnUsers)

	http.HandleFunc("/index", index)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func urlParsing() {
	var urlString = "https://pdf.hana-ci.com/compress"
	var uri, err = url.Parse(urlString)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Current Url is %s\n", uri.String())
	fmt.Printf("Protocol is %s\n", uri.Scheme)
	fmt.Printf("Host is %s\n", uri.Host)
	fmt.Printf("Path is %s\n", uri.Path)
}

func jsonLearn() {
	var jsonString = `{"name": "John", "age": 30, "city": "New York"}`
	var jsonString2 = `[
		{"name": "Jane", "age": 25, "city": "Los Angeles"},
		{"name": "Mike", "age": 32, "city": "Chicago"},
		{"name": "Sara", "age": 28, "city": "Miami"}
	]`

	_ = []byte(jsonString)
	var jsonData2 = []byte(jsonString2)

	var data []User

	var err = json.Unmarshal(jsonData2, &data)
	if err != nil {
		fmt.Println("Unmarshal error:", err)
		return
	}

	fmt.Println("Name:", data[2].Name)
	fmt.Println("Age:", data[2].Age)
	fmt.Println("City:", data[2].City)
}

func WebAPILearnUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		var result, err = json.Marshal(data)
			if err != nil {
				http.Error(w, "Error marshaling data", http.StatusInternalServerError)
				return
			}

			w.Write(result)
			return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func WebAPILearnUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		var id = r.FormValue("id")
			var result []byte
			var err error

			for _, each := range data {
				if each.ID == id {
					result, err = json.Marshal(each)
					if err != nil {
						http.Error(w, "Error marshaling data", http.StatusInternalServerError)
						return
					}
					w.Write(result)
					return
				}
			}

			http.Error(w, "User not found", http.StatusNotFound)
			return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func HttpRequestLearn()(changelogResponse, error) {
	var err error
	var client = &http.Client{}
	var data changelogResponse

	request, err := http.NewRequest("GET", baseSvr+"/api/v1/version/fetch", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return data, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+secretsauce)
	request.Header.Set("User-Agent", "Golang-HTTP-Client/1.0")
	request.Header.Set("X-Requested-With", "XMLHttpRequest")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error making request:", err)
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return data, err
	}

	return data, nil
}

func callHttpRequestLearn() {
	data, err := HttpRequestLearn()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, change := range data.Data {
		fmt.Printf("Release Date: %s\n", change.ReleaseDate)
		fmt.Printf("Version: %s\n", change.Version)
		fmt.Printf("Changelog:",)
		for _, changelog := range change.Changelog {
			fmt.Printf(" - %s\n", changelog)
		}
	}
}

func dbTest() {
	connConfig, err := pgx.ParseURI(databaseUrl)
	if err != nil {
		fmt.Println("Error parsing database URL:", err)
		return
	}

	conn, err := pgx.Connect(connConfig)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer conn.Close()

	var username, password string
	err = conn.QueryRow("select name, password from users").Scan(&username, &password)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query Row Failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(username, password)
}