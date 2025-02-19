package config

type Session struct {
	Secret   string `env:"SESSION_SECRET,required"`
	Lifetime uint8  `env:"SESSION_LIFETIME" envDefault:"120"` // number of minutes before it expires
}
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
	Session  Session
	Database Database
	Mail     Mail
}
