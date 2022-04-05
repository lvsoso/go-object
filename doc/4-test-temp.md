```shell
(base) lv@lv:go-object$ echo -n "this object will have only 1 instance" | openssl dgst -sha256 -binary | base64
aWKQ2BipX94sb+h3xdTbWYAu1yzjn5vyFG2SOwUQIXY=


(base) lv@lv:go-object$ curl -v 10.29.1.2:12345/objects/test4_1 -XPUT -d "this object will have only 1 instance" -H "Digest:SHA-256=aWKQ2BipX94sb+h3xdTbWYAu1yzjn5vyFG2SOwUQIXY="
*   Trying 10.29.1.2:12345...
* TCP_NODELAY set
* Connected to 10.29.1.2 (10.29.1.2) port 12345 (#0)
> PUT /objects/test4_1 HTTP/1.1
> Host: 10.29.1.2:12345
> User-Agent: curl/7.65.3
> Accept: */*
> Digest:SHA-256=aWKQ2BipX94sb+h3xdTbWYAu1yzjn5vyFG2SOwUQIXY=
> Content-Length: 37
> Content-Type: application/x-www-form-urlencoded
> 
* upload completely sent off: 37 out of 37 bytes
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Tue, 05 Apr 2022 06:17:50 GMT
< Content-Length: 0
< 
* Connection #0 to host 10.29.1.2 left intact
(base) lv@lv:go-object$ curl -v 10.29.1.2:12345/objects/test4_2 -XPUT -d "this object will have only 1 instance" -H "Digest:SHA-256=aWKQ2BipX94sb+h3xdTbWYAu1yzjn5vyFG2SOwUQIXY="
*   Trying 10.29.1.2:12345...
* TCP_NODELAY set
* Connected to 10.29.1.2 (10.29.1.2) port 12345 (#0)
> PUT /objects/test4_2 HTTP/1.1
> Host: 10.29.1.2:12345
> User-Agent: curl/7.65.3
> Accept: */*
> Digest:SHA-256=aWKQ2BipX94sb+h3xdTbWYAu1yzjn5vyFG2SOwUQIXY=
> Content-Length: 37
> Content-Type: application/x-www-form-urlencoded
> 
* upload completely sent off: 37 out of 37 bytes
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Tue, 05 Apr 2022 06:18:16 GMT
< Content-Length: 0
< 
* Connection #0 to host 10.29.1.2 left intact
(base) lv@lv:go-object$ curl -v 10.29.1.2:12345/objects/test4_3 -XPUT -d "this object will have only 1 instance" -H "Digest:SHA-256=aWKQ2BipX94sb+h3xdTbWYAu1yzjn5vyFG2SOwUQIXY="
*   Trying 10.29.1.2:12345...
* TCP_NODELAY set
* Connected to 10.29.1.2 (10.29.1.2) port 12345 (#0)
> PUT /objects/test4_3 HTTP/1.1
> Host: 10.29.1.2:12345
> User-Agent: curl/7.65.3
> Accept: */*
> Digest:SHA-256=aWKQ2BipX94sb+h3xdTbWYAu1yzjn5vyFG2SOwUQIXY=
> Content-Length: 37
> Content-Type: application/x-www-form-urlencoded
> 
* upload completely sent off: 37 out of 37 bytes
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Tue, 05 Apr 2022 06:18:24 GMT
< Content-Length: 0
< 
* Connection #0 to host 10.29.1.2 left intact
(base) lv@lv:go-object$ curl -v 10.29.1.2:12345/objects/test4_4 -XPUT -d "this object will have only 1 instance" -H "Digest:SHA-256=aWKQ2BipX94sb+h3xdTbWYAu1yzjn5vyFG2SOwUQIXY="
*   Trying 10.29.1.2:12345...
* TCP_NODELAY set
* Connected to 10.29.1.2 (10.29.1.2) port 12345 (#0)
> PUT /objects/test4_4 HTTP/1.1
> Host: 10.29.1.2:12345
> User-Agent: curl/7.65.3
> Accept: */*
> Digest:SHA-256=aWKQ2BipX94sb+h3xdTbWYAu1yzjn5vyFG2SOwUQIXY=
> Content-Length: 37
> Content-Type: application/x-www-form-urlencoded
> 
* upload completely sent off: 37 out of 37 bytes
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Tue, 05 Apr 2022 06:18:32 GMT
< Content-Length: 0
< 
* Connection #0 to host 10.29.1.2 left intact

(base) lv@lv:go-object$ curl -v 10.29.1.2:12345/locate/aWKQ2BipX94sb+h3xdTbWYAu1yzjn5vyFG2SOwUQIXY=
*   Trying 10.29.1.2:12345...
* TCP_NODELAY set
* Connected to 10.29.1.2 (10.29.1.2) port 12345 (#0)
> GET /locate/aWKQ2BipX94sb+h3xdTbWYAu1yzjn5vyFG2SOwUQIXY= HTTP/1.1
> Host: 10.29.1.2:12345
> User-Agent: curl/7.65.3
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Tue, 05 Apr 2022 06:37:48 GMT
< Content-Length: 17
< Content-Type: text/plain; charset=utf-8
< 
* Connection #0 to host 10.29.1.2 left intact
"10.29.1.5:12345"

(base) lv@lv:go-object$ curl 10.29.1.2:12345/objects/test4_2
this object will have only 1 instance

(base) lv@lv:go-object$ curl 10.29.1.2:12345/objects/test4_1
this object will have only 1 instance
```