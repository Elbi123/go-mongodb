package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// type MyHandler struct{}

// func (mh *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	h := r.Header
// 	fmt.Fprintln(w, h)
// }

// type Hello struct{}

// func (mh *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	h := r.Header
// 	fmt.Fprintln(w, h)
// }

func hello(w http.ResponseWriter, r *http.Request) {
	// h := r.Header
	length := r.ContentLength
	body := make([]byte, length)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func header(w http.ResponseWriter, r *http.Request) {
	h := r.Header

	fmt.Fprintln(w, h)
}

func body(w http.ResponseWriter, r *http.Request) {
	length := r.ContentLength
	body := make([]byte, length)

	r.Body.Read(body)

	fmt.Fprintln(w, string(body))
}

func processForm(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	r.ParseForm()

	fmt.Fprintln(w, r.Form["name"])
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)

	file, header, err := r.FormFile("myFile")
	// file, err := fileHeader.Open()

	if err != nil {
		fmt.Fprintln(w, "error while retrieving file")
		panic(err)
		// return
	}

	defer file.Close()

	fmt.Printf("File Name: %+v\n", header.Filename)
	fmt.Printf("File Size: %+v\n", header.Size)
	fmt.Printf("MIME Header: %+v\n", header.Header)

	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")

	if err != nil {
		fmt.Println(err)
	}

	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)

	fmt.Fprintf(w, "File uploaded successfully")
}

func setCookies(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "This is the first cookie created",
		HttpOnly: true,
	}

	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "This is the second cookie created",
		HttpOnly: true,
	}

	w.Header().Set("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String())
}

func anotherWayToSetCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "another_first_cookie",
		Value:    "This is the another first cookie created",
		HttpOnly: true,
	}

	c2 := http.Cookie{
		Name:     "another_second_cookie",
		Value:    "This is the another second cookie created",
		HttpOnly: true,
	}

	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func ServeAPI() error {

	server := http.Server{
		Addr: "127.0.0.1:8087",
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/header", header)
	http.HandleFunc("/body", body)
	http.HandleFunc("/process", processForm)
	http.HandleFunc("/file", uploadFile)
	http.HandleFunc("/setCookies", setCookies)
	http.HandleFunc("/anotherwaytosetcookie", anotherWayToSetCookie)

	return server.ListenAndServe()
}
