package bd

import (
	"api-documentos/modelos"
)

//Insertar Orden de Compra
func InserOrden(orden modelos.OrdenCompra)(respuesta modelos.RespuestaBasica){
	query, err:= db.Prepare("call usp_InsOrdenCompra(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)") //Son 33 + 28 = 61 Signos 
	if revisarError(err) {
		respuesta= modelos.RespuestaBasica{Codigo: 400, Mensaje: "Hubo un error al conectar con la base de datos"}
		return
	}
	_, err = query.Exec(
		orden.Fecha,
		orden.NoOrden,
		orden.FechaEntrega,
		orden.Proveedor,
		orden.EjecutivoP,
		orden.TelefonoP,
		orden.EmailP,
		orden.ClienteNoP,
		orden.DireccionP,
		orden.ColoniaP,
		orden.CiudadP,
		orden.CpP,
		orden.TelefonoPe,
		orden.FacturaA,
		orden.RfcF,
		orden.DireccionF,
		orden.ColoniaF,
		orden.CiudadF,
		orden.CpF,
		orden.TelefonoF,
		orden.LugarEntrega,
		orden.ClaveProducto,
		orden.KGLT,
		orden.Descripcion,
		orden.Cantidad,
		orden.PrecioUnit,
		orden.Importe,
		orden.Descuento,

		orden.ClaveProducto2,
		orden.KGLT2,
		orden.Descripcion2,
		orden.Cantidad2,
		orden.PrecioUnit2,
		orden.Importe2,
		orden.Descuento2,
		orden.ClaveProducto3,
		orden.KGLT3,
		orden.Descripcion3,
		orden.Cantidad3,
		orden.PrecioUnit3,
		orden.Importe3,
		orden.Descuento3,
		orden.ClaveProducto4,
		orden.KGLT4,
		orden.Descripcion4,
		orden.Cantidad4,
		orden.PrecioUnit4,
		orden.Importe4,
		orden.Descuento4,
		orden.ClaveProducto5,
		orden.KGLT5,
		orden.Descripcion5,
		orden.Cantidad5,
		orden.PrecioUnit5,
		orden.Importe5,
		orden.Descuento5,

		orden.TotalKGLT,
		orden.Subtotal,
		orden.Descuentos,
		orden.IVA,
		orden.Total,
	)
	if revisarError(err) {
		respuesta= modelos.RespuestaBasica{Codigo: 400, Mensaje: "Hubo un error al ejecutar la sentencia en la Base de Datos"}
		return
	}
	respuesta= modelos.RespuestaBasica{Codigo: 200, Mensaje: "Inserción realizada correctamente"}
	return
}

//Seleccionar todas Ordenes
func SelectOrdenes() (respuesta modelos.RespuestaGetOrdenes) {
	var (
		orden   modelos.OrdenCompra
		ordenes []modelos.OrdenCompra
	)
	if Err != nil {
		respuesta = modelos.RespuestaGetOrdenes{Codigo: 400, Mensaje: "Hubo un error al conectar con la base de datos", Error: Err.Error()}
		return
	}
	query, err := db.Query("SELECT * FROM ordencompramx")
	if err != nil {
		respuesta = modelos.RespuestaGetOrdenes{Codigo: 400, Mensaje: "Hubo un error en la consulta", Error: Err.Error()}
		return
	}
	for query.Next() {
		Err = query.Scan(
			&orden.IdOrden,
			&orden.Fecha,
			&orden.NoOrden,
			&orden.FechaEntrega,
			&orden.Proveedor,
			&orden.EjecutivoP,
			&orden.TelefonoP,
			&orden.EmailP,
			&orden.ClienteNoP,
			&orden.DireccionP,
			&orden.ColoniaP,
			&orden.CiudadP,
			&orden.CpP,
			&orden.TelefonoPe,
			&orden.FacturaA,
			&orden.RfcF,
			&orden.DireccionF,
			&orden.ColoniaF,
			&orden.CiudadF,
			&orden.CpF,
			&orden.TelefonoF,
			&orden.LugarEntrega,
			&orden.ClaveProducto,
			&orden.KGLT,
			&orden.Descripcion,
			&orden.Cantidad,
			&orden.PrecioUnit,
			&orden.Importe,
			&orden.Descuento,

			&orden.ClaveProducto2,
			&orden.KGLT2,
			&orden.Descripcion2,
			&orden.Cantidad2,
			&orden.PrecioUnit2,
			&orden.Importe2,
			&orden.Descuento2,
			&orden.ClaveProducto3,
			&orden.KGLT3,
			&orden.Descripcion3,
			&orden.Cantidad3,
			&orden.PrecioUnit3,
			&orden.Importe3,
			&orden.Descuento3,
			&orden.ClaveProducto4,
			&orden.KGLT4,
			&orden.Descripcion4,
			&orden.Cantidad4,
			&orden.PrecioUnit4,
			&orden.Importe4,
			&orden.Descuento4,
			&orden.ClaveProducto5,
			&orden.KGLT5,
			&orden.Descripcion5,
			&orden.Cantidad5,
			&orden.PrecioUnit5,
			&orden.Importe5,
			&orden.Descuento5,

			&orden.TotalKGLT,
			&orden.Subtotal,
			&orden.Descuentos,
			&orden.IVA,
			&orden.Total,
		)
		ordenes = append(ordenes, orden)
		if Err != nil {
			respuesta = modelos.RespuestaGetOrdenes{Codigo: 400, Mensaje: "Hubo un error al recibir un dato", Error: Err.Error()}
			return
		}
	}
	if len(ordenes) == 0 {
		respuesta = modelos.RespuestaGetOrdenes{Codigo: 201, Mensaje: "No se encontraron registros en la base de datos", Error: Err.Error()}
		return
	}
	respuesta = modelos.RespuestaGetOrdenes{Codigo: 200, Mensaje: "Datos enviados correctamente", Datos: ordenes}
	return
  }
