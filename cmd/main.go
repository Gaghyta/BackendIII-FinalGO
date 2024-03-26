package main

import (
	"database/sql"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/Gaghyta/BackendIIIFinalGO/cmd/server/handler"
	odontologoStore "github.com/Gaghyta/BackendIIIFinalGO/pkg/store"
	pacienteStore "github.com/Gaghyta/BackendIIIFinalGO/pkg/store"
	turnoStore "github.com/Gaghyta/BackendIIIFinalGO/pkg/store"

	odontologos "github.com/Gaghyta/BackendIIIFinalGO/internal/odontologos"
	pacientes "github.com/Gaghyta/BackendIIIFinalGO/internal/pacientes"
	"github.com/Gaghyta/BackendIIIFinalGO/internal/turnos"
	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	DBUsername string `json:"db_username"`
	DBPassword string `json:"db_password"`
	DBHost     string `json:"db_host"`
	DBPort     string `json:"db_port"`
	DBName     string `json:"db_name"`
}

func main() {

	bd, err := sql.Open("mysql", "root:digital@tcp(localhost:3306)/turnos-odontologia")
	if err != nil {
		log.Fatal("Error conectando a la base de datos:", err)
	}
	defer bd.Close()

	storageOdontologo := odontologoStore.NewOdontologoSqlStore(bd)
	repoOdontologos := odontologos.NewRepository(storageOdontologo)
	serviceOdontologos := odontologos.NewService(repoOdontologos)
	odontologoHandler := handler.NewOdontologoHandler(serviceOdontologos)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	odontologos := r.Group("/odontologos")

	{
		odontologos.GET("/:odontologo_id", odontologoHandler.GetByID())
		odontologos.POST("", odontologoHandler.Post())
		odontologos.DELETE("/:odontologo_id", odontologoHandler.DeleteByID())
		odontologos.PATCH("/:odontologo_id", odontologoHandler.Patch())
		odontologos.PUT("/:odontologo_id", odontologoHandler.Put())
	}

	storagePaciente := pacienteStore.NewPacienteSqlStore(bd)
	repoPacientes := pacientes.NewRepository(storagePaciente)
	servicePacientes := pacientes.NewService(repoPacientes)
	pacienteHandler := handler.NewPacienteHandler(servicePacientes)

	pacientes := r.Group("/pacientes")

	{
		pacientes.GET(":paciente_id", pacienteHandler.GetByID())
		pacientes.POST("", pacienteHandler.Post())
		pacientes.DELETE(":paciente_id", pacienteHandler.DeleteByID())
		pacientes.PATCH(":paciente_id", pacienteHandler.Patch())
		pacientes.PUT(":paciente_id", pacienteHandler.Put())
	}

	storageTurno := turnoStore.NewTurnoSqlStore(bd)
	repoTurno := turnos.NewRepository(storageTurno)
	serviceTurnos := turnos.NewService(repoTurno)
	turnosHandler := handler.NewTurnoHandler(serviceTurnos,servicePacientes,serviceOdontologos)

	turnos := r.Group("/turnos")

	{
		//turnos.GET(":turno_id", turnosHandler.GetByID())
		turnos.GET(":dni", turnosHandler.GetByDNI())
		turnos.POST("", turnosHandler.Post())
		turnos.POST("/dni-matricula", turnosHandler.PostWithDniAndMatricula())
		turnos.DELETE(":turno_id", turnosHandler.DeleteByID())
		turnos.PATCH(":turno_id", turnosHandler.Patch())
		turnos.PUT(":turno_id", turnosHandler.Put())
	}

	r.Run(":8080")

}
