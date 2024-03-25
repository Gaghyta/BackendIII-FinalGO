package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/Gaghyta/BackendIIIFinalGO/cmd/server/handler"
	odontologoStore "github.com/Gaghyta/BackendIIIFinalGO/pkg/store"
	pacienteStore "github.com/Gaghyta/BackendIIIFinalGO/pkg/store"
	turnoStore "github.com/Gaghyta/BackendIIIFinalGO/pkg/store"

	odontologos "github.com/Gaghyta/BackendIIIFinalGO/internal/odontologos"
	pacientes "github.com/Gaghyta/BackendIIIFinalGO/internal/pacientes"
	"github.com/Gaghyta/BackendIIIFinalGO/internal/turnos"
	//"github.com/go-sql-driver/mysql"
)

type Config struct {
	DBUsername string `json:"db_username"`
	DBPassword string `json:"db_password"`
	DBHost     string `json:"db_host"`
	DBPort     string `json:"db_port"`
	DBName     string `json:"db_name"`
}

func main() {

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

	storageOdontologo := odontologoStore.NewOdontologoSqlStore(db)
	repoOdontologos := odontologos.NewRepository(storageOdontologo)
	serviceOdontologos := odontologos.NewService(repoOdontologos)
	odontologoHandler := handler.NewOdontologoHandler(serviceOdontologos)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	odontologos := r.Group("/odontologos")

	{
		odontologos.GET(":odontologo_id", odontologoHandler.GetByID())
		odontologos.POST("", odontologoHandler.Post())
		odontologos.DELETE(":odontologo_id", odontologoHandler.DeleteByID())
		odontologos.PATCH(":odontologo_id", odontologoHandler.Patch())
		odontologos.PUT(":odontologo_id", odontologoHandler.Put())
	}

	storagePaciente := pacienteStore.NewPacienteSqlStore(db)
	repoPacientes := pacientes.NewRepository(storagePaciente)
	servicePacientes := pacientes.NewService(repoPacientes)
	pacienteHandler := handler.NewPacienteHandler(servicePacientes)

	pacientes := r.Group("/pacientes")

	{
		pacientes.GET(":id", pacienteHandler.GetByID())
		pacientes.POST("", pacienteHandler.Post())
		pacientes.DELETE(":id", pacienteHandler.DeleteByID())
		pacientes.PATCH(":id", pacienteHandler.Patch())
		pacientes.PUT(":id", pacienteHandler.Put())
	}

	storageTurno := turnoStore.NewTurnoSqlStore(db)
	repoTurno := turnos.NewRepository(storageTurno)
	serviceTurnos := turnos.NewService(repoTurno)
	turnosHandler := handler.NewTurnoHandler(serviceTurnos)

	turnos := r.Group("/turnos")

	{
		turnos.GET(":id", turnosHandler.GetByID())
		turnos.POST("", turnosHandler.Post())
		turnos.DELETE(":id", turnosHandler.DeleteByID())
		turnos.PATCH(":id", turnosHandler.Patch())
		turnos.PUT(":id", turnosHandler.Put())
	}

	r.Run(":8080")

}
