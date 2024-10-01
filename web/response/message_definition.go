package response

const (
	MSG_SUCCESS = "success"

	MSG_NOT_SUPPORT_FUNCTION              = "not support function"
	MSG_AUTH_PLATFORM_IS_EMPTY            = "platform-name is null"
	MSG_AUTH_API_KEY_IS_EMPTY             = "api-key is null"
	MSG_AUTH_PLATFORM_NOT_EXIST           = "platform-name is not exist"
	MSG_CURRENCY_NOT_SUPPORT              = "currency is not support"
	MSG_INVALID_ACCOUNT_OR_PASSWORD       = "invalid account or password"
	MSG_LOBBY_URL_ERROR                   = "occur lobby url error"
	MSG_GAME_URL_ERROR                    = "occur game url error"
	MSG_INVALID_TIME_FORMAT               = "Time format error"
	MSG_TIME_INTERVAL_IS_TOO_LONG         = "Time range is too long.(Limit: %s)"
	MSG_INVALID_TRANSFER_TYPE             = "transfer type is not support"
	MSG_TRANSFER_FAILED_PLAYER_IS_PLAYING = "transfer out failed, player is playing"
	MSG_AUTH_PROVIDER_IS_EMPTY            = "provider-code is null"
	MSG_AUTH_PROVIDER_NOT_EXIST           = "provider-code is not exist"
	MSG_ACCOUNT_IS_NOT_ACTIVE             = "account is not active"
)

// api
const (
	MSG_PARAMETER_ERROR               = "invalid parameters"
	MSG_ILLEGAL_PARAMETER             = "illegal parameters"
	MSG_API_CENTER_ERROR              = "api center response error"
	MSG_PARSE_API_RESPONSE_ERROR      = "parse api response error"
	MSG_PARSE_API_RESPONSE_DATA_ERROR = "parse api response data error"
	MSG_API_INTERNAL_SERVER_ERROR     = "Internal server error"
	MSG_API_REQUEST_TOO_FREQUENTLY    = "request too frequently"
)

// ws
const (
	MSG_WS_QUERY_PARAMETER_ERROR = "invalid ws query parameters"
	MSG_WS_IS_CLOSED             = "ws is closed"
)

// player
const (
	MSG_PLAYER_RELOGIN                   = "player relogin"
	MSG_PLAYER_RECOVER                   = "player is recover to the game"
	MSG_PLAYER_MONEY_NOT_ENOUGH          = "player's money is not enough"
	MSG_PLAYER_MOENY_IS_OUT_OF_MAX_BET   = "player's bet money is out of max bet"
	MSG_PLAYER_CURRENT_ROUND_HAS_BET     = "player current round has bet"
	MSG_PLAYER_CURRENT_ROUND_NOT_BET     = "player current round not bet"
	MSG_PLAYER_LAST_ROUND_NOT_BET        = "player last round not bet"
	MSG_PLAYER_IS_IN_ANOTHER_GAME        = "player is in another game: %s, room id: %s, room level: %s"
	MSG_PLAYER_ACTION_IN_WRONG_STATE     = "player can't use this action in the state"
	MSG_PLAYER_REQUEST_TOO_FREQUENTLY    = "player request too frequently"
	MSG_PLAYER_ACTOION_ALREADY_DONE      = "player's action already done"
	MSG_PLAYER_ACTOION_IS_NOT_ALLOW      = "player's action is NOT allow"
	MSG_PLAYER_EVENT_IS_NOT_SUPPOERED    = "event is not supported"
	MSG_PLAYER_BET_IS_MUTUALLY_EXCLUSIVE = "player bet is mutually exclusive"
	MSG_PLAYER_IS_IN_ANOTHER_ROOM_LEVEL  = "player is in another room level, game: %s, room level: %s"
	MSG_SYSTEM_IS_MAINTENANCE            = "system maintenance"
	MSG_PLAYER_DISABLE_CARD              = "player can't disable all bingo card"
)

// game
const (
	MSG_ROOM_IS_FINISH        = "room is finish"
	MSG_GAME_IS_NOT_INIT      = "game is not init"
	MSG_NOT_IN_STATE          = "not in %s state, current state: %s"
	MSG_GAME_IN_NOT_BET_STATE = "not in bet state"
	MSG_IS_NOT_FREE_GAME      = "game is not free game"
	MSG_IS_NOT_FEATURE_GAME   = "game is not feature game"
	MSG_DUPLICATE_TOKEN       = "player is duplicate"
	MSG_BET_AREA_IS_LOCK      = "bet area is lock"

	MSG_SINGLE_WALLET_IS_RETRYING_WITHDRAW    = "retrying withdraw"
	MSG_SINGLE_WALLET_MAX_RETRY_LIMIT_REACHED = "the max retry limit reached"
)

// internal
const (
	MSG_MARSHAL_ERROR      = "marshal error"
	MSG_UNMARSHAL_ERROR    = "unmarshal error"
	MSG_CONVERT_TYPE_ERROR = "convert type error"
	MSG_DATA_NOT_FOUND     = "%s doesn't exist %s"
	MSG_INTERNAL_ERROR     = "internal error"
)

// database
const (
	MSG_DATABASE_ABNORMAL      = "database abnormal"
	MSG_DATABASE_NOT_FIND_DATA = "database can not find data"
)

