package repository

import (
    "Vuelos/models"
    "github.com/jmoiron/sqlx"
)

type ReservaRepository interface {
    GetAll() ([]models.Reserva, error)
    GetByID(id int) (models.Reserva, error)
    Create(reserva models.Reserva) error
    Update(reserva models.Reserva) error
    Delete(id int) error
}

type reservaRepository struct {
    db *sqlx.DB
}

func NewReservaRepository(db *sqlx.DB) ReservaRepository {
    return &reservaRepository{db: db}
}

func (r *reservaRepository) GetAll() ([]models.Reserva, error) {
    var reservas []models.Reserva
    err := r.db.Select(&reservas, "SELECT * FROM reserva")
    return reservas, err
}

func (r *reservaRepository) GetByID(id int) (models.Reserva, error) {
    var reserva models.Reserva
    err := r.db.Get(&reserva, "SELECT * FROM reserva WHERE id=$1", id)
    return reserva, err
}

func (r *reservaRepository) Create(reserva models.Reserva) error {
    _, err := r.db.NamedExec(`INSERT INTO reserva (pasajero_id, vuelo_id, fecha) VALUES (:pasajero_id, :vuelo_id, :fecha)`, &reserva)
    return err
}

func (r *reservaRepository) Update(reserva models.Reserva) error {
    _, err := r.db.NamedExec(`UPDATE reserva SET pasajero_id=:pasajero_id, vuelo_id=:vuelo_id, fecha=:fecha WHERE id=:id`, &reserva)
    return err
}

func (r *reservaRepository) Delete(id int) error {
    _, err := r.db.Exec("DELETE FROM reserva WHERE id=$1", id)
    return err
}
