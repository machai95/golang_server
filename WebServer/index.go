package main

import (
	"GoServer/WebServer/godatabase"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const maxUploadSize = 2 * 1024 * 1024 * 1024 * 1024 // 2 mb
const uploadPath = "./upload"

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./static/index.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
	}
}

func uploadFileHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			renderError(w, "FILE_TOO_BIG", 1000)
			fmt.Println("BIGfile")
		}
		// parse and validate file and post parameters
		fileType := r.PostFormValue("type")
		file, _, err := r.FormFile("uploadFile")
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}
		defer file.Close()
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}

		// check file type, detectcontenttype only needs the first 512 bytes
		filetype := http.DetectContentType(fileBytes)
		switch filetype {
		case "image/jpeg", "image/jpg":
		case "image/gif", "image/png":
		case "application/pdf":
			break
		default:
			renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
			return
		}
		fileName := randToken(12)
		fileEndings, err := mime.ExtensionsByType(fileType)
		if err != nil {
			renderError(w, "CANT_READ_FILE_TYPE", http.StatusInternalServerError)
			return
		}
		newPath := filepath.Join(uploadPath, fileName+fileEndings[0])
		fmt.Printf("FileType: %s, File: %s\n", fileType, newPath)

		// write file
		newFile, err := os.Create(newPath)
		if err != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		defer newFile.Close() // idempotent, okay to call twice
		if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("SUCCESS"))
	})
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// upload logic
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		//fmt.Println(r)
		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
			fmt.Println("BIGfile")
		}
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}

}

// upload logic
func upload1(w http.ResponseWriter, r *http.Request) {
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))

	t, _ := template.ParseFiles("upload.gtpl")
	t.Execute(w, token)
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		//t, _ := template.ParseFiles("login.gtpl")
		//t.Execute(w, nil)
	} else {
		// fmt.Println("hello")
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		// fmt.Printf("Kieu username %T: ", r.Form["username"])
		// fmt.Printf("Kieu password %T: ", r.Form["password"])
		fmt.Println("So sanh user trong database")
		// if r.Form["username"][0] == "macduyhai" || r.Form["username"][0] == "sontv" {
		// 	upload1(w, r)
		// 	fmt.Println("ok")
		// } else {
		// 	http.FileServer(http.Dir("./static"))
		// }
		// fmt.Printf("Len %v: ", len(r.Form["username"]))
		// fmt.Printf("Len-1 %v: ", len(r.Form["username"])-1)
		// fmt.Println("username:", r.Form["username"][len(r.Form["username"])-1])
		// fmt.Println("password:", r.Form["password"][len(r.Form["username"])-1])
		// fmt.Printf("Kieu username %T: ", r.Form["username"][len(r.Form["username"])-1])
		// fmt.Printf("Kieu password %T: ", r.Form["password"][len(r.Form["username"])-1])
		if godatabase.CheckUser(r.Form["username"][len(r.Form["username"])-1], r.Form["password"][len(r.Form["username"])-1]) == true {
			upload1(w, r)
			fmt.Println("ok")
		} else {
			http.FileServer(http.Dir("./static"))
			fmt.Println("Sai username or password")
		}
		fmt.Println("end function login")

	}
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		//t, _ := template.ParseFiles("login.gtpl")
		//t.Execute(w, nil)
	} else {
		// fmt.Println("hello")
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		_ = godatabase.InsertDB(r.Form["username"][len(r.Form["username"])-1], r.Form["password"][len(r.Form["username"])-1])
		fmt.Println("end function login")

	}
}
func Download(w http.ResponseWriter, r *http.Request) {

}
func main() {
	godatabase.CreateDB()
	// http.HandleFunc("/", index) // setting router rule
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/login", login)
	http.HandleFunc("/signin", SignIn)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/download", Download)

	fs := http.FileServer(http.Dir("./upload"))
	http.Handle("/downloads/", http.StripPrefix("/downloads", fs))
	err := http.ListenAndServe(":8080", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
