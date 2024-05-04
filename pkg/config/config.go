package config

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
	"os"
)

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}

func WorkingDir() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Working directory: ", dir)
	}
}
