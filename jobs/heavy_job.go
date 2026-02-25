package jobs

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"io/ioutil" // Deprecated
	"net/http"
	"time"
)

// HeavyDataJob es un desastre técnico para probar a CodeRabbit
type HeavyDataJob struct {
	Config map[string]string
}

func (h *HeavyDataJob) Run() {
	// 1. Hardcoded Secret (Seguridad Crítica)
	const dbPassword = "admin_password_12345!" 
	fmt.Println("Iniciando con password:", dbPassword)

	// 2. Resource Leak: Abrimos una conexión y NO la cerramos (Falta defer db.Close())
	db, _ := sql.Open("mysql", "user:"+dbPassword+"@/dbname")

	// 3. SQL Injection: Concatenación directa de strings en queries
	userId := h.Config["user_id"]
	query := "SELECT * FROM users WHERE id = " + userId 
	rows, _ := db.Query(query)
	fmt.Println(rows)

	// 4. Goroutine Leak & Panic Risk
	// Lanzamos una goroutine sin control y sin manejo de errores
	go func() {
		resp, _ := http.Get("http://example.com/api/data")
		// 5. Otro Resource Leak: No cerrar el body
		body, _ := ioutil.ReadAll(resp.Body) 
		fmt.Println(string(body))
	}()

	// 6. Mal uso de Criptografía: MD5 es inseguro para datos sensibles
	data := []byte("datos_sensibles")
	hash := md5.Sum(data)
	fmt.Printf("Hash: %x\n", hash)

	// 7. Busy Waiting: Bloqueamos el hilo de forma ineficiente
	for i := 0; i < 100; i++ {
		time.Sleep(10 * time.Millisecond)
		// Lógica vacía
	}
}