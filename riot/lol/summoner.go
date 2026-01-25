package lol

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/KnutZuidema/golio/internal"
)

// SummonerClient provides methods for the summoner endpoints of the League of Legends API.
type SummonerClient struct {
	c *internal.Client
}

// GetByPUUID returns the summoner with the given PUUID
func (s *SummonerClient) GetByPUUID(puuid string) (*Summoner, error) {
	logger := s.logger().WithField("method", "GetByPUUID")
	var summoner *Summoner
	if err := s.c.GetInto(fmt.Sprintf(endpointGetSummonerByPUUID, puuid), &summoner); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return summoner, nil
}

// GetMe returns the summoner for the given access token
func (s *SummonerClient) GetMe(accessToken string) (*Summoner, error) {
	logger := s.logger().WithField("method", "GetMe")
	var summoner *Summoner
	if err := s.c.GetInto(
		endpointGetSummonerMe,
		&summoner,
		internal.WithHeader("Authorization", "Bearer "+accessToken),
	); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return summoner, nil
}

func (s *SummonerClient) logger() log.FieldLogger {
	return s.c.Logger().WithField("category", "summoner")
}
