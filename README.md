# translater
this is a small tool used for translate an english article sentence by sentence.

Todo:
1. add UT, remove test() from the main code
2. add configure file, to used the main without rebuild

run:
1. install bregydoc/gtranslate(https://github.com/bregydoc/gtranslate) into your host
2. download the transFromGoogle.go file
3. go run transFromGoogle.go

make sure your host can connect to google directly.
or else you need to setup a proxy.
```
set HTTP_PROXY=http://yourproxyIp:port
```
