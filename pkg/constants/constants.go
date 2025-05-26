package constants

import "os"

const (
	API = "/api/v1"

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
)

type CtxKeyAppConfig struct{}
