package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "value")
	bookHotel(ctx)
}

// convenção, sempre colocar o ctx como primeiro parametro da funcao.
func bookHotel(ctx context.Context) {
	token := ctx.Value("key")
	fmt.Println(token)
}
