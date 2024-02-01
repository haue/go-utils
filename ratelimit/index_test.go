package ratelimit

import (
	"log"
	"testing"
)

func TestRateLimit(t *testing.T) {
	id := "aa"
	ruleName := "test"
	rule := rules[ruleName]
	for i := 0; i <= 20; i++ {
		allow := check(ruleName, id)
		if !allow {
			log.Println("访问量超出,其剩余访问次数情况如下:", rule.RemainingVisits(id))
		} else {
			log.Println("允许访问,其剩余访问次数情况如下:", rule.RemainingVisits(id))
		}
	}
}
