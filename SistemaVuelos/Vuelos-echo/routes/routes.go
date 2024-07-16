package routes

import (
    "github.com/labstack/echo/v4"
    "Vuelos/handlers"
)

func InitRoutes(e *echo.Echo, pasajeroHandler *handlers.PasajeroHandler, vueloHandler *handlers.VueloHandler, reservaHandler *handlers.ReservaHandler, graphQLHandler *handlers.GraphQLHandler, webSocketHandler *handlers.WebSocketHandler) {
    e.GET("/pasajeros", pasajeroHandler.GetAllPasajeros)
    e.GET("/pasajeros/:id", pasajeroHandler.GetPasajero)
    e.POST("/pasajeros", pasajeroHandler.CreatePasajero)
    e.PUT("/pasajeros/:id", pasajeroHandler.UpdatePasajero)
    e.DELETE("/pasajeros/:id", pasajeroHandler.DeletePasajero)

    e.GET("/vuelos", vueloHandler.GetAllVuelos)
    e.GET("/vuelos/:id", vueloHandler.GetVuelo)
    e.POST("/vuelos", vueloHandler.CreateVuelo)
    e.PUT("/vuelos/:id", vueloHandler.UpdateVuelo)
    e.DELETE("/vuelos/:id", vueloHandler.DeleteVuelo)

    e.GET("/reservas", reservaHandler.GetAllReservas)
    e.GET("/reservas/:id", reservaHandler.GetReserva)
    e.POST("/reservas", reservaHandler.CreateReserva)
    e.PUT("/reservas/:id", reservaHandler.UpdateReserva)
    e.DELETE("/reservas/:id", reservaHandler.DeleteReserva)

    e.POST("/graphql", graphQLHandler.ServeHTTP)
    e.GET("/ws", webSocketHandler.ServeHTTP)
}
