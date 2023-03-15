https://go.dev/play/p/Ly4d_EVtiCs

## time formatting and parsing
https://gobyexample.com/time-formatting-parsing
```go
t := time.Now()
print(t.Format(time.RFC3339)) // 2014-04-15T18:00:15-07:00

t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")

pring(t.Format("Mon Jan _2 15:04:05 2006"))
```

## Different timezone

```go
p := fmt.Println
t := time.Now()
p(t)
p(t.UTC())
loc, _ := time.LoadLocation("Asia/Shanghai")
p(t.In(loc))
```