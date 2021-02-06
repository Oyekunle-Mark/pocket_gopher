package main

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
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

var (
	authSetupOnce sync.Once
	httpClient    *http.Client
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

func makeRequest(req *http.Request, params url.Values) (*http.Response, error) {
	authSetupOnce.Do(func() {
		setupTwitterAuth()
		httpClient = &http.Client{
			Transport: &http.Transport{
				Dial: dial,
			},
		}
	})

	formEnc := params.Encode()
	req.Header.Set("Content-Type", "application/x-www-form- urlencoded")
	req.Header.Set("Content-Length", strconv.Itoa(len(formEnc)))
	req.Header.Set("Authorization", authClient.AuthorizationHeader(creds, "POST",
		req.URL, params))

	return httpClient.Do(req)
}

type tweet struct {
	Text string
}

func readFromTwitter(votes chan<- string) {
	options, err := loadOption()

	if err != nil {
		log.Println("failed to load options:", err)
		return
	}

	u, err := url.Parse("https://stream.twitter.com/1.1/statuses/filter.json")

	if err != nil {
		log.Println("creating filter request failed:", err)
		return
	}

	query := make(url.Values)
	query.Set("track", strings.Join(options, ","))

	req, err := http.NewRequest("POST", u.String(), strings.NewReader(query.Encode()))

	if err != nil {
		log.Println("creating filter request failed:", err)
		return
	}

	response, err := makeRequest(req, query)

	if err != nil {
		log.Println("making request failed:", err)
		return
	}

	reader := response.Body
	decoder := json.NewDecoder(reader)

	for {
		var t tweet

		if err := decoder.Decode(&t); err != nil {
			break
		}

		for _, option := range options {
			if strings.Contains(
				strings.ToLower(t.Text),
				strings.ToLower(option)) {
				log.Println("vote:", option)
				votes <- option
			}
		}
	}
}
