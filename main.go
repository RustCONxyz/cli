package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/RustCONxyz/rustcon-go"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"golang.org/x/term"
)

const version = "0.0.1"

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print only the version",
	}

	cli.VersionPrinter = func(cCtx *cli.Context) {
		fmt.Printf("version=%s\n", cCtx.App.Version)
	}

	app := &cli.App{
		Name:    "rustcon",
		Usage:   "Connect to your Rust servers via RCON",
		Version: version,
		Action: func(cCtx *cli.Context) error {
			interrupt := make(chan os.Signal, 1)
			signal.Notify(interrupt, os.Interrupt)

			done := make(chan struct{})

			connectionDetails := cCtx.Args().First()
			if connectionDetails == "" {
				return fmt.Errorf("missing connection details")
			}

			host, port, err := ParseConnectionDetails(connectionDetails)
			if err != nil {
				return err
			}

			fmt.Print("RCON password: ")
			bytepw, err := term.ReadPassword(int(syscall.Stdin))
			if err != nil {
				return err
			}

			connection := &rustcon.RconConnection{
				IP:       host,
				Port:     port,
				Password: string(bytepw),
				OnConnected: func() {
					ClearScreen()
					color.Green("Connected to server")
				},
				OnMessage: func(message *rustcon.Message) {
					if message.Message == "" {
						return
					}

					if message.Type == "Error" {
						color.Red(message.Message)
					} else if message.Type == "Warning" {
						color.Yellow(message.Message)
					} else {
						fmt.Println(message.Message)
					}
				},
				OnChatMessage: func(chatMessage *rustcon.ChatMessage) {
					color.Blue("[%s] %s: %s\n", FormatTimestamp(chatMessage.Time, "15:04"), chatMessage.Username, chatMessage.Message)
				},
			}

			if err := connection.Connect(); err != nil {
				return err
			}

			go func() {
				scanner := bufio.NewScanner(os.Stdin)
				for scanner.Scan() {
					fmt.Print("\033[1A\033[K")

					input := scanner.Text()
					if len(input) == 0 {
						continue
					}

					color.Green("> " + input)

					if err := connection.SendCommand(input); err != nil {
						log.Fatal(err)
					}
				}
			}()

			for {
				select {
				case <-done:
					return nil
				case <-interrupt:
					err := connection.Disconnect()
					if err != nil {
						return nil
					}
					select {
					case <-done:
					case <-time.After(time.Second):
					}
					return nil
				}
			}
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
