package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/shawnsey/LegoMOC/LegoStore/common"
	"github.com/shawnsey/LegoMOC/LegoStore/database/daos"
	"github.com/shawnsey/LegoMOC/LegoStore/handler"
)

func (app *App) loadRoutes() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	r.Group(func(router chi.Router) {
		router.Route("/orders", app.loadOrderRoutes)
	})

	r.Group(func(router chi.Router) {
		r.Route("/creations", app.loadCreationRoutes)
	})
	

	app.router = r
}

func (app *App) loadOrderRoutes(r chi.Router) {
	orderHandler := handler.OrderHandler{
		OrderDao: daos.NewOrderPsqlDao(app.DB),
		}
	
	r.Post("/", orderHandler.Create)
	r.Get("/", orderHandler.List)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(common.ValidateParams)
			r.Get("/", orderHandler.Get)
			r.Put("/", orderHandler.Update)
			r.Delete("/", orderHandler.DeleteById)
		})
}

func (app *App) loadCreationRoutes(r chi.Router) {
	creationHandler := handler.CreationsHandler{
		CreationsDao: &daos.CreationsPsqlDao{
			Client: app.DB,
		},
	}

	r.Post("/", creationHandler.Create)
	r.Get("/", creationHandler.List)
	r.Get("/search", creationHandler.Search)
	r.Route("/{id}", func(r chi.Router) {
		r.Use(common.ValidateParams)
		r.Get("/", creationHandler.GetById)
		r.Put("/", creationHandler.UpdateById)
		r.Delete("/", creationHandler.DeleteById)
	})


}
