package commands

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/dipperin/dipperin-core/third-party/log"
	"github.com/urfave/cli"
	"os"
	"strings"
)

type Console struct {
	Prompt  *prompt.Prompt
	paused bool
}


func Executor(c *cli.Context) prompt.Executor {
	return func(command string) {
		if command == "" {
			return
		} else if command == "exit" {
			os.Exit(0)
		}
		cmdArgs := strings.Split(strings.TrimSpace(command)," ")
		log.Info("the cmdArgs is:","cmdArgs",cmdArgs)
		if len(cmdArgs) == 0 {
			return
		} else if len(cmdArgs) == 1 && cmdArgs[0] != "-h" && cmdArgs[0] != "--help"{
			fmt.Println("Please assign the method you want to call!")
			return
		}
		s := []string{os.Args[0]}
		s = append(s, cmdArgs...)

		log.Info("the Executor run s is:","s",s)
		err := c.App.Run(s)
		if err !=nil{
			log.Info("Executor", "err", err)
		}
	}
}

func CgoProCliCompleter(d prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

// NewConsole return *Console , completer Completer
func NewConsole(executor prompt.Executor, completer prompt.Completer) *Console {
	c := &Console{}

	// New prompt and load history input
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix("> "),
		prompt.OptionTitle("cgoProCli"),
		prompt.OptionPrefixTextColor(prompt.Yellow),
		prompt.OptionPreviewSuggestionTextColor(prompt.Blue),
		prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
		prompt.OptionSuggestionBGColor(prompt.DarkGray),
	)

	c.Prompt = p

	return c
}
