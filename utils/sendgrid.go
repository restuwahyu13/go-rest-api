package util

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendGridMail(name, email, subject, token string) {
	from := mail.NewEmail("admin", "admin@unindra.com")
	to := mail.NewEmail(name, email)
	// plainTextContent := "register new account"
	htmlContent := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
			<title>Tech Soft</title>
			<meta charset="utf-8">
			<meta name="viewport" content="width=device-width, initial-scale=1">
			<style type="text/css">
					html, body {
							font-family: "Roboto Thin";
							clear: both;
							margin: auto;
							padding: auto;
					}
					.container {
							position: relative !important;
							margin: auto !important;
							text-align: center;
							width: 650px;
							height: 650px;
							border-radius: 10px;
					}
					.container .card {
							justify-content: center;
							align-content: center;
							position: relative;
							top: 30px;
					}
					.container .logo {
							background: #07cd10;
							width: 360px;
							height: 40px;
							border-radius: 10px;
							position: relative;
							margin: auto;
							text-align: center;
							font-weight: bold;
					}
					.logo h4 a {
							font-size: 22px;
							text-align: center;
							line-height: 40px;
							color: #f5f5f5;
							opacity: 1;
					}
					.card-title {
							font-size: 19px;
							line-height: 30px;
					}
					.card-subtitle {
							font-size: 17px;
					}
					.list-accout {
							font-size: 16px;
							position: relative;
							top: 20px;
					}
					.button-spotify {
							position: relative;
							justify-content: center;
					}
					a {
							display: inline-block;
							position: relative;
							margin: auto;
							text-decoration: none;
							color: white;
							font-weight: bold;
							outline: none;
							box-shadow: none;
							font-size: 15px;
					}
					button {
							width: 150px;
							height: 50px;
							border-radius: 10px;
							background: #07cd10;
							position: relative;
							top: 10px;
							font-weight: bold;
							font-size: 18px;
							color: #f5f5f5;
							opacity: 1;
					}
					.text-content {
							font-size: 16px;
							word-wrap: break-word;
							position: relative;
							top: 55px;
					}
					.footer-logo{
							background: #07cd10;
							color: #f5f5f5;
							width: 360px;
							height: 40px;
							border-radius: 15px;
							position: relative;
							margin: auto;
							text-align: center;
							top: 3vh;
							font-weight: bold;
					}
					.footer {
							font-size: 17px;
							line-height: 40px;
							text-align: center;
							opacity: 1;
					}
			</style>
	</head>
	<body>
	<div class="container">
			<div class="card">
					<div class="logo">
							<h4><a href="">Campus Portal.inc</a></h4>
					</div>
					<div class="card-body">
							<p class="card-title"><strong>Hello Dear ${to}</strong></p>
									<p class="card-subtitle"><strong>Kepada user YTH </strong>silahkan konfirmasi account anda:
							</p>
					<div class="text-content">
							<button>
								<a href="${CLIENT_URL}/api/v1/user/activation/${token}">Activation Account</a>
							</button>
							<div class="footer-logo">
									<span class="footer">&copy; ${new Date().getFullYear()} Campus, Inc All Right Reserved</span>
									</div>
							</div>
					</div>
			</div>
		</body>
</html>
	`

	message := mail.NewSingleEmail(from, subject, to, "", htmlContent)
	client := sendgrid.NewSendClient(GodotEnv("SG_API_KEY"))
	client.Send(message)
}
