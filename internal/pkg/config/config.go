package config

import (
	"net/url"
	"time"
)

type App struct {
	Url url.URL `env:"APP_URL,required"`
}
type Session struct {
	Secret   string        `env:"SESSION_SECRET,required"`
	Lifetime time.Duration `env:"SESSION_LIFETIME" envDefault:"120m"` // the duration before it expires(with unit)
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
	App      App
	Session  Session
	Database Database
	Mail     Mail
}
