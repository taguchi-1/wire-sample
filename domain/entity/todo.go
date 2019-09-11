package entity

// Todo todo
type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// TodoGetRequest get request
type TodoGetRequest struct {
	ID string `json:"id"`
}

// TodoResponse reponse
type TodoResponse struct {
	Todo Todo `json:"todo"`
}

// NewTodoRequest new request
func NewTodoRequest(id string) *TodoGetRequest {
	return &TodoGetRequest{ID: id}
}

// NewTodoResponse new response
func NewTodoResponse(todo *Todo) *TodoResponse {
	return &TodoResponse{Todo: *todo}
}
