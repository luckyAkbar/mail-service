package utils

import (
	"fmt"
	"mail-service/internal/db"
	"mail-service/internal/models"
)

func RegisterSingleMailingList(receipient, mailerName, body string, priority, userID int) (int, error) {
	mailingList := &models.MailingList{
		Receipient: receipient,
		MailerName: mailerName,
		Payload:    body,
		Priority:   priority,
		UserID:     userID,
	}

	if db.DB.Create(mailingList).Error != nil {
		return 0, fmt.Errorf("failed to register to mailing queue for receipient: %s", receipient)
	}

	return int(mailingList.ID), nil
}
