package handlers

import (
    "net/http"
    "Vuelos/models"
    "Vuelos/repository"
    "strconv"

    "github.com/labstack/echo/v4"
)

type ReservaHandler struct {
    Repo repository.ReservaRepository
}

func (h *ReservaHandler) GetAllReservas(c echo.Context) error {
    reservas, err := h.Repo.GetAll()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusOK, reservas)
}

func (h *ReservaHandler) GetReserva(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    reserva, err := h.Repo.GetByID(id)
    if err != nil {
        return c.JSON(http.StatusNotFound, err)
    }
    return c.JSON(http.StatusOK, reserva)
}

func (h *ReservaHandler) CreateReserva(c echo.Context) error {
    var reserva models.Reserva
    if err := c.Bind(&reserva); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }
    if err := h.Repo.Create(reserva); err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusCreated, reserva)
}

func (h *ReservaHandler) UpdateReserva(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    var reserva models.Reserva
    if err := c.Bind(&reserva); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }
    reserva.ID = id
    if err := h.Repo.Update(reserva); err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusOK, reserva)
}

func (h *ReservaHandler) DeleteReserva(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := h.Repo.Delete(id); err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.NoContent(http.StatusNoContent)
}
