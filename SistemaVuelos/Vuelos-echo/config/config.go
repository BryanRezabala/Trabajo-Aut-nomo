package config

import (
    "github.com/go-pg/pg/v10"
)

// Connect establece una conexión con la base de datos PostgreSQL
func Connect() *pg.DB {
    db := pg.Connect(&pg.Options{
        User:     "postgres", // Nombre de usuario de PostgreSQL
        Password: "1234",     // Contraseña del usuario de PostgreSQL
        Database: "BD_Tablas",       // Nombre de la base de datos
    })
    return db
}
