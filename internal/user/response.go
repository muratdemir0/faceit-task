package user

type DefaultResponse struct {
	Data interface{} `json:"data"`
}

type Response struct {
	Users []User `json:"users"`
}

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Nickname  string `json:"nickname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Country   string `json:"country"`
}
