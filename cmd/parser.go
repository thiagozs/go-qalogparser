package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thiagozs/go-qalogparser/internal/parserqa"
)

var (
	input  string
	output string
)

var parserCmd = &cobra.Command{
	Use:   "parser",
	Short: "A parser for QA log files",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("parser called")

		if err := handlerVars(cmd); err != nil {
			fmt.Println(err)
			return
		}

		popts := []parserqa.Options{parserqa.WithFileName(input)}
		p, err := parserqa.NewParserQA(popts...)
		if err != nil {
			fmt.Println(err)
			return
		}

		matches, err := p.Parse()
		if err != nil {
			fmt.Println(err)
			return
		}

		for i, matche := range matches {
			fmt.Println("Game_", i)

			players := p.ConsolidatePlayers(matche)
			for _, player := range players {
				fmt.Println("  - Player:", player)
			}

			kills := p.ConsolidateKills(matche)
			fmt.Println("  - Total kills:", kills)

			killbyplayers := p.ConsolidateKillByPlayers(matche)

			for player, kill := range killbyplayers {
				fmt.Println("  - Total kills by", player, "-", kill)
			}

			killbyworld := p.ConsolidateKillByWorldPlayers(matche)

			for player, kill := range killbyworld {
				fmt.Println("  - Total kills by world -", player, "-", kill)
			}

			killbymods := p.ConsolidateKillByMod(matche)

			for player, kill := range killbymods {
				fmt.Println("  - Total kills by mod -", player, "-", kill)
			}

		}

	},
}

func init() {
	parserCmd.Flags().StringVarP(&input, "input", "i", "", "Input file")
	parserCmd.Flags().StringVarP(&output, "output", "o", "", "Output file")
}

func handlerVars(cmd *cobra.Command) error {
	if input == "" {
		return fmt.Errorf("input file is required")
	}

	if output == "" {
		return fmt.Errorf("output file is required")
	}

	return nil
}
