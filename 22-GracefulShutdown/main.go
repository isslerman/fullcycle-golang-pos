package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server := &http.Server{Addr: ":3030"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(4 * time.Second)
		w.Write([]byte("Hello Boss"))
	})

	// solta em paralelo o server.
	go func() {
		fmt.Println("Server is running at localhost:3000")
		if err := server.ListenAndServe(); err != nil && http.ErrServerClosed != err { // só dá o panic se o server closed ñ acontecer.
			log.Fatalf("Could not listen on %s: %v\n", server.Addr, err)
		}
	}()

	// cria um canal do tipo os.Signal com um buffer de tamanho 1.
	// O tipo os.Signal é frequentemente usado para capturar sinais do sistema operacional,
	// como sinais de interrupção (Ctrl+C) ou outros sinais relacionados ao término de um programa.
	stop := make(chan os.Signal, 1)
	// registra os canais que desejamos receber notificações do os.signal.
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	// programa para até receber um os.signal.
	<-stop

	// criamos um contexto com timeout de 5 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("Shutting down server...")
	// tenta dar um shutdown aguardando 5 segundos do contexto.
	if err := server.Shutdown(ctx); err != nil {
		// caso não consiga, log fatal
		log.Fatalf("Could not gracefully shutdown the server: %v \n", err)
	}
	// FIM
	fmt.Println("Server stopped")
}
