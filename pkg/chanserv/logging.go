package chanserv

func (c *Chanserv) debugf(message string, args ...interface{}) {
	c.logger.Debugf(message, args...)
}

func (c *Chanserv) infof(message string, args ...interface{}) {
	c.logger.Infof(message, args...)
}

func (c *Chanserv) warnf(message string, args ...interface{}) {
	c.logger.Warnf(message, args...)
}

func (c *Chanserv) errorf(message string, args ...interface{}) {
	c.logger.Errorf(message, args...)
}
