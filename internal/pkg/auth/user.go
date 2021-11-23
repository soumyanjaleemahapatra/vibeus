package auth

type User struct {
	ID         int    `json:"id,primary_key"`
	Password   string `json:"password"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}

var user = User{
	ID:       1,
	Password: "password",
}
