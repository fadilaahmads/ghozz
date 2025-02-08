package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli/v2"
  "gofuzzer/internal/fuzzer"
	"gofuzzer/internal/tor"
)

func main()  {
  fmt.Println("GoFuzzer")
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
    },
    Action: func(c *cli.Context) error {
      target:= c.String("target")
      wordlistFile := c.String("wordlist")
      wordlist, err := fuzzer.ReadWordlist(wordlistFile)
      if err != nil {
        return err
      }
      fuzzer.Fuzz(target, wordlist, tor)
      return nil
    },
  }

  err = app.Run(os.Args)
  if err != nil {
    fmt.Errorf("Error running app: %v", err)
  }
}
