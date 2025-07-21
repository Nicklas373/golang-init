package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"text/template"
)

type M map[string]interface{}

type Info struct {
	Affiliation string
	Address string
}

type Person struct {
	Name string
	Gender string
	Hobbies []string
	Info Info
}

func actionHandler() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = Person{
			Name:    "John Doe",
			Gender: "Male",
			Hobbies: []string{"Reading", "Traveling", "Cooking"},
			Info: Info{
				Affiliation: "Golang Enthusiasts",
				Address:     "123 Golang St, Go City, GO 12345",
			},
		}

		var temp = template.Must(template.ParseFiles("views/view.html"))
		if err := temp.Execute(w, person); err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
		}
	})
	fmt.Println("Local Server start at localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func (t Info) GetAffiliationDetailInfo() string {
    return "have 31 divisions"
}

func partialRender() {
	var temp, err = template.ParseGlob("views/*.html")
	if err != nil {
		panic(err.Error())
	}	

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var data = M{
			"title":   "Welcome to Golang Server",
			"message": "This is a simple web server in Go",
		}

		err = temp.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		var data = M{
			"title":   "About Golang Server",
			"message": "This is the about page of the Golang server",
		}

		err = temp.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Local Server start at localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var message string = "Welcome to Golang Server !"
	w.Write([]byte(message))
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	var message string = "Hello from Golang !"
	w.Write([]byte(message))
}

func staticHandler() {
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	fmt.Println("Local Server start at localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func emptyHandler(w http.ResponseWriter, r *http.Request) {
	// This handler does nothing
	var filepath string = path.Join("views", "index.html")
	var temp,err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "Welcome to Golang Server",
		"message": "This is a simple web server in Go",
	}

	err = temp.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

func learnHTTPMethods() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			w.Write([]byte("Received a POST request"))
		case http.MethodGet:
			w.Write([]byte("Received a GET request"))
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Local Server start at localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func learnFormValue() {
	http.HandleFunc("/", routeIndexGet)
	http.HandleFunc("/submit", routeIndexPost)

	fmt.Println("Local Server start at localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func routeIndexGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var temp,err = template.ParseFiles("newview.html")

		if err != nil {
			http.Error(w, "Error executing template form: "+err.Error(), http.StatusInternalServerError)
			return
		}

		err = temp.ExecuteTemplate(w, "form", nil)
		if err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "Bad Request", http.StatusBadRequest)
}

func routeIndexPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var temp, err = template.ParseFiles("newview.html")

		if err != nil {
			http.Error(w, "Error executing template form: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		var name = r.FormValue("name")
		var message = r.FormValue("message")

		var data = map[string]string{"name": name, "message": message}

		if err := temp.ExecuteTemplate(w, "result", data); err != nil {
			http.Error(w, "Error executing template result", http.StatusInternalServerError)
			return
		}
		return
	}

	http.Error(w, "Bad Request", http.StatusBadRequest)
}

func learnUploadFileMain() {
	http.HandleFunc("/", routeHandlerUploadFileMain)
	http.HandleFunc("/upload", routeHandlerUploadFile)

	fmt.Println("Local Server start at localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func routeHandlerUploadFileMain(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var filepath string = path.Join("views", "file.html")
	var temp = template.Must(template.ParseFiles(filepath))
	var err = temp.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func routeHandlerUploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, "Error parsing form: "+err.Error(), http.StatusBadRequest)
		return
	}

	alias := r.FormValue("alias")

	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer uploadedFile.Close()

	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, "Error getting current directory: "+err.Error(), http.StatusInternalServerError)
		return
	}

	filename := handler.Filename
	if alias != "" {
		filename = alias + path.Ext(handler.Filename)
	}

	fileLocation := path.Join(dir, "uploads", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error creating file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		http.Error(w, "Error saving file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Oke"))
}

func main() {
	// http.HandleFunc("/", indexHandler)
	// http.HandleFunc("/index", indexHandler)
	// http.HandleFunc("/welcome", welcomeHandler)

	// var address = "localhost:8080"
	// fmt.Printf("Server started at %s\n", address)
	// err := http.ListenAndServe(address, nil)
	// if err != nil {
	// 	fmt.Println("Error Starting server due: ", err)
	// }
	learnUploadFileMain()
}