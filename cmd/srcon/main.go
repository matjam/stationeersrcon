package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"

	"github.com/jaytaylor/html2text"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	app             = kingpin.New("srcon", "a CLI interface to Stationeers dedicated server RCON.")
	serverIP        = app.Flag("ip", "Server IP to connect to.").Default("127.0.0.1").IP()
	port            = app.Flag("port", "Port to connect to").Default("27500").String()
	rconPassword    = app.Flag("password", "Password to use for the RCON command.").Required().String()
	statusCmd       = app.Command("status", "Fetch the status of the Stationeers server.")
	saveCmd         = app.Command("save", "Saves the game to a specified file.")
	saveCmdArg      = saveCmd.Arg("savefile", "Filename to save the game to.").Required().String()
	shutdownCmd     = app.Command("shutdown", "Shutdowns the server")
	shutdownMsg     = shutdownCmd.Flag("message", "message to send to all players").Short('m').String()
	shutdownTimeout = shutdownCmd.Flag("timeout", "timeout in seconds").Short('t').String()
	noticeCmd       = app.Command("notice", "Sends a notice to all players")
	noticeCmdArg    = noticeCmd.Arg("message", "Message to send.").Required().String()
	banCmd          = app.Command("ban", "Bans a player for a specific time.")
	banPlayer       = banCmd.Arg("steamId", "Steam ID to ban").Required().String()
	banTimeout      = banCmd.Arg("timeout", "How long to ban for. Timeout is a double in hours. 0.5 is 30 minutes, 0 is infinite.").Required().String()
	unbanCmd        = app.Command("unban", "Remove a player from the ban list.")
	unbanPlayer     = unbanCmd.Arg("steamID", "Steam ID to remove from the ban").Required().String()
	kickCmd         = app.Command("kick", "Kick player from the server.")
	kickPlayer      = kickCmd.Arg("steamID", "Steam ID to kick.").Required().String()
	clearallCmd     = app.Command("clearall", "Delete disconnected players from the server.")
)

func main() {
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: cookieJar,
	}

	cmd := kingpin.MustParse(app.Parse(os.Args[1:]))

	rconLogin(client)

	switch cmd {
	case statusCmd.FullCommand():
		rconStatus(client)
	case saveCmd.FullCommand():
		rconSave(client)
	case shutdownCmd.FullCommand():
		rconShutdown(client)
	case noticeCmd.FullCommand():
		rconNotice(client)
	case banCmd.FullCommand():
		rconBan(client)
	case unbanCmd.FullCommand():
		rconUnban(client)
	case kickCmd.FullCommand():
		rconKick(client)
	case clearallCmd.FullCommand():
		rconClearall(client)
	}

}

func rconStatus(c *http.Client) {
	rconExec(c, "status")
}

func rconSave(c *http.Client) {
	rconExec(c, fmt.Sprintf("save %s", *saveCmdArg))
}

func rconShutdown(c *http.Client) {
	command := fmt.Sprintf("shutdown")

	if len(*shutdownMsg) > 0 {
		command += fmt.Sprintf(" -m \"%s\"", *shutdownMsg)
	}

	if len(*shutdownTimeout) > 0 {
		command += fmt.Sprintf(" -t %s", *shutdownTimeout)
	}

	rconExec(c, command)
}

func rconNotice(c *http.Client) {
	rconExec(c, fmt.Sprintf("notice \"%s\"", *noticeCmdArg))
}

func rconBan(c *http.Client) {
	rconExec(c, fmt.Sprintf("ban %s %s", *banPlayer, *banTimeout))
}

func rconUnban(c *http.Client) {
	rconExec(c, fmt.Sprintf("unban %s", *unbanPlayer))
}

func rconKick(c *http.Client) {
	rconExec(c, fmt.Sprintf("kick %s", *kickPlayer))
}

func rconClearall(c *http.Client) {
	rconExec(c, "clearall")
}

func rconLogin(c *http.Client) {
	escapedCommand := url.PathEscape(fmt.Sprintf("login %s", *rconPassword))

	request := fmt.Sprintf("http://%s:%s/console/run?command=%s", serverIP.String(), *port, escapedCommand)
	_, err := c.Get(request)

	if err != nil {
		fmt.Println("error: ", err.Error())
		os.Exit(1)
	}
}

func rconExec(c *http.Client, command string) {
	escapedCommand := url.PathEscape(command)

	request := fmt.Sprintf("http://%s:%s/console/run?command=%s", serverIP.String(), *port, escapedCommand)
	resp, err := c.Get(request)

	if err != nil {
		fmt.Println("error: ", err.Error())
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)
	text, err := html2text.FromString(string(body), html2text.Options{PrettyTables: true})

	fmt.Println(string(text))
}
