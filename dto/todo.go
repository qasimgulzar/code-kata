package dto

// Todo represents a task fetched from a TODO management service.
type Todo struct {
	UserId    int    `json:"userId"`    // UserId is the identifier for the user who owns the TODO item.
	Id        int    `json:"id"`        // Id is the unique identifier for the TODO item.
	Title     string `json:"title"`     // Title is the name of the TODO item.
	Completed bool   `json:"completed"` // Completed indicates whether the TODO item has been completed.
}
