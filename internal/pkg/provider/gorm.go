package provider

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func NewGorm(l *zap.Logger, conf *viper.Viper) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.GetString("postgres.host"),
		conf.GetString("postgres.port"),
		conf.GetString("postgres.user"),
		conf.GetString("postgres.password"),
		conf.GetString("postgres.db"),
	)

	var (
		db  *gorm.DB
		err error
	)
	retryErr := retry(l, 3, time.Second, func() error {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		return err
	})
	if retryErr != nil {
		l.Error(retryErr.Error())
	}

	return db
}

func retry(l *zap.Logger, attempts int, sleep time.Duration, f func() error) (err error) {
	for i := 0; i < attempts; i++ {
		if i > 0 {
			l.Info(fmt.Sprintf("retrying after error: %v", err))
			time.Sleep(sleep)
			sleep *= 2
		}
		err = f()
		if err == nil {
			return nil
		}
	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}
