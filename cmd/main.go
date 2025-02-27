package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli/v2"
  "ghozz/internal/fuzzer"
	"ghozz/internal/tor"
)

func main()  {
  tor, err := tor.SetupTOR()
  if err != nil {
    fmt.Errorf("Error setting up TOR proxy: %v", err)
  }
   
  app := &cli.App{
    Name: "GoFuzzer",
    Usage: "Directory Fuzzing With TOR Support",
    Flags: []cli.Flag{
      &cli.StringFlag{Name: "target", Required: true},
      &cli.StringFlag{Name: "wordlist", Required: true},
      &cli.StringFlag{Name: "output", Required: false, Usage: "File to save output"},
      &cli.StringFlag{Name: "hide code", Required: false, Usage: "Filter output http code"},
    },
    Action: func(c *cli.Context) error {
      target:= c.String("target")
      wordlistFile := c.String("wordlist")
      outputFile:= c.String("output")
      hideCode:= c.String("hide-code")

      wordlist, err := fuzzer.ReadWordlist(wordlistFile)
      if err != nil {
        return err
      }
      fuzzer.Fuzz(target, wordlist, hideCode, nil,  tor, outputFile)
      return nil
    },
  }

  err = app.Run(os.Args)
  if err != nil {
    fmt.Errorf("Error running app: %v", err)
  }
}
