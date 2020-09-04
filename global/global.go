package global

import (
	"baseapi/config"

	"github.com/go-redis/redis"
	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	BA_DB     *gorm.DB
	BA_REDIS  *redis.Client
	BA_CONFIG config.Server
	BA_VP     *viper.Viper
	BA_LOG    *oplogging.Logger
)
