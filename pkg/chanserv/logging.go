package chanserv

import (
	"hash/fnv"

	"github.com/fatih/color"
)

func (c *Chanserv) debugf(message string, args ...interface{}) {
	c.logger.Debugf(message, withChanservLoggerArguments(c, args)...)
}

func (c *Chanserv) infof(message string, args ...interface{}) {
	c.logger.Infof(message, withChanservLoggerArguments(c, args)...)
}

func (c *Chanserv) warnf(message string, args ...interface{}) {
	c.logger.Warnf(message, withChanservLoggerArguments(c, args)...)
}

func (c *Chanserv) errorf(message string, args ...interface{}) {
	c.logger.Errorf(message, withChanservLoggerArguments(c, args)...)
}

func withChanservLoggerArguments(c *Chanserv, args []interface{}) []interface{} {
	return append(
		[]interface{}{
			"clientID",
			c.clientID,
			"nick",
			coloured(c.nick),
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
