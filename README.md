[![Build Status](https://travis-ci.org/matjam/stationeersrcon.svg?branch=master)](https://travis-ci.org/matjam/stationeersrcon)
[![Documentation](https://camo.githubusercontent.com/915b7be44ada53c290eb157634330494ebe3e30a/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f676f6c616e672f6764646f3f7374617475732e737667)](https://godoc.org/github.com/matjam/stationeersrcon)

# stationeersrcon
A cli to the stationeers rcon for Stationeers dedicated servers.

```
usage: srcon [<flags>] <command> [<args> ...]

a CLI interface to Stationeers dedicated server RCON.

Flags:
  -h, --help               Show context-sensitive help (also try --help-long and --help-man).
  -c, --config=CONFIG      Path to a json configuration file for the tool.
      --ip=IP              Server IP to connect to.
      --port=PORT          Port to connect to
      --password=PASSWORD  Password to use for the RCON command.

Commands:
  help [<command>...]
    Show help.

  status
    Fetch the status of the Stationeers server.

  save <savefile>
    Saves the game to a specified file.

  shutdown [<flags>]
    Shutdowns the server

  notice <message>
    Sends a notice to all players

  ban <steamId> <timeout>
    Bans a player for a specific time.

  unban <steamID>
    Remove a player from the ban list.

  kick <steamID>
    Kick player from the server.

  clearall
    Delete disconnected players from the server.

  hungerRate <rate>
    Set the hungerRate on the server to the given number.
```

## Building and Installing

```
go get ./...
go install ./...
```

This will build and install the binary into your $GOBIN directory. This could be copied anywhere in your path.

# Prebuilt Binary Releases

You can find the latest release here:

  https://github.com/matjam/stationeersrcon/releases/tag/v1.2.0

Download the binary for your platform and just drop it in your path somewhere.
