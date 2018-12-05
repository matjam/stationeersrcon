package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	srcon "github.com/matjam/stationeersrcon"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	app             = kingpin.New("srcon", "a CLI interface to Stationeers dedicated server RCON.")
	configFlag      = app.Flag("config", "Path to a json configuration file for the tool.").Short('c').String()
	serverIP        = app.Flag("ip", "Server IP to connect to.").IP()
	port            = app.Flag("port", "Port to connect to").String()
	rconPassword    = app.Flag("password", "Password to use for the RCON command.").String()
	statusCmd       = app.Command("status", "Fetch the status of the Stationeers server.")
	saveCmd         = app.Command("save", "Saves the game to a specified file.")
	saveCmdArg      = saveCmd.Arg("savefile", "Filename to save the game to.").Required().String()
	shutdownCmd     = app.Command("shutdown", "Shutdowns the server")
	shutdownMsg     = shutdownCmd.Flag("message", "message to send to all players").Short('m').String()
	shutdownTimeout = shutdownCmd.Flag("timeout", "timeout in seconds").Short('t').Int()
	noticeCmd       = app.Command("notice", "Sends a notice to all players")
	noticeCmdArg    = noticeCmd.Arg("message", "Message to send.").Required().String()
	banCmd          = app.Command("ban", "Bans a player for a specific time.")
	banPlayer       = banCmd.Arg("steamId", "Steam ID to ban").Required().String()
	banTimeout      = banCmd.Arg("timeout", "How long to ban for. Timeout is a double in hours. 0.5 is 30 minutes, 0 is infinite.").Required().Int()
	unbanCmd        = app.Command("unban", "Remove a player from the ban list.")
	unbanPlayer     = unbanCmd.Arg("steamID", "Steam ID to remove from the ban").Required().String()
	kickCmd         = app.Command("kick", "Kick player from the server.")
	kickPlayer      = kickCmd.Arg("steamID", "Steam ID to kick.").Required().String()
	clearallCmd     = app.Command("clearall", "Delete disconnected players from the server.")
	hungerRateCmd   = app.Command("hungerRate", "Set the hungerRate on the server to the given number.")
	hungerRate      = hungerRateCmd.Arg("rate", "The rate to use. Use 0 to disable hunger completely.").Required().Int()
)

var config *Config
var client *srcon.Client

func main() {
	var resp string
	var err error
	app.HelpFlag.Short('h')

	cmd := kingpin.MustParse(app.Parse(os.Args[1:]))

	client = rconLogin()

	switch cmd {
	case statusCmd.FullCommand():
		resp, err = rconStatus()
	case saveCmd.FullCommand():
		resp, err = rconSave()
	case shutdownCmd.FullCommand():
		resp, err = rconShutdown()
	case noticeCmd.FullCommand():
		resp, err = rconNotice()
	case banCmd.FullCommand():
		resp, err = rconBan()
	case unbanCmd.FullCommand():
		resp, err = rconUnban()
	case kickCmd.FullCommand():
		resp, err = rconKick()
	case clearallCmd.FullCommand():
		resp, err = rconClearall()
	case hungerRateCmd.FullCommand():
		resp, err = rconHungerRate()
	}

	if err != nil {
		log.Fatalf("Error processing command: \n%v", err)
	}

	fmt.Println(resp)
}

func rconStatus() (string, error) {
	return client.Status()
}

func rconSave() (string, error) {
	return client.Save(*saveCmdArg)
}

func rconShutdown() (string, error) {
	return client.Shutdown(*shutdownMsg, *shutdownTimeout)
}

func rconNotice() (string, error) {
	return client.Notice(*noticeCmdArg)
}

func rconBan() (string, error) {
	return client.Ban(*banPlayer, *banTimeout)
}

func rconUnban() (string, error) {
	return client.Unban(*unbanPlayer)
}

func rconKick() (string, error) {
	return client.Kick(*kickPlayer)
}

func rconClearall() (string, error) {
	return client.ClearAll()
}

func rconHungerRate() (string, error) {
	return client.HungerRate(*hungerRate)
}

func rconLogin() *srcon.Client {
	config = loadConfig()
	if len(config.Password) == 0 && len(*rconPassword) == 0 {
		// No password provided, so we need to prompt for one
		config.Password = loginPrompt()
	}

	if len(config.Password) == 0 && len(*rconPassword) > 0 {
		// We don't have a saved password, but one was provided.
		config.Password = *rconPassword
	}

	if serverIP.String() != "<nil>" {
		config.Hostname = serverIP.String()
	}

	if len(config.Hostname) == 0 {
		config.Hostname = "127.0.0.1"
	}

	if len(*port) > 0 {
		config.Port = *port
	}

	if len(config.Port) == 0 {
		// default
		config.Port = "27500"
	}

	// By this point we should have a valid Config, so save it.
	saveConfig(config)

	port, err := strconv.Atoi(config.Port)
	if err != nil {
		log.Fatalf("Error converting port: \n%v", err)
	}

	client := srcon.New(config.Hostname, port)
	err = client.Login(config.Password)

	if err != nil {
		log.Fatalf("Error logging into RCON: \n%v", err)
	}

	return client
}
