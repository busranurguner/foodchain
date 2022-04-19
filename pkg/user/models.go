package user

type GetAllRequest struct {
	Username string `json:"username" `
	Email    string `json:"email" `
	Role     string `json:"role" `
	Password string `json:"password" `
	Refresh  string `json:"refresh"`
}
type UpdateRequest struct {
	ID       string `bson:"_id"`
	Password string `json:"password"`
}

type DeleteRequest struct {
	ID string `bson:"_id"`
}
