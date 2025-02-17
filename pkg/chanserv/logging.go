package chanserv

func (c *Chanserv) debugf(message string, args ...interface{}) {
	c.config.Logger.Debugf(message, args...)
}

func (c *Chanserv) infof(message string, args ...interface{}) {
	c.config.Logger.Infof(message, args...)
}

func (c *Chanserv) warnf(message string, args ...interface{}) {
	c.config.Logger.Warnf(message, args...)
}

func (c *Chanserv) errorf(message string, args ...interface{}) {
	c.config.Logger.Errorf(message, args...)
}
