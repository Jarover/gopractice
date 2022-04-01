package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Jarover/gopractice/internal/app/config"
	"github.com/Jarover/gopractice/internal/app/listfile"
)

func main() {
	start := time.Now()
	log.Println(config.VersionStr())

	const wantExt = ".go"
	ctx := context.Background()
	ctx = context.WithValue(ctx, "deep", 2)
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1, syscall.SIGUSR2)

	//Обработать сигнал SIGUSR1
	waitCh := make(chan struct{})
	go func() {
		res, err := listfile.FindFiles(ctx, wantExt)
		if err != nil {
			log.Printf("Error on search: %v\n", err)
			os.Exit(1)
		}
		for _, f := range res {
			fmt.Printf("\tName: %s\t\t Path: %s\n", f.Name, f.Path)
		}
		waitCh <- struct{}{}
	}()

	go func() {
		s := <-sigCh

		switch s {
		case syscall.SIGINT, syscall.SIGTERM:
			log.Println("Signal received, terminate...", s)
			log.Println("Deep: " + ctx.Value("deep").(string))
			cancel()
		case syscall.SIGUSR1:
			fmt.Println("usr1", s)
		case syscall.SIGUSR2:
			fmt.Println("usr2", s)
		default:
			fmt.Println("other", s)

		}

	}()
	//Дополнительно: Ожидание всех горутин перед завершением
	<-waitCh

	log.Println("Done")

	log.Printf("%v: %v\n", "Время работы программы", time.Since(start))

}
