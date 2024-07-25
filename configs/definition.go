package configs

import (
	"sync"

	"gopkg.in/ini.v1"
)

const (
	SECTION_SYSTEM                    = "system"
	SECTION_WEB_API                   = "web_api"
	SECTION_MOCK                      = "mock"
	SECTION_LOG                       = "log"
	SECTION_DATABASE_REAL_TIME        = "database_real_time"
	SECTION_DATABASE_RECORD           = "database_record"
	SECTION_DATABASE_SCRIPT           = "database_script"
	SECTION_DATABASE_REAL_TIME_SYSTEM = "database_real_time_system"
	SECTION_MONGO_DATABASE_RECORD     = "mongo_database_record"
	SECTION_CACHE                     = "cache"
	SECTION_BOT                       = "bot"
	SECTION_DATABASE_REAL_TIME_READ   = "database_real_time_read"
	SECTION_DATABASE_RECORD_READ      = "database_record_read"
	SECTION_DATABASE_SCRIPT_READ      = "database_script_read"
	SECTION_DATABASE_REAL_TIME_WRITE  = "database_real_time_write"
	SECTION_DATABASE_RECORD_WRITE     = "database_record_write"

	SYSTEM_HTTP_PORT             string = "http_port"
	SYSTEM_APP_MODE              string = "app_mode"
	SYSTEM_ID                    string = "id"
	SYSTEM_MATCH_WORKER_AMOUNT   string = "match_worker_amount"
	SYSTEM_USE_CACHE             string = "use_cache"
	SYSTEM_ENABLE_BET_CHECK      string = "enble_bet_check"
	SYSTEM_WEBSOCKET_ENCODE_MODE string = "websocket_encode_mode"
	SYSTEM_WEBSOCKET_DECODE_MODE string = "websocket_decode_mode"
	SYSTEN_ENABLE_BATTLE_BOT     string = "enable_battle_bot"
	SYSTEM_ENABLE_HUNDRED_BOT    string = "enable_hundred_bot"
	SYSTEM_ENABLE_LOBBY_OVERALL  string = "enable_lobby_overall"
	SYSTEM_EXCLUDE_CURRENCIES    string = "exclude_currencies"

	WEB_API_API_CENTER_DOMAIN             string = "api_center_domain"
	WEB_API_ALLOW_CORS_DOMAIN             string = "allowed_cors_domain"
	WEB_API_ALLOW_CORS_HEADERS            string = "allowed_cors_headers"
	WEB_API_JWT_SECRET_KEY_DURATION_HOURS string = "jwt_secret_key_duration_hours"
	WEB_API_AUTH_TOKEN_DURATION_HOURS     string = "auth_token_duration_hours"
	WEB_API_ENABLE_TRUST_IP               string = "enable_trust_ip"
	WEB_API_MIDDLE_PATH                   string = "middle_path"

	MOCK_PLAYER                   = "mock_player"
	MOCK_ENABLE_MOCK_PLAYER       = "enable_mock_player"
	MOCK_ENABLE_MOCK_RESULT       = "enable_mock_result"
	MOCK_CURRENCY                 = "TEST"
	MOCK_ENABLE_MOCK_GAME_SETTING = "enable_mock_game_setting"

	CACHE_ENGINE     = "engine"
	CACHE_PREFIX_KEY = "prefix_key"
	CACHE_HOST       = "host"
	CACHE_PORT       = "port"
	CACHE_DB_NUM     = "db_num"
	CACHE_PASSWORD   = "password"
	CACHE_USE_TLS    = "use_tls"

	LOG_FILE_PATH        = "file_path"
	LOG_FILE             = "file"
	LOG_ENABLE_STD_OUT   = "enable_std_out"
	LOG_ENABLE_DEBUG_LOG = "enable_debug_log"
	LOG_KEYS             = "log_keys"
	LOG_PANIC_TO_FILE    = "panic_to_file"
	LOG_ENABLE_DAILY     = "enable_daily"

	MODE_PRODUCTION  string = "production"
	MODE_DEVELOPMENT string = "development"
	MODE_DEBUG       string = "debug"

	DB_ACCOUNT                  string = "account"
	DB_PASSWORD                 string = "password"
	DB_HOST                     string = "host"
	DB_PORT                     string = "port"
	DB_NAME                     string = "db_name"
	DB_MAX_OPEN_CONNECTIONS     string = "max_open_connections"
	DB_MAX_IDLE_CONNECTIONS     string = "max_idle_connections"
	DB_MAX_CONNECTIONS_LIFETIME string = "max_connections_lifetime"

	BOT_ENABLE_TELEGRAM_ALARM_RTP       string = "enable_telegram_alarm_rtp"
	BOT_TELEGRAM_ALARM_RTP_CHAT_ROOM_ID string = "telegram_alarm_rtp_chat_room_id"
	BOT_TELEGRAM_ALARM_RTP_BOT_TOKEN    string = "telegram_alarm_rtp_bot_token"
	BOT_ENABLE_TELEGRAM_ALARM           string = "enable_telegram_alarm"
	BOT_TELEGRAM_ALARM_CHAT_ROOM_ID     string = "telegram_alarm_chat_room_id"
	BOT_TELEGRAM_ALARM_BOT_TOKEN        string = "telegram_alarm_bot_token"

	LOG_DEFAULT_PATH = "storage/logs"
	LOG_DEFAULT_FILE = "system.log"

	YES = "yes"
	NO  = "no"

	ENABLE_GUEST = "enable_guest"

	ENABLE_WS_UNSPECIFIED int = 0
	ENABLE_WS_BASE64      int = 1
	ENABLE_WS_MSG_PACK    int = 2
)

var configMutex sync.RWMutex
var envConfig *ini.File
