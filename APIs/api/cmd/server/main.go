package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/wendellnd/graduate-go-expert-classes/api/configs"
	_ "github.com/wendellnd/graduate-go-expert-classes/api/docs"
	"github.com/wendellnd/graduate-go-expert-classes/api/internal/entity"
	"github.com/wendellnd/graduate-go-expert-classes/api/internal/infra/database"
	"github.com/wendellnd/graduate-go-expert-classes/api/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title           Go Expert API Example
// @version         1.0
// @description     Product API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Wendell Dourado
// @contact.url    wendelldourado.dev
// @contact.email  wendelldourado.wn@gmail.com

// @license.name   Full Cycle License
// @license.url    wendelldourado.dev

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{}, &entity.Product{})

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, cfg.TokenAuth, cfg.JWTExpiresIn)

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(cfg.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Get("/", productHandler.GetProducts)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJWT)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)
}

//func LogRequest(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		println(r.Method, r.URL.Path)
//		next.ServeHTTP(w, r)
//	})
//}
