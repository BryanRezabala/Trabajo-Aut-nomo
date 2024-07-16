package graphql

import (
    "github.com/graphql-go/graphql"
    "Vuelos/repository"
)

type Resolver struct {
    PasajeroRepo repository.PasajeroRepository
    VueloRepo    repository.VueloRepository
    ReservaRepo  repository.ReservaRepository
}

var pasajeroType = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "Pasajero",
        Fields: graphql.Fields{
            "id": &graphql.Field{
                Type: graphql.Int,
            },
            "nombre": &graphql.Field{
                Type: graphql.String,
            },
            "apellido": &graphql.Field{
                Type: graphql.String,
            },
            "email": &graphql.Field{
                Type: graphql.String,
            },
        },
    },
)

var vueloType = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "Vuelo",
        Fields: graphql.Fields{
            "id": &graphql.Field{
                Type: graphql.Int,
            },
            "origen": &graphql.Field{
                Type: graphql.String,
            },
            "destino": &graphql.Field{
                Type: graphql.String,
            },
            "fecha": &graphql.Field{
                Type: graphql.String,
            },
        },
    },
)

var reservaType = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "Reserva",
        Fields: graphql.Fields{
            "id": &graphql.Field{
                Type: graphql.Int,
            },
            "pasajero_id": &graphql.Field{
                Type: graphql.Int,
            },
            "vuelo_id": &graphql.Field{
                Type: graphql.Int,
            },
            "fecha": &graphql.Field{
                Type: graphql.String,
            },
        },
    },
)

func (r *Resolver) InitSchema() *graphql.Schema {
    fields := graphql.Fields{
        "pasajeros": &graphql.Field{
            Type: graphql.NewList(pasajeroType),
            Resolve: func(params graphql.ResolveParams) (interface{}, error) {
                return r.PasajeroRepo.GetAll()
            },
        },
        "vuelos": &graphql.Field{
            Type: graphql.NewList(vueloType),
            Resolve: func(params graphql.ResolveParams) (interface{}, error) {
                return r.VueloRepo.GetAll()
            },
        },
        "reservas": &graphql.Field{
            Type: graphql.NewList(reservaType),
            Resolve: func(params graphql.ResolveParams) (interface{}, error) {
                return r.ReservaRepo.GetAll()
            },
        },
    }

    rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
    schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
    schema, _ := graphql.NewSchema(schemaConfig)

    return &schema
}
