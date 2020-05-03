package main

import (
	"fmt"
	"net/http"

	"lenslocked.com/controllers"
	"lenslocked.com/middleware"
	"lenslocked.com/models"

	"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "lenslocked_dev"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>There is no try. 404</h1>")
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	services, err := models.NewServices(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer services.Close()
	services.AutoMigrate()

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(services.User)
	galleriesC := controllers.NewGalleries(services.Gallery)

	requireUserMw := middleware.RequireUser{
		UserService: services.User,
	}

	newGallery := requireUserMw.Apply(galleriesC.New)
	createGallery := requireUserMw.ApplyFn(galleriesC.Create)

	r := mux.NewRouter()

	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/faq", staticC.Faq).Methods("GET")

	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.Handle("/login", usersC.LoginView).Methods("GET")
	r.HandleFunc("/login", usersC.Login).Methods("POST")
	r.HandleFunc("/cookietest", usersC.CookieTest).Methods("GET")

	r.Handle("/galleries/new", newGallery).Methods("GET")
	r.HandleFunc("/galleries", createGallery).Methods("POST")

	nf := http.HandlerFunc(notFound)
	r.NotFoundHandler = nf

	http.ListenAndServe(":3000", r)
}
