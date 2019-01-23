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
	"net/http"
	//"gopkg.in/yaml.v2"
)

type listMediaLive struct {
	Number       int
	Name         string
	homepage     string
	download     string
	wikipedia    string
	distrowatch  string
	sizeMin      int
	sizeMax      int
	stableVer    string
	lastRelease  string
	Target       []int
	os           byte
	State        byte
	media        []int
	architecture []int
	note         string
}

var (
	lmlDB []listMediaLive
	// maps
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
	fcss := http.FileServer(http.Dir("./assets/css"))
	http.Handle("/assets/css/", http.StripPrefix("/assets/css/", fcss))
	fjs := http.FileServer(http.Dir("./assets/js"))
	http.Handle("/assets/js/", http.StripPrefix("/assets/js/", fjs))
	fimg := http.FileServer(http.Dir("./assets/img"))
	http.Handle("/assets/img/", http.StripPrefix("/assets/img/", fimg))

	serveHTTP()
	fmt.Println(" ")
}

func serveHTTP() {
	http.HandleFunc("/", httpHandler)
	http.HandleFunc("/lm", lmHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func httpHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("StartPage GET start page +++")

	parsedTemplate, err := template.ParseFiles("assets/http/index.html", "assets/http/header.html",
		"assets/http/footer.html", "assets/http/nav.html")
	if err != nil {
		fmt.Println("I don`t parse static files assets/http/index.html")
	}
	err = parsedTemplate.Execute(w, lmlDB)
	if err != nil {
		fmt.Println("Error arsedTemplate.Execute in StartPage : ", err)
		return
	}
}

func lmHandler(w http.ResponseWriter, r *http.Request) {
	var htmlPage string
	htmlPage = htmlPage + "<!DOCTYPE html> <html> <head> <meta charset=\"utf-8\" />"
	htmlPage = htmlPage + "<meta name=\"Live Media List\" content=\"width=device-width, initial-scale=1.0\">"
	htmlPage = htmlPage + "<title>LML</title>"
	htmlPage = htmlPage + "<link href=\"/assets/css/bootstrap.min.css\" rel=\"stylesheet\">"
	htmlPage = htmlPage + "<link rel=\"shortcut icon\" href=\"/assets/img/home36.png\" type=\"image\">"
	htmlPage = htmlPage + "<script src=\"./assets/js/bootstrap.min.js\"></script></head>"
	htmlPage = htmlPage + "<body>"
	htmlPage = htmlPage + "<nav class=\"navbar navbar-expand-lg navbar-light bg-light\">"
	htmlPage = htmlPage + "<a class=\"navbar-brand\" href=\"#\"> &nbsp LML</a>"
	htmlPage = htmlPage + "<button class=\"navbar-toggler\" type=\"button\" data-toggle=\"collapse\" data-target=\"#navbarNavAltMarkup\" "
	htmlPage = htmlPage + "aria-controls=\"navbarNavAltMarkup\" aria-expanded=\"false\" aria-label=\"Toggle navigation\"> "
	htmlPage = htmlPage + "<span class=\"navbar-toggler-icon\"></span></button> "
	htmlPage = htmlPage + "<div class=\"collapse navbar-collapse\" id=\"navbarNavAltMarkup\">"
	htmlPage = htmlPage + "<div class=\"navbar-nav\">"
	htmlPage = htmlPage + "<a class=\"nav-item nav-link active\" href=\"/\">Home <span class=\"sr-only\">(current)</span></a>"
	htmlPage = htmlPage + "<a class=\"nav-item nav-link\" href=\"#\">News</a>"
	htmlPage = htmlPage + "<a class=\"nav-item nav-link\" href=\"#\">Info</a>"
	htmlPage = htmlPage + "<a class=\"nav-item nav-link\" href=\"#\">Tests</a>"
	htmlPage = htmlPage + "<a class=\"nav-item nav-link\" href=\"#\">About</a>"
	htmlPage = htmlPage + "<a class=\"nav-item nav-link\" href=\"#\">Profile</a>"
	htmlPage = htmlPage + "<a class=\"nav-item nav-link\" href=\"../auth/sign-up\">Sign Up</a>"
	htmlPage = htmlPage + "<a class=\"nav-item nav-link\" href=\"../auth/login\">Login</a>"
	htmlPage = htmlPage + "<a class=\"nav-item nav-link disabled\" alifn=\"left\" href=\"#\">Disabled</a>"
	htmlPage = htmlPage + "</div> </div> </nav>"
	htmlPage = htmlPage + "<h1 class=\"display-1\">" + lmlDB[0].Name + "</h1>\" "
	htmlPage = htmlPage + "<p>Привет, мир</p>"
	htmlPage = htmlPage + "</body>"
	htmlPage = htmlPage + "<footer class=\"page-footer font-small gray pt-4\">"
	htmlPage = htmlPage + "<div class=\"footer-copyright text-center py-3\">© 2019 Copyright:"
	htmlPage = htmlPage + "<a href=\"https://www.orbis.com.ua/\"> Trident</a>"
	htmlPage = htmlPage + "</div> </footer> </html>"

	//</html>"
	fmt.Fprintf(w, htmlPage)
}
