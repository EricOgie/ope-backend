package responsedto

type FlutterWaveResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    Link   `json:"link"`
}

type Link struct {
	Link string
}
