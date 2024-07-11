package main

import (
    "Vuelos-echo/config"
    "Vuelos-echo/controllers"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
    db := config.Connect()
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    pasajeroController := &controllers.PasajeroController{DB: db}
    reservaController := &controllers.ReservaController{DB: db}
    vueloController := &controllers.VueloController{DB: db}

    e.GET("/pasajeros", pasajeroController.GetAll)
    e.POST("/pasajeros", pasajeroController.Create)
    e.GET("/pasajeros/:id", pasajeroController.GetById)
    e.PUT("/pasajeros/:id", pasajeroController.Update)
    e.DELETE("/pasajeros/:id", pasajeroController.Delete)

    e.GET("/reservas", reservaController.GetAll)
    e.POST("/reservas", reservaController.Create)
    e.GET("/reservas/:id", reservaController.GetById)
    e.PUT("/reservas/:id", reservaController.Update)
    e.DELETE("/reservas/:id", reservaController.Delete)

    e.GET("/vuelos", vueloController.GetAll)
    e.POST("/vuelos", vueloController.Create)
    e.GET("/vuelos/:id", vueloController.GetById)
    e.PUT("/vuelos/:id", vueloController.Update)
    e.DELETE("/vuelos/:id", vueloController.Delete)

    e.Logger.Fatal(e.Start(":8080"))
}
