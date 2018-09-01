# arp-scan
arp-scan

## Usage
```$xslt
$ sudo go run main.go -h
Usage of main.go:
  -i string
        Please specify interface
  -n string
        Please specify network address including cider
  -t duration
        Please specify the request timeout time in milliseconds (default 1000ms) (default 1s)

```
Arp scan for network (timeout is 50 msec)
```$xslt
$ sudo go run main.go -i eth0 -n 192.168.1.0/24 -t 50ms
```