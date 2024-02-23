package ratelimit

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/yudeguang/ratelimit"
)

type PeriodTimes struct {
	Seconds int
	Times   int
}

var (
	rules map[string]*ratelimit.Rule = make(map[string]*ratelimit.Rule)
)

func init() {
	rulesConfig := make(map[string][]PeriodTimes)
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Fatal error resources file: %s \n", err.Error())
	}
	if err := viper.UnmarshalKey("ratelimit", &rulesConfig); err != nil {
		fmt.Printf("unable to decode into struct %s \n", err.Error())
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err := viper.UnmarshalKey("ratelimit", &rulesConfig); err != nil {
			fmt.Printf("unable to decode into struct %s \n", err.Error())
		}
	})
	viper.WatchConfig()
	for name, item := range rulesConfig {
		rule := ratelimit.NewRule()
		for _, item2 := range item {
			rule.AddRule(time.Second*time.Duration(item2.Seconds), item2.Times)
		}
		rules[name] = rule
	}
}

func Check(ruleName string, id string) bool {
	rule := rules[ruleName]
	return rule.AllowVisit(id)
}
func GetRule(ruleName string) *ratelimit.Rule {
	return rules[ruleName]
}
