
```shell
(base) lv@lv:tmp$ dd if=/dev/zero of=/tmp/file bs=1M count=100
记录了100+0 的读入
记录了100+0 的写出
104857600 bytes (105 MB, 100 MiB) copied, 0.0820504 s, 1.3 GB/s
(base) lv@lv:tmp$ echo "123" >> /tmp/file 


(base) lv@lv:tmp$ openssl dgst -sha256 -binary /tmp/file | base64
HF3K/ujco5VuPD+mVKIxEm5XZ3UM1AElNmwN0XW7dDI=

(base) lv@lv:tmp$ curl -v 10.29.1.2:12345/objects/test777 -XPUT --data-binary @/tmp/file -H "Digest:SHA-256=HF3K/ujco5VuPD+mVKIxEm5XZ3UM1AElNmwN0XW7dDI="
*   Trying 10.29.1.2:12345...
* TCP_NODELAY set
* Connected to 10.29.1.2 (10.29.1.2) port 12345 (#0)
> PUT /objects/test777 HTTP/1.1
> Host: 10.29.1.2:12345
> User-Agent: curl/7.65.3
> Accept: */*
> Digest:SHA-256=HF3K/ujco5VuPD+mVKIxEm5XZ3UM1AElNmwN0XW7dDI=
> Content-Length: 104857604
> Content-Type: application/x-www-form-urlencoded
> Expect: 100-continue
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Tue, 05 Apr 2022 16:54:58 GMT
< Content-Length: 0
< Connection: close
< 


(base) lv@lv:tmp$ docker exec -it deploy_n3_1 ls -lh /tmp/objects
total 28K
-rw-r--r-- 1 root root 25K Apr  5 16:52 'HF3K%2Fujco5VuPD+mVKIxEm5XZ3UM1AElNmwN0XW7dDI=.3.nuZ60C0UYRIjDbF8T%2FZQWEkrkJ2aniZiQtVOf10Sbgo='

(base) lv@lv:tmp$ curl -v 10.29.1.2:12345/objects/test777 -o /tmp/output
*   Trying 10.29.1.2:12345...
* TCP_NODELAY set
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0* Connected to 10.29.1.2 (10.29.1.2) port 12345 (#0)
> GET /objects/test777 HTTP/1.1
> Host: 10.29.1.2:12345
> User-Agent: curl/7.65.3
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Tue, 05 Apr 2022 17:00:56 GMT
< Content-Type: application/octet-stream
< Transfer-Encoding: chunked
< 
{ [519 bytes data]
100  100M    0  100M    0     0   159M      0 --:--:-- --:--:-- --:--:--  159M
* Connection #0 to host 10.29.1.2 left intact
(base) lv@lv:tmp$ diff -s output file 
檔案 output 和 file 相同


(base) lv@lv:tmp$ curl -v 10.29.1.2:12345/objects/test777 -H "Accept-Encoding:gzip" -o /tmp/output2.gz
*   Trying 10.29.1.2:12345...
* TCP_NODELAY set
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0* Connected to 10.29.1.2 (10.29.1.2) port 12345 (#0)
> GET /objects/test777 HTTP/1.1
> Host: 10.29.1.2:12345
> User-Agent: curl/7.65.3
> Accept: */*
> Accept-Encoding:gzip
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Content-Encoding: gzip
< Date: Tue, 05 Apr 2022 17:02:09 GMT
< Transfer-Encoding: chunked
< 
{ [3988 bytes data]
100   99k    0   99k    0     0  83837      0 --:--:--  0:00:01 --:--:-- 83769
* Connection #0 to host 10.29.1.2 left intact
(base) lv@lv:tmp$ gunzip output2.gz 
gzip: output2 already exists; do you wish to overwrite (y or n)? y

(base) lv@lv:tmp$ diff -s file output2 
檔案 file 和 output2 相同


```


