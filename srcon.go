package srcon

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/jaytaylor/html2text"
)

// Client is a handle to a Stationeers rcon server.
type Client struct {
	*http.Client
	hostName string
	port     int
}

// Status fetches the status of the stationeers server.
func (c *Client) Status() (string, error) {
	return c.Exec("status")
}

// Save saves the game to the given file.
func (c *Client) Save(file string) (string, error) {
	return c.Exec(fmt.Sprintf("save %s", file))
}

// Shutdown terminates the stationeers server immediately.
func (c *Client) Shutdown(message string, timeout int) (string, error) {
	command := fmt.Sprintf("shutdown")

	if len(message) > 0 {
		command += fmt.Sprintf(" -m \"%s\"", message)
	}

	if timeout > 0 {
		command += fmt.Sprintf(" -t %d", timeout)
	}

	return c.Exec(command)
}

// Notice sends a message to all connected clients.
func (c *Client) Notice(message string) (string, error) {
	return c.Exec(fmt.Sprintf("notice \"%s\"", message))
}

// Ban a given player for a timeout number of seconds.
func (c *Client) Ban(player string, timeout int) (string, error) {
	return c.Exec(fmt.Sprintf("ban %s %d", player, timeout))
}

// Unban a given player.
func (c *Client) Unban(player string) (string, error) {
	return c.Exec(fmt.Sprintf("unban %s", player))
}

// Kick a given player.
func (c *Client) Kick(player string) (string, error) {
	return c.Exec(fmt.Sprintf("kick %s", player))
}

// ClearAll disconnected players.
func (c *Client) ClearAll() (string, error) {
	return c.Exec("clearall")
}

// HungerRate sets the global hunger rate for the server.
func (c *Client) HungerRate(rate int) (string, error) {
	return c.Exec(fmt.Sprintf("hungerRate %d", rate))
}

// New will create a new connection handle for connecting to a Stationeers RCON server.
func New(hostName string, port int) *Client {
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: cookieJar,
	}

	c := Client{client, hostName, port}
	return &c
}

// Login to the stationeers rcon server.
func (c *Client) Login(rconPassword string) error {

	escapedCommand := url.PathEscape(fmt.Sprintf("login %s", rconPassword))

	request := fmt.Sprintf("http://%s:%d/console/run?command=%s", c.hostName, c.port, escapedCommand)
	_, err := c.Get(request)

	if err != nil {
		return err
	}

	return nil
}

// Exec executes a command on the stationeers RCON.
func (c *Client) Exec(command string) (string, error) {
	escapedCommand := url.PathEscape(command)

	request := fmt.Sprintf("http://%s:%d/console/run?command=%s", c.hostName, c.port, escapedCommand)
	resp, err := c.Get(request)

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	text, err := html2text.FromString(string(body), html2text.Options{PrettyTables: true})

	return string(text), nil
}
