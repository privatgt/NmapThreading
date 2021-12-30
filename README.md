# NmapThreading
Nmap multithreading with golang
## Why this small program
Nmap is great port scanner but it is slow and run in one thread. I tried make it faster with power of golang's multithreading
## build
To build just write
```go build main.go```
## run
To run built program just write
```./main {ip to scan} {number of threads}```
after it will automatically run nmap with ```-A -sT``` and scan all 65535 ports with multiple threads
