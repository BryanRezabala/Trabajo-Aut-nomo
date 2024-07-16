package models

type Reserva struct {
    ID         int    `db:"id" json:"id"`
    PasajeroID int    `db:"pasajero_id" json:"pasajero_id"`
    VueloID    int    `db:"vuelo_id" json:"vuelo_id"`
    Fecha      string `db:"fecha" json:"fecha"`
}
