package jobs

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// User representa la estructura de datos limpia
type User struct {
	ID    int
	Email string
}

// CleanProcessorJob demuestra buenas prácticas de concurrencia y recursos
type CleanProcessorJob struct {
	DB     *sql.DB
	UserID int
}

// Run ejecuta el proceso de forma segura
func (c *CleanProcessorJob) Run() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Printf("[%s] 🚀 Iniciando procesamiento del usuario %d\n", 
		time.Now().Format(time.RFC3339), c.UserID)

	query := "SELECT id, email FROM users WHERE id = ?"
	row := c.DB.QueryRowContext(ctx, query, c.UserID)

	var u User
	if err := row.Scan(&u.ID, &u.Email); err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("ℹ️ No se encontró el usuario")
			return
		}
		fmt.Printf("❌ Error al consultar DB: %v\n", err)
		return
	}

	processData(u)
}

func processData(u User) {
	fmt.Printf("✅ Procesando datos para el usuario: %s\n", u.Email)
}