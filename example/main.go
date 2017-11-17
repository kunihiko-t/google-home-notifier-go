package main

import (
	"context"
	"github.com/kunihiko-t/google-home-notifier-go"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	//Set your google home's IP Address & port
	n, err := notifier.NewClient(ctx, "192.168.3.16", 8009)
	if err != nil {
		panic(err)
	}
	defer n.Close()
	//Play text
	n.Notify("あいうえお", "ja")
	time.Sleep(3 * time.Second)

	//Play music via URL
	err = n.Play("http://www.sample-videos.com/audio/mp3/crowd-cheering.mp3")
	if err != nil {
		log.Println(err)
	}
	time.Sleep(6 * time.Second)
	//Stop
	err = n.Stop()
	if err != nil {
		log.Println(err)
	}
	//Quit app
	err = n.Quit()
	if err != nil {
		log.Println(err)
	}
}
