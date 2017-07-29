package main

// Payload struct
type Payload struct {
	Session    string  `json:"session"`
	MatchList  []Match `json:"match_list"`
	SummonerID int64   `json:"summoner_id"`
	AccountID  int64   `json:"account_id"`
	Phase      string  `json:"phase"`
	ChannelID  string  `json:"channel_id"`
	Role       string  `json:"role"`
}

// Match  struct
type Match struct {
	Timestamp int64 `json:"timestamp"`
	MatchID   int64 `json:"match_id"`
}
