package v1

import "time"

type AddExpensePayload struct {
	Amount      float64   `json:"amount" binding:"required"`
	Category    string    `json:"category" binding:"required"`
	Description string    `json:"description"`
	Mode        string    `json:"mode"`
	Date        time.Time `json:"date"`
}
