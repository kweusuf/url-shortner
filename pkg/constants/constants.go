package constants

import (
	"os"
	"time"
)

const (
	API_V1 = "/api/v1"

	SEPARATOR = "/"

	CONF_FILE = "file"
	IN_MEM_DB = "inmemdb"
	DB        = "db"

	Config = "config"
	Id     = "id"

	// InternalConfFile = utils.GetCWD() + string(os.PathSeparator) + "data" + string(os.PathSeparator) + "conf.yml"

	ContentType                           = "Content-Type"
	ApplicationJSONContentTypeWithCharset = "application/json; charset=utf-8"
	ApplicationJSONContentType            = "application/json"
	Bearer                                = "Bearer "
	Authorization                         = "Authorization"

	PathSeparator = string(os.PathSeparator)

	BaseShortenedUrl = "https://my-domain.com/s/%s" //TODO: check if this can be moved to config

	ExpirationInterval = 24 * time.Hour   //TODO: make this configurable
	CleanupInterval    = 15 * time.Minute //TODO: make this configurable
)

type CtxKeyAppConfig struct{}
