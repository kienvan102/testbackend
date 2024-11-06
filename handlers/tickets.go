// handlers/tickets.go
package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "testbackend/models"
    "testbackend/repositories"
    "database/sql"
)

var db *sql.DB

// NewTicketHandler initializes the handler with a database connection
func NewTicketHandler(database *sql.DB) {
    db = database
}

func GetTickets(c *gin.Context) {
    tickets, err := repositories.GetAllTickets(db)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tickets"})
        return
    }
    c.JSON(http.StatusOK, tickets)
}

func CreateTicket(c *gin.Context) {
    var ticket models.Ticket
    if err := c.ShouldBindJSON(&ticket); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := repositories.InsertTicket(db, &ticket); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create ticket"})
        return
    }

    c.JSON(http.StatusCreated, ticket)
}
