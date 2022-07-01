package log

import (
	"context"
	"io"
	"log"

	//	"log/syslog"
	"os"

	"gentree/config"
)

type key int

const logKey key = 93

//parseFacility configures the syslog facility
/*
func parseFacility(fac string) syslog.Priority {
	switch fac {
	case "kern":
		return syslog.LOG_KERN
	case "user":
		return syslog.LOG_USER
	case "mail":
		return syslog.LOG_MAIL
	case "daemon":
		return syslog.LOG_DAEMON
	case "auth":
		return syslog.LOG_AUTH
	case "syslog":
		return syslog.LOG_SYSLOG
	case "lpr":
		return syslog.LOG_LPR
	case "news":
		return syslog.LOG_NEWS
	case "uucp":
		return syslog.LOG_UUCP
	case "cron":
		return syslog.LOG_CRON
	case "authpriv":
		return syslog.LOG_AUTHPRIV
	case "ftp":
		return syslog.LOG_FTP
	case "local0":
		return syslog.LOG_LOCAL0
	case "local1":
		return syslog.LOG_LOCAL1
	case "local2":
		return syslog.LOG_LOCAL2
	case "local3":
		return syslog.LOG_LOCAL3
	case "local4":
		return syslog.LOG_LOCAL4
	case "local5":
		return syslog.LOG_LOCAL5
	case "local6":
		return syslog.LOG_LOCAL6
	case "local7":
		return syslog.LOG_LOCAL7
	}
	return syslog.LOG_KERN
}
*/

//Writer sets the syslog
func Writer(c *config.Config) (io.Writer, error) {
	//	if c.Log.Logger == "syslog" {
	//		facility := parseFacility(c.Log.Facility)
	//		return syslog.New(facility,	c.Log.Program)
	//	}
	return os.Stdout, nil
}

//FromConfig starts the logger
func FromConfig(c *config.Config) *log.Logger {
	//Inicializar aqui o debug e syslog posteriormente
	writer, err := Writer(c)
	if err != nil {
		log.Print(err)
		writer = os.Stdout
	}
	return log.New(writer, "", log.Ldate|log.Ltime|log.Lshortfile)
}

// NewContext generates a new Context storing the logger into its values.
// Thats helpfull if you need to transfer the logger inside the context
// to another function.
func NewContext(ctx context.Context, l *log.Logger) context.Context {
	return context.WithValue(ctx, logKey, l)
}

// FromContext retrieves a *logger previously added to the context by the NewContext func.
// It returns nil if no logger is found.
func FromContext(ctx context.Context) (*log.Logger, bool) {
	s, ok := ctx.Value(logKey).(*log.Logger)
	return s, ok
}
