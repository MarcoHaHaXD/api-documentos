package modelos

type RespuestaGetDevoluciones struct {
	Codigo  int    			`json:"codigo"`
	Mensaje string      	`json:"mensaje"`
	Datos   []Devolucion	`json:"datos"`
	Error   string      	`json:"error"`
}
