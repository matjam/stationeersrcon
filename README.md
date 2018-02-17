# stationeersrcon
A cli to the stationeers rcon for Stationeers dedicated servers.

```
usage: srcon --password=PASSWORD [<flags>] <command> [<args> ...]

a CLI interface to Stationeers dedicated server RCON.

Flags:
  --help               Show context-sensitive help (also try --help-long and --help-man).
  --ip=127.0.0.1       Server IP to connect to.
  --port="27500"       Port to connect to
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
```

## Building

```
go get github.com/jaytaylor/html2text
go get gopkg.in/alecthomas/kingpin.v2
make
```

Only tested on MacOS. Should work fine on Linux. Windows you might have to build manually with `go build -o srcon cmd/srcon/*.go` unless you run the build inside a bash shell of some kind.
