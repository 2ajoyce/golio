package lol

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/KnutZuidema/golio/internal"
)

// LeagueClient provides methods for league endpoints of the League of Legends API.
type LeagueClient struct {
	c *internal.Client
}

// GetChallenger returns the current Challenger league for the Region
func (l *LeagueClient) GetChallenger(queue queue) (*LeagueList, error) {
	logger := l.logger().WithField("method", "GetChallenger")
	var list *LeagueList
	if err := l.c.GetInto(fmt.Sprintf(endpointGetChallengerLeague, queue), &list); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return list, nil
}

// GetGrandmaster returns the current Grandmaster league for the Region
func (l *LeagueClient) GetGrandmaster(queue queue) (*LeagueList, error) {
	logger := l.logger().WithField("method", "GetGrandmaster")
	var list *LeagueList
	if err := l.c.GetInto(fmt.Sprintf(endpointGetGrandmasterLeague, queue), &list); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return list, nil
}

// GetMaster returns the current Master league for the Region
func (l *LeagueClient) GetMaster(queue queue) (*LeagueList, error) {
	logger := l.logger().WithField("method", "GetMaster")
	var list *LeagueList
	if err := l.c.GetInto(fmt.Sprintf(endpointGetMasterLeague, queue), &list); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return list, nil
}

// ListBySummoner returns all leagues a summoner with the given ID is in
func (l *LeagueClient) ListBySummoner(summonerID string) ([]*LeagueItem, error) {
	logger := l.logger().WithField("method", "ListBySummoner")
	var leagues []*LeagueItem
	if err := l.c.GetInto(fmt.Sprintf(endpointGetLeaguesBySummoner, summonerID), &leagues); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return leagues, nil
}

// ListByPuuid returns all leagues a summoner with the given puuid is in
func (l *LeagueClient) ListByPuuid(puuid string) ([]*LeagueItem, error) {
	logger := l.logger().WithField("method", "ListByPuuid")
	var leagues []*LeagueItem
	if err := l.c.GetInto(fmt.Sprintf(endpointGetLeaguesByPuuid, puuid), &leagues); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return leagues, nil
}

// ListPlayers returns all players with a league specified by its queue, tier and division (defaults to page 1)
func (l *LeagueClient) ListPlayers(queue queue, tier tier, division division) ([]*LeagueItem, error) {
	return l.ListPlayersWithPage(queue, tier, division, 1)
}

// ListPlayersWithPage returns all players with a league specified by its queue, tier, division, and page number
func (l *LeagueClient) ListPlayersWithPage(queue queue, tier tier, division division, page int) ([]*LeagueItem, error) {
	logger := l.logger().WithField("method", "ListPlayersWithPage")
	var leagues []*LeagueItem
	pageOpt := internal.WithQueryParam("page", fmt.Sprintf("%d", page))
	if err := l.c.GetInto(fmt.Sprintf(endpointGetLeagues, queue, tier, division), &leagues, pageOpt); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return leagues, nil
}

// Get returns a ranked league with the specified ID
func (l *LeagueClient) Get(leagueID string) (*LeagueList, error) {
	logger := l.logger().WithField("method", "Get")
	var leagues *LeagueList
	if err := l.c.GetInto(fmt.Sprintf(endpointGetLeague, leagueID), &leagues); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return leagues, nil
}

func (l *LeagueClient) logger() log.FieldLogger {
	return l.c.Logger().WithField("category", "league")
}
