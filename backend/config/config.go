package config

var Config = map[string]string{
	"PORT":                   "8080",
	"JWT_KEY":                "5201314",
	"SSL_ENABLE":             "false",
	"SSL_CERT":               "cert.pem",
	"SSL_KEY":                "key.pem",
	"DB_USER":                "floodguard",
	"DB_HOST":                "localhost",
	"DB_NAME":                "floodguard",
	"DB_PASSWORD":            "floodguard",
	"SMTP_SERVER":            "",
	"SMTP_PORT":              "",
	"SMTP_MAIL":              "",
	"SMTP_PASSWORD":          "",
	"PERMISSION_USER":        "false",
	"PERMISSION_REGION":      "false",
	"PERMISSION_FLOODEVENT":  "false",
	"PERMISSION_HISTORYDATA": "false",
	"PERMISSION_NOTICE":      "false",
	"PERMISSION_COMMENT":     "false",
	"PERMISSION_SENSOR":      "false",
}
