```go
timer1 := time.NewTimer(2 * time.Second)
<-timer1.C
fmt.Println("Timer 1 fired")
```