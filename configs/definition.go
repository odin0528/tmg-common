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
	SECTION_DATABASE_REAL_TIME_SLAVE  = "database_real_time_slave"
	SECTION_DATABASE_RECORD           = "database_record"
	SECTION_DATABASE_RECORD_SLAVE     = "database_record_slave"
	SECTION_DATABASE_SCRIPT           = "database_script"
	SECTION_DATABASE_SCRIPT_SLAVE     = "database_script_slave"
	SECTION_DATABASE_REAL_TIME_SYSTEM = "database_real_time_system"
	SECTION_MONGO_DATABASE_RECORD     = "mongo_database_record"
	SECTION_CACHE                     = "cache"
	SECTION_BOT                       = "bot"
	SECTION_DATABASE_REAL_TIME_READ   = "database_real_time_read"
	SECTION_DATABASE_RECORD_READ      = "database_record_read"
	SECTION_DATABASE_SCRIPT_READ      = "database_script_read"
	SECTION_DATABASE_REAL_TIME_WRITE  = "database_real_time_write"
	SECTION_DATABASE_RECORD_WRITE     = "database_record_write"
	SECTION_GEOIP                     = "geoip"
	SECTION_RABBIT_MQ                 = "rabbit_mq"
	SECTION_AWS_S3                    = "aws_s3"

	SYSTEM_HTTP_HOST                       string = "host"
	SYSTEM_HTTP_PORT                       string = "http_port"
	SYSTEM_APP_MODE                        string = "app_mode"
	SYSTEM_ID                              string = "id"
	SYSTEM_MATCH_WORKER_AMOUNT             string = "match_worker_amount"
	SYSTEM_USE_CACHE                       string = "use_cache"
	SYSTEM_ENABLE_BET_CHECK                string = "enble_bet_check"
	SYSTEM_WEBSOCKET_ENCODE_MODE           string = "websocket_encode_mode"
	SYSTEM_WEBSOCKET_DECODE_MODE           string = "websocket_decode_mode"
	SYSTEN_ENABLE_BATTLE_BOT               string = "enable_battle_bot"
	SYSTEM_ENABLE_SLOT_BATTLE_BOT          string = "enable_slot_battle_bot"
	SYSTEM_ENABLE_HUNDRED_BOT              string = "enable_hundred_bot"
	SYSTEM_ENABLE_LOBBY_OVERALL            string = "enable_lobby_overall"
	SYSTEM_ENABLE_AI_AGENT                 string = "enable_ai_agent"
	SYSTEM_EXCLUDE_CURRENCIES              string = "exclude_currencies"
	SYSTEM_HUNDRED_ROOM_TYPE               string = "hundred_room_type"
	SYSTEM_ENABLE_PRESET_ROOM              string = "enable_preset_room"
	SYSTEM_ENABLE_HUNDRED_ROOM_GROUP_IDS   string = "enable_hundred_room_group_ids"
	SYSTEM_ENABLE_DEMO_TOOL                string = "enable_demo_tool"
	SYSTEM_ENABLE_WATCHDOG                 string = "enable_watchdog"
	SYSTEM_SERVICE_NAME                    string = "service_name"
	SYSTEM_SERVICE_NAME_WOW_GAMING         string = "wow_gaming"
	SYSTEM_SERVICE_NAME_AI_LIVE_CASINO     string = "ai_live_casino"
	SYSTEM_SERVICE_NAME_MAGIC_POKER        string = "magic_poker"
	SYSTEM_ENABLE_CHECK_JP_THRESHOLD       string = "enable_check_jp_threshold"
	SYSTEM_ENABLE_TEST_MODE                string = "enable_test_mode"
	SYSTEM_ENABLE_RTP_TRACKING             string = "enable_rtp_tracking"
	SYSTEM_ENV                             string = "env"
	SYSTEM_BET_RECORD_SPLIT_START_DATETIME string = "bet_record_split_start_datetime"
	SYSTEM_NAMESPACE                       string = "namespace"
	SYSTEM_SERVICE_NAMESPACE               string = "service_namespace"
	SYSTEM_PRODUCT_NAME                    string = "product_name"

	WEB_API_API_CENTER_DOMAIN                  string = "api_center_domain"
	WEB_API_GAME_SERVER_DOMAIN                 string = "game_server_domain"
	WEB_API_SCRIPT_SERVER_DOMAIN               string = "script_server_domain"
	WEB_API_ALLOW_CORS_DOMAIN                  string = "allowed_cors_domain"
	WEB_API_SLOT_MACHINE_DOMAIN                string = "slot_machine_domain"
	WEB_API_ALLOW_CORS_HEADERS                 string = "allowed_cors_headers"
	WEB_API_CORS_ENABLE                        string = "cors_enable"
	WEB_API_JWT_SECRET_KEY_DURATION_HOURS      string = "jwt_secret_key_duration_hours"
	WEB_API_AUTH_TOKEN_DURATION_HOURS          string = "auth_token_duration_hours"
	WEB_API_ENABLE_TRUST_IP                    string = "enable_trust_ip"
	WEB_API_MIDDLE_PATH                        string = "middle_path"
	WEB_API_MIDDLE_PATH_WOW_GAMING             string = "wow_gaming"
	WEB_API_MIDDLE_PATH_AI_LIVE_CASINO         string = "ai_live_casino"
	WEB_API_CALC_GAME_CURRENT_RTP_PERIOD_DAY   string = "calc_game_current_rtp_period_day"
	WEB_API_GAME_SERVER_IPS                    string = "game_server_ips"
	WEB_API_APP_PROFILE                        string = "app_profile"
	WEB_API_APP_PROFILE_WOW_GAMING_PROD        string = "wow_gaming_prod"
	WEB_API_APP_PROFILE_WOW_GAMING_STAGING     string = "wow_gaming_staging"
	WEB_API_APP_PROFILE_AI_LIVE_CASINO_PROD    string = "ai_live_casino_prod"
	WEB_API_APP_PROFILE_AI_LIVE_CASINO_STAGING string = "ai_live_casino_staging"
	WEB_API_ENABLE_AUTH_API_CENTER_TOKEN       string = "enable_auth_api_center_token"

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

	CACHE_KEY_CLEAN_REDIS_MUTEX_MAP_TIME_MIN  = "clean_redis_mutex_map_time_min"
	CACHE_KEY_EXPIRY_REDIS_MUTEX_MAP_TIME_MIN = "expiry_redis_mutex_map_time_min"

	LOG_FILE_PATH          = "file_path"
	LOG_FILE               = "file"
	LOG_ENABLE_STD_OUT     = "enable_logger_std_out"
	LOG_ENABLE_DEBUG_LOG   = "enable_debug_log"
	LOG_KEYS               = "log_keys"
	LOG_PANIC_TO_FILE      = "panic_to_file"
	LOG_ENABLE_DAILY       = "enable_daily"
	LOG_ENABLE_FILE_OUTPUT = "enable_file_output"

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
	DB_SLAVE_ENABLE_REAL_TIME   string = "db_slave_enable_real_time"
	DB_SLAVE_ENABLE_RECORD      string = "db_slave_enable_record"
	DB_SLAVE_ENABLE_SCRIPT      string = "db_slave_enable_script"

	BOT_ENABLE_TELEGRAM_ALARM_RTP               string = "enable_telegram_alarm_rtp"
	BOT_TELEGRAM_ALARM_RTP_CHAT_ROOM_ID         string = "telegram_alarm_rtp_chat_room_id"
	BOT_TELEGRAM_ALARM_RTP_BOT_TOKEN            string = "telegram_alarm_rtp_bot_token"
	BOT_ENABLE_TELEGRAM_ALARM                   string = "enable_telegram_alarm"
	BOT_TELEGRAM_ALARM_CHAT_ROOM_ID             string = "telegram_alarm_chat_room_id"
	BOT_TELEGRAM_ALARM_BOT_TOKEN                string = "telegram_alarm_bot_token"
	BOT_ENABLE_TELEGRAM_ALARM_BUY_FEATURE       string = "enable_telegram_alarm_buy_feature"
	BOT_TELEGRAM_ALARM_BUY_FEATURE_CHAT_ROOM_ID string = "telegram_alarm_buy_feature_chat_room_id"
	BOT_TELEGRAM_ALARM_BUY_FEATURE_BOT_TOKEN    string = "telegram_alarm_buy_feature_bot_token"
	BOT_ENABLE_TELEGRAM_ALARM_INTERNAL          string = "enable_telegram_alarm_internal"
	BOT_TELEGRAM_ALARM_INTERNAL_CHAT_ROOM_ID    string = "telegram_alarm_internal_chat_room_id"
	BOT_TELEGRAM_ALARM_INTERNAL_BOT_TOKEN       string = "telegram_alarm_internal_bot_token"

	MQ_PROTOCOL                                 string = "protocol"
	MQ_USERNAME                                 string = "username"
	MQ_PASSWORD                                 string = "password"
	MQ_HOST                                     string = "host"
	MQ_PORT                                     string = "port"
	MQ_PRODUCER_DEFAULT_COUNT                   string = "producer_default_count"
	MQ_CONSUMER_DEFAULT_COUNT                   string = "consumer_default_count"
	MQ_CONSUMER_BATCH_SIZE_SETTLEMENT           string = "consumer_batch_size_settlement"
	MQ_CONSUMER_BATCH_TIMEOUT_SECOND_SETTLEMENT string = "consumer_batch_timeout_second_settlement"
	MQ_CONSUMER_BATCH_SIZE_RECORD               string = "consumer_batch_size_record"
	MQ_CONSUMER_BATCH_TIMEOUT_SECOND_RECORD     string = "consumer_batch_timeout_second_record"
	MQ_CONSUMER_BATCH_SIZE_REPORT               string = "consumer_batch_size_report"
	MQ_CONSUMER_BATCH_TIMEOUT_SECOND_REPORT     string = "consumer_batch_timeout_second_report"

	AWS_S3_ACCESS_KEY_ID          string = "aws_s3_access_key_id"
	AWS_S3_SECRET_ACCESS_KEY      string = "aws_s3_secret_access_key"
	AWS_S3_REGION                 string = "aws_s3_region"
	AWS_S3_BUCKET                 string = "aws_s3_bucket"
	AWS_S3_FOLDER_PLAYER_PORTRAIT string = "aws_s3_folder_player_portrait"

	ENABLE_CHECK_BET_RECORD_LASTET_UPDATED_AT string = "enable_check_bet_record_lastet_updated_at"

	GEOIP_FILE_PATH string = "file_path"

	LOG_DEFAULT_PATH = "storage/logs"
	LOG_DEFAULT_FILE = "system.log"

	FEATURE_DISABLE_KICK_PLAYER = "disable_kick_player"

	YES = "yes"
	NO  = "no"

	ENABLE_GUEST = "enable_guest"

	ENABLE_WS_UNSPECIFIED      int = 0
	ENABLE_WS_BASE64           int = 1
	ENABLE_WS_MSG_PACK         int = 2
	ENABLE_WS_MIX_BASE64_SHIFT int = 3

	GAME_EANBLE_FORCED_CONTROL_TYPE string = "enable_forced_control_type"

	SPIN_MODE string = "spin_mode"

	SPIN_MODE_NATURAL int = 0
	SPIN_MODE_SCRIPT  int = 1
)

var (
	configMutex sync.RWMutex
	envConfig   *ini.File
)
