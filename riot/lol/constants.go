package lol

const (
	endpointBase                                = "/lol"
	endpointMasteryBase                         = endpointBase + "/champion-mastery/v4"
	endpointMasteriesBase                       = endpointMasteryBase + "/champion-masteries"
	endpointGetChampionMasteriesByPuuid         = endpointMasteriesBase + "/by-puuid/%s"
	endpointGetChampionMasteryByPuuid           = endpointMasteriesBase + "/by-puuid/%s/by-champion/%s"
	endpointGetChampionMasteriesTopByPuuid      = endpointMasteriesBase + "/by-puuid/%s/top?count=%d"
	endpointGetChampionMasteryTotalScoreByPuuid = endpointMasteryBase + "/scores/by-puuid/%s"
	endpointChallengesBase                      = endpointBase + "/challenges/v1"
	endpointChallengesBaseChallenges            = endpointChallengesBase + "/challenges"
	endpointChallengesConfig                    = endpointChallengesBaseChallenges + "/config"
	endpointChallengesPercentiles               = endpointChallengesBaseChallenges + "/percentiles"
	endpointChallengesConfigByChallengeID       = endpointChallengesBaseChallenges + "/%d/config"
	endpointChallengesLeaderboardsBase          = endpointChallengesBaseChallenges + "/%d/leaderboards"
	endpointChallengesLeaderboards              = endpointChallengesLeaderboardsBase + "/by-level/%s?limit=%d"
	endpointChallengesPercentilesByChallengeID  = endpointChallengesBaseChallenges + "/%d/percentiles"
	endpointChallengesPlayerDataByPUUID         = endpointChallengesBase + "/player-data/%s"
	endpointPlatformBase                        = endpointBase + "/platform/v3"
	endpointGetFreeChampionRotation             = endpointPlatformBase + "/champion-rotations"
	endpointLeagueBase                          = endpointBase + "/league/v4"
	endpointGetChallengerLeague                 = endpointLeagueBase + "/challengerleagues/by-queue/%s"
	endpointGetGrandmasterLeague                = endpointLeagueBase + "/grandmasterleagues/by-queue/%s"
	endpointGetMasterLeague                     = endpointLeagueBase + "/masterleagues/by-queue/%s"
	endpointGetLeaguesByPuuid                   = endpointLeagueBase + "/entries/by-puuid/%s"
	endpointGetLeagues                          = endpointLeagueBase + "/entries/%s/%s/%s"
	endpointGetLeague                           = endpointLeagueBase + "/leagues/%s"
	endpointLeagueExpBase                       = endpointBase + "/league-exp/v4"
	endpointGetLeagueExpEntries                 = endpointLeagueExpBase + "/entries/%s/%s/%s"
	endpointStatusBase                          = endpointBase + "/status/v4"
	endpointGetStatus                           = endpointStatusBase + "/platform-data"
	endpointMatchBase                           = endpointBase + "/match/v5"
	endpointGetMatchIDsBase                     = endpointMatchBase + "/matches/by-puuid"
	endpointGetMatchIDs                         = endpointGetMatchIDsBase + "/%s/ids?start=%d&count=%d"
	endpointGetMatch                            = endpointMatchBase + "/matches/%s"
	endpointGetMatchTimeline                    = endpointMatchBase + "/matches/%s/timeline"
	endpointGetMatchReplays                     = endpointMatchBase + "/matches/by-puuid/%s/replays"
	endpointSummonerBase                        = endpointBase + "/summoner/v4"
	endpointGetSummonerByPUUID                  = endpointSummonerBase + "/summoners/by-puuid/%s"
	endpointGetSummonerMe                       = endpointSummonerBase + "/summoners/me"
	endpointSpectatorBase                       = endpointBase + "/spectator/v5"
	endpointGetCurrentGame                      = endpointSpectatorBase + "/active-games/by-puuid/%s"
	endpointGetFeaturedGames                    = endpointSpectatorBase + "/featured-games"
	endpointTournamentStubBase                  = endpointBase + "/tournament-stub/v5"
	endpointCreateStubTournamentCodes           = endpointTournamentStubBase + "/codes?count=%d&tournamentId=%d"
	endpointGetStubLobbyEvents                  = endpointTournamentStubBase + "/lobby-events/by-code/%s"
	endpointCreateStubTournamentProvider        = endpointTournamentStubBase + "/providers"
	endpointCreateStubTournament                = endpointTournamentStubBase + "/tournaments"
	endpointTournamentBase                      = endpointBase + "/tournament/v5"
	endpointCreateTournamentCodes               = endpointTournamentBase + "/codes?count=%d&tournamentId=%d"
	endpointGetLobbyEvents                      = endpointTournamentBase + "/lobby-events/by-code/%s"
	endpointCreateTournamentProvider            = endpointTournamentBase + "/providers"
	endpointCreateTournament                    = endpointTournamentBase + "/tournaments"
	endpointGetTournament                       = endpointTournamentBase + "/codes/%s"
	endpointUpdateTournament                    = endpointTournamentBase + "/codes/%s"
	endpointGetThirdPartyCode                   = endpointBase + "/platform/v4/third-party-code/by-puuid/%s"
)

type identification string

const (
	identificationName       identification = "name"
	identificationAccountID  identification = "account"
	identificationPUUID      identification = "puuid"
	identificationSummonerID identification = "summonerID"
)

type queue string

// All possible queues
const (
	QueueRankedSolo            queue = "RANKED_SOLO_5x5"
	QueueRankedFlex            queue = "RANKED_FLEX_SR"
	QueueRankedTwistedTreeline queue = "RANKED_FLEX_TT"
)

type tier string

// All possible Tiers
const (
	TierIron        tier = "IRON"
	TierBronze      tier = "BRONZE"
	TierSilver      tier = "SILVER"
	TierGold        tier = "GOLD"
	TierPlatinum    tier = "PLATINUM"
	TierEmerald     tier = "EMERALD"
	TierDiamond     tier = "DIAMOND"
	TierMaster      tier = "MASTER"
	TierGrandMaster tier = "GRANDMASTER"
	TierChallenger  tier = "CHALLENGER"
)

type division string

// All possible divisions
const (
	DivisionOne   division = "I"
	DivisionTwo   division = "II"
	DivisionThree division = "III"
	DivisionFour  division = "IV"
)

var (
	// Queues is a list of all available queue types
	Queues = []queue{
		QueueRankedSolo,
		QueueRankedFlex,
		QueueRankedTwistedTreeline,
	}

	// Tiers is a list of all available tiers
	Tiers = []tier{
		TierIron,
		TierBronze,
		TierSilver,
		TierGold,
		TierPlatinum,
		TierEmerald,
		TierDiamond,
	}

	// Divisions is a list of all available divisions
	Divisions = []division{
		DivisionOne,
		DivisionTwo,
		DivisionThree,
		DivisionFour,
	}
)
