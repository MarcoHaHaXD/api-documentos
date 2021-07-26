package modelos

type Usuario struct {
	IdUsuario	int32  `json:"idUsuario"`
	Usuario		string `json:"usuario"`
	Contrasenia	string `json:"contrasenia"`
}