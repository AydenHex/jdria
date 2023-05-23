package api

type TestDataResponse struct {
	Name      string `json:"name"`
	Firstname string `json:"firstName"`
}

type TestResponse struct {
	Message string           `json:"message"`
	Data    TestDataResponse `json:"data"`
}
