package i18n

import "fmt"

func init() {
	loadI18nConfig()
}

func T(language string, targetKey string) string {
	return "XXX"
}

func loadI18nConfig() {
	fmt.Println("------------   loadI18nConfig   -----------")
}