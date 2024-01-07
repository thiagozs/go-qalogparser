package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thiagozs/go-qalogparser/internal/parserqa"
)

type KindOutput int

const (
	KindOutputText KindOutput = iota
	KindOutputJSON
)

func (k KindOutput) String() string {
	return [...]string{"text", "json"}[k]
}

var (
	input      string
	output     string
	kind       string
	kindSelect KindOutput
)

var parserCmd = &cobra.Command{
	Use:   "parser",
	Short: "A parser for Quake Arena log files",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if err := handlerVars(cmd); err != nil {
			fmt.Println(err)
			return
		}

		popts := []parserqa.Options{parserqa.WithFileName(input)}
		p, err := parserqa.NewParserQA(popts...)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}

		switch kindSelect {
		case KindOutputText:
			if err := p.StdoutText(); err != nil {
				fmt.Printf(err.Error())
				return
			}
		case KindOutputJSON:
			if err := p.StdoutJSON(); err != nil {
				fmt.Printf(err.Error())
				return
			}
		}

	},
}

func init() {
	parserCmd.Flags().StringVarP(&input, "input", "i", "", "Input file")
	parserCmd.Flags().StringVarP(&output, "output", "o", "", "Output file")
	parserCmd.Flags().StringVarP(&kind, "kind", "k", "json", "Kind of output (json, text)")
}

func handlerVars(cmd *cobra.Command) error {
	if input == "" {
		return fmt.Errorf("input file is required")
	}

	switch kind {
	case KindOutputText.String():
		kindSelect = KindOutputText
	case KindOutputJSON.String():
		kindSelect = KindOutputJSON
	default:
		return fmt.Errorf("kind of output is invalid")
	}

	return nil
}
