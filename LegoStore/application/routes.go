package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/shawnsey/LegoMOC/LegoStore/database/daos"
	"github.com/shawnsey/LegoMOC/LegoStore/handler"
	// jwtmiddleware "github.com/shawnsey/LegoMOC/LegoStore/middleware"
)

func (app *App) loadRoutes() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", app.loadOrderRoutes)
	router.Route("/creations", app.loadCreationRoutes)

	app.router = router
}

func (app *App) loadOrderRoutes(router chi.Router) {
	orderHandler := handler.OrderHandler{
		OrderDao: daos.NewOrderPsqlDao(app.DB),
		}
	
	router.Post("/", orderHandler.Create)
	router.Get("/", orderHandler.List)
	router.Get("/{id}", orderHandler.Get)
	router.Put("/{id}", orderHandler.Update)
	router.Delete("/{id}", orderHandler.DeleteById)

}

func (app *App) loadCreationRoutes(router chi.Router) {
	creationHandler := handler.CreationsHandler{
		CreationsDao: &daos.CreationsPsqlDao{
			Client: app.DB,
		},
	}

	router.Post("/", creationHandler.Create)
	router.Get("/", creationHandler.List)
	router.Get("/search", creationHandler.Search)
	router.Get("/{id}", creationHandler.GetById)
	router.Put("/{id}", creationHandler.UpdateById)
	router.Delete("/{id}", creationHandler.DeleteById)
}
