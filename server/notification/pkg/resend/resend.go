package resend

import (
	"github.com/resend/resend-go/v2"
	"qezde/notification/internal/config"
)

type Client struct {
	Client  *resend.Client
	Sender  string
	Subject string
}

func New(config config.ResendConfig) *Client {
	return &Client{
		Client:  resend.NewClient(config.APIKey),
		Sender:  config.Sender,
		Subject: config.Subject,
	}
}
