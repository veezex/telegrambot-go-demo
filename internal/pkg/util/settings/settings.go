package settings

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type Config struct {
	TelegramBotApiKey  string
	GrpcServerPort     uint64
	RestServerPort     uint64
	DbHost             string
	DbPort             uint64
	DbUser             string
	DbPassword         string
	DbName             string
	KafkaBrokers       []string
	KafkaTopic         string
	KafkaGroupId       string
	DebugPort          uint64
	RedisHost          string
	RedisPubSubChannel string
	RedisPort          uint64
	RedisPassword      string
	RedisDb            uint64
}

type impl struct {
	storage map[string]string
}

type Settings interface {
	getString(key string) string
	getNumber(key string) uint64
	GetConfig() Config
}

func New(envFile string) Settings {
	envs, err := godotenv.Read(envFile)
	if err != nil {
		logrus.Fatal("Unable to load .env file, you can create it as a copy from .env.example")
	}

	return &impl{
		storage: envs,
	}
}

func (i *impl) GetConfig() Config {
	return Config{
		TelegramBotApiKey:  i.getString("TELEGRAM_BOT_API_KEY"),
		GrpcServerPort:     i.getNumber("GRPC_SERVER_PORT"),
		RestServerPort:     i.getNumber("REST_SERVER_PORT"),
		DbHost:             i.getString("DB_HOST"),
		DbPort:             i.getNumber("DB_PORT"),
		DbUser:             i.getString("DB_USER"),
		DbPassword:         i.getString("DB_PASSWORD"),
		DbName:             i.getString("DB_NAME"),
		KafkaBrokers:       i.getList("KAFKA_BROKERS"),
		KafkaTopic:         i.getString("KAFKA_TOPIC"),
		KafkaGroupId:       i.getString("KAFKA_GROUP_ID"),
		DebugPort:          i.getNumber("DEBUG_PORT"),
		RedisHost:          i.getString("REDIS_HOST"),
		RedisPubSubChannel: i.getString("REDIS_PUBSUB_CHANNEL"),
		RedisPort:          i.getNumber("REDIS_PORT"),
		RedisDb:            i.getNumber("REDIS_DB"),
		RedisPassword:      i.getString("REDIS_PASSWORD"),
	}
}

func (i *impl) getList(key string) []string {
	str := i.getString(key)
	return strings.Split(str, ";")
}

func (i *impl) getString(key string) string {
	result, ok := i.storage[key]
	if !ok {
		logrus.Fatalf("Undefined <%v> env variable", key)
	}

	return result
}

func (i *impl) getNumber(key string) uint64 {
	val := i.getString(key)

	num, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		logrus.Fatalf("field <%v> value <%v> is not a number", key, val)
	}

	return num
}
