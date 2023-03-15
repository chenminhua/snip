```go
go func() {
    defer func() {
        if err := recover(); err != nil {
            fmt.Printf("[SafeGo] Panic: %+v", err)
        }
    }()

    ...
}()
```

