package models

import (
    "github.com/go-pg/pg/v10"
    "github.com/go-pg/pg/v10/orm"
)

// Pasajero representa la tabla 'pasajero' en la base de datos
type Pasajero struct {
    ID        int64  `json:"id"`       // ID del pasajero
    Nombre    string `json:"nombre"`   // Nombre del pasajero
    Apellido  string `json:"apellido"` // Apellido del pasajero
    Email     string `json:"email"`    // Email del pasajero
}

// Reserva representa la tabla 'reserva' en la base de datos
type Reserva struct {
    ID         int64  `json:"id"`         // ID de la reserva
    PasajeroID int64  `json:"pasajero_id"`// ID del pasajero asociado
    VueloID    int64  `json:"vuelo_id"`   // ID del vuelo asociado
    Fecha      string `json:"fecha"`      // Fecha de la reserva
}

// Vuelo representa la tabla 'vuelo' en la base de datos
type Vuelo struct {
    ID      int64  `json:"id"`      // ID del vuelo
    Origen  string `json:"origen"`  // Origen del vuelo
    Destino string `json:"destino"` // Destino del vuelo
    Fecha   string `json:"fecha"`   // Fecha del vuelo
}

// CreateSchema crea las tablas en la base de datos si no existen
func CreateSchema(db *pg.DB) error {
    models := []interface{}{
        (*Pasajero)(nil),
        (*Reserva)(nil),
        (*Vuelo)(nil),
    }

    for _, model := range models {
        err := db.Model(model).CreateTable(&orm.CreateTableOptions{
            IfNotExists: true, // Crear la tabla solo si no existe
        })
        if err != nil {
            return err
        }
    }
    return nil
}
