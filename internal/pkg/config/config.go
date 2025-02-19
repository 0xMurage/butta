package config

type Database struct {
	Url string `env:"DATABASE_URL,notEmpty"`
}

type MailGun struct {
	Domain   string `env:"MAILGUN_DOMAIN"`
	Secret   string `env:"MAILGUN_SECRET"`
	Endpoint string `env:"MAILGUN_ENDPOINT" envDefault:"https://api.eu.mailgun.net/v3"`
}

type Mail struct {
	MailGun
	From string `env:"MAIL_FROM"`
}

type Config struct {
	Database Database
	Mail     Mail
}
