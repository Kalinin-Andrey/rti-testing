package db

import (
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/component"
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/price"
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/product"
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/ruleapplicability"
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/user"
	"github.com/Kalinin-Andrey/rti-testing/internal/pkg/session"
)

func (db *DB) AutoMigrateAll() {
	db.DB().AutoMigrate(
		&user.User{},
		&session.Session{},
		&product.Product{},
		&component.Component{},
		&price.Price{},
		&ruleapplicability.RuleApplicability{},
		)
}
