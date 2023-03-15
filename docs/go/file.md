https://go.dev/play/p/DF2Wo8nDKaF

## read file
dat, err := os.ReadFile("/tmp/dat")
check(err)
fmt.Print(string(dat))

## open file
f, err := os.Open("/tmp/dat")
b1 := make([]byte, 5)
n1, err := f.Read(b1)
check(err)
fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

## Seek and Peek