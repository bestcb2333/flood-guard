package config

var defaultConfig = map[string]any{
	"port":                   8080,
	"JwtKey":                 "5201314",
	"db.user":                "floodguard",
	"db.host":                "localhost",
	"db.dbname":              "floodguard",
	"db.password":            "floodguard",
	"useSSL":                 true,
	"ssl.cert":               "mcax.cn.pem",
	"ssl.key":                "mcax.cn.key",
	"permission.user":        false,
	"permission.region":      false,
	"permission.floodevent":  false,
	"permission.historydata": false,
	"permission.notice":      false,
	"permission.comment":     false,
}
