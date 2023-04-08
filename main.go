package main

import (
	"log"

	"github.com/matiaslenci/inventory-manager/settings"
	"go.uber.org/fx"
)

func main() {
	//Crea una nueva instancia de la aplicación, que se utiliza para configurar y ejecutar la aplicación.
	app := fx.New(
		//Define los proveedores de dependencias, en este caso "settings.New", que es una función que crea una nueva instancia de la estructura de configuración definida en el paquete "settings".
		fx.Provide(
			settings.New,
		),
		// Define los invocadores de dependencias con una función anónima que toma la estructura de configuración creada anteriormente y la registra en el registro de logs.
		fx.Invoke(
			func(s *settings.Settings) {
				log.Println(s)
			},
		),
	)

	//Ejecuta la aplicación configurada. Cuando se ejecuta, la aplicación se encargará de resolver las dependencias y ejecutar los invocadores de dependencias en el orden apropiado.
	app.Run()
}
