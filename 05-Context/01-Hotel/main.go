package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	bookHotel(ctx)
	ctx2 := context.Background()
	ctx2, cancel2 := context.WithTimeout(ctx2, time.Second*1)
	defer cancel2()
	bookHotel(ctx2)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached.")
		return
	case <-time.After(2 * time.Second):
		fmt.Println("Hotel booked.")
		return

	}

}
