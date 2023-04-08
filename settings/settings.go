package settings

import (
	//agregamos _ para la libreria no sea requerida
	_ "embed"
	"encoding/json"
)

/*utilizamos la libreria embed desde los comentarios para
inyectar el archivo settings.json en el slice settingsFile */

//go:embed settings.json
var settingsFile []byte

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Settings struct {
	Port int            `json:"port"`
	DB   DatabaseConfig `json:"database"`
}

// New crea una nueva instancia de Settings con valores predeterminados.
// Los valores predeterminados se cargan desde un archivo json.
func New() (*Settings, error) {
	// Crea una nueva instancia de Settings
	var s Settings
	// Deserializa el archivo de configuración json en la nueva instancia
	err := json.Unmarshal(settingsFile, &s)
	// Verifica si hay errores
	if err != nil {
		// Devuelve el error si se produjo alguno
		return nil, err
	}
	// Devuelve la nueva instancia de Settings si no se produjeron errores
	return &s, nil
}
