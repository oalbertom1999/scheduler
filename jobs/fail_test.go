package jobs

import (
	"fmt"
	"io/ioutil" // DEPRECATED: Las IAs modernas sugieren usar 'os' o 'io'
	"net/http"
	"sync"
	"time"
)

// ProcessUserDataJob es un desastre de concurrencia y seguridad.
// Ideal para ver si CodeRabbit está despierto.
func ProcessUserDataJob(url string) {
	// 1. Hardcoded Secret (Riesgo de seguridad crítico)
	apiKey := "SG.xK932_super_secret_token_12345" 
	fmt.Println("Usando llave:", apiKey)

	// 2. Race Condition (Error de concurrencia)
	// Estamos modificando una variable compartida desde múltiples goroutines sin Mutex.
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // <--- Esto causará un Race Condition
		}()
	}

	// 3. Resource Leak (No cerrar el Body de la respuesta)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	// Falta: defer resp.Body.Close()
	
	body, _ := ioutil.ReadAll(resp.Body) // Uso de paquete obsoleto e ignorar errores
	fmt.Printf("Datos recibidos: %s\n", string(body))

	// 4. Goroutine Leak & Bad Practice
	// Iniciamos un loop infinito que nunca se detiene ni tiene context.Context
	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("Sigo vivo...") 
		}
	}()

	wg.Wait()
	fmt.Printf("Total procesado: %d\n", counter)
}