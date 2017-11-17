package notifier

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/url"

	"github.com/barnybug/go-cast"
	"github.com/barnybug/go-cast/controllers"
)

// Notifier is a google-home-notifier-go client
type Notifier struct {
	client *cast.Client
	ctx    context.Context
}

// NewClient makes a connection and create a client
func NewClient(ctx context.Context, host string, port int) (*Notifier, error) {
	ips, err := net.LookupIP(host)
	if err != nil {
		return nil, err
	}
	client := cast.NewClient(ips[0], port)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	log.Println("Connected")
	n := &Notifier{client: client, ctx: ctx}
	return n, nil
}

// Notify sends a message to google home
func (n *Notifier) Notify(text string, language string) error {
	baseURL := "https://translate.google.com/translate_tts?ie=UTF-8&q=%s&tl=%s&client=tw-ob"
	u := fmt.Sprintf(baseURL, url.QueryEscape(text), url.QueryEscape(language))
	return n.Play(u)
}

//Play sound via URL
func (n *Notifier) Play(url string) error {
	media, err := n.client.Media(n.ctx)
	if err != nil {
		return err
	}
	contentType := "audio/mpeg"
	item := controllers.MediaItem{
		ContentId:   url,
		StreamType:  "BUFFERED",
		ContentType: contentType,
	}
	_, err = media.LoadMedia(n.ctx, item, 0, true, map[string]interface{}{})
	return err
}

//Stop sound
func (n *Notifier) Stop() error {
	if !n.client.IsPlaying(n.ctx) {
		return nil
	}
	media, err := n.client.Media(n.ctx)
	if err != nil {
		return err
	}
	_, err = media.Stop(n.ctx)
	return err
}

// Quit
func (n *Notifier) Quit() error {
	receiver := n.client.Receiver()
	_, err := receiver.QuitApp(n.ctx)
	return err
}

// Close connection
func (n *Notifier) Close() {
	n.client.Close()
}
