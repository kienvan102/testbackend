package models


type Ticket struct {
    ID       int    `json:"id"`
    Title    string `json:"title"`
    Status   string `json:"status"` // e.g., "To Do", "In Progress", "Done"
    Priority string `json:"priority"`
}


