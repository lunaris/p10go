package operserv

import (
	"hash/fnv"

	"github.com/fatih/color"
)

func (o *Operserv) debugf(message string, args ...interface{}) {
	o.logger.Debugf(message, withOperservLoggerArguments(o, args)...)
}

func (o *Operserv) infof(message string, args ...interface{}) {
	o.logger.Infof(message, withOperservLoggerArguments(o, args)...)
}

func (o *Operserv) warnf(message string, args ...interface{}) {
	o.logger.Warnf(message, withOperservLoggerArguments(o, args)...)
}

func (o *Operserv) errorf(message string, args ...interface{}) {
	o.logger.Errorf(message, withOperservLoggerArguments(o, args)...)
}

func withOperservLoggerArguments(o *Operserv, args []interface{}) []interface{} {
	return append(
		[]interface{}{
			"clientID",
			o.clientID,
			"nick",
			coloured(o.nick),
		},
		args...,
	)
}

func coloured(s string) string {
	return colourFor(s).Sprint(s)
}

func colourFor(s string) *color.Color {
	hash := fnv.New32a()
	hash.Write([]byte(s))
	sum := hash.Sum32()

	r := int((sum >> 16) & 0xFF) // Get bits 16-23
	g := int((sum >> 8) & 0xFF)  // Get bits 8-15
	b := int(sum & 0xFF)         // Get bits 0-7

	return color.RGB(r, g, b)
}
