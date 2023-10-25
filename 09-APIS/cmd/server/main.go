package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/isslerman/goexpert/09-APIS/configs"
	_ "github.com/isslerman/goexpert/09-APIS/docs"
	"github.com/isslerman/goexpert/09-APIS/internal/entity"
	"github.com/isslerman/goexpert/09-APIS/internal/infra/database"
	"github.com/isslerman/goexpert/09-APIS/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title           Go Expert API Example
// @version         1.0
// @description     Product API with authentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Marcos Issler
// @contact.url    http://www.bosshouse.com.br
// @contact.email  marcos@bosshouse.com.br

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, configs.TokenAuth, configs.JWTExpiresIn)

	r := chi.NewRouter()

	// Middleware list - https://github.com/go-chi/chi#middlewares
	r.Use(middleware.Logger)
	// não deixa o server morrer
	r.Use(middleware.Recoverer)
	// Product
	r.Route("/products", func(r chi.Router) {
		// usando esse midware para passar o token para o contexto do http.
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		// esse valida e ve se ele é válido
		// todas chamadas do grupo /products precisam tem um token valido para serem acessadas.
		// ou seja: tornamos o acesso /products privado a quem estiver logado.
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	// User
	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
		r.Post("/generate_token", userHandler.GetJWT)
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)
}
