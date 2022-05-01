package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	sib "github.com/sendinblue/APIv3-go-library/lib"
	"github.com/sirupsen/logrus"
)

func PgConnString() string {
	host := os.Getenv("PGHOST")
	db := os.Getenv("PGDATABASE")
	user := os.Getenv("PGUSER")
	pw := os.Getenv("PGPASSWORD")
	port := os.Getenv("PGPORT")

	if os.Getenv("ENV") == "production" {
		host = "host.docker.internal" // in the production, will be using docker.
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pw, db, port)

	return dsn
}

func ServerPort() string {
	port := os.Getenv("SERVER_PORT")

	if port == "" {
		return "5000" // default port
	}

	return port
}

func LogLevel() string {
	level := os.Getenv("LOG_LEVEL")

	if level == "" {
		return "debug"
	}

	return level
}

func GetSIBKey() string {
	cfg := os.Getenv("SIB_API_KEY")

	if cfg == "" {
		logrus.Error("SIB_API_KEY is not set. Mailing utility is disabled")
	}

	return cfg
}

func SIBClient() *sib.APIClient {
	cfg := sib.NewConfiguration()
	cfg.AddDefaultHeader("api-key", GetSIBKey())

	return sib.NewAPIClient(cfg)
}

func SIBRegisteredEmail() string {
	return os.Getenv("SIB_REGISTERED_EMAIL")
}

func SMTPServer() string {
	cfg := os.Getenv("SMTP_SERVER")

	if cfg == "" {
		logrus.Error("SMTP_SERVER value is not set. Mailing utility is broken")
	}

	return cfg
}

func SMPTPort() int {
	cfg, err := strconv.Atoi(os.Getenv("SMPT_PORT"))
	if err != nil {
		return 587
	}

	return cfg
}

func SMTPKey() string {
	cfg := os.Getenv("SMTP_KEY")
	if cfg == "" {
		logrus.Error("SMTP_SERVER value is not set. Mailing utility is broken")
	}

	return cfg
}

func SMTPAddress() string {
	return fmt.Sprintf("%s:%d", SMTPServer(), SMPTPort())
}

func FreeMailerWorkerSleepDuration() time.Duration {
	cfg := os.Getenv("FREE_MAILER_WORKER_SLEEP_DURATION_SEC")
	if cfg == "" {
		return DEFAULT_FREE_MAILER_WORKER_SLEEP_DURATION
	}

	sec, err := strconv.Atoi(cfg)
	if err != nil {
		return DEFAULT_FREE_MAILER_WORKER_SLEEP_DURATION
	}

	return time.Duration(sec * int(time.Second))
}

func FreeMailProcessingLimit() int {
	cfg, err := strconv.Atoi(os.Getenv("FREE_MAIL_PROCESSING_LIMIT_PER_DURATION"))
	if err != nil {
		return DEFAULT_FREE_EMAIL_LIST_QUERY_LIMIT
	}

	return cfg
}

func PrivateKey() string {
	cfg := os.Getenv("PRIVATE_KEY")
	if cfg == "" {
		logrus.Panic("PRIVATE_KEY is not provided")
	}

	return cfg
}

func IvKey() string {
	cfg := os.Getenv("IV_KEY")
	if cfg == "" {
		logrus.Panic("IV_KEY is not provided")
	}

	return cfg
}

func GetFreeMailUserEmailAddress() string {
	cfg := os.Getenv("FREE_MAIL_USER_EMAIL_ADDRESS")
	if cfg == "" {
		logrus.Warn("Free mail user are registered with empty string")
	}

	return cfg
}

func GetFreeMailUserPassword() string {
	return os.Getenv("FREE_MAIL_USER_PASSWORD")
}

func GetEmailConfirmationCode() string {
	cfg := os.Getenv("EMAIL_ACCOUNT_CONFIRMATION_LINK")
	if cfg == "" {
		return DEFAULT_EMAIL_CONFIRMATION_LINK
	}

	return cfg
}
