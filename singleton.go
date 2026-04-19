package main

import "sync"

/**
スレッドセーフなシングルトンパターンです。
**/

type ConfigManager struct {
	AppConfig map[string]string
}

var (
	once     sync.Once
	instance *ConfigManager
)

func GetInstance() *ConfigManager {
	once.Do(func() {
		instance = &ConfigManager{
			AppConfig: map[string]string{
				"env":     "prod",
				"port":    "8080",
				"timeout": "30s",
			},
		}
	})

	return instance
}

func Singleton() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			config := GetInstance()
			println(config.AppConfig["env"])
		}()
	}
	wg.Wait()
}
