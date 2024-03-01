package gorm_gms

import (
	"gorm.io/gorm"
)

// NoopTxOption go-mysql-server not support transaction, so use noopTx to hack
type NoopTxOption struct {
}

func (t *NoopTxOption) Apply(c *gorm.Config) error {

	if p := c.ConnPool; p != nil {
		c.ConnPool = &noopTx{ConnPool: p}
	}

	return nil
}

func (t *NoopTxOption) AfterInitialize(c *gorm.DB) error {

	if p := c.ConnPool; p != nil {
		c.ConnPool = &noopTx{ConnPool: p}
	}
	if p := c.Statement.ConnPool; p != nil {
		c.Statement.ConnPool = &noopTx{ConnPool: p}
	}
	return nil
}

// noopTx by-pass transaction
//
//	 go-mysql-server not support transaction, executing multiple-sql in actual,
//		but gorm will wait for lock when committing tx, it will be locked in this case
type noopTx struct {
	gorm.ConnPool
}

func (t *noopTx) Commit() error {
	return nil
}
func (t *noopTx) Rollback() error {
	return nil
}
