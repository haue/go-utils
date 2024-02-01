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
type RuleConfig struct {
	Name  string
	Rules []PeriodTimes
}

var (
	config []RuleConfig
	rules  map[string]*ratelimit.Rule
)

func init() {
	rules = make(map[string]*ratelimit.Rule)
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Fatal error resources file: %s \n", err.Error())
	}
	if err := viper.UnmarshalKey("ratelimit", &config); err != nil {
		fmt.Printf("unable to decode into struct %s \n", err.Error())
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err := viper.UnmarshalKey("ratelimit", &config); err != nil {
			fmt.Printf("unable to decode into struct %s \n", err.Error())
		}
	})
	viper.WatchConfig()
	for _, item := range config {
		rule := ratelimit.NewRule()
		name := item.Name
		for _, item2 := range item.Rules {
			rule.AddRule(time.Second*time.Duration(item2.Seconds), item2.Times)
		}
		rules[name] = rule
	}
}

func Check(ruleName string, id string) bool {
	rule := rules[ruleName]
	return rule.AllowVisit(id)
}
