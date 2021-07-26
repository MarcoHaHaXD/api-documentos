package modelos

type RespuestaError struct {
	Codigo  int    `json:"codigo"`
	Mensaje string `json:"mensaje"`
	Error   error  `json:"error"`
}

type RespuestaBasica struct {
	Codigo  int    `json:"codigo"`
	Mensaje string `json:"mensaje"`
	Error   string `json:"error"`
}