// LML - LiveMediaList
// Service for choice LiveMedia
// Roadmap
// 1 - message at start programm
// 1 - Pattern
// 2 - Read start data from file and write to DB
// 3 - Read config file
// 4 - Read file variable
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	//"gopkg.in/yaml.v2"
)

type listMediaLive struct {
	number       int
	name         string
	homepage     string
	download     string
	wikipedia    string
	distrowatch  string
	sizeMin      int
	sizeMax      int
	stableVer    string
	lastRelease  string
	target       []int
	os           byte
	state        byte
	media        []int
	architecture []int
	note         string
}

var (
	lmlDB        []listMediaLive
	State        = map[string]byte{}
	Target       = map[string]byte{}
	OS           = map[string]byte{}
	Media        = map[string]byte{}
	Architecture = map[string]byte{}
)

//	State: Target: OS: Media: Architecture:

func init() {
	fmt.Println(" ")
	fmt.Println("Start Init func: ")
	fmt.Println(" ")
	readValueDb()
	readStartDb()

}

func main() {

	fmt.Println(" ")
	fmt.Println("Start Main func: ")
	fmt.Println(" ")

	// для отдачи сервером статичных файлов из папки public/static
	//fs := http.FileServer(http.Dir("./public/static"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))
	fcss := http.FileServer(http.Dir("./assets/css"))
	http.Handle("/assets/css/", http.StripPrefix("/assets/css/", fcss))
	fjs := http.FileServer(http.Dir("./assets/js"))
	http.Handle("/assets/js/", http.StripPrefix("/assets/js/", fjs))

	serveHTTP()
	fmt.Println(" ")
}

func serveHTTP() {
	http.HandleFunc("/", httpHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func httpHandler(w http.ResponseWriter, r *http.Request) {

	//t := template.New("some template")              // Create a template.
	//t, _ = t.ParseFiles("./assets/http/index.html") // Parse template file.
	//user := GetUser()                             // Get current user infomration.
	//t.Execute(w) // merge.
	//t.Execute(os.Stdout, nil)

	//fmt.Fprintf(w, "Hello")
	//log.Printf("request from %s: %s %q", r.RemoteAddr, r.Method, r.URL)
	//fmt.Fprintf(w, "go-daemon: %q", html.EscapeString(r.URL.Path))

	//logrus.Println("StartPage GET start page +++")
	log.Println("StartPage GET start page +++")
	// person := Person{ID: "1", Name: "Foo"}
	parsedTemplate, err := template.ParseFiles("assets/http/index.html", "assets/http/header.html",
		"assets/http/footer.html", "assets/http/nav.html")
	if err != nil {
		log.Println("I don`t parse static files assets/http/index.html")
	}
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("Error arsedTemplate.Execute in StartPage : ", err)
		return
	}

	//fmt.Println("Start Page: ")
	//w.Write([]byte("hello"))

}
