package model



type Course struct{
	Email         string `json:"email"`
	Name          string  `json:"name"`
}

type SignUpInput struct{
    Email	string `json:"email"`
	Id      int `json:"id"`
	SignupDate string `json:"signup_date"`
}