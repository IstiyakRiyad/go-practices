package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	fmt.Println(time.Now())

	fmt.Printf("Value of myKey : %s\n", ctx.Value("myKey"))
	doAnotherCtx := context.WithValue(ctx, "myKey", "AnotherValue")

	doAnother(doAnotherCtx)
	fmt.Printf("Value of myKey : %s\n", ctx.Value("myKey"))
}

func doAnother(ctx context.Context) {
	fmt.Printf("The do another function, value : %s\n", ctx.Value("myKey"))
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "myKey", "myValue")

	doSomething(ctx)
}


