package parserqa

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	// Example log content
	logContent := `0:00 InitGame: \sv_floodProtect\1
        20:34 ClientUserinfoChanged: 2 n\Isgalamido\t\0
        20:54 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT`

	reader := strings.NewReader(logContent)

	parser, err := NewParserQA()
	assert.NoError(t, err)

	matches, err := parser.Parse(reader)
	assert.NoError(t, err)

	assert.Len(t, matches, 1, "One match parsed")
	assert.Contains(t, matches[0]["ClientUserinfoChanged"][0], "Isgalamido", "The parsed data should contain player 'Isgalamido'")
	assert.Contains(t, matches[0]["Kill"][0], "MOD_TRIGGER_HURT", "The parsed data should contain kill info 'MOD_TRIGGER_HURT'")
}

func TestConsolidatePlayers(t *testing.T) {
	// Mock match data
	matchData := map[string][]string{
		"ClientUserinfoChanged": {
			"20:34 ClientUserinfoChanged: 2 n\\Isgalamido\\t\\0",
			"21:15 ClientUserinfoChanged: 3 n\\Dono da Bola\\t\\0",
			"22:07 ClientUserinfoChanged: 2 n\\Isgalamido\\t\\0",
		},
	}

	parser, err := NewParserQA()
	assert.NoError(t, err)

	players := parser.ConsolidatePlayers(matchData)

	assert.Len(t, players, 2, "There should be 2 unique players")
	assert.Contains(t, players, "Isgalamido", "Players should include 'Isgalamido'")
	assert.Contains(t, players, "Dono da Bola", "Players should include 'Dono da Bola'")
}

func TestConsolidateKills(t *testing.T) {
	matchData := map[string][]string{
		"Kill": {
			"20:54 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT",
			"21:07 Kill: 1022 2 22: <world> killed Isgalamido by MOD_FALLING",
			"22:40 Kill: 2 2 7: Isgalamido killed Isgalamido by MOD_ROCKET_SPLASH",
		},
	}

	parser, err := NewParserQA()
	assert.NoError(t, err)

	totalKills := parser.ConsolidateKills(matchData)

	assert.Equal(t, 3, totalKills, "Total kills should be correctly counted")
}

func TestConsolidateKillByPlayers(t *testing.T) {
	matchData := map[string][]string{
		"Kill": {
			"20:54 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT",
			"21:07 Kill: 1022 3 22: <world> killed Dono da Bola by MOD_FALLING",
			"22:40 Kill: 2 4 7: Zeh killed Mocinha by MOD_ROCKET_SPLASH",
			"22:50 Kill: 4 2 7: Mocinha killed Zeh by MOD_ROCKET_SPLASH",
		},
	}

	parser, err := NewParserQA()
	assert.NoError(t, err)

	killsByPlayers := parser.ConsolidateKillByPlayers(matchData)

	expectedKills := map[string]int{
		"Mocinha": 1,
		"Zeh":     1,
	}

	assert.Equal(t, expectedKills, killsByPlayers, "Kills should be correctly counted for each player")
}

func TestConsolidateKillByWorldPlayers(t *testing.T) {
	matchData := map[string][]string{
		"Kill": {
			"20:54 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT",
			"21:07 Kill: 1022 3 22: <world> killed Dono da Bola by MOD_FALLING",
			"22:40 Kill: 1022 4 22: <world> killed Zeh by MOD_LAVA",
			"22:40 Kill: 2 4 7: Zeh killed Mocinha by MOD_ROCKET_SPLASH",
			"22:50 Kill: 4 2 7: Mocinha killed Zeh by MOD_ROCKET_SPLASH",
		},
	}

	parser, err := NewParserQA()
	assert.NoError(t, err)

	killsByWorld := parser.ConsolidateKillByWorldPlayers(matchData)

	expectedKills := map[string]int{
		"Isgalamido":   1,
		"Dono da Bola": 1,
		"Zeh":          1,
	}

	assert.Equal(t, expectedKills, killsByWorld, "World kills should be correctly counted for each player")
}

func TestConsolidateKillByMod(t *testing.T) {
	matchData := map[string][]string{
		"Kill": {
			"20:54 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT",
			"21:07 Kill: 1022 3 22: <world> killed Dono da Bola by MOD_FALLING",
			"22:40 Kill: 2 4 7: Zeh killed Mocinha by MOD_ROCKET_SPLASH",
			"22:50 Kill: 4 2 7: Mocinha killed Zeh by MOD_SHOTGUN",
		},
	}

	parser, err := NewParserQA()
	assert.NoError(t, err)

	killsByMod := parser.ConsolidateKillByMod(matchData)

	expectedKills := map[string]int{
		"MOD_TRIGGER_HURT":  1,
		"MOD_FALLING":       1,
		"MOD_ROCKET_SPLASH": 1,
		"MOD_SHOTGUN":       1,
	}

	assert.Equal(t, expectedKills, killsByMod, "Kills should be correctly counted for each MOD")
}

func TestStdoutJSON(t *testing.T) {

	parser, err := NewParserQA(WithFileName("testdata/qa.log"), WithBuffered())
	assert.NoError(t, err)

	err = parser.StdoutJSON()
	assert.NoError(t, err)

	assert.Contains(t, parser.GetBufferString(), "games")
	assert.Contains(t, parser.GetBufferString(), "game_0")
	assert.Contains(t, parser.GetBufferString(), "total_kills")
	assert.Contains(t, parser.GetBufferString(), "players")
	assert.Contains(t, parser.GetBufferString(), "kills")
	assert.Contains(t, parser.GetBufferString(), "kills_by_means")
	assert.Contains(t, parser.GetBufferString(), "MOD_ROCKET")
	assert.Contains(t, parser.GetBufferString(), "MOD_TRIGGER_HURT")
	assert.Contains(t, parser.GetBufferString(), "MOD_FALLING")
	assert.Contains(t, parser.GetBufferString(), "Dono da Bola")
	assert.Contains(t, parser.GetBufferString(), "Mocinha")
}

func TestStdoutText(t *testing.T) {

	parser, err := NewParserQA(WithFileName("testdata/qa.log"), WithBuffered())
	assert.NoError(t, err)

	err = parser.StdoutText()
	assert.NoError(t, err)

	assert.Contains(t, parser.GetBufferString(), "- Player: Dono da Bola")
	assert.Contains(t, parser.GetBufferString(), "- Player: Mocinha")
	assert.Contains(t, parser.GetBufferString(), "- Player: Isgalamido")
	assert.Contains(t, parser.GetBufferString(), "- Player: Zeh")
	assert.Contains(t, parser.GetBufferString(), "----------")
	assert.Contains(t, parser.GetBufferString(), "- Total kills: 4")
	assert.Contains(t, parser.GetBufferString(), "- Total kills by Mocinha - 1")
	assert.Contains(t, parser.GetBufferString(), "- Total kills by world - Zeh - 2")
	assert.Contains(t, parser.GetBufferString(), "- Total kills by world - Dono da Bola - 1")
	assert.Contains(t, parser.GetBufferString(), "Game_0")
	assert.Contains(t, parser.GetBufferString(), "MOD_ROCKET")
	assert.Contains(t, parser.GetBufferString(), "MOD_TRIGGER_HURT")
	assert.Contains(t, parser.GetBufferString(), "MOD_FALLING")
	assert.Contains(t, parser.GetBufferString(), "Dono da Bola")
	assert.Contains(t, parser.GetBufferString(), "Mocinha")
}
