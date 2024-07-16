package main

import (
    "Vuelos/db"
    "Vuelos/handlers"
    "Vuelos/repository"
    "Vuelos/routes"
    "Vuelos/graphql"

    "net/http"
    "github.com/gorilla/websocket"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
    db.Init() // Inicializar la conexión a la base de datos
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    pasajeroRepo := repository.NewPasajeroRepository(db.Conectar())
    pasajeroHandler := &handlers.PasajeroHandler{Repo: pasajeroRepo}
    vueloRepo := repository.NewVueloRepository(db.Conectar())
    vueloHandler := &handlers.VueloHandler{Repo: vueloRepo}
    reservaRepo := repository.NewReservaRepository(db.Conectar())
    reservaHandler := &handlers.ReservaHandler{Repo: reservaRepo}

    resolver := &graphql.Resolver{
        PasajeroRepo: pasajeroRepo,
        VueloRepo:    vueloRepo,
        ReservaRepo:  reservaRepo,
    }

    schema := resolver.InitSchema()
    graphQLHandler := &handlers.GraphQLHandler{Schema: schema}

    webSocketHandler := &handlers.WebSocketHandler{
        Upgrader: websocket.Upgrader{
            ReadBufferSize:  1024,
            WriteBufferSize: 1024,
            CheckOrigin: func(r *http.Request) bool {
                return true
            },
        },
    }

    routes.InitRoutes(e, pasajeroHandler, vueloHandler, reservaHandler, graphQLHandler, webSocketHandler)

    e.Logger.Fatal(e.Start(":8080"))
}
