package parserqa

import (
	"bufio"
	"os"
	"strings"
)

type ParserQA struct {
	params *ParserQAParams
}

func NewParserQA(opts ...Options) (*ParserQA, error) {
	params, err := newParserQAParams(opts...)

	if err != nil {
		return nil, err
	}

	return &ParserQA{
		params: params,
	}, nil
}

func (p *ParserQA) Parse() ([]map[string][]string, error) {
	var matches []map[string][]string
	var currentMatch map[string][]string

	f, err := os.OpenFile(p.params.GetFileName(), os.O_RDONLY, 0644)
	if err != nil {
		return matches, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "InitGame:") {
			if currentMatch != nil {
				// Save the previous match
				matches = append(matches, currentMatch)
			}
			// Start a new match
			currentMatch = make(map[string][]string)
		}

		if currentMatch != nil {
			if strings.Contains(line, "ClientUserinfoChanged:") {
				currentMatch["ClientUserinfoChanged"] = append(currentMatch["ClientUserinfoChanged"], line)
			} else if strings.Contains(line, "Kill:") {
				currentMatch["Kill"] = append(currentMatch["Kill"], line)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Append the last match if exists
	if currentMatch != nil {
		matches = append(matches, currentMatch)
	}

	return matches, nil
}

func (p *ParserQA) ConsolidatePlayers(matches map[string][]string) []string {
	players := make(map[string]int)

	for _, match := range matches {
		for _, line := range match {
			if !strings.Contains(line, "ClientUserinfoChanged:") {
				continue
			}
			player := strings.Split(line, " n\\")[1]
			player = strings.Split(player, "\\t\\")[0]
			player = strings.TrimSpace(player)

			players[player] = 1
		}
	}

	result := []string{}

	for player := range players {
		result = append(result, player)
	}

	return result
}

func (p *ParserQA) ConsolidateKills(matches map[string][]string) int {
	total := 0

	for _, match := range matches {
		for _, line := range match {
			if !strings.Contains(line, "Kill:") {
				continue
			}
			total++
		}
	}

	return total
}

func (p *ParserQA) ConsolidateKillByPlayers(matches map[string][]string) map[string]int {
	players := make(map[string]int)

	for _, match := range matches {
		for _, line := range match {
			if !strings.Contains(line, "Kill:") {
				continue
			}
			player := strings.Split(line, "killed ")[1]
			player = strings.Split(player, " by")[0]
			player = strings.TrimSpace(player)

			players[player]++
		}
	}

	return players
}

func (p *ParserQA) ConsolidateKillByWorldPlayers(matches map[string][]string) map[string]int {
	players := make(map[string]int)

	for _, match := range matches {
		for _, line := range match {
			if !strings.Contains(line, "Kill:") {
				continue
			}

			if !strings.Contains(line, "<world>") {
				continue
			}

			player := strings.Split(line, "killed ")[1]
			player = strings.Split(player, " by")[0]
			player = strings.TrimSpace(player)

			players[player]++

		}
	}

	return players
}

func (p *ParserQA) ConsolidateKillByMod(matches map[string][]string) map[string]int {
	mods := make(map[string]int)

	for _, match := range matches {
		for _, line := range match {
			if !strings.Contains(line, "Kill:") {
				continue
			}

			if !strings.Contains(line, "MOD_") {
				continue
			}

			mod := strings.Split(line, "by ")[1]
			mod = strings.TrimSpace(mod)
			mods[mod]++
		}
	}

	return mods
}
