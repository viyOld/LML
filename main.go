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

var (
	lml struct {
		Field        []string   // is slice Fields database
		State        []string   // State is slice livecikle distrib
		Target       []string   // Target is slice pupose distrib
		OS           []string   // OS is slice OS
		Media        []string   // Media is slice units for distribution
		Architecture []string   // Architecture is slice processor architecture
		License      []string   // License is slice licenses
		DB           [][]string // Data base
	}
)

func init() {
	fmt.Print("\n \n ***** Start Init func ***** \n")

	//***** readConf() читає конфігурацію з файлу

	readValueDb() // читає з файлу можливі значення полей в слайси

	if _, err := os.Stat("./db/DB.txt"); os.IsNotExist(err) {
		readStartDb() // читає з файлу БД яка була на http://livecdlist.com/
	} else {
		readDB() // читає БД з файла
	}

	writeLMLdb() // пише БД в файл та робить копію попередньої

	// checkLMLdb()  //перевірка посилань дистрибутивів

	fmt.Print("\n \n ***** Init func succseful finish ***** \n")
}

func main() {

	fmt.Print("\n \n ***** Start Main func ***** \n")

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

// GetLML is create array for one distributiv
func GetLML(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DistribPage GET page ****")
	ID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(ID)
	if err != nil {
		panic(err)
	}
	DistInfo := make([][]string, 2)
	for i, v := range lml.Field {
		DistInfo[0] = append(DistInfo[0], v)
		DistInfo[1] = append(DistInfo[1], lml.DB[id-1][i])
	}
	// fmt.Println(ID)

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

	err = parsedTemplate.Execute(w, DistInfo)
	if err != nil {
		panic(err)
	}
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
	err = parsedTemplate.Execute(w, lml)
	if err != nil {
		fmt.Println("Error arsedTemplate.Execute in StartPage : ", err)
		return
	}
}

func checkLMLdb() {
	// for ii, v := range lmlDB {
	// 	if v.State == 0 { // Dead is dead
	// 		continue
	// 	}
	// 	// if v.Homepage == "" {
	// 	// 	fmt.Println("For distributiva: " + v.Name + " homepage = nill")
	// 	// 	continue
	// 	// }
	// 	go checkURL(ii)
	// }

}

func checkURL(ii int) {
	fmt.Println("Check: " + "lmlDB[ii].Name")
	// resp, err := http.Get(lmlDB[ii].Homepage)

	// if err != nil {
	// 	fmt.Println("For distributiva: "+lmlDB[ii].Name+" error connect with homepage", err)
	// 	//continue
	// 	return
	// }
	// defer resp.Body.Close()

	// if resp.StatusCode != 200 {
	// 	fmt.Println("For distributiva: "+lmlDB[ii].Name+" error. http-статус: ", resp.StatusCode)
	// 	//continue
	// 	return
	// }

	// fmt.Println("For distributiva: "+lmlDB[ii].Name+" homepage online. http-статус: ", resp.StatusCode)

}

// type listMediaLive struct {
// 	Number       int
// 	Name         string
// 	Homepage     string
// 	Download     string
// 	Wikipedia    string
// 	Distrowatch  string
// 	SizeMin      int
// 	SizeMax      int
// 	StableVer    string
// 	LastRelease  string
// 	Target       []int
// 	OS           byte
// 	BasedOS      string
// 	License      byte
// 	Language     string // Primary Language(s):
// 	State        byte
// 	Media        []int
// 	Architecture []int
// 	Note         string
// 	Rating       int
// }
