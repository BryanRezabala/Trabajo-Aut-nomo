package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "Vuelos-echo/config"
    "Vuelos-echo/controllers"
    "Vuelos-echo/models"
    "github.com/labstack/echo/v4"
    "github.com/stretchr/testify/assert"
)

func Setup() *echo.Echo {
    db := config.Connect()
    e := echo.New()

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

    return e
}

func TestCreatePasajero(t *testing.T) {
    e := Setup()

    pasajero := models.Pasajero{
        Nombre:   "Luis",
        Apellido: "Martinez",
        Email:    "luis.martinez@example.com",
    }
    jsonPasajero, _ := json.Marshal(pasajero)

    req := httptest.NewRequest(http.MethodPost, "/pasajeros", bytes.NewBuffer(jsonPasajero))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    e.ServeHTTP(rec, req)

    assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestGetPasajero(t *testing.T) {
    e := Setup()

    req := httptest.NewRequest(http.MethodGet, "/pasajeros/1", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    c.SetParamNames("id")
    c.SetParamValues("1")

    e.ServeHTTP(rec, req)
    assert.Equal(t, http.StatusOK, rec.Code)
}

func TestCreateReserva(t *testing.T) {
    e := Setup()

    reserva := models.Reserva{
        PasajeroID: 1,
        VueloID:    1,
        Fecha:      "2023-08-01",
    }
    jsonReserva, _ := json.Marshal(reserva)

    req := httptest.NewRequest(http.MethodPost, "/reservas", bytes.NewBuffer(jsonReserva))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    e.ServeHTTP(rec, req)

    assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestGetReserva(t *testing.T) {
    e := Setup()

    req := httptest.NewRequest(http.MethodGet, "/reservas/1", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    c.SetParamNames("id")
    c.SetParamValues("1")

    e.ServeHTTP(rec, req)
    assert.Equal(t, http.StatusOK, rec.Code)
}

func TestCreateVuelo(t *testing.T) {
    e := Setup()

    vuelo := models.Vuelo{
        Origen:  "Madrid",
        Destino: "Barcelona",
        Fecha:   "2023-08-01",
    }
    jsonVuelo, _ := json.Marshal(vuelo)

    req := httptest.NewRequest(http.MethodPost, "/vuelos", bytes.NewBuffer(jsonVuelo))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    e.ServeHTTP(rec, req)

    assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestGetVuelo(t *testing.T) {
    e := Setup()

    req := httptest.NewRequest(http.MethodGet, "/vuelos/1", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    c.SetParamNames("id")
    c.SetParamValues("1")

    e.ServeHTTP(rec, req)
    assert.Equal(t, http.StatusOK, rec.Code)
}
