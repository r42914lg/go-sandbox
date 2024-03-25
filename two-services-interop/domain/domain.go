package domain

type RequestBody struct {
	Value int `json:"value"`
}

type ServiceTwoResponse struct {
	StrValue string `json:"str_val"`
}

type ServiceOneRequestBody struct {
	IntValue int `json:"value"`
}
