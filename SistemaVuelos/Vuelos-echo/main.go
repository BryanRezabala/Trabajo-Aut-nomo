package main

import (
    "log"
    "Vuelos-echo/config"
    "Vuelos-echo/controllers"
    "Vuelos-echo/models"
    "github.com/labstack/echo/v4"
)

func main() {
    // Conectar a la base de datos
    db := config.Connect()
    defer db.Close()

    // Crear el esquema de la base de datos si no existe
    err := models.CreateSchema(db)
    if err != nil {
        log.Fatalf("Error creating schema: %v", err)
    }

    // Inicializar Echo
    e := echo.New()

    // Configurar controladores
    pasajeroController := &controllers.PasajeroController{DB: db}
    reservaController := &controllers.ReservaController{DB: db}
    vueloController := &controllers.VueloController{DB: db}

    // Definir rutas
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

    // Iniciar el servidor
    e.Start(":8080")
}
