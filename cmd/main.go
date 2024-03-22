package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/Gaghyta/BackendIIIFinalGO/cmd/server/handler"
	"github.com/Gaghyta/BackendIIIFinalGO/pkg/store"

	"github.com/Gaghyta/BackendIIIFinalGO/internal/odontologos"
	//"github.com/go-sql-driver/mysql"
	//"github.com/Gaghyta/BackendIIIFinalGO/internal/pacientes"
)

type Config struct {
	DBUsername string `json:"db_username"`
	DBPassword string `json:"db_password"`
	DBHost     string `json:"db_host"`
	DBPort     string `json:"db_port"`
	DBName     string `json:"db_name"`
}

func main() {

	/*bd, err := sql.Open("mysql", "root:yokit@11@tcp(localhost:3306)/turnero_odontologos")
	if err != nil {
		log.Fatal(err)
	}*/

	configFile, err := os.Open("/Users/gaghy/Desktop/BackendIIIFinalGO/config/config.json")
	if err != nil {
		log.Fatal("Error abriendo el archivo de configuración:", err)
	}
	defer configFile.Close()

	// Decodificar el archivo de configuración en una estructura Config
	var config Config
	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		log.Fatal("Error decodificando el archivo de configuración:", err)
	}

	// Construir la cadena de conexión
	dbConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	// Abrir la conexión a la base de datos
	db, err := sql.Open("mysql", dbConnectionString)
	if err != nil {
		log.Fatal("Error conectando a la base de datos:", err)
	}
	defer db.Close()

	storage := store.NewSqlStore(db)
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
		odontologos.GET(":odontologo_id", odontologoHandler.GetByID())
		odontologos.POST("", odontologoHandler.Post())
		//odontologos.DELETE(":id", odontologoHandler.Delete())
		//odontologos.PATCH(":id", odontologoHandler.Patch())
		//odontologos.PUT(":id", odontologoHandler.Put())
	}

	r.Run(":8080")

}
