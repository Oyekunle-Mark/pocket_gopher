package main

import (
	"io"
	"log"
	"net"
	"time"

	"github.com/joeshaw/envdecode"
	"github.com/matryer/go-oauth/oauth"
)

var conn net.Conn
var reader io.ReadCloser
var (
	authClient *oauth.Client
	creds      *oauth.Credentials
)

func dial(network, address string) (net.Conn, error) {
	if conn != nil {
		conn.Close()
		conn = nil
	}

	netConnection, err := net.DialTimeout(network, address, 5*time.Second)

	if err != nil {
		return nil, err
	}

	return netConnection, nil
}

func closeConnection() {
	if conn != nil {
		conn.Close()
	}

	if reader != nil {
		reader.Close()
	}
}

func setupTwitterAuth() {
	var ts struct {
		ConsumerKey    string `env:"SP_TWITTER_KEY,required"`
		ConsumerSecret string `env:"SP_TWITTER_SECRET,required"`
		AccessToken    string `env:"SP_TWITTER_ACCESSTOKEN,required"`
		AccessSecret   string `env:"SP_TWITTER_ACCESSSECRET,required"`
	}

	if err := envdecode.Decode(&ts); err != nil {
		log.Fatalln(err)
	}

	creds = &oauth.Credentials{
		Token:  ts.AccessToken,
		Secret: ts.AccessSecret,
	}

	authClient = &oauth.Client{
		Credentials: oauth.Credentials{
			Token:  ts.ConsumerKey,
			Secret: ts.ConsumerSecret,
		},
	}
}
