// LML - LiveMediaList - Service for choice LiveMedia

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	//"gopkg.in/yaml.v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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

	lml struct {
		Field        []string // is slice Fields database
		State        []string // State is slice livecikle distrib
		Target       []string // Target is slice pupose distrib
		OS           []string // OS is slice OS
		Media        []string // Media is slice units for distribution
		Architecture []string // Architecture is slice processor architecture
		License      []string // License is slice licenses
		db           [][]string
	}
)

func init() {
	fmt.Println(" ")
	fmt.Println("Start Init func: ")
	fmt.Println(" ")

	// readConf() читає конфігурацію з файлу //**************************
	readValueDb() // читає з файлу можливі значення полей в мапи
	// -----------------------------------------------------------------------
	for ind, vol := range lml.License {
		println("index = ", ind, " volume = ", vol)
	}
	fmt.Println(" ")
	// -----------------------------------------------------------------------
	//filename := "./db/DB.txt"
	////readStartDb() // читає з файлу БД яка була на http://livecdlist.com/
	readStartDb2()

	for _, v := range lml.db {
		for j, vv := range v {
			println("number = ", j, "volume = ", vv)
		}
	}
	fmt.Println(" ")
	for i, v := range lml.db[1] {
		println("number = ", i, "volume = ", v)
	}
	// -----------------------------------------------------------------------
	// if _, err := os.Stat("./db/DB.txt"); os.IsNotExist(err) {
	// 	readStartDb() // читає з файлу БД яка була на http://livecdlist.com/
	// } else {
	// 	readDB() // читає БД з файла
	// }

	// writeLMLdb()  // пише БД в файл та робить копію попередньої

	// checkLMLdb()  //перевірка посилань дистрибутивів

}

func main() {

	fmt.Println(" ")
	fmt.Println("Start Main func: ")
	fmt.Println(" ")

	// для отдачи сервером статичных файлов из папки public/static
	// fcss := http.FileServer(http.Dir("./assets/css"))
	// http.Handle("/assets/css/", http.StripPrefix("/assets/css/", fcss))
	// fjs := http.FileServer(http.Dir("./assets/js"))
	// http.Handle("/assets/js/", http.StripPrefix("/assets/js/", fjs))
	// fimg := http.FileServer(http.Dir("./assets/img"))
	// http.Handle("/assets/img/", http.StripPrefix("/assets/img/", fimg))
	//

	serveHTTP()
	fmt.Println(" ")
}

func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}

func serveHTTP() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	//r.Use(middleware.Recoverer)
	//r.Use(middleware.URLFormat)
	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "assets/js")
	fileServer(r, "/js/", http.Dir(filesDir))
	filesDir = filepath.Join(workDir, "assets/css")
	fileServer(r, "/css/", http.Dir(filesDir))
	filesDir = filepath.Join(workDir, "assets/img")
	fileServer(r, "/img", http.Dir(filesDir))

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	r.Get("/", httpHandler)
	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("hello world"))
	// })

	r.Route("/lml", func(r chi.Router) {
		//r.With(paginate).Get("/", ListArticles)
		//r.Post("/", CreateArticle)       // POST /articles
		//r.Get("/search", SearchArticles) // GET /articles/search
		//r.Get("/{id}", GetLML)

		r.Route("/{id}", func(r chi.Router) {
			//r.Use(ArticleCtx)            // Load the *Article on the request context
			r.Get("/", GetLML) // GET /articles/123
			//r.Put("/", UpdateArticle)    // PUT /articles/123
			//r.Delete("/", DeleteArticle) // DELETE /articles/123
		})

		// GET /articles/whats-up
		//r.With(ArticleCtx).Get("/{articleSlug:[a-z-]+}", GetArticle)
	})

	//http.HandleFunc("/", httpHandler)
	//http.HandleFunc("/lm", lmHandler)
	//http.ListenAndServe("127.0.0.1:8080", r)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}

// GetLML is
func GetLML(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DistribPage GET page ****")
	ID := chi.URLParam(r, "id")
	// ctx := r.Context()
	// key := ctx.Value("key").(string)
	fmt.Println(ID)
	parsedTemplate, err := template.ParseFiles(
		"assets/http/index.html",
		"assets/http/nav.html",
		"assets/http/header.html",
		"assets/http/body_d.html",
		"assets/http/footer.html",
	)
	if err != nil {
		panic(err)
	}
	x, err := strconv.Atoi(ID)
	if err != nil {
		panic(err)
	}
	err = parsedTemplate.Execute(w, lmlDB[x])
	if err != nil {
		panic(err)
	}

	//w.Write([]byte("get lml " + val + id + " " + key + "\n"))
	//w.Write([]byte("get lml " + " " + id + "\n"))
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
