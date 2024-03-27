package main

import (
	"database/sql"
	"log"

	"github.com/Gaghyta/BackendIIIFinalGO/cmd/server/handler"
	odontologoStore "github.com/Gaghyta/BackendIIIFinalGO/pkg/store"
	pacienteStore "github.com/Gaghyta/BackendIIIFinalGO/pkg/store"
	turnoStore "github.com/Gaghyta/BackendIIIFinalGO/pkg/store"
	"github.com/gin-gonic/gin"

	odontologos "github.com/Gaghyta/BackendIIIFinalGO/internal/odontologos"
	pacientes "github.com/Gaghyta/BackendIIIFinalGO/internal/pacientes"
	"github.com/Gaghyta/BackendIIIFinalGO/internal/turnos"
	_ "github.com/go-sql-driver/mysql"

	"github.com/Gaghyta/BackendIIIFinalGO/pkg/middleware"
)

type Config struct {
	DBUsername string `json:"db_username"`
	DBPassword string `json:"db_password"`
	DBHost     string `json:"db_host"`
	DBPort     string `json:"db_port"`
	DBName     string `json:"db_name"`
}

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/turnos-odontologia")
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
		odontologos.GET("/:odontologo_id", odontologoHandler.GetByID())
		odontologos.POST("", middleware.Authenticate(), odontologoHandler.Post())
		odontologos.DELETE("/:odontologo_id", middleware.Authenticate(), odontologoHandler.DeleteByID())
		odontologos.PATCH("/:odontologo_id", middleware.Authenticate(), odontologoHandler.Patch())
		odontologos.PUT("/:odontologo_id", middleware.Authenticate(), odontologoHandler.Put())
	}

	storagePaciente := pacienteStore.NewPacienteSqlStore(db)
	repoPacientes := pacientes.NewRepository(storagePaciente)
	servicePacientes := pacientes.NewService(repoPacientes)
	pacienteHandler := handler.NewPacienteHandler(servicePacientes)

	pacientes := r.Group("/pacientes")

	{
		pacientes.GET(":paciente_id", pacienteHandler.GetByID())
		pacientes.POST("", middleware.Authenticate(), pacienteHandler.Post())
		pacientes.DELETE(":paciente_id", middleware.Authenticate(), pacienteHandler.DeleteByID())
		pacientes.PATCH(":paciente_id", middleware.Authenticate(), pacienteHandler.Patch())
		pacientes.PUT(":paciente_id", middleware.Authenticate(), pacienteHandler.Put())
	}

	storageTurno := turnoStore.NewTurnoSqlStore(db)
	repoTurno := turnos.NewRepository(storageTurno)
	serviceTurnos := turnos.NewService(repoTurno)
	turnosHandler := handler.NewTurnoHandler(serviceTurnos, servicePacientes, serviceOdontologos)

	turnos := r.Group("/turnos")

	{
		//turnos.GET(":turno_id", turnosHandler.GetByID())
		turnos.GET(":dni", turnosHandler.GetByDNI())
		turnos.POST("", middleware.Authenticate(), turnosHandler.Post())
		turnos.POST("/dni-matricula", middleware.Authenticate(), turnosHandler.PostWithDniAndMatricula())
		turnos.DELETE(":turno_id", middleware.Authenticate(), turnosHandler.DeleteByID())
		turnos.PATCH(":turno_id", middleware.Authenticate(), turnosHandler.Patch())
		turnos.PUT(":turno_id", middleware.Authenticate(), turnosHandler.Put())
	}

	r.Run(":8080")

}
