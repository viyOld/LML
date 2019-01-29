// LML - LiveMediaList - Service for choice LiveMedia

package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	//"gopkg.in/yaml.v2"
)

type listMediaLive struct {
	Number       int
	Name         string
	Homepage     string
	Download     string
	Wikipedia    string
	Distrowatch  string
	SizeMin      int
	SizeMax      int
	StableVer    string
	LastRelease  string
	Target       []int
	OS           byte
	BasedOS      string
	License      byte
	Language     string // Primary Language(s):
	State        byte
	Media        []int
	Architecture []int
	Note         string
	Rating       int
}

var (
	lmlDB []listMediaLive
	// State is map livecikle distrib
	State = map[string]byte{}
	// Target is map pupose distrib
	Target       = map[string]byte{}
	OS           = map[string]byte{}
	Media        = map[string]byte{}
	Architecture = map[string]byte{}
	License      = map[string]byte{}
)

//	State: Target: OS: Media: Architecture:

func init() {
	fmt.Println(" ")
	fmt.Println("Start Init func: ")
	fmt.Println(" ")
	// readConf() читає конфігурацію з файлу
	readValueDb() // читає з файлу можливі значення полей в мапи
	readStartDb() // читає з файлу БД яка була на http://livecdlist.com/
	writeLMLdb()  // пише БД в файл та робить копію попередньої
	// readDB() читає БД з файла
	// checkLMLdb()  //перевірка посилань дистрибутивів

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

	parsedTemplate, err := template.ParseFiles(
		"assets/http/index.html",
		"assets/http/header.html",
		"assets/http/footer.html",
		"assets/http/body.html",
		"assets/http/nav.html",
	)
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
	htmlPage = htmlPage + "<h1 class=\"display-3\">" + lmlDB[0].Name + "</h1>"
	htmlPage = htmlPage + "<div class=\"card\" >" //style=\"width: 18rem;\"
	htmlPage = htmlPage + "<ul class=\"list-group list-group-flush\">"
	htmlPage = htmlPage + "<li class=\"list-group-item\">" + "Homepage: " + lmlDB[0].Homepage + "</li>"
	htmlPage = htmlPage + "<li class=\"list-group-item\">" + "Download: " + lmlDB[0].Download + "</li>"
	htmlPage = htmlPage + "<li class=\"list-group-item\">" + "Wikipedia: " + lmlDB[0].Wikipedia + "</li>"
	htmlPage = htmlPage + "<li class=\"list-group-item\">" + "Distrowatch: " + lmlDB[0].Distrowatch + "</li>"
	htmlPage = htmlPage + "</ul> </div>"
	//htmlPage = htmlPage +

	//htmlPage = htmlPage + "<p>Привет, мир</p>"
	htmlPage = htmlPage + "</body>"
	htmlPage = htmlPage + "<footer class=\"footer footer-fixed-bottom font-small gray pt-4\">"
	//footer mt-auto py-3 page-   navbar-fixed-bottom
	htmlPage = htmlPage + "<div class=\"footer-copyright text-center py-3\">© 2019 Copyright:"
	htmlPage = htmlPage + "<a href=\"https://www.orbis.com.ua/\"> Trident</a>"
	htmlPage = htmlPage + "</div> </footer> </html>"

	//</html>"
	fmt.Fprintf(w, htmlPage)
}

func writeLMLdb() {

	filename := "./db/DB.txt"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
	} else {
		if err := os.Rename(filename, "./db/DB.bak"); err != nil {
			panic(err)
		}
		if _, err = os.Create(filename); err != nil {
			panic(err)
		}
	}

	fileDB, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer fileDB.Close()

	fileDB.WriteString("#comments time" + "\n" + "\n")
	for i, v := range lmlDB {
		if _, err = fileDB.WriteString("Number: " + strconv.Itoa(v.Number) + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("Name: " + v.Name + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("Homepage: " + v.Homepage + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("Download: " + v.Download + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("Wikipedia: " + v.Wikipedia + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("Distrowatch: " + v.Distrowatch + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("SizeMin: " + strconv.Itoa(v.SizeMin) + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("SizeMax: " + strconv.Itoa(v.SizeMax) + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("StableVer: " + v.StableVer + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("LastRelease: " + v.LastRelease + "\n"); err != nil {
			panic(err)
		}
		str := ""
		for _, val := range lmlDB[i].Target {
			str = str + " " + strconv.Itoa(val)
		}
		if _, err = fileDB.WriteString("Target: " + str + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("OS: " + strconv.Itoa(int(v.OS)) + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("BasedOS: " + v.BasedOS + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("License: " + strconv.Itoa(int(v.License)) + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("Language: " + v.Language + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("State: " + strconv.Itoa(int(v.State)) + "\n"); err != nil {
			panic(err)
		}
		str = ""
		for _, val := range lmlDB[i].Media {
			str = str + " " + strconv.Itoa(val)
		}
		if _, err = fileDB.WriteString("Media: " + str + "\n"); err != nil {
			panic(err)
		}
		str = ""
		for _, val := range lmlDB[i].Architecture {
			str = str + " " + strconv.Itoa(val)
		}
		if _, err = fileDB.WriteString("Architecture: " + str + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("Note: " + v.Note + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("Rating: " + strconv.Itoa(v.Rating) + "\n"); err != nil {
			panic(err)
		}
		fileDB.WriteString("\n")
	}
}

func checkLMLdb() {
	for ii, v := range lmlDB {
		if v.State == 0 { // Dead is dead
			continue
		}
		// if v.Homepage == "" {
		// 	fmt.Println("For distributiva: " + v.Name + " homepage = nill")
		// 	continue
		// }
		go checkURL(ii)
	}

}

func checkURL(ii int) {
	fmt.Println("Check: " + lmlDB[ii].Name)
	resp, err := http.Get(lmlDB[ii].Homepage)

	if err != nil {
		fmt.Println("For distributiva: "+lmlDB[ii].Name+" error connect with homepage", err)
		//continue
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("For distributiva: "+lmlDB[ii].Name+" error. http-статус: ", resp.StatusCode)
		//continue
		return
	}

	fmt.Println("For distributiva: "+lmlDB[ii].Name+" homepage online. http-статус: ", resp.StatusCode)

}
