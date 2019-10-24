package redis_client

import (
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
	projectPath := os.Getenv("GOPATH") + "/src/github.com/Evrard-Nil/middleware"
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
