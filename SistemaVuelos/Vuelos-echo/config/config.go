package config

import (
    "github.com/go-pg/pg/v10"
    "log"
)

// Connect establece una conexión con la base de datos PostgreSQL
func Connect() *pg.DB {
    db := pg.Connect(&pg.Options{
        User:     "postgres",
        Password: "1234",
        Database: "BD_Vuelos",
    })

    // Verificar la conexión a la base de datos
    _, err := db.Exec("SELECT 1")
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }

    return db
}
