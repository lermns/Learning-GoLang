package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Tarea struct {
	ID          int    `json:"id"`
	Titulo      string `json:"titulo"`
	Descripcion string `json:"descripcion"`
	Completado  bool   `json:"completado"`
}

var tareas = []Tarea{}
var nextID = 1

func main() {
	router := gin.Default()

	// Configurar CORS para React
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	// Rutas
	router.GET("/tareas", obtenerTareas)
	router.GET("/tareas/:id", obtenerTarea)
	router.POST("/tareas", crearTarea)
	router.PUT("/tareas/:id", actualizarTarea)
	router.PATCH("/tareas/:id/completar", completarTarea)
	router.DELETE("/tareas/:id", eliminarTarea)

	router.Run(":8080")
}

// GET /tareas - Obtener todas las tareas
func obtenerTareas(c *gin.Context) {
	c.JSON(http.StatusOK, tareas)
}

// GET /tareas/:id - Obtener una tarea específica
func obtenerTarea(c *gin.Context) {
	id := c.Param("id")

	for _, tarea := range tareas {
		if string(rune(tarea.ID)) == id {
			c.JSON(http.StatusOK, tarea)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Tarea no encontrada"})
}

// POST /tareas - Crear nueva tarea
func crearTarea(c *gin.Context) {
	var nuevaTarea Tarea

	if err := c.BindJSON(&nuevaTarea); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	nuevaTarea.ID = nextID
	nextID++
	nuevaTarea.Completado = false

	tareas = append(tareas, nuevaTarea)
	c.JSON(http.StatusCreated, nuevaTarea)
}

// PUT /tareas/:id - Actualizar tarea completa
func actualizarTarea(c *gin.Context) {
	id := c.Param("id")
	var tareaActualizada Tarea

	if err := c.BindJSON(&tareaActualizada); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, tarea := range tareas {
		if string(rune(tarea.ID)) == id {
			tareaActualizada.ID = tarea.ID
			tareas[i] = tareaActualizada
			c.JSON(http.StatusOK, tareaActualizada)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Tarea no encontrada"})
}

// PATCH /tareas/:id/completar - Marcar como completada
func completarTarea(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Completado bool `json:"completado"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, tarea := range tareas {
		if string(rune(tarea.ID)) == id {
			tareas[i].Completado = body.Completado
			c.JSON(http.StatusOK, tareas[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Tarea no encontrada"})
}

// DELETE /tareas/:id - Eliminar tarea
func eliminarTarea(c *gin.Context) {
	id := c.Param("id")

	for i, tarea := range tareas {
		if string(rune(tarea.ID)) == id {
			tareas = append(tareas[:i], tareas[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Tarea eliminada"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Tarea no encontrada"})
}
