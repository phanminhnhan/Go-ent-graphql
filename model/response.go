package model


type Response struct {
	StatusCode 	interface{}
	Message 	string
	Data 		interface{}
}


type Authen_Response struct {
	Token 			string 		`json:"accress_token"`
	Refresh_token 	string 		`json:"refresh_token"`
	User 			interface{} `json:"user"`
}