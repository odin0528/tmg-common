package cache

const (
	WOW_MUTEX_PREFIX string = "wow_mutex:"

	BET_RECORD_KEY string = "bet_record"

	MUTEX_DURATION_SECOND = 20

	JWT_SECRET_KEY          string = "jwt_secret_key"
	JWT_SECRET_KEY_PREVIOUS string = "jwt_secret_key_previous"
	JWT_ACCOUNT_TOKEN       string = "jwt_account_token_"
	JWT_TOKEN_ACCOUNT       string = "jwt_token_account_"
)

type PlayingRoomInfo struct {
	PfAccountList []string `json:"pf_account_list"`
	Oid           int      `json:"oid"`
}

type SingleWalletWithdrawCacheInfo struct {
	RecordID  string  `json:"record_id"`
	PfID      string  `json:"pf_id"`
	PfSubPfID string  `json:"pf_sub_pf_id"`
	PfAccount string  `json:"pf_account"`
	Currency  string  `json:"currency"`
	GameCode  string  `json:"game_code"`
	Oid       int     `json:"oid"`
	Money     float64 `json:"money"`
	RoundID   string  `json:"round_id"`
	RoomNo    string  `json:"room_no"`
	CreatedAt int64   `json:"created_at"`
}

type SWRoundCacheInfo struct {
	WithdrawRecordIDs    []string `json:"withdraw_record_id_list"`
	DepositRecordID      string   `json:"deposit_record_id"`
	PfSubPfID            string   `json:"pf_sub_pf_id"`
	PfAccount            string   `json:"pf_account"`
	Currency             string   `json:"currency"`
	RoundID              string   `json:"round_id"`
	RoomNo               string   `json:"room_no"`
	GameCode             string   `json:"game_code"`
	Oid                  int      `json:"oid"`
	TotalMoney           float64  `json:"money"`
	RepeatTimes          int      `json:"repeat_times"`
	IsFinishedSettlement bool     `json:"is_finished_settlement"`
}

// TODO 再整理 package
type WalletCacheInfo struct {
	WithdrawMoney float64 `json:"withdraw_money"`
}
