package configs

import (
	"sync"

	"gopkg.in/ini.v1"
)

const (
	// Section
	SECTION_SYSTEM  = "system"
	SECTION_WEB_API = "web_api"
	SECTION_MOCK    = "mock"
	SECTION_LOG     = "log"

	// Key
	SYSTEM_HTTP_PORT           string = "http_port"
	SYSTEM_APP_MODE            string = "app_mode"
	SYSTEM_ID                  string = "id"
	SYSTEM_MATCH_WORKER_AMOUNT string = "match_worker_amount"

	WEB_API_API_CENTER_DOMAIN string = "api_center_domain"
	WEB_API_ALLOW_CORS_DOMAIN string = "allowed_cors_domain"

	MOCK_PLAYER             = "mock_player"
	MOCK_ENABLE_MOCK_PLAYER = "enable_mock_player"
	MOCK_ENABLE_MOCK_RESULT = "enable_mock_result"

	LOG_FILE_PATH        = "file_path"
	LOG_FILE             = "file"
	LOG_ENABLE_STD_OUT   = "enable_std_out" //value rename
	LOG_ENABLE_DEBUG_LOG = "enable_debug_log"
	LOG_KEYS             = "log_keys"
	LOG_PANIC_TO_FILE    = "panic_to_file"

	// Value
	MODE_PRODUCTION  string = "production"
	MODE_DEVELOPMENT string = "development"
	MODE_DEBUG       string = "debug"

	LOG_DEFAULT_PATH = "storage/logs"
	LOG_DEFAULT_FILE = "system.log"

	YES = "yes"
	NO  = "no"
)

var configMutex sync.RWMutex
var envConfig *ini.File
