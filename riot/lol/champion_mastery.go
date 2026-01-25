package lol

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/KnutZuidema/golio/internal"
)

// ChampionMasteryClient provides methods for the champion mastery endpoints of the
// League of Legends API.
type ChampionMasteryClient struct {
	c *internal.Client
}

// List returns information about masteries for the summoner with the given ID
func (c *ChampionMasteryClient) List(summonerID string) ([]*ChampionMastery, error) {
	logger := c.logger().WithField("method", "List")
	var masteries []*ChampionMastery
	if err := c.c.GetInto(
		fmt.Sprintf(endpointGetChampionMasteries, summonerID),
		&masteries,
	); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return masteries, nil
}

// Get returns information about the mastery of the champion with the given ID the summoner with the
// given ID has
func (c *ChampionMasteryClient) Get(summonerID, championID string) (*ChampionMastery, error) {
	logger := c.logger().WithField("method", "Get")
	var mastery *ChampionMastery
	if err := c.c.GetInto(
		fmt.Sprintf(endpointGetChampionMastery, summonerID, championID),
		&mastery,
	); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return mastery, nil
}

// GetTotal returns the accumulated mastery score of all champions played by the summoner with the
// given ID
func (c *ChampionMasteryClient) GetTotal(summonerID string) (int, error) {
	logger := c.logger().WithField("method", "GetTotal")
	var score int
	if err := c.c.GetInto(fmt.Sprintf(endpointGetChampionMasteryTotalScore, summonerID), &score); err != nil {
		logger.Debug(err)
		return 0, err
	}
	return score, nil
}

// ListByPuuid returns information about masteries for the summoner with the given PUUID
func (c *ChampionMasteryClient) ListByPuuid(puuid string) ([]*ChampionMastery, error) {
	logger := c.logger().WithField("method", "ListByPuuid")
	var masteries []*ChampionMastery
	if err := c.c.GetInto(
		fmt.Sprintf(endpointGetChampionMasteriesByPuuid, puuid),
		&masteries,
	); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return masteries, nil
}

// GetByPuuid returns information about the mastery of the champion with the given ID for the summoner with the given PUUID
func (c *ChampionMasteryClient) GetByPuuid(puuid, championID string) (*ChampionMastery, error) {
	logger := c.logger().WithField("method", "GetByPuuid")
	var mastery *ChampionMastery
	if err := c.c.GetInto(
		fmt.Sprintf(endpointGetChampionMasteryByPuuid, puuid, championID),
		&mastery,
	); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return mastery, nil
}

// GetTopByPuuid returns the top champion masteries for the summoner with the given PUUID
func (c *ChampionMasteryClient) GetTopByPuuid(puuid string, count int) ([]*ChampionMastery, error) {
	logger := c.logger().WithField("method", "GetTopByPuuid")
	var masteries []*ChampionMastery
	if err := c.c.GetInto(
		fmt.Sprintf(endpointGetChampionMasteriesTopByPuuid, puuid, count),
		&masteries,
	); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return masteries, nil
}

// GetTotalByPuuid returns the accumulated mastery score of all champions played by the summoner with the given PUUID
func (c *ChampionMasteryClient) GetTotalByPuuid(puuid string) (int, error) {
	logger := c.logger().WithField("method", "GetTotalByPuuid")
	var score int
	if err := c.c.GetInto(fmt.Sprintf(endpointGetChampionMasteryTotalScoreByPuuid, puuid), &score); err != nil {
		logger.Debug(err)
		return 0, err
	}
	return score, nil
}

func (c *ChampionMasteryClient) logger() log.FieldLogger {
	return c.c.Logger().WithField("category", "champion mastery")
}
