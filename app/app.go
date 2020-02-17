package app
 
import (
	"fmt"
	"log"
	"net/http"
 
	"github.com/firdaus-git/restapi/app/handler"
	"github.com/firdaus-git/restapi/app/model"
	"github.com/firdaus-git/restapi/config"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)
 
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}
 
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)
 
	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}
 
	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}
 
func (a *App) setRouters() {
	a.Get("/api/employees", a.GetAllEmployees)
	a.Post("/api/employees", a.CreateEmployee)
	a.Get("/api/employees/{title}", a.GetEmployee)
	a.Put("/api/employees/{title}", a.UpdateEmployee)
	a.Delete("/api/employees/{title}", a.DeleteEmployee)
	a.Put("/api/employees/{title}/disable", a.DisableEmployee)
	a.Put("/api/employees/{title}/enable", a.EnableEmployee)
}
 
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}
 
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}
 
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}
 
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}
 
func (a *App) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	handler.GetAllEmployees(a.DB, w, r)
}
 
func (a *App) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	handler.CreateEmployee(a.DB, w, r)
}
 
func (a *App) GetEmployee(w http.ResponseWriter, r *http.Request) {
	handler.GetEmployee(a.DB, w, r)
}
 
func (a *App) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	handler.UpdateEmployee(a.DB, w, r)
}
 
func (a *App) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	handler.DeleteEmployee(a.DB, w, r)
}
 
func (a *App) DisableEmployee(w http.ResponseWriter, r *http.Request) {
	handler.DisableEmployee(a.DB, w, r)
}
 
func (a *App) EnableEmployee(w http.ResponseWriter, r *http.Request) {
	handler.EnableEmployee(a.DB, w, r)
}
 
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}