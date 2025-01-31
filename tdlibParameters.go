package tdlib

// TdlibParameters Contains parameters for TDLib initialization
type TdlibParameters struct {
	tdCommon
	UseTestDc              bool   `json:"use_test_dc"`              // If set to true, the Telegram test environment will be used instead of the production environment
	DatabaseDirectory      string `json:"database_directory"`       // The path to the directory for the persistent database; if empty, the current working directory will be used
	FilesDirectory         string `json:"files_directory"`          // The path to the directory for storing files; if empty, database_directory will be used
	UseFileDatabase        bool   `json:"use_file_database"`        // If set to true, information about downloaded and uploaded files will be saved between application restarts
	UseChatInfoDatabase    bool   `json:"use_chat_info_database"`   // If set to true, the library will maintain a cache of users, basic groups, supergroups, channels and secret chats. Implies use_file_database
	UseMessageDatabase     bool   `json:"use_message_database"`     // If set to true, the library will maintain a cache of chats and messages. Implies use_chat_info_database
	UseSecretChats         bool   `json:"use_secret_chats"`         // If set to true, support for secret chats will be enabled
	ApiId                  int32  `json:"api_id"`                   // Application identifier for Telegram API access, which can be obtained at https://my.telegram.org
	ApiHash                string `json:"api_hash"`                 // Application identifier hash for Telegram API access, which can be obtained at https://my.telegram.org
	SystemLanguageCode     string `json:"system_language_code"`     // IETF language tag of the user's operating system language; must be non-empty
	DeviceModel            string `json:"device_model"`             // Model of the device the application is being run on; must be non-empty
	SystemVersion          string `json:"system_version"`           // Version of the operating system the application is being run on. If empty, the version is automatically detected by TDLib
	ApplicationVersion     string `json:"application_version"`      // Application version; must be non-empty
	EnableStorageOptimizer bool   `json:"enable_storage_optimizer"` // If set to true, old files will automatically be deleted
	IgnoreFileNames        bool   `json:"ignore_file_names"`        // If set to true, original file names will be ignored. Otherwise, downloaded files will be saved under names as close as possible to the original name
}

// MessageType return the string telegram-type of TdlibParameters
func (tdlibParameters *TdlibParameters) MessageType() string {
	return "tdlibParameters"
}

// NewTdlibParameters creates a new TdlibParameters
//
// @param useTestDc If set to true, the Telegram test environment will be used instead of the production environment
// @param databaseDirectory The path to the directory for the persistent database; if empty, the current working directory will be used
// @param filesDirectory The path to the directory for storing files; if empty, database_directory will be used
// @param useFileDatabase If set to true, information about downloaded and uploaded files will be saved between application restarts
// @param useChatInfoDatabase If set to true, the library will maintain a cache of users, basic groups, supergroups, channels and secret chats. Implies use_file_database
// @param useMessageDatabase If set to true, the library will maintain a cache of chats and messages. Implies use_chat_info_database
// @param useSecretChats If set to true, support for secret chats will be enabled
// @param apiId Application identifier for Telegram API access, which can be obtained at https://my.telegram.org
// @param apiHash Application identifier hash for Telegram API access, which can be obtained at https://my.telegram.org
// @param systemLanguageCode IETF language tag of the user's operating system language; must be non-empty
// @param deviceModel Model of the device the application is being run on; must be non-empty
// @param systemVersion Version of the operating system the application is being run on. If empty, the version is automatically detected by TDLib
// @param applicationVersion Application version; must be non-empty
// @param enableStorageOptimizer If set to true, old files will automatically be deleted
// @param ignoreFileNames If set to true, original file names will be ignored. Otherwise, downloaded files will be saved under names as close as possible to the original name
func NewTdlibParameters(useTestDc bool, databaseDirectory string, filesDirectory string, useFileDatabase bool, useChatInfoDatabase bool, useMessageDatabase bool, useSecretChats bool, apiId int32, apiHash string, systemLanguageCode string, deviceModel string, systemVersion string, applicationVersion string, enableStorageOptimizer bool, ignoreFileNames bool) *TdlibParameters {
	tdlibParametersTemp := TdlibParameters{
		tdCommon:               tdCommon{Type: "tdlibParameters"},
		UseTestDc:              useTestDc,
		DatabaseDirectory:      databaseDirectory,
		FilesDirectory:         filesDirectory,
		UseFileDatabase:        useFileDatabase,
		UseChatInfoDatabase:    useChatInfoDatabase,
		UseMessageDatabase:     useMessageDatabase,
		UseSecretChats:         useSecretChats,
		ApiId:                  apiId,
		ApiHash:                apiHash,
		SystemLanguageCode:     systemLanguageCode,
		DeviceModel:            deviceModel,
		SystemVersion:          systemVersion,
		ApplicationVersion:     applicationVersion,
		EnableStorageOptimizer: enableStorageOptimizer,
		IgnoreFileNames:        ignoreFileNames,
	}

	return &tdlibParametersTemp
}
