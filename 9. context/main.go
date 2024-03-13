package main

import (
	"context"
	"fmt"
	"time"
)

func doAnother(ctx context.Context, printCh <- chan int) {
	for {
		select {
		case <- ctx.Done() : 
			if err := ctx.Err() ; err != nil {
				fmt.Printf("do Another error: %s\n", err)
			}
			fmt.Println("do Another finished")
			return;

		case num := <- printCh :
			fmt.Println("do Another ", num)
		}
	}
//	fmt.Printf("The do another function, value : %s\n", ctx.Value("myKey"))
}

func doSomething(ctx context.Context) {
	ctx, cancelCtx := context.WithCancel(ctx)

	printCh := make(chan int)
	go doAnother(ctx, printCh)

	for i := 1; i <= 3; i++ {
		printCh <- i
	}

	cancelCtx()
	time.Sleep(300 * time.Millisecond)

	fmt.Println("Do something finished")

//	fmt.Println(time.Now())
//
//	fmt.Printf("Value of myKey : %s\n", ctx.Value("myKey"))
//	doAnotherCtx := context.WithValue(ctx, "myKey", "AnotherValue")
//
//	doAnother(doAnotherCtx)
//	fmt.Printf("Value of myKey : %s\n", ctx.Value("myKey"))
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "myKey", "myValue")

	doSomething(ctx)
}


