package logs

import (
	"encoding/json"
	"fmt"
	"log"
	"mgmt/common/configs"
	"os"
	"reflect"
	"strings"
	"syscall"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogs() {
	filePath := configs.Get(configs.SECTION_LOG, configs.LOG_FILE_PATH, configs.LOG_DEFAULT_PATH)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if createErr := os.MkdirAll(filePath, 0755); createErr != nil {
			log.Panicln("Failed to create log file path.")
		}
	}

	files := configs.Get(configs.SECTION_LOG, configs.LOG_FILE, configs.LOG_DEFAULT_FILE)
	fileList := strings.Split(files, ",")

	for _, file := range fileList {
		if file == LOG_FILE_SYSTEM {
			systemLogger = newLogger(fmt.Sprintf("%s/%s", filePath, LOG_FILE_SYSTEM))
		}

		if file == LOG_FILE_RECORD {
			recordLogger = newLogger(fmt.Sprintf("%s/%s", filePath, LOG_FILE_RECORD))
		}

		if file == LOG_FILE_CMS {
			cmsLogger = newLogger(fmt.Sprintf("%s/%s", filePath, LOG_FILE_CMS))
		}
	}

	if isEnableDebugLog() {
		level.SetLevel(zap.DebugLevel)
	} else {
		level.SetLevel(zap.InfoLevel)
	}

	initSystemLog()
	initPanicLog()
}

func isEnableDebugLog() bool {
	return configs.YES == configs.Get(configs.SECTION_LOG, configs.LOG_ENABLE_DEBUG_LOG, configs.NO)
}

func initSystemLog() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Turn on panic log in linux environment
func initPanicLog() {
	if configs.NO == configs.Get(configs.SECTION_LOG, configs.LOG_PANIC_TO_FILE, configs.NO) {
		return
	}

	fileName := configs.Get(configs.SECTION_LOG, configs.LOG_FILE_PATH, configs.LOG_DEFAULT_PATH) + "/" + configs.Get(configs.SECTION_LOG, "panic_file", "panic.log")
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Panicln(err)
		return
	}
	if err = syscall.Dup2(int(file.Fd()), int(os.Stderr.Fd())); err != nil {
		log.Panicln(err)
		return
	}
}

