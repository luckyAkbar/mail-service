package handler

type SuccessRegisterFreeEmail struct {
	OK            bool   `json:"ok"`
	Receipient    string `json:"receipient"`
	MailingListID int    `json:"mailingListID"`
}

type FailedRegisterEmail struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}
