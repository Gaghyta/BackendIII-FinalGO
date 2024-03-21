package main

import (
	//"encoding/json"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/Gaghyta/BackendIIIFinalGO/pkg/store"

	"github.com/Gaghyta/BackendIIIFinalGO/cmd/handler"

	"github.com/Gaghyta/BackendIIIFinalGO/internal/odontologos"
	//"github.com/go-sql-driver/mysql"
	//"github.com/Gaghyta/BackendIIIFinalGO/internal/pacientes"
)

func main() {

	bd, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/turnero_odontologos")
	if err != nil {
		log.Fatal(err)
	}

	storage := store.NewSqlStore(bd)
	repo := odontologos.NewRepository(storage)
	service := odontologos.NewService(repo)
	odontologoHandler := handler.NewProductHandler(service)

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
