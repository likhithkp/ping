package ses

import (
	"context"
	"fmt"

	"github.com/likhithkp/ping/utils/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

type Mailer struct {
	client *ses.Client
	sender *string
}

func NewMailer(client *ses.Client, env *config.Env) *Mailer {
	return &Mailer{client: client, sender: &env.SenderEmail}
}

func (m *Mailer) SendEmail(ctx context.Context, to, otp string) error {
	input := &ses.SendEmailInput{
		Source: m.sender,
		Destination: &types.Destination{
			ToAddresses: []string{*m.sender},
		},
		Message: &types.Message{
			Subject: &types.Content{
				Data:    OtpEmailBodyHtml(otp),
				Charset: aws.String("UTF-8"),
			},
			Body: &types.Body{
				Html: &types.Content{
					Data:    OtpEmailBodyText(otp),
					Charset: aws.String("UTF-8"),
				},
				Text: &types.Content{
					Data:    OtpEmailBodyText(otp),
					Charset: aws.String("UTF-8"),
				},
			},
		},
	}

	_, err := m.client.SendEmail(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}

func OtpEmailBodyText(otp string) *string {
	s := fmt.Sprintf(`## Requested OTP

Please use the 6-digit OTP below.

<strong style="font-size: 130%%">%v</strong>

If you didn’t request this, you can ignore this email.

Thanks,
***Fit Near***`, otp)

	return aws.String(s)
}

func OtpEmailBodyHtml(otp string) *string {
	s := `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Reset your password and login</title>
		<!--[if mso]><style type="text/css">body, table, td, a { font-family: Arial, Helvetica, sans-serif !important; }</style><![endif]-->
	</head>
	<body style="font-family: Helvetica, Arial, sans-serif; margin: 0px; padding: 0px; background-color: #ffffff;">
		<table role="presentation"
			style="width: 100%%; border-collapse: collapse; border: 0px; border-spacing: 0px; font-family: Arial, Helvetica, sans-serif; background-color: rgb(255, 255, 255);">
			<tbody>
				<tr>
					<td align="center" style="padding: 1rem 2rem; vertical-align: top; width: 100%%;">
						<table role="presentation" style="max-width: 600px; border-collapse: collapse; border: 0px; border-spacing: 0px; text-align: left;">
							<tbody>
								<tr>
									<td style="padding: 40px 0px 0px;">
										<div style="text-align: center;">
											<div style="padding-bottom: 20px;"><img src="https://i.ibb.co/Qbnj4mz/logo.png" alt="Fit Near" style="width: 80px;"></div>
										</div>
										<div style="padding: 20px; background-color: rgb(237, 237, 237);">
											<div style="color: rgb(19, 17, 18); text-align: center;">
												<h3>Reset Password</h3>
												<p style="padding-bottom: 16px">Please use the 6-digit OTP below to reset password.</p>
												<center style="padding-top: 8px; padding-bottom: 20px;"><strong style="font-size: 170%%;">%v</strong></center>

												<p style="padding-bottom: 16px">If you didn’t request this, you can ignore this email.</p>

												<p style="padding-bottom: 16px">Thanks,<br><em><strong>Fit Near</strong></em></p>
											</div>
										</div>
										<div style="padding-top: 20px; color: rgb(254, 98, 98); text-align: center;">
											<p style="padding-bottom: 16px">Made with ♥ in Netherlands</p>
										</div>
									</td>
								</tr>
							</tbody>
						</table>
					</td>
				</tr>
			</tbody>
		</table>
	</body>
</html>`
	res := fmt.Sprintf(s, otp)
	return aws.String(res)
}
