package vindinium

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func (c *Client) post(uri string, params map[string]string, timeout time.Duration) error {
	p := url.Values{}
	for key := range params {
		p.Set(key, params[key])
	}

	dialFunc := func(network, addr string) (net.Conn, error) {
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
	fmt.Println("Connecting to " + c.url)
	if c.mode == "arena" {
		fmt.Println("Waiting for other players to join...")
	}
	uri := c.url + "/api/" + c.mode
	params := map[string]string{
		"key":   c.key,
		"turns": strconv.Itoa(c.turns),
		"map":   "m1",
	}
	// client should timeout after 10 minutes
	return c.post(uri, params, time.Duration(10*time.Minute))
}

func (c *Client) move(direction string) error {
	params := map[string]string{
		"key": c.key,
		"dir": direction,
	}
	return c.post(c.state.PlayURL, params, time.Duration(15*time.Second))
}

func (c *Client) isFinished() bool {
	return c.state.Game.Finished
}

func (c *Client) Play() {
	fmt.Println("Playing at " + c.state.ViewURL)
	var bot Bot
	switch c.bot {
	case "fighter":
		bot = FighterBot{}
	case "slow":
		bot = SlowBot{}
	default:
		bot = RandomBot{}
	}
	i := 0
	for {
    direction := bot.Move(c.state)
		if err := c.move(direction); err != nil {
			break
		}
		if c.isFinished() {
			break
		}
		i++
		fmt.Print("\rTaking turn " + strconv.Itoa(i))
	}
	fmt.Println("\rFinished " + strconv.Itoa(i) + " turn(s)")
}
