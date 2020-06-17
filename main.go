//package main
//
//import (
//	"fmt"
//	"i18n"
//	"io/ioutil"
//	"log"
//)
//
//func main() {
//	fmt.Println("====================")
//	result := i18n.T("en", "hello")
//	fmt.Println("result => ", result)
//	fmt.Println("====================")
//}


package main

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"reflect"
	_ "reflect"
	"strings"
)

var I18nConfigMap = map[interface{}]interface{}{}

func main() {
	buffer, err := ioutil.ReadFile("./config/locales/zh-CN.yml")
	err = yaml.Unmarshal(buffer, &I18nConfigMap)
	if err != nil {
		log.Fatalf(err.Error())
	}

	//log.Printf("I18nConfigMap => %v", I18nConfigMap)
	result, err := T("err.msg.kitty.dkdk")
	if err != nil {
		log.Fatalf("Fail to T, err: %v", err)
	}
	log.Println("result => ", result)
}

func T(targetKey string) (interface{}, error) {
	if targetKey == "" {
		return "", errors.New("targetKey should not be an empty string")
	}

	keys := strings.Split(targetKey, ".")
	configMap := I18nConfigMap
	for index, key := range keys {
		config, ok := configMap[key]
		if !ok {
			return "", errors.New(fmt.Sprintf("targetKey: %v does Not exist", targetKey))
		}
		configType := reflect.TypeOf(config)
		//log.Printf("config: %v, configType.Kind(): %v", config, configType.Kind())
		if configType.Kind() == reflect.String {
			if index == len(keys)-1 {
				return reflect.ValueOf(config), nil
			}
		} else if configType.Kind() == reflect.Map {
			valueOfConfig := reflect.ValueOf(config)
			configMap = valueOfConfig.Interface().(map[interface{}]interface{})
		}else {
			return "", errors.New("configMap type error")
		}
	}

	return "", errors.New(fmt.Sprintf("targetKey: %v does Not exist", targetKey))
}

