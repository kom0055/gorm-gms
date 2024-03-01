# gorm-gms

gorm options for go-mysql-server

GMS(go-mysql-server) is good implementation for mysql-interface, so we most use it in unit tests.

But it is not thread safe and not support transaction, you could use this package to avoid those panic. 


## Usage

```go
	db, err := gorm.Open(cfg,
		&Chain{
			Options: []gorm.Option{
				&ConnectionOption{},
				&NoopTxOption{},
			},
		})
```