// auth error
const (
	MSG_AUTH_UNAUTHORIZED         = "unauthrized"
	MSG_AUTH_API_KEY_FAILED       = "auth api key failed"
	MSG_AUTH_GENERATE_TOKEN_ERROR = "generate token failed"
	MSG_AUTH_TOKEN_ERROR          = "auth token error"
	MSG_AUTH_GET_TOKEN_ERROR      = "get token error"
	MSG_AUTH_DELETE_TOKEN_ERROR   = "delete token error"
	MSG_AUTH_TRUST_IP_FAILED      = "not allowed ip(%s)."
	MSG_AUTH_CONTACT_ADMIN        = "please contact administrator"
)

// cache
const (
	MSG_CACHE_KEY_COLLISION = "cache key collision"
	MSG_CACHE_PUT_ERROR     = "fail to put cache"
	MSG_CACHE_GET_ERROR     = "fail to get cache"
)

// CMS
const (
	MSG_PERMISSION_IS_NOT_ALLOW                 = "權限不足"
	MSG_CMS_SUPPORT_CURRENCY_NEED_MORE_THAN_ONE = "支援幣別需要大於1個"
	MSG_CMS_TOKEN_IS_EMPTY                      = "Token驗證錯誤"
	MSG_CMS_PLAYER_BALANCE_IS_NOT_ENOUGH        = "玩家餘額不足"
	MSG_CMS_SINGLE_WALLET_CAN_NOT_TRANSFER      = "單一錢包不支援轉帳"
	MSG_CMS_PARAMETERS_IS_REQUIRED              = "參數錯誤"
	MSG_CMS_TIME_FORMAT_ERROR                   = "時間格式錯誤"
	MSG_CMS_TIME_CAN_NOT_MORE_THAN_N_DAY        = "查詢時間超過N天"
	MSG_CMS_CURRENCY_ABNORMAL_QUANTITY          = "幣別數量異常"
	MSG_CMS_CURRENCY_NOT_SUPPORT                = "幣別不支援"
	MSG_CMS_SYST_CAN_NOT_DEPOSIT                = "無法對總控進行買分"
	MSG_CMS_SYST_CAN_NOT_LOGIN_IN_AGENT_SITE    = "請至總控後台登入"
	MSG_CMS_AUTH_TRUST_IP_FAILED                = "not allowed ip(%s)."
	MSG_CMS_TRANSFER_FAILED_PLAYER_IS_PLAYING   = "transfer out failed, player is playing"

	MSG_CMS_ROLE_IS_EXIST = "角色已存在"

	MSG_CMS_CMS_USER_IS_DISABLED          = "用戶已禁用"
	MSG_CMS_CMS_USER_OTP_FAIL             = "OTP驗證失敗"
	MSG_CMS_CMS_USER_CAN_NOT_OPERATE_SELF = "不可對自身操作"
	MSG_CMS_CMS_USER_IS_EXIST             = "帳戶已存在"
	MSG_CMS_CMS_USER_ACCOUNT_PWD_ERROR    = "用戶名或密碼錯誤"
	MSG_CMS_CMS_USER_OLD_PWD_ERROR        = "密碼輸入有誤請重新輸入"

	MSG_AGENT_MONEY_IS_NOT_ENOUGH                      = "餘額不足"
	MSG_CMS_AGENT_IS_EXIST                             = "代理已存在"
	MSG_CMS_AGENT_CAN_NOT_EDIT_SELF_COMMERCIAL_MODE    = "不可修改自身合作模式請諮詢上級代理或客服"
	MSG_CMS_AGENT_CURRENT_OCCUPY_NEED_MORE_THAN_PARENT = "當前新代理點位不可低於上級點位"
	MSG_CMS_AGENT_CAN_NOT_EDIT_SELF_CURRENCY           = "不可修改自身幣別請諮詢上級代理或客服"
	MSG_CMS_AGENT_COMMERCIAL_MODE_SUPPORT_TYPE         = "合作模式僅支援:買分網和信用網"
	MSG_CMS_AGENT_WALLET_MODE_SUPPORT_TYPE             = "錢包模式僅支援:額轉錢包和單一錢包"
	MSG_CMS_AGENT_CAN_NOT_LOGIN_IN_STST_SITE           = "發生錯誤請諮詢客服"
	MSG_CMS_AGENT_NOT_SUPPORT_CURRENCY                 = "代理不支援該幣別"
	MSG_CMS_AGENT_TRANSFER_RECORD_IS_NOT_DEPOSIT       = "此筆非上分紀錄"
	MSG_CMS_AGENT_IP_IS_EXIST                          = "IP已存在"
	MSG_CMS_AGENT_FILE_EXPORT_FAIL                     = "匯出失敗"
	MSG_CMS_AGENT_FILE_UPLOAD_FAIL                     = "上傳失敗"
	MSG_CMS_AGENT_CAN_NOT_OPERATE_SELF                 = "不可修改自身代理帳號狀態請諮詢上級代理或客服"

	MSG_CMS_SERVER_ERROR_BY_DATABASE        = "伺服器錯誤"
	MSG_CMS_SERVER_ERROR_BY_REDIS           = "伺服器錯誤"
	MSG_CMS_SERVER_ERROR_BY_JSON            = "伺服器錯誤"
	MSG_CMS_SERVER_ERROR_BY_FILE_TRAVERSE   = "伺服器錯誤"
	MSG_CMS_SERVER_ERROR_BY_FILE_PERMISSION = "伺服器錯誤"
)

const (
	MSG_SYSTEM_MAINTENANCE_MODE = "game in maintenance"
)

// Mailbox
const (
	MSG_MAILBOX_AWARD_IS_REDEEMED  = "Award is redeemed"
	MSG_MAILBOX_AWARD_IS_NOT_EXIST = "Award is not exist"
)
