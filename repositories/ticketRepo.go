package repositories

import (
    "database/sql"
    "testbackend/models"
)

func GetAllTickets(db *sql.DB) ([]models.Ticket, error) {
    rows, err := db.Query("SELECT id, title, status, priority FROM tickets")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tickets []models.Ticket
    for rows.Next() {
        var ticket models.Ticket
        if err := rows.Scan(&ticket.ID, &ticket.Title, &ticket.Status, &ticket.Priority); err != nil {
            return nil, err
        }
        tickets = append(tickets, ticket)
    }
    return tickets, nil
}

func InsertTicket(db *sql.DB, ticket *models.Ticket) error {
    query := "INSERT INTO tickets (title, status, priority) VALUES ($1, $2, $3) RETURNING id"
    return db.QueryRow(query, ticket.Title, ticket.Status, ticket.Priority).Scan(&ticket.ID)
}
