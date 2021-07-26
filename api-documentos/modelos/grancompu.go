package modelos

type Devolucion struct{
	IdDevolución	int32 	`json:"idDevolucion"`
	Fecha			string	`json:"fecha"`
	NoReferencia	string 	`json:"noReferencia"`
	EmpresaCompra	string 	`json:"empresaCompra"`
	NombreApellidos	string 	`json:"nombreApellidos"`
	Direccion		string 	`json:"direccion"`
	Estado			string	`json:"estado"`
	Ciudad			string	`json:"ciudad"`
	Cp 				int		`json:"cp"`
	Telefono		string	`json:"telefono"`
	Email			string	`json:"email"`
	DireccionOpc	string 	`json:"direccionOpc"`
	EstadoOpc		string 	`json:"estadoOpc"`
	CiudadOpc		string 	`json:"ciudadOpc"`
	CpOpc			int 	`json:"cpOpc"`	
	Procedimiento	string 	`json:"procedimiento"`
	DescripcionEnt	string 	`json:"descripcionEnt"`
	MarcaEnt		string	`json:"marcaEnt"`
	ModeloEnt		string	`json:"modeloEnt"`
	SerieEnt		string	`json:"serieEnt"`
	Cantidad		int		`json:"cantidad"`
	PrecioEnt		float64	`json:"precioEnt"`
	DiferenciaEnt	int		`json:"diferenciaEnt"`
	DescripcionSal	string	`json:"descripcionSal"`
	MarcaSal	 	string 	`json:"marcaSal"`
	ModeloSal		string 	`json:"modeloSal"`
	SerieSal		string	`json:"serieSal"`
	CantidadSal		int		`json:"cantidadSal"`
	PrecioSal		float64	`json:"precioSal"`
	DiferenciaSal	int		`json:"diferenciaSal"`
	Observaciones	string	`json:"observaciones"`
}

type Cotizacion struct{
	IdGranCompu 	int32	`json:"idGranCompu"`
	Direccion		string	`json:"direccion"`
	Telefonos		string	`json:"telefonos"`
	Correo			string	`json:"correo"`
	SitioWeb		string	`json:"sitioweb"`
	EspVentas		string	`json:"espVentas"`
}