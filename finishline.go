package main

import (
	"fmt"
	"log"
	"os"
	// "os/exec"

	"github.com/mitsuse/pushbullet-go"
	"github.com/mitsuse/pushbullet-go/requests"
)

func main() {
	// Set up the pushbullet token
	token := os.Getenv("PB_TOKEN")
	if token == "" {
		log.Fatal("PB_TOKEN is undefined. Please set your Pushbullet API token.")
	}

	hn, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	n := requests.NewNote()
	n.Title = "Job's done!"
	n.Body = fmt.Sprintf("Job on %s completed", hn)

	pb := pushbullet.New(token)

	// Send the note via Pushbullet.
	if _, err := pb.PostPushesNote(n); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		return
	}
}
