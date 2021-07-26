package modelos

type RespuestaGetOrdenes struct {
	Codigo  int    			`json:"codigo"`
	Mensaje string      	`json:"mensaje"`
	Datos   []OrdenCompra	`json:"datos"`
	Error   string      	`json:"error"`
}

type OrdenCompra struct{
	IdOrden		 	int32	`json:"idOrden"`
	Fecha			string	`json:"fecha"`
	NoOrden			int32	`json:"noOrden"`
	FechaEntrega	string	`json:"fechaEntrega"`
	Proveedor		string	`json:"proveedor"`
	EjecutivoP		string	`json:"ejecutivoP"`
	TelefonoP		string	`json:"telefonoP"`
	EmailP			string	`json:"emailP"`
	ClienteNoP		int32	`json:"clienteNoP"`
	DireccionP		string	`json:"direccionP"`
	ColoniaP		string	`json:"coloniaP"`
	CiudadP			string	`json:"ciudadP"`
	CpP				string	`json:"cpP"`
	TelefonoPe		string	`json:"telefonoPe"`
	FacturaA		string	`json:"facturaA"`
	RfcF			string	`json:"rfcF"`
	DireccionF		string	`json:"direccionF"`
	ColoniaF		string	`json:"coloniaF"`
	CiudadF			string	`json:"ciudadF"`
	CpF				string	`json:"cpF"`
	TelefonoF		string	`json:"telefonoF"`
	LugarEntrega	string	`json:"lugarEntrega"`
	ClaveProducto	string	`json:"claveProducto"`
	KGLT			string	`json:"kglt"`
	Descripcion		string	`json:"descripcion"`
	Cantidad		int32	`json:"cantidad"`
	PrecioUnit		float32	`json:"precioUnit"`
	Importe			float32	`json:"importe"`
	Descuento		float32	`json:"descuento"`

	ClaveProducto2	string	`json:"claveProducto2"`
	KGLT2			string	`json:"kglt2"`
	Descripcion2	string	`json:"descripcion2"`
	Cantidad2		int32	`json:"cantidad2"`
	PrecioUnit2		float32	`json:"precioUnit2"`
	Importe2		float32	`json:"importe2"`
	Descuento2		float32	`json:"descuento2"`
	ClaveProducto3	string	`json:"claveProducto3"`
	KGLT3			string	`json:"kglt3"`
	Descripcion3	string	`json:"descripcion3"`
	Cantidad3		int32	`json:"cantidad3"`
	PrecioUnit3		float32	`json:"precioUnit3"`
	Importe3		float32	`json:"importe3"`
	Descuento3		float32	`json:"descuento3"`
	ClaveProducto4	string	`json:"claveProducto4"`
	KGLT4			string	`json:"kglt4"`
	Descripcion4	string	`json:"descripcion4"`
	Cantidad4		int32	`json:"cantidad4"`
	PrecioUnit4		float32	`json:"precioUnit4"`
	Importe4		float32	`json:"importe4"`
	Descuento4		float32	`json:"descuento4"`
	ClaveProducto5	string	`json:"claveProducto5"`
	KGLT5			string	`json:"kglt5"`
	Descripcion5	string	`json:"descripcion5"`
	Cantidad5		int32	`json:"cantidad5"`
	PrecioUnit5		float32	`json:"precioUnit5"`
	Importe5		float32	`json:"importe5"`
	Descuento5		float32	`json:"descuento5"`

	TotalKGLT		int32	`json:"totalKGLT"`
	Subtotal		float32	`json:"subtotal"`
	Descuentos		float32	`json:"descuentos"`
	IVA				float32	`json:"iva"`
	Total			float32	`json:"total"`
}