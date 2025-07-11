package connections

import (
	"github.com/manicar2093/gomancer/bootstrap/core/httphealthcheck"
	"log"

	"time"

	_ "github.com/jackc/pgx/v5"
	"github.com/manicar2093/gormpager"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	DatabaseConnectionConfig struct {
		Dns          string `env:"DATABASE_URL" validate:"required|isURL"`
		MaxIdleConns int    `env:"DB_MAX_IDLE_CONNS" envDefault:"1" validate:"required"`
		MaxOpenConns int    `env:"DB_MAX_OPEN_CONNS" envDefault:"1" validate:"required"`
		MinPageSize  uint   `env:"MIN_PAGE_SIZE" envDefault:"1" validate:"required"`
		MaxPageSize  uint   `env:"MAX_PAGE_SIZE" envDefault:"100" validate:"required"`
	}

	ConnWrapper struct {
		*gormpager.GormPager
	}
)

func GetGormConnection(config DatabaseConnectionConfig) *ConnWrapper {
	gormDB, err := gorm.Open(postgres.Open(config.Dns))
	if err != nil {
		log.Panicln(err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Panic(err)
	}
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return &ConnWrapper{
		GormPager: gormpager.WrapGormDBWithOptions(gormDB, gormpager.Options{
			PageSizeLowerLimit: config.MinPageSize,
			PageSizeUpperLimit: config.MaxPageSize,
		}),
	}
}

func (c *ConnWrapper) ServiceHealth() (httphealthcheck.HealthStatusData, error) {
	defaultDb, err := c.DB.DB()
	if err != nil {
		return httphealthcheck.HealthStatusData{}, err
	}

	if err := defaultDb.Ping(); err != nil {
		return httphealthcheck.HealthStatusData{
			Error: err,
		}, nil
	}

	return httphealthcheck.HealthStatusData{
		IsAvailable: true,
	}, nil
}

func (c *ConnWrapper) ServiceName() string {
	return "database_connection"
}
