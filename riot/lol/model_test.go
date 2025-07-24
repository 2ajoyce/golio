package lol

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/KnutZuidema/golio/api"
	"github.com/KnutZuidema/golio/datadragon"
	"github.com/KnutZuidema/golio/internal"
	"github.com/KnutZuidema/golio/internal/mock"
	"github.com/KnutZuidema/golio/static"
)

func TestLeagueList_GetRank(t *testing.T) {
	tests := []struct {
		name    string
		i       int
		entries []*LeagueItem
		want    *LeagueItem
	}{
		{
			name: "get rank",
			i:    0,
			entries: []*LeagueItem{
				{
					LeaguePoints: 0,
				},
				{
					LeaguePoints: 1,
				},
			},
			want: &LeagueItem{
				LeaguePoints: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				l := &LeagueList{
					Entries: tt.entries,
				}
				require.Equal(t, tt.want, l.GetRank(tt.i))
			},
		)
	}
}

func TestChampionInfo_GetChampionsForNewPlayers(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   ChampionInfo
		want    []datadragon.ChampionDataExtended
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(
				map[string]datadragon.ChampionData{
					"1": {Key: "1", ID: "1", Name: "champion1"},
					"2": {Key: "2", ID: "2", Name: "champion2"},
				},
			),
			model: ChampionInfo{
				FreeChampionIDsForNewPlayers: []int{1, 2},
			},
			want: []datadragon.ChampionDataExtended{
				{ChampionData: datadragon.ChampionData{ID: "1", Name: "champion1", Key: "1"}},
				{ChampionData: datadragon.ChampionData{ID: "2", Name: "champion2", Key: "2"}},
			},
		},
		{
			name: "unknown error",
			doer: mock.NewStatusMockDoer(999),
			model: ChampionInfo{
				FreeChampionIDsForNewPlayers: []int{1},
			},
			wantErr: api.Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
				got, err := test.model.GetChampionsForNewPlayers(client)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestChampionInfo_GetChampions(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   ChampionInfo
		want    []datadragon.ChampionDataExtended
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(
				map[string]datadragon.ChampionData{
					"1": {Key: "1", ID: "1", Name: "champion1"},
					"2": {Key: "2", ID: "2", Name: "champion2"},
				},
			),
			model: ChampionInfo{
				FreeChampionIDs: []int{1, 2},
			},
			want: []datadragon.ChampionDataExtended{
				{ChampionData: datadragon.ChampionData{ID: "1", Name: "champion1", Key: "1"}},
				{ChampionData: datadragon.ChampionData{ID: "2", Name: "champion2", Key: "2"}},
			},
		},
		{
			name: "unknown error",
			doer: mock.NewStatusMockDoer(999),
			model: ChampionInfo{
				FreeChampionIDs: []int{1},
			},
			wantErr: api.Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
				got, err := test.model.GetChampions(client)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestChampionMastery_GetSummoner(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   ChampionMastery
		want    *Summoner
		wantErr error
	}
	tests := []test{
		{
			name:  "valid",
			doer:  mock.NewJSONMockDoer(Summoner{ID: "id"}, 200),
			model: ChampionMastery{SummonerID: "id"},
			want:  &Summoner{ID: "id"},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := internal.NewClient(api.RegionKorea, "key", test.doer, log.StandardLogger())
				got, err := test.model.GetSummoner(NewClient(client))
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestChampionMastery_GetChampion(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   ChampionMastery
		want    datadragon.ChampionDataExtended
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(
				map[string]datadragon.ChampionData{
					"1": {Key: "1", ID: "1", Name: "champion"},
				},
			),
			model: ChampionMastery{ChampionID: 1},
			want: datadragon.ChampionDataExtended{
				ChampionData: datadragon.ChampionData{ID: "1", Name: "champion", Key: "1"},
			},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
				got, err := test.model.GetChampion(client)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestLeagueItem_GetSummoner(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   LeagueItem
		want    *Summoner
		wantErr error
	}
	tests := []test{
		{
			name:  "valid",
			doer:  mock.NewJSONMockDoer(Summoner{ID: "id"}, 200),
			model: LeagueItem{SummonerID: "id"},
			want:  &Summoner{ID: "id"},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := internal.NewClient(api.RegionKorea, "key", test.doer, log.StandardLogger())
				got, err := test.model.GetSummoner(NewClient(client))
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestMatchInfo_GetQueue(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   MatchInfo
		want    static.Queue
		wantErr error
	}
	tests := []test{
		{
			name:  "valid",
			doer:  mock.NewJSONMockDoer([]static.Queue{{ID: 1}}, 200),
			model: MatchInfo{QueueID: 1},
			want:  static.Queue{ID: 1},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := static.NewClient(test.doer, log.StandardLogger())
				got, err := test.model.GetQueue(client)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestMatchInfo_GetMap(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   MatchInfo
		want    static.Map
		wantErr error
	}
	tests := []test{
		{
			name:  "valid",
			doer:  mock.NewJSONMockDoer([]static.Map{{ID: 1}}, 200),
			model: MatchInfo{MapID: 1},
			want:  static.Map{ID: 1},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := static.NewClient(test.doer, log.StandardLogger())
				got, err := test.model.GetMap(client)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestMatchInfo_GetGameType(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   MatchInfo
		want    static.GameType
		wantErr error
	}
	tests := []test{
		{
			name:  "valid",
			doer:  mock.NewJSONMockDoer([]static.GameType{{Type: "type"}}, 200),
			model: MatchInfo{GameType: "type"},
			want:  static.GameType{Type: "type"},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := static.NewClient(test.doer, log.StandardLogger())
				got, err := test.model.GetGameType(client)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestMatchInfo_GetGameMode(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   MatchInfo
		want    static.GameMode
		wantErr error
	}
	tests := []test{
		{
			name:  "valid",
			doer:  mock.NewJSONMockDoer([]static.GameMode{{Mode: "type"}}, 200),
			model: MatchInfo{GameMode: "type"},
			want:  static.GameMode{Mode: "type"},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := static.NewClient(test.doer, log.StandardLogger())
				got, err := test.model.GetGameMode(client)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestMatchInfo_GetEndOfGameResult(t *testing.T) {
	tests := []struct {
		name     string
		model    MatchInfo
		expected string
	}{
		{
			name:     "game complete",
			model:    MatchInfo{EndOfGameResult: "GameComplete"},
			expected: "GameComplete",
		},
		{
			name:     "surrender",
			model:    MatchInfo{EndOfGameResult: "Surrender"},
			expected: "Surrender",
		},
		{
			name:     "empty result",
			model:    MatchInfo{},
			expected: "",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equal(t, tt.expected, tt.model.EndOfGameResult)
			},
		)
	}
}

func TestParticipant_GetSummoner(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   Participant
		want    *Summoner
		wantErr error
	}
	tests := []test{
		{
			name:  "valid",
			doer:  mock.NewJSONMockDoer(Summoner{ID: "id"}, 200),
			model: Participant{SummonerID: "id"},
			want:  &Summoner{ID: "id"},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := internal.NewClient(api.RegionKorea, "key", test.doer, log.StandardLogger())
				got, err := test.model.GetSummoner(NewClient(client))
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestParticipant_GetProfileIcon(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   Participant
		want    datadragon.ProfileIcon
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(
				map[string]datadragon.ProfileIcon{
					"champion": {ID: 1},
				},
			),
			model: Participant{ProfileIcon: 1},
			want:  datadragon.ProfileIcon{ID: 1},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
				got, err := test.model.GetProfileIcon(client)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestTeamBan_GetChampion(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   TeamBan
		want    datadragon.ChampionDataExtended
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(
				map[string]datadragon.ChampionData{
					"1": {Key: "1", ID: "1", Name: "champion"},
				},
			),
			model: TeamBan{ChampionID: 1},
			want: datadragon.ChampionDataExtended{
				ChampionData: datadragon.ChampionData{ID: "1", Name: "champion", Key: "1"},
			},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
				got, err := test.model.GetChampion(client)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestParticipant_GetChampion(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   Participant
		want    datadragon.ChampionDataExtended
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(
				map[string]datadragon.ChampionData{
					"1": {Key: "1", ID: "1", Name: "champion"},
				},
			),
			model: Participant{ChampionID: 1},
			want: datadragon.ChampionDataExtended{
				ChampionData: datadragon.ChampionData{ID: "1", Name: "champion", Key: "1"},
			},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
				got, err := test.model.GetChampion(client)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestParticipant_GetSpell1(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   Participant
		want    datadragon.SummonerSpell
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(
				map[string]datadragon.SummonerSpell{
					"champion": {Key: "1"},
				},
			),
			model: Participant{Summoner1ID: 1},
			want:  datadragon.SummonerSpell{Key: "1"},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
				got, err := test.model.GetSpell1(client)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestParticipant_GetSpell2(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   Participant
		want    datadragon.SummonerSpell
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(
				map[string]datadragon.SummonerSpell{
					"champion": {Key: "2"},
				},
			),
			model: Participant{Summoner2ID: 2},
			want:  datadragon.SummonerSpell{Key: "2"},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
				got, err := test.model.GetSpell2(client)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestParticipant_GetItem0(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   Participant
		want    datadragon.Item
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(
				map[string]datadragon.Item{
					"1": {},
				},
			),
			model: Participant{Item0: 1},
			want:  datadragon.Item{ID: "1"},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
				got, err := test.model.GetItem0(client)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestParticipant_GetItem1(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   Participant
		want    datadragon.Item
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(
				map[string]datadragon.Item{
					"1": {},
				},
			),
			model: Participant{Item1: 1},
			want:  datadragon.Item{ID: "1"},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
				got, err := test.model.GetItem1(client)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestParticipant_GetItem2(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   Participant
		want    datadragon.Item
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(
				map[string]datadragon.Item{
					"1": {},
				},
			),
			model: Participant{Item2: 1},
			want:  datadragon.Item{ID: "1"},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
				got, err := test.model.GetItem2(client)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestParticipant_GetItem3(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   Participant
		want    datadragon.Item
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(
				map[string]datadragon.Item{
					"1": {},
				},
			),
			model: Participant{Item3: 1},
			want:  datadragon.Item{ID: "1"},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
				got, err := test.model.GetItem3(client)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestParticipant_GetItem4(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   Participant
		want    datadragon.Item
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(
				map[string]datadragon.Item{
					"1": {},
				},
			),
			model: Participant{Item4: 1},
			want:  datadragon.Item{ID: "1"},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
				got, err := test.model.GetItem4(client)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestParticipant_GetItem5(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   Participant
		want    datadragon.Item
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(
				map[string]datadragon.Item{
					"1": {},
				},
			),
			model: Participant{Item5: 1},
			want:  datadragon.Item{ID: "1"},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
				got, err := test.model.GetItem5(client)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestParticipant_GetItem6(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   Participant
		want    datadragon.Item
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(
				map[string]datadragon.Item{
					"1": {},
				},
			),
			model: Participant{Item6: 1},
			want:  datadragon.Item{ID: "1"},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
				got, err := test.model.GetItem6(client)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestBannedChampion_GetChampion(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   BannedChampion
		want    datadragon.ChampionDataExtended
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(
				map[string]datadragon.ChampionData{
					"1": {Key: "1", ID: "1", Name: "champion"},
				},
			),
			model: BannedChampion{ChampionID: 1},
			want: datadragon.ChampionDataExtended{
				ChampionData: datadragon.ChampionData{ID: "1", Name: "champion", Key: "1"},
			},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
				got, err := test.model.GetChampion(client)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestCurrentGameParticipant_GetChampion(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   CurrentGameParticipant
		want    datadragon.ChampionDataExtended
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(
				map[string]datadragon.ChampionData{
					"1": {Key: "1", ID: "1", Name: "champion"},
				},
			),
			model: CurrentGameParticipant{ChampionID: 1},
			want: datadragon.ChampionDataExtended{
				ChampionData: datadragon.ChampionData{ID: "1", Name: "champion", Key: "1"},
			},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
				got, err := test.model.GetChampion(client)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestCurrentGameParticipant_GetSpell1(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   CurrentGameParticipant
		want    datadragon.SummonerSpell
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(
				map[string]datadragon.SummonerSpell{
					"champion": {Key: "1"},
				},
			),
			model: CurrentGameParticipant{Spell1ID: 1},
			want:  datadragon.SummonerSpell{Key: "1"},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
				got, err := test.model.GetSpell1(client)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestCurrentGameParticipant_GetSpell2(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   CurrentGameParticipant
		want    datadragon.SummonerSpell
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(
				map[string]datadragon.SummonerSpell{
					"champion": {Key: "2"},
				},
			),
			model: CurrentGameParticipant{Spell2ID: 2},
			want:  datadragon.SummonerSpell{Key: "2"},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
				got, err := test.model.GetSpell2(client)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestGameInfo_GetMatch(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   GameInfo
		want    *Match
		wantErr error
	}
	tests := []test{
		{
			name:  "valid",
			doer:  mock.NewJSONMockDoer(Match{Metadata: &MatchMetadata{MatchID: "NA1_1"}}, 200),
			model: GameInfo{GameID: 1},
			want:  &Match{Metadata: &MatchMetadata{MatchID: "NA1_1"}},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := internal.NewClient(api.RegionKorea, "key", test.doer, log.StandardLogger())
				got, err := test.model.GetMatch(NewClient(client))
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestMatchTimeline_StructFields(t *testing.T) {
	mt := MatchTimeline{
		Metadata: MetadataTimeline{
			DataVersion:  "1",
			MatchID:      "match1",
			Participants: []string{"a", "b"},
		},
		Info: InfoTimeline{
			EndOfGameResult: "Win",
			FrameInterval:   1000,
			GameID:          123,
			Participants:    []ParticipantTimeline{{ParticipantID: 1, PUUID: "puuid1"}},
			Frames: []FramesTimeline{{
				Timestamp: 100,
				Events:    []EventsTimeline{{Type: "CHAMPION_KILL", Timestamp: 100}},
				ParticipantFrames: map[string]ParticipantFrame{
					"1": {ParticipantID: 1, Level: 10},
				},
			}},
		},
	}
	assert.Equal(t, "1", mt.Metadata.DataVersion)
	assert.Equal(t, "match1", mt.Metadata.MatchID)
	assert.Equal(t, "Win", mt.Info.EndOfGameResult)
	assert.Equal(t, int64(123), mt.Info.GameID)
	assert.Equal(t, 1, mt.Info.Participants[0].ParticipantID)
	assert.Equal(t, 100, mt.Info.Frames[0].Timestamp)
	assert.Equal(t, "CHAMPION_KILL", mt.Info.Frames[0].Events[0].Type)
	assert.Equal(t, 1, mt.Info.Frames[0].ParticipantFrames["1"].ParticipantID)
}

func TestEventsTimeline_StructFields(t *testing.T) {
	pos := EventPosition{X: 13958, Y: 3889}
	victim := VictimDamage{
		Basic:          func() *bool { b := true; return &b }(),
		MagicDamage:    func() *int { i := 61; return &i }(),
		Name:           func() *string { s := "Rammus"; return &s }(),
		ParticipantID:  func() *int { i := 6; return &i }(),
		PhysicalDamage: func() *int { i := 73; return &i }(),
		SpellName:      func() *string { s := "rammusbasicattack"; return &s }(),
		SpellSlot:      func() *int { i := 0; return &i }(),
		TrueDamage:     func() *int { i := 0; return &i }(),
		Type:           func() *string { s := "MONSTER"; return &s }(),
	}
	et := EventsTimeline{
		Timestamp:            1006818,
		Type:                 "CHAMPION_KILL",
		RealTimestamp:        func() *int64 { i := int64(1006818); return &i }(),
		Position:             &pos,
		VictimDamageDealt:    []VictimDamage{victim},
		VictimDamageReceived: []VictimDamage{victim},
	}
	assert.Equal(t, int64(1006818), et.Timestamp)
	assert.Equal(t, "CHAMPION_KILL", et.Type)
	assert.NotNil(t, et.Position)
	assert.Equal(t, 13958, et.Position.X)
	assert.Len(t, et.VictimDamageDealt, 1)
	assert.Len(t, et.VictimDamageReceived, 1)
}

func TestParticipantFrame_StructFields(t *testing.T) {
	pf := ParticipantFrame{
		ChampionStats:            ChampionStats{Health: 1847, Armor: 67, AttackDamage: 80, MagicResist: 41, MovementSpeed: 397},
		CurrentGold:              1189,
		DamageStats:              DamageStats{TotalDamageDone: 43528, MagicDamageDone: 26216, PhysicalDamageDone: 15222, TrueDamageDone: 2090, MagicDamageTaken: 0, PhysicalDamageTaken: 3791, TotalDamageTaken: 4102, TrueDamageTaken: 311},
		GoldPerSecond:            0,
		JungleMinionsKilled:      0,
		Level:                    11,
		MinionsKilled:            115,
		ParticipantID:            1,
		Position:                 ParticipantPosition{X: 1020, Y: 7019},
		TimeEnemySpentControlled: 131193,
		TotalGold:                4889,
		XP:                       7566,
	}
	assert.Equal(t, 1847, pf.ChampionStats.Health)
	assert.Equal(t, 67, pf.ChampionStats.Armor)
	assert.Equal(t, 80, pf.ChampionStats.AttackDamage)
	assert.Equal(t, 41, pf.ChampionStats.MagicResist)
	assert.Equal(t, 397, pf.ChampionStats.MovementSpeed)
	assert.Equal(t, 43528, pf.DamageStats.TotalDamageDone)
	assert.Equal(t, 26216, pf.DamageStats.MagicDamageDone)
	assert.Equal(t, 15222, pf.DamageStats.PhysicalDamageDone)
	assert.Equal(t, 2090, pf.DamageStats.TrueDamageDone)
	assert.Equal(t, 0, pf.DamageStats.MagicDamageTaken)
	assert.Equal(t, 3791, pf.DamageStats.PhysicalDamageTaken)
	assert.Equal(t, 4102, pf.DamageStats.TotalDamageTaken)
	assert.Equal(t, 311, pf.DamageStats.TrueDamageTaken)
	assert.Equal(t, 0, pf.GoldPerSecond)
	assert.Equal(t, 0, pf.JungleMinionsKilled)
	assert.Equal(t, 11, pf.Level)
	assert.Equal(t, 115, pf.MinionsKilled)
	assert.Equal(t, 1, pf.ParticipantID)
	assert.Equal(t, 1020, pf.Position.X)
	assert.Equal(t, 7019, pf.Position.Y)
	assert.Equal(t, 131193, pf.TimeEnemySpentControlled)
	assert.Equal(t, 4889, pf.TotalGold)
	assert.Equal(t, 7566, pf.XP)
}

func TestMetadataTimeline_StructFields(t *testing.T) {
	mt := MetadataTimeline{
		DataVersion:  "2",
		MatchID:      "NA1_1",
		Participants: []string{"a", "b"},
	}
	assert.Equal(t, "2", mt.DataVersion)
	assert.Equal(t, "NA1_1", mt.MatchID)
	assert.Equal(t, []string{"a", "b"}, mt.Participants)
}

func TestInfoTimeline_StructFields(t *testing.T) {
	it := InfoTimeline{
		EndOfGameResult: "Win",
		FrameInterval:   1000,
		GameID:          123,
		Participants:    []ParticipantTimeline{{ParticipantID: 1, PUUID: "puuid1"}},
		Frames:          []FramesTimeline{{Timestamp: 100}},
	}
	assert.Equal(t, "Win", it.EndOfGameResult)
	assert.Equal(t, int64(1000), it.FrameInterval)
	assert.Equal(t, int64(123), it.GameID)
	assert.Equal(t, 1, it.Participants[0].ParticipantID)
	assert.Equal(t, 100, it.Frames[0].Timestamp)
}

func TestParticipantTimeline_StructFields(t *testing.T) {
	pt := ParticipantTimeline{ParticipantID: 1, PUUID: "puuid1"}
	assert.Equal(t, 1, pt.ParticipantID)
	assert.Equal(t, "puuid1", pt.PUUID)
}

func TestFramesTimeline_StructFields(t *testing.T) {
	ft := FramesTimeline{
		Timestamp: 100,
		Events:    []EventsTimeline{{Type: "CHAMPION_KILL"}},
		ParticipantFrames: map[string]ParticipantFrame{
			"1": {ParticipantID: 1, Level: 11},
		},
	}
	assert.Equal(t, 100, ft.Timestamp)
	assert.Equal(t, "CHAMPION_KILL", ft.Events[0].Type)
	assert.Equal(t, 1, ft.ParticipantFrames["1"].ParticipantID)
	assert.Equal(t, 11, ft.ParticipantFrames["1"].Level)
}

func TestEventPosition_StructFields(t *testing.T) {
	ep := EventPosition{X: 13958, Y: 3889}
	assert.Equal(t, 13958, ep.X)
	assert.Equal(t, 3889, ep.Y)
}

func TestVictimDamage_StructFields(t *testing.T) {
	b := true
	md := 61
	n := "Rammus"
	pid := 6
	pd := 73
	sn := "rammusbasicattack"
	ss := 0
	td := 0
	typ := "MONSTER"
	vd := VictimDamage{
		Basic:          &b,
		MagicDamage:    &md,
		Name:           &n,
		ParticipantID:  &pid,
		PhysicalDamage: &pd,
		SpellName:      &sn,
		SpellSlot:      &ss,
		TrueDamage:     &td,
		Type:           &typ,
	}
	assert.Equal(t, true, *vd.Basic)
	assert.Equal(t, 61, *vd.MagicDamage)
	assert.Equal(t, "Rammus", *vd.Name)
	assert.Equal(t, 6, *vd.ParticipantID)
	assert.Equal(t, 73, *vd.PhysicalDamage)
	assert.Equal(t, "rammusbasicattack", *vd.SpellName)
	assert.Equal(t, 0, *vd.SpellSlot)
	assert.Equal(t, 0, *vd.TrueDamage)
	assert.Equal(t, "MONSTER", *vd.Type)
}

func TestChampionStats_StructFields(t *testing.T) {
	cs := ChampionStats{Health: 1847, Armor: 67, AttackDamage: 80, MagicResist: 41, MovementSpeed: 397}
	assert.Equal(t, 1847, cs.Health)
	assert.Equal(t, 67, cs.Armor)
	assert.Equal(t, 80, cs.AttackDamage)
	assert.Equal(t, 41, cs.MagicResist)
	assert.Equal(t, 397, cs.MovementSpeed)
}

func TestDamageStats_StructFields(t *testing.T) {
	ds := DamageStats{
		MagicDamageDone:     26216,
		PhysicalDamageDone:  15222,
		TotalDamageDone:     43528,
		TrueDamageDone:      2090,
		MagicDamageTaken:    0,
		PhysicalDamageTaken: 3791,
		TotalDamageTaken:    4102,
		TrueDamageTaken:     311,
	}
	assert.Equal(t, 26216, ds.MagicDamageDone)
	assert.Equal(t, 15222, ds.PhysicalDamageDone)
	assert.Equal(t, 43528, ds.TotalDamageDone)
	assert.Equal(t, 2090, ds.TrueDamageDone)
	assert.Equal(t, 0, ds.MagicDamageTaken)
	assert.Equal(t, 3791, ds.PhysicalDamageTaken)
	assert.Equal(t, 4102, ds.TotalDamageTaken)
	assert.Equal(t, 311, ds.TrueDamageTaken)
}

func TestParticipantPosition_StructFields(t *testing.T) {
	pp := ParticipantPosition{X: 1020, Y: 7019}
	assert.Equal(t, 1020, pp.X)
	assert.Equal(t, 7019, pp.Y)
}

type dataDragonResponse struct {
	Type    string
	Format  string
	Version string
	Data    interface{}
}

func dataDragonResponseDoer(object interface{}) internal.Doer {
	return mock.NewJSONMockDoer(
		dataDragonResponse{
			Data: object,
		}, 200,
	)
}
