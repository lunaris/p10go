package client

func (c *P10Client) debugf(message string, args ...interface{}) {
	c.config.Logger.Debugf(message, args...)
}

func (c *P10Client) infof(message string, args ...interface{}) {
	c.config.Logger.Infof(message, args...)
}

func (c *P10Client) warnf(message string, args ...interface{}) {
	c.config.Logger.Warnf(message, args...)
}

func (c *P10Client) errorf(message string, args ...interface{}) {
	c.config.Logger.Errorf(message, args...)
}
