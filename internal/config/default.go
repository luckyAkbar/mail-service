package config

import "time"

var (
	DEFAULT_FREE_MAILER_WORKER_SLEEP_DURATION = 60 * time.Minute
	DEFAULT_FREE_EMAIL_LIST_QUERY_LIMIT       = 5
	DEFAULT_EMAIL_CONFIRMATION_LINK           = "https://mail.service.luckyakbar.tech/account/confirmation?code="
)
