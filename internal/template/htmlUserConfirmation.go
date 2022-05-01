package template

import "fmt"

// GenerateHTMLUserConfirmationTemplate will generate email template which will act
// as confirmation on user email address. Supply the parameter with full link
// to redirect user
func GenerateHTMLUserConfirmationTemplate(link string) string {
	return fmt.Sprintf(
		`<!DOCTYPE html>
			<html>
				<head>
					<title>Email Account User Confirmation</title>
				</head>
				<body>
					<h1>Please Confirm Your Email as a User of <strong>Lucky's Mail Service</strong></h1>
					<p>Hei there! You receive this email because youre registering to my service using this email</p>
					<p>Click on the button below to confirm this is your registered email</p>
					<p>If you wasn't regitering with this email, simply just do nothing, and nothing will happen on your email too</p>
					<a href="%s">
						<button>Confirm My Email</button>
					</a>
				</body>
			</html>
	`, link)
}
