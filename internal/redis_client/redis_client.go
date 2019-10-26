package redis_client

import (
	"go/build"
	"io/ioutil"
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
	"gopkg.in/yaml.v2"
)

//RedisConf store configuration for redis client
type RedisConf struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
}

//GetConf Retrieve conf from configs folder
func GetConf() RedisConf {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	projectPath := gopath + "/src/github.com/Evrard-Nil/middleware"
	yamlFile, err := ioutil.ReadFile(projectPath + "/configs/conf_redis.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	c := RedisConf{}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

//ConnectToRedis returns a connection to redis according to conf
func ConnectToRedis(conf RedisConf) redis.Conn {
	client, err := redis.Dial("tcp", conf.Host, redis.DialPassword(conf.Password))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Succesfully connected to Redis at %s\n", conf.Host)
	}
	return client
}

// InitRedisPool returns a pool according to redis conf
// and allows us to connect with multiple instances to redis server
func InitRedisPool(conf RedisConf) *redis.Pool {
	return &redis.Pool{
		MaxIdle:   12,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			client, err := redis.Dial("tcp", conf.Host, redis.DialPassword(conf.Password))
			if err != nil {
				log.Fatal(err)
				return nil, err
			}
			log.Printf("Succesfully connected to Redis at %s\n", conf.Host)
			return client, err
		},
	}
}
