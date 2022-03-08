package repository

import (
	"ShoppingBasket/ent"
	"ShoppingBasket/ent/migrate"
	"ShoppingBasket/external"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "go-micro.dev/v4/logger"
)

// ClientProvider takes external.Host and provider with initialize client
func ClientProvider(h *external.Host) (*ent.Client, error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", h.User, h.Password, h.Address, h.Port, h.Name)
	//client, err := ent.Open(h.Type, "root:root@tcp(localhost:3306)/EventCatalogDB?parseTime=true")
	client, err := ent.Open(h.Type, connString)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	//defer client.Close()
	ctx := context.Background()
	// Run migration.
	if err := client.Schema.Create(ctx, migrate.WithDropIndex(true), migrate.WithDropColumn(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client, nil
}
