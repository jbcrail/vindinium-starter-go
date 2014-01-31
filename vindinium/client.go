package vindinium

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Client struct {
	url   string
	key   string
	mode  string
	bot   string
	turns int

	state State
}

func NewClient(url, key, mode, bot string, turns int) *Client {
	return &Client{
		url:   url,
		key:   key,
		mode:  mode,
		bot:   bot,
		turns: turns,
	}
}

func (c *Client) post(uri string, params map[string]string) error {
	p := url.Values{}
	for key := range params {
		p.Set(key, params[key])
	}

	// client should timeout after 10 minutes
	dialFunc := func(network, addr string) (net.Conn, error) {
		timeout := time.Duration(10 * time.Minute)
		return net.DialTimeout(network, addr, timeout)
	}

	transport := http.Transport{Dial: dialFunc}
	client := http.Client{Transport: &transport}

	response, err := client.PostForm(uri, p)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(bytes.NewReader(data))
	if err := decoder.Decode(&c.state); err != nil {
		return err
	}

	return nil
}

func (c *Client) Connect() error {
	uri := c.url + "/api/" + c.mode
	params := map[string]string{
		"key":   c.key,
		"turns": strconv.Itoa(c.turns),
		"map":   "m1",
	}
	return c.post(uri, params)
}

func (c *Client) move(direction string) {
}

func (c *Client) isFinished() bool {
	return c.state.Game.Finished
}

func (c *Client) Play() {
	for !c.isFinished() {
	}
}
