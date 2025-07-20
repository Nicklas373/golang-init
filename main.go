package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
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

type UserCredentials struct {
	ID int 
	Name string
	Email string
	Email_verified_at string
	Password string
	Remember_token string 
	Created_at string
	Updated_at string
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

type pdfCompress struct {
	Filename string
	Filesize string
	CompMethod string 
}

func main() {
	runtime.GOMAXPROCS(2) // Set jumlah goroutine yang berjalan bersamaan
	sqlExec()
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

func sqlConnect() (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQuery() {
	db, err := sqlConnect()
	if err != nil {
		fmt.Println("Error connection: ", err)
		return
	}
	defer db.Close()

	var userName string = "eureka"
	rows, err := db.Query("select name, email, created_at from users where name = $1", userName)
	if err != nil {
		fmt.Println("Error Query: ", err)
		return
	}
	defer rows.Close()

	var result []UserCredentials

	for rows.Next() {
		var each = UserCredentials{}
		var err = rows.Scan(&each.Name, &each.Email, &each.Created_at)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.Name)
		fmt.Println(each.Email)
		fmt.Println(each.Created_at)
	}
}

func sqlQueryRow() {
	var db, err = sqlConnect()
	if err != nil {
		fmt.Println("Error koneksi ke database: ", err)
		return
	}
	defer db.Close()

	var result = pdfCompress{}
	var fileName string = "1._Cover.pdf"
	err = db.QueryRow(`select "fileName", "fileSize", "compMethod" from "pdfCompress" where "fileName" = $1`, fileName).Scan(&result.Filename, &result.Filesize, &result.CompMethod)
	if err != nil {
		fmt.Println("Error queryrow: ", err)
		return
	}
	fmt.Printf("Filename: %s\n Filesize: %s\n Compression Method: %v\n", result.Filename, result.Filesize, result.CompMethod)
}

func sqlPrepare() {
	db, err := sqlConnect()
	if err != nil {
		fmt.Println("Error koneksi ke database: ", err)
		return 
	}
	defer db.Close()

	prp, err := db.Prepare(`select "fileName", "fileSize", "compMethod" from "pdfCompress" where "fileName" = $1`)
	if err != nil {
		fmt.Println("Prepare error: ", err)
		return
	}

	var result = pdfCompress{}
	prp.QueryRow("Ebook_Panduan_ClearOS_6_by_Andimicro.pdf").Scan(&result.Filename, &result.Filesize, &result.CompMethod)
	fmt.Printf("Filename: %s\n Filesize: %s\n Compression Method: %v\n", result.Filename, result.Filesize, result.CompMethod)

	var result2 = pdfCompress{}
	prp.QueryRow("estatement_04012024.pdf").Scan(&result2.Filename, &result2.Filesize, &result2.CompMethod)
	fmt.Printf("Filename: %s\n Filesize: %s\n Compression Method: %v\n", result2.Filename, result2.Filesize, result2.CompMethod)

	var result3 = pdfCompress{}
	prp.QueryRow("pdfWatermark_bcae1fbc.pdf").Scan(&result3.Filename, &result3.Filesize, &result3.CompMethod)
	fmt.Printf("Filename: %s\n Filesize: %s\n Compression Method: %v\n", result3.Filename, result3.Filesize, result3.CompMethod)
}

func sqlExec() {
	db, err := sqlConnect()
	if err != nil {
		fmt.Println("Error koneksi ke database: ", err)
		return 
	}
	defer db.Close()

	GUID := uuid.New()
	_, err = db.Exec("insert into sessions values ($1, $2, $3, $4, $5, $6)", GUID, nil, "10.0.0.1", "MyGoClient/1.0", "GOLANG", "1752985976")
	if err != nil {
		fmt.Println("Exec error: ", err)
		return
	}
	fmt.Println("Insert berhasil gan, hehe")

	_, err = db.Exec("update sessions set ip_address = $1 where id = $2", "10.0.0.2", GUID)
	if err != nil {
		fmt.Println("Exec error: ", err)
		return
	}
	fmt.Println("Update berhasil gan, hehe")

	_, err = db.Exec("delete from sessions where ip_address = $1", "10.0.0.2")
	if err != nil {
		fmt.Println("Exec error: ", err)
		return
	}
	fmt.Println("Delete berhasil gan, hehe")
}