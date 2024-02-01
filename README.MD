# ratelimit

速率限制

使用：

1. config.yaml中添加以下示例节点

   ```
   ratelimit:
     - name: test
       rules:
         - seconds: 60
           times: 20
         - seconds: 3600
           times: 200
         - seconds: 86400
           times: 400
   ```

2. 示例

   ```
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
   ```

   