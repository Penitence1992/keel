package loader

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func LoadEnvOrFatal(envname string) (v string) {
	v = os.Getenv(envname)
	if v == "" {
		log.Fatalf("无法获取变量:%s的内容", envname)
	}
	log.Printf("env name : %s , value : %s\n", envname, v)
	return
}

func LoadEnvOrDefault(envname string, dval string) (v string)  {
	v = os.Getenv(envname)
	if v == "" {
		v = dval
	}
	return
}
