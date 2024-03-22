package main

import (
	//"encoding/json"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/Gaghyta/BackendIIIFinalGO/pkg/store/odontologoStore"
	"github.com/Gaghyta/BackendIIIFinalGO/internal/odontologo"
	"github.com/Gaghyta/BackendIIIFinalGO/cmd/server/handler"

	_ "github.com/go-sql-driver/mysql"
	//"github.com/Gaghyta/BackendIIIFinalGO/internal/pacientes"
)

func main() {

	bd, err := sql.Open("mysql", "root:digital@tcp(localhost:3306)/turnero_odontologos")
	if err != nil {
		log.Fatal(err)
	}

	storageOdontologo := odontologoStore.NewSqlStore(bd)
	repoOdontologos := odontologo.NewRepository(storageOdontologo)
	serviceOdontologos := odontologo.NewService(repoOdontologos)
	odontologoHandler := handler.NewProductHandler(serviceOdontologos)


	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	//pacientes := r.Group("/pacientes")

	/* {
		pacientes.GET(":id", pacienteHandler.GetByID())
		pacientes.POST("", pacienteHandler.Post())
		pacientes.DELETE(":id", pacienteHandler.Delete())
		pacientes.PATCH(":id", pacienteHandler.Patch())
		pacientes.PUT(":id", pacienteHandler.Put())
	}
	*/
	odontologos := r.Group("/odontologos")

	{
		odontologos.GET(":id", odontologoHandler.GetByID())
		odontologos.POST("", odontologoHandler.Post())
		//odontologos.DELETE(":id", odontologoHandler.Delete())
		//odontologos.PATCH(":id", odontologoHandler.Patch())
		//odontologos.PUT(":id", odontologoHandler.Put())
	}

	r.Run(":8080")

}
