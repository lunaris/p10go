package client

func (c *P10Client) debugf(message string, args ...interface{}) {
	c.logger.Debugf(message, args...)
}

func (c *P10Client) infof(message string, args ...interface{}) {
	c.logger.Infof(message, args...)
}

func (c *P10Client) warnf(message string, args ...interface{}) {
	c.logger.Warnf(message, args...)
}

func (c *P10Client) errorf(message string, args ...interface{}) {
	c.logger.Errorf(message, args...)
}
