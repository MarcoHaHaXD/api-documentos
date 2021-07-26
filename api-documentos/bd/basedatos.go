package bd

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Err     error
	db      *sql.DB
	SignKey = []byte("esta clave hay que cambiarla")
)

func NuevaConexionBD() *sql.DB {
	//db, Err = sql.Open("mysql", "root:@tcp(localhost)/documentos")
	//db, Err = sql.Open("mysql", "project-documentos:Documentos_in06@tcp(189.236.90.166:3306)/Documentacion")
	//db, Err = sql.Open("mysql", "project-documentos:Documentos_in06@tcp(189.236.109.161:3306)/Documentacion")
	db, Err = sql.Open("mysql", "intel:Intel06!@tcp(74.208.31.248:3306)/Documentacion")
	//Err = db.Ping()
	revisarError(Err)
	return db
}

func CerrarBD() {
	db.Close()
}

func revisarError(err error) (pass bool) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}
