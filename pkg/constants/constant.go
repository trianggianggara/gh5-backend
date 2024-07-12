package constants

const (
	APP       = "APP"
	PORT      = "PORT"
	ENV       = "ENV"
	VERSION   = "VERSION"
	HOST      = "HOST"
	SCHEME    = "SCHEME"
	JWT_KEY   = "JWT_KEY"
	LOG_LEVEL = "LOG_LEVEL"
	WORK_DIR  = "WORK_DIR"

	CRON_ENABLED = "CRON_ENABLED"

	SENTRY_DSN = "SENTRY_DSN"
)

const (
	DB_DEFAULT_CREATED_BY = "system"
	DB_HOST               = "DB_HOST"
	DB_USER               = "DB_USER"
	DB_PASS               = "DB_PASS"
	DB_PORT               = "DB_PORT"
	DB_NAME               = "DB_NAME"
	DB_CHARSET            = "DB_CHARSET"
	DB_PARSE_TIME         = "DB_PARSE_TIME"
	DB_LOC                = "DB_LOC"
	DB_SSLMODE            = "DB_SSLMODE"
	DB_TZ                 = "DB_TZ"
	DB_GH5_BACKEND        = "gh5-backend"

	SERVICE_ACCOUNT_FILENAME = "SERVICE_ACCOUNT_FILENAME"

	GOOGLE_PROJECT_ID  = "GOOGLE_PROJECT_ID"
	GOOGLE_BUCKET_NAME = "GOOGLE_BUCKET_NAME"

	MIGRATION_ENABLED = "MIGRATION_ENABLED"
	SEEDER_ENABLED    = "SEEDER_ENABLED"
)

type (
	contextKey string
	reqIDKey   string
)

const (
	CONTEXT_KEY contextKey = "context_key"
	REQ_ID_KEY  reqIDKey   = "req_id_key"
)
