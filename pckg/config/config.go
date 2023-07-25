package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// holds the application config
type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
	InfoLog       *log.Logger
	InPoduction   bool
	Session       *scs.SessionManager
}
