package main

import "time"

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 4)
		c1 <- 1
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- 2
	}()

	// o select vai aguardar quem chega primeiro, o ch1, ch2 ou se passar 3 segundos, timeout pelo 3.
	// podemos usar por exemplo para pegar o dados de duas apis e usar a que chegar primeiro.
	// ou dar timeout.
	select {
	case msg1 := <-c1:
		println("received", msg1)
	case msg2 := <-c2:
		println("received", msg2)
	case <-time.After(time.Second * 3):
		println("timeout")
		// default acaba sendo usado se os canais não tiverem dados.
		// default:
		// 	println("default")
	}

	// Exemplo de looping infinito recebendo mensagens do kafka e rabbitmq
	// for {
	// 	select {
	// 	case msg1 := <-c1: // kafka
	// 		println("received", msg1)
	// 	case msg2 := <-c2: // rabbitmq
	// 		println("received", msg2)
	// 	}
	// }
}
