package models

type Reserva struct {
    ID         int64  `json:"id"`
    PasajeroID int64  `json:"pasajero_id" validate:"required"`
    VueloID    int64  `json:"vuelo_id" validate:"required"`
    Fecha      string `json:"fecha" validate:"required"`
}
