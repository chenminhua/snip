We might want a server to gracefully shutdown when it receives a SIGTERM.

Go singal notification works by sending os.Singal on a channel. We'll create a channel to receive these notification.

```go 
sigs := make(chan os.Singal, 1)
signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
done := make(chan bool, 1)

go func() {
    sig := <- sigs
    fmt.Println()
    fmt.Println(sig)
    done<-true
}()
fmt.Println("awaiting signal")
<-done
```