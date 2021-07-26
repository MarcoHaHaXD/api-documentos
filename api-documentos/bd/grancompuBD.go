package bd

import (
	"api-documentos/modelos"
)

//Obtener todos las devoluciones
func TodosDevoluciones() (respuesta modelos.RespuestaGetDevoluciones) {
	var (
		devolucion   modelos.Devolucion
		devoluciones []modelos.Devolucion
	)
	if Err != nil {
		respuesta = modelos.RespuestaGetDevoluciones{Codigo: 400, Mensaje: "Hubo un error al conectar con la base de datos", Error: Err.Error()}
		return
	}
	query, err := db.Query("SELECT * FROM devolucion;")
	if err != nil {
		respuesta = modelos.RespuestaGetDevoluciones{Codigo: 400, Mensaje: "Hubo un error en la consulta", Error: Err.Error()}
		return
	}
	for query.Next() {
		Err = query.Scan(
			&devolucion.IdDevolución,
			&devolucion.Fecha,
			&devolucion.NoReferencia,
			&devolucion.EmpresaCompra,
			&devolucion.NombreApellidos,
			&devolucion.Direccion,
			&devolucion.Estado,
			&devolucion.Ciudad,
			&devolucion.Cp,
			&devolucion.Telefono,
			&devolucion.Email,
			&devolucion.DireccionOpc,
			&devolucion.EstadoOpc,
			&devolucion.CiudadOpc,
			&devolucion.CpOpc,
			&devolucion.Procedimiento,
			&devolucion.DescripcionEnt,
			&devolucion.MarcaEnt,
			&devolucion.ModeloEnt,
			&devolucion.SerieEnt,
			&devolucion.Cantidad,
			&devolucion.PrecioEnt,
			&devolucion.DiferenciaEnt,
			&devolucion.DescripcionSal,
			&devolucion.MarcaSal,
			&devolucion.ModeloSal,
			&devolucion.SerieSal,
			&devolucion.CantidadSal,
			&devolucion.PrecioSal,
			&devolucion.DiferenciaSal,
			&devolucion.Observaciones,
		)
		devoluciones = append(devoluciones, devolucion)
		if Err != nil {
			respuesta = modelos.RespuestaGetDevoluciones{Codigo: 400, Mensaje: "Hubo un error al recibir un dato", Error: Err.Error()}
			return
		}
	}
	if len(devoluciones) == 0 {
		respuesta = modelos.RespuestaGetDevoluciones{Codigo: 201, Mensaje: "No se encontraron registros en la base de datos", Error: Err.Error()}
		return
	}
	respuesta = modelos.RespuestaGetDevoluciones{Codigo: 200, Mensaje: "Datos enviados correctamente", Datos: devoluciones}
	return
}
func InserCorizacion( cotizacion modelos.Cotizacion)(respuesta modelos.RespuestaBasica){
	query, err:= db.Prepare("call usp_InsCotizacion(?,?,?,?,?)")
	if revisarError(err) {
		respuesta= modelos.RespuestaBasica{Codigo: 400, Mensaje: "Hubo un error al conectar con la base de datos"}
		return
	}
	_, err = query.Exec(
		cotizacion.Direccion,
		cotizacion.Telefonos,
		cotizacion.Correo,
		cotizacion.SitioWeb,
		cotizacion.EspVentas,
	)
	if revisarError(err) {
		respuesta= modelos.RespuestaBasica{Codigo: 400, Mensaje: "Hubo un error al ejecutar la sentencia en la Base de Datos"}
		return
	}
	respuesta= modelos.RespuestaBasica{Codigo: 200, Mensaje: "Inserción realizada correctamente"}
	return
}