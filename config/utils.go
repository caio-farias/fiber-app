package config

import (
	"os"
	"strings"
	"text/template"
)

type String string
func (s String) Format(data map[string]interface{}) (out string, err error) {
    t := template.Must(template.New("").Parse(string(s)))
    builder := &strings.Builder{}
    if err = t.Execute(builder, data); err != nil {
        return
    }
    out = builder.String()
    return
}

func MountDbURL() string {
	const template = "host={{.dbHost}} user={{.dbUser}} password={{.dbPassword}} dbname={{.dbName}} port={{.dbPort}} sslmode=disable"
	data := map[string]interface {} {
		"dbHost": os.Getenv("USERS_DB_HOSTNAME"),
		"dbPort": os.Getenv("USERS_DB_PORT"),
		"dbUser": os.Getenv("USERS_DB_USER"),
		"dbPassword": os.Getenv("USERS_DB_PASSWORD"),
		"dbName": os.Getenv("USERS_DB_NAME"),
	}
	s, _ := String(template).Format(data)
	return s 
}