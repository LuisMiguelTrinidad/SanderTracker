package logging

// StatusColor retorna el código ANSI basado en el código HTTP.
func StatusColor(status int) string {
	switch status {
	case 500:
		return "\x1b[31m" // rojo
	case 400:
		return "\x1b[35m" // magenta
	case 300:
		return "\x1b[33m" // amarillo
	case 200:
		return "\x1b[32m" // verde
	case 100:
		return "\x1b[34m" // azul
	default:
		return "\x1b[0m" // reset
	}
}

// MethodColor retorna el código ANSI basado en el método HTTP.
func MethodColor(method string) string {
	switch method {
	case "GET":
		return "\x1b[36m" // cian
	case "POST":
		return "\x1b[32m" // verde
	case "PUT":
		return "\x1b[33m" // amarillo
	case "DELETE":
		return "\x1b[31m" // rojo
	default:
		return "\x1b[0m"
	}
}

// LevelColor retorna el código ANSI basado en el nivel del log.
func LevelColor(level string) string {
	switch level {
	case "DEBUG":
		return "\x1b[36m" // cian
	case "INFO":
		return "\x1b[32m" // verde
	case "WARN":
		return "\x1b[33m" // amarillo
	case "ERROR":
		return "\x1b[31m" // rojo
	case "FATAL":
		return "\x1b[35m" // magenta
	default:
		return "\x1b[0m"
	}
}
