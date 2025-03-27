package seeders

import "gorm.io/gorm"

type Registry struct {
	db *gorm.DB
}

type IseederRegistry interface {
	Run()
}

func NewSeederRegistry(db *gorm.DB) IseederRegistry {
	return &Registry{db: db}
}

func (s *Registry) Run() {
	RunRoleSeeder(s.db)
	RunUserSeeder(s.db)
}
