package dropbox

import (
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
)

type Client struct {
	token string
}

func New(token) {
	config := dropbox.Config{
		token
	}
}
