package gorm_gms

import "gorm.io/gorm"

type Chain []gorm.Option

func (t Chain) Apply(c *gorm.Config) error {
	for i := range t {
		opt := t[i]
		if err := opt.Apply(c); err != nil {
			return err
		}
	}

	return nil
}

func (t Chain) AfterInitialize(c *gorm.DB) error {
	for i := range t {
		opt := t[i]
		if err := opt.AfterInitialize(c); err != nil {
			return err
		}
	}
	return nil
}
