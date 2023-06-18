package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/hashicorp/consul/api"
	"sync"

	"github.com/spf13/viper"
	"github.com/taoti888/user/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Agent  *api.Agent
	DB     *gorm.DB
	REDIS  *redis.ClusterClient
	CONFIG config.Server
	VP     *viper.Viper
	LOG    *zap.Logger

	lock sync.RWMutex
)
