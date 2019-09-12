package entity

// User User
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// UserGetRequest get request
type UserGetRequest struct {
	ID string `json:"id"`
}

// UserResponse reponse
type UserResponse struct {
	User User `json:"User"`
}

// NewUserRequest new request
func NewUserRequest(id string) *UserGetRequest {
	return &UserGetRequest{ID: id}
}

// NewUserResponse new response
func NewUserResponse(User *User) *UserResponse {
	return &UserResponse{User: *User}
}
