package gorm_gms

import (
	"gorm.io/gorm"
)

const (
	defaultMaxIdleConns = 1
	defaultMaxOpenConns = 1
)

// ConnectionOption  go-mysql-server not thread-safe, set max-connection=1
type ConnectionOption struct {
	o gorm.Option
}

func (t *ConnectionOption) Apply(c *gorm.Config) error {

	return nil
}

func (t *ConnectionOption) AfterInitialize(c *gorm.DB) error {
	if c == nil {
		return nil
	}
	conn, err := c.DB()
	if err != nil {
		return err
	}
	conn.SetMaxIdleConns(defaultMaxIdleConns)
	conn.SetMaxOpenConns(defaultMaxOpenConns)

	return nil
}
