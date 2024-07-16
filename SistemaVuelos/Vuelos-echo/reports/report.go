package reports

import (
    "bytes"
    "github.com/jung-kurt/gofpdf"
    "Vuelos/models"
    "strconv"
)

func GenerateReport(pasajeros []models.Pasajero, vuelos []models.Vuelo, reservas []models.Reserva) ([]byte, error) {
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    pdf.SetFont("Arial", "B", 16)
    pdf.Cell(40, 10, "Reporte de Vuelos")

    pdf.SetFont("Arial", "", 12)
    pdf.Ln(12)

    pdf.Cell(40, 10, "Pasajeros:")
    pdf.Ln(10)
    for _, pasajero := range pasajeros {
        pdf.Cell(40, 10, pasajero.Nombre+" "+pasajero.Apellido)
        pdf.Ln(10)
    }

    pdf.Ln(10)
    pdf.Cell(40, 10, "Vuelos:")
    pdf.Ln(10)
    for _, vuelo := range vuelos {
        pdf.Cell(40, 10, vuelo.Origen+" -> "+vuelo.Destino+" ("+vuelo.Fecha+")")
        pdf.Ln(10)
    }

    pdf.Ln(10)
    pdf.Cell(40, 10, "Reservas:")
    pdf.Ln(10)
    for _, reserva := range reservas {
        pdf.Cell(40, 10, "Reserva ID: "+strconv.Itoa(reserva.ID)+", Pasajero ID: "+strconv.Itoa(reserva.PasajeroID)+", Vuelo ID: "+strconv.Itoa(reserva.VueloID)+", Fecha: "+reserva.Fecha)
        pdf.Ln(10)
    }

    var buf bytes.Buffer
    err := pdf.Output(&buf)
    if err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}