func newLogger(filepath string) *zap.Logger {
	fWriter, _, err := zap.Open(filepath)
	if err != nil {
		fmt.Println("newLogger err:", err.Error())
		os.Exit(1)
	}

	writer := zapcore.AddSync(fWriter)

	fileEncoderConfig := zapcore.EncoderConfig{
		TimeKey:        TIME_KEY,
		LevelKey:       LEVEL_KEY,
		NameKey:        NAME_KEY,
		CallerKey:      CALLER_KEY,
		MessageKey:     MESSAGE_KEY,
		StacktraceKey:  STACKTRACE_KEY,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	stdEncoderConfig := zapcore.EncoderConfig{
		TimeKey:        TIME_KEY,
		LevelKey:       LEVEL_KEY,
		NameKey:        NAME_KEY,
		CallerKey:      CALLER_KEY,
		MessageKey:     MESSAGE_KEY,
		StacktraceKey:  STACKTRACE_KEY,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	level = zap.NewAtomicLevel()

	if isEnableDebugLog() {
		level.SetLevel(zap.DebugLevel)
	}

	cores := make([]zapcore.Core, 0)
	cores = append(cores, zapcore.NewCore(
		zapcore.NewJSONEncoder(fileEncoderConfig),
		writer,
		level,
	))

	if configs.MODE_DEVELOPMENT == configs.Get(configs.SECTION_SYSTEM, configs.SYSTEM_APP_MODE, configs.MODE_PRODUCTION) {
		if configs.YES == configs.Get(configs.SECTION_LOG, configs.LOG_ENABLE_STD_OUT, configs.NO) {
			cores = append(cores, zapcore.NewCore(
				zapcore.NewJSONEncoder(stdEncoderConfig),
				zapcore.Lock(os.Stdout),
				level,
			))
		}
	}

	core := zapcore.NewTee(cores...)
	caller := zap.AddCaller()

	return zap.New(core, caller, zap.AddCallerSkip(1))
}

func getLogger(logType string) *zap.Logger {
	var outputLogger *zap.Logger
	switch logType {
	case LOG_TYPE_CMS:
		outputLogger = cmsLogger
	case LOG_TYPE_RECORD:
		outputLogger = recordLogger
	default:
		outputLogger = systemLogger
	}

	return outputLogger
}

func Debug(logType, logKey, msg string, payload interface{}, fields ...Field) {
	if isEnableDebugLog() {
		zapFields := transferMappingToFields(logKey, fields)
		zapFields = append(zapFields,
			zap.Any(FIELD_KEY_PAYLOAD, payload),
			zap.String(FIELD_KEY_FUNC_NAME, getFuncCallerName(BASE_SKIP_LAYER)), //Skip "Debug" layer
			zap.String(FIELD_KEY_FUNC_CALLER_STACK, combinFuncCallerName(getAllFuncCallerNameList(BASE_SKIP_LAYER))),
		)
		if log := getLogger(logType); log != nil {
			log.Debug(msg, zapFields...)
		}
	}
}

func Info(logType, logKey, msg string, payload interface{}, fields ...Field) {
	zapFields := transferMappingToFields(logKey, fields)
	zapFields = append(zapFields,
		zap.Any(FIELD_KEY_PAYLOAD, payload),
		zap.String(FIELD_KEY_FUNC_NAME, getFuncCallerName(BASE_SKIP_LAYER)), //Skip "Info" layer
	)

	if log := getLogger(logType); log != nil {
		log.Info(msg, zapFields...)
	}
}

func Error(logType, logKey, msg string, payload interface{}, fields ...Field) {
	zapFields := transferMappingToFields(logKey, fields)
	zapFields = append(zapFields,
		zap.Any(FIELD_KEY_PAYLOAD, payload),
		zap.String(FIELD_KEY_FUNC_NAME, getFuncCallerName(BASE_SKIP_LAYER)), //Skip "Error" layer
		zap.String(FIELD_KEY_FUNC_CALLER_STACK, combinFuncCallerName(getAllFuncCallerNameList(BASE_SKIP_LAYER))),
	)

	if logger := getLogger(logType); nil != logger {
		logger.Error(msg, zapFields...)
	}
}

func Record(logKey string, betRecordList interface{}) {
	info := map[string]interface{}{
		FIELD_KEY_BET_RECORD: betRecordList,
	}
	zapFields := transferMappingToFields(logKey, []Field{})
	zapFields = append(zapFields,
		zap.Any(FIELD_KEY_PAYLOAD, info),
		zap.String(FIELD_KEY_FUNC_NAME, getFuncCallerName(1)),
	)

	if log := getLogger(LOG_TYPE_RECORD); log != nil {
		log.Info(SAVE_BET_RECORD, zapFields...)
	}
}

func transferMappingToFields(logKey string, fields []Field) []zap.Field {
	zapFields := []zap.Field{
		zap.String(FIELD_KEY_LOG_KEY, logKey),
	}
	for _, field := range fields {
		if nil != field.Data {
			switch t := reflect.TypeOf(field.Data).Kind(); t {
			case reflect.Int:
				zapFields = append(zapFields, zap.Int(field.Name, field.Data.(int)))
			case reflect.String:
				zapFields = append(zapFields, zap.String(field.Name, field.Data.(string)))
			case reflect.Bool:
				zapFields = append(zapFields, zap.Bool(field.Name, field.Data.(bool)))
			default:
				if jsonData, err := json.Marshal(field.Data); err == nil {
					zapFields = append(zapFields, zap.String(field.Name, string(jsonData)))
				} else {
					systemLogger.Error(err.Error())
				}
			}
		} else {
			zapFields = append(zapFields, zap.Any(field.Name, field.Data))
		}
	}

	return zapFields
}
