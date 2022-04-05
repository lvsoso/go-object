```shell
(base) lv@lv:tmp$ dd if=/dev/urandom of=/tmp/file bs=1000 count=100
记录了100+0 的读入
记录了100+0 的写出
100000 bytes (100 kB, 98 KiB) copied, 0.00063364 s, 158 MB/s

(base) lv@lv:tmp$ openssl dgst -sha256 -binary /tmp/file | base64
BpkFdeDXkbcAlO4iSq9pgZXiJ1QXpPkFiKHWhEHuwV8=

(base) lv@lv:tmp$ curl -v 10.29.1.2:12345/objects/test6 -XPOST -H "Digest:SHA-256=BpkFdeDXkbcAlO4iSq9pgZXiJ1QXpPkFiKHWhEHuwV8=" -H "Size: 100000"
*   Trying 10.29.1.2:12345...
* TCP_NODELAY set
* Connected to 10.29.1.2 (10.29.1.2) port 12345 (#0)
> POST /objects/test6 HTTP/1.1
> Host: 10.29.1.2:12345
> User-Agent: curl/7.65.3
> Accept: */*
> Digest:SHA-256=BpkFdeDXkbcAlO4iSq9pgZXiJ1QXpPkFiKHWhEHuwV8=
> Size: 100000
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 201 Created
< Location: /temp/eyJOYW1lIjoidGVzdDYiLCJTaXplIjoxMDAwMDAsIkhhc2giOiJCcGtGZGVEWGtiY0FsTzRpU3E5cGdaWGlKMVFYcFBrRmlLSFdoRUh1d1Y4PSIsIlNlcnZlcnMiOlsiMTAuMjkuMS42OjEyMzQ1IiwiMTAuMjkuMS4xMDoxMjM0NSIsIjEwLjI5LjEuMTE6MTIzNDUiLCIxMC4yOS4xLjU6MTIzNDUiLCIxMC4yOS4xLjQ6MTIzNDUiLCIxMC4yOS4xLjc6MTIzNDUiXSwiVXVpZHMiOlsiMjYwYjNiMWYtZTFkYi00YzBiLTljNDItMDBkODFjNjI0ZGIyIiwiZTQ2ZGEwOGMtZTA3My00ZDBiLTkyOGQtZDE3ZjRlZGY2MDBiIiwiZTcyMzQxMWEtZjAxNy00M2QwLTlkYjgtMDBmNmUwOWM0MGJlIiwiNDRjOGU0NDAtMjMwYS00NDg2LWFjOGItNjgzNmI2YmU1ZGIxIiwiMmFiMzcwNmYtNjFjYy00YjUwLWFkNGUtOTBhMWJhNTA3MTQ3IiwiNjMyMzFmY2UtODM0Zi00M2Y3LTgxYzgtZTMxYjMzNGU5NjE1Il19
< Date: Tue, 05 Apr 2022 15:32:44 GMT
< Content-Length: 0
< 
* Connection #0 to host 10.29.1.2 left intact

(base) lv@lv:tmp$ curl -I 10.29.1.2:12345/temp/eyJOYW1lIjoidGVzdDYiLCJTaXplIjoxMDAwMDAsIkhhc2giOiJCcGtGZGVEWGtiY0FsTzRpU3E5cGdaWGlKMVFYcFBrRmlLSFdoRUh1d1Y4PSIsIlNlcnZlcnMiOlsiMTAuMjkuMS42OjEyMzQ1IiwiMTAuMjkuMS4xMDoxMjM0NSIsIjEwLjI5LjEuMTE6MTIzNDUiLCIxMC4yOS4xLjU6MTIzNDUiLCIxMC4yOS4xLjQ6MTIzNDUiLCIxMC4yOS4xLjc6MTIzNDUiXSwiVXVpZHMiOlsiMjYwYjNiMWYtZTFkYi00YzBiLTljNDItMDBkODFjNjI0ZGIyIiwiZTQ2ZGEwOGMtZTA3My00ZDBiLTkyOGQtZDE3ZjRlZGY2MDBiIiwiZTcyMzQxMWEtZjAxNy00M2QwLTlkYjgtMDBmNmUwOWM0MGJlIiwiNDRjOGU0NDAtMjMwYS00NDg2LWFjOGItNjgzNmI2YmU1ZGIxIiwiMmFiMzcwNmYtNjFjYy00YjUwLWFkNGUtOTBhMWJhNTA3MTQ3IiwiNjMyMzFmY2UtODM0Zi00M2Y3LTgxYzgtZTMxYjMzNGU5NjE1Il19
HTTP/1.1 200 OK
Content-Length: 0
Date: Tue, 05 Apr 2022 15:35:18 GMT


(base) lv@lv:tmp$ dd if=/tmp/file of=/tmp/first bs=1000 count=50
记录了50+0 的读入
记录了50+0 的写出
50000 bytes (50 kB, 49 KiB) copied, 0.000237204 s, 211 MB/s

(base) lv@lv:tmp$ curl -v -XPUT --data-binary @/tmp/first 10.29.1.2:12345/temp/eyJOYW1lIjoidGVzdDYiLCJTaXplIjoxMDAwMDAsIkhhc2giOiJCcGtGZGVEWGtiY0FsTzRpU3E5cGdaWGlKMVFYcFBrRmlLSFdoRUh1d1Y4PSIsIlNlcnZlcnMiOlsiMTAuMjkuMS42OjEyMzQ1IiwiMTAuMjkuMS4xMDoxMjM0NSIsIjEwLjI5LjEuMTE6MTIzNDUiLCIxMC4yOS4xLjU6MTIzNDUiLCIxMC4yOS4xLjQ6MTIzNDUiLCIxMC4yOS4xLjc6MTIzNDUiXSwiVXVpZHMiOlsiMjYwYjNiMWYtZTFkYi00YzBiLTljNDItMDBkODFjNjI0ZGIyIiwiZTQ2ZGEwOGMtZTA3My00ZDBiLTkyOGQtZDE3ZjRlZGY2MDBiIiwiZTcyMzQxMWEtZjAxNy00M2QwLTlkYjgtMDBmNmUwOWM0MGJlIiwiNDRjOGU0NDAtMjMwYS00NDg2LWFjOGItNjgzNmI2YmU1ZGIxIiwiMmFiMzcwNmYtNjFjYy00YjUwLWFkNGUtOTBhMWJhNTA3MTQ3IiwiNjMyMzFmY2UtODM0Zi00M2Y3LTgxYzgtZTMxYjMzNGU5NjE1Il19
*   Trying 10.29.1.2:12345...
* TCP_NODELAY set
* Connected to 10.29.1.2 (10.29.1.2) port 12345 (#0)
> PUT /temp/eyJOYW1lIjoidGVzdDYiLCJTaXplIjoxMDAwMDAsIkhhc2giOiJCcGtGZGVEWGtiY0FsTzRpU3E5cGdaWGlKMVFYcFBrRmlLSFdoRUh1d1Y4PSIsIlNlcnZlcnMiOlsiMTAuMjkuMS42OjEyMzQ1IiwiMTAuMjkuMS4xMDoxMjM0NSIsIjEwLjI5LjEuMTE6MTIzNDUiLCIxMC4yOS4xLjU6MTIzNDUiLCIxMC4yOS4xLjQ6MTIzNDUiLCIxMC4yOS4xLjc6MTIzNDUiXSwiVXVpZHMiOlsiMjYwYjNiMWYtZTFkYi00YzBiLTljNDItMDBkODFjNjI0ZGIyIiwiZTQ2ZGEwOGMtZTA3My00ZDBiLTkyOGQtZDE3ZjRlZGY2MDBiIiwiZTcyMzQxMWEtZjAxNy00M2QwLTlkYjgtMDBmNmUwOWM0MGJlIiwiNDRjOGU0NDAtMjMwYS00NDg2LWFjOGItNjgzNmI2YmU1ZGIxIiwiMmFiMzcwNmYtNjFjYy00YjUwLWFkNGUtOTBhMWJhNTA3MTQ3IiwiNjMyMzFmY2UtODM0Zi00M2Y3LTgxYzgtZTMxYjMzNGU5NjE1Il19 HTTP/1.1
> Host: 10.29.1.2:12345
> User-Agent: curl/7.65.3
> Accept: */*
> Content-Length: 50000
> Content-Type: application/x-www-form-urlencoded
> Expect: 100-continue
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 100 Continue
* We are completely uploaded and fine
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Tue, 05 Apr 2022 15:36:43 GMT
< Content-Length: 0
< 
* Connection #0 to host 10.29.1.2 left intact

(base) lv@lv:tmp$ curl -I 10.29.1.2:12345/temp/eyJOYW1lIjoidGVzdDYiLCJTaXplIjoxMDAwMDAsIkhhc2giOiJCcGtGZGVEWGtiY0FsTzRpU3E5cGdaWGlKMVFYcFBrRmlLSFdoRUh1d1Y4PSIsIlNlcnZlcnMiOlsiMTAuMjkuMS42OjEyMzQ1IiwiMTAuMjkuMS4xMDoxMjM0NSIsIjEwLjI5LjEuMTE6MTIzNDUiLCIxMC4yOS4xLjU6MTIzNDUiLCIxMC4yOS4xLjQ6MTIzNDUiLCIxMC4yOS4xLjc6MTIzNDUiXSwiVXVpZHMiOlsiMjYwYjNiMWYtZTFkYi00YzBiLTljNDItMDBkODFjNjI0ZGIyIiwiZTQ2ZGEwOGMtZTA3My00ZDBiLTkyOGQtZDE3ZjRlZGY2MDBiIiwiZTcyMzQxMWEtZjAxNy00M2QwLTlkYjgtMDBmNmUwOWM0MGJlIiwiNDRjOGU0NDAtMjMwYS00NDg2LWFjOGItNjgzNmI2YmU1ZGIxIiwiMmFiMzcwNmYtNjFjYy00YjUwLWFkNGUtOTBhMWJhNTA3MTQ3IiwiNjMyMzFmY2UtODM0Zi00M2Y3LTgxYzgtZTMxYjMzNGU5NjE1Il19
HTTP/1.1 200 OK
Content-Length: 32000
Date: Tue, 05 Apr 2022 15:37:57 GMT

(base) lv@lv:tmp$ dd if=/tmp/file of=/tmp/second bs=1000 skip=32 count=68
记录了68+0 的读入
记录了68+0 的写出
68000 bytes (68 kB, 66 KiB) copied, 0.000967396 s, 70.3 MB/s

curl -v -XPUT --data-binary @/tmp/second -H "range:bytes=32000-" 10.29.1.2:12345/temp/eyJOYW1lIjoidGVzdDYiLCJTaXplIjoxMDAwMDAsIkhhc2giOiJCcGtGZGVEWGtiY0FsTzRpU3E5cGdaWGlKMVFYcFBrRmlLSFdoRUh1d1Y4PSIsIlNlcnZlcnMiOlsiMTAuMjkuMS42OjEyMzQ1IiwiMTAuMjkuMS4xMDoxMjM0NSIsIjEwLjI5LjEuMTE6MTIzNDUiLCIxMC4yOS4xLjU6MTIzNDUiLCIxMC4yOS4xLjQ6MTIzNDUiLCIxMC4yOS4xLjc6MTIzNDUiXSwiVXVpZHMiOlsiMjYwYjNiMWYtZTFkYi00YzBiLTljNDItMDBkODFjNjI0ZGIyIiwiZTQ2ZGEwOGMtZTA3My00ZDBiLTkyOGQtZDE3ZjRlZGY2MDBiIiwiZTcyMzQxMWEtZjAxNy00M2QwLTlkYjgtMDBmNmUwOWM0MGJlIiwiNDRjOGU0NDAtMjMwYS00NDg2LWFjOGItNjgzNmI2YmU1ZGIxIiwiMmFiMzcwNmYtNjFjYy00YjUwLWFkNGUtOTBhMWJhNTA3MTQ3IiwiNjMyMzFmY2UtODM0Zi00M2Y3LTgxYzgtZTMxYjMzNGU5NjE1Il19

(base) lv@lv:tmp$ curl -v -XPUT --data-binary @/tmp/second -H "range:bytes=32000-" 10.29.1.2:12345/temp/eyJOYW1lIjoidGVzdDYiLCJTaXplIjoxMDAwMDAsIkhhc2giOiJCcGtGZGVEWGtiY0FsTzRpU3E5cGdaWGlKMVFYcFBrRmlLSFdoRUh1d1Y4PSIsIlNlcnZlcnMiOlsiMTAuMjkuMS42OjEyMzQ1IiwiMTAuMjkuMS4xMDoxMjM0NSIsIjEwLjI5LjEuMTE6MTIzNDUiLCIxMC4yOS4xLjU6MTIzNDUiLCIxMC4yOS4xLjQ6MTIzNDUiLCIxMC4yOS4xLjc6MTIzNDUiXSwiVXVpZHMiOlsiMjYwYjNiMWYtZTFkYi00YzBiLTljNDItMDBkODFjNjI0ZGIyIiwiZTQ2ZGEwOGMtZTA3My00ZDBiLTkyOGQtZDE3ZjRlZGY2MDBiIiwiZTcyMzQxMWEtZjAxNy00M2QwLTlkYjgtMDBmNmUwOWM0MGJlIiwiNDRjOGU0NDAtMjMwYS00NDg2LWFjOGItNjgzNmI2YmU1ZGIxIiwiMmFiMzcwNmYtNjFjYy00YjUwLWFkNGUtOTBhMWJhNTA3MTQ3IiwiNjMyMzFmY2UtODM0Zi00M2Y3LTgxYzgtZTMxYjMzNGU5NjE1Il19
*   Trying 10.29.1.2:12345...
* TCP_NODELAY set
* Connected to 10.29.1.2 (10.29.1.2) port 12345 (#0)
> PUT /temp/eyJOYW1lIjoidGVzdDYiLCJTaXplIjoxMDAwMDAsIkhhc2giOiJCcGtGZGVEWGtiY0FsTzRpU3E5cGdaWGlKMVFYcFBrRmlLSFdoRUh1d1Y4PSIsIlNlcnZlcnMiOlsiMTAuMjkuMS42OjEyMzQ1IiwiMTAuMjkuMS4xMDoxMjM0NSIsIjEwLjI5LjEuMTE6MTIzNDUiLCIxMC4yOS4xLjU6MTIzNDUiLCIxMC4yOS4xLjQ6MTIzNDUiLCIxMC4yOS4xLjc6MTIzNDUiXSwiVXVpZHMiOlsiMjYwYjNiMWYtZTFkYi00YzBiLTljNDItMDBkODFjNjI0ZGIyIiwiZTQ2ZGEwOGMtZTA3My00ZDBiLTkyOGQtZDE3ZjRlZGY2MDBiIiwiZTcyMzQxMWEtZjAxNy00M2QwLTlkYjgtMDBmNmUwOWM0MGJlIiwiNDRjOGU0NDAtMjMwYS00NDg2LWFjOGItNjgzNmI2YmU1ZGIxIiwiMmFiMzcwNmYtNjFjYy00YjUwLWFkNGUtOTBhMWJhNTA3MTQ3IiwiNjMyMzFmY2UtODM0Zi00M2Y3LTgxYzgtZTMxYjMzNGU5NjE1Il19 HTTP/1.1
> Host: 10.29.1.2:12345
> User-Agent: curl/7.65.3
> Accept: */*
> range:bytes=32000-
> Content-Length: 68000
> Content-Type: application/x-www-form-urlencoded
> Expect: 100-continue
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 100 Continue
* We are completely uploaded and fine
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Tue, 05 Apr 2022 15:39:37 GMT
< Content-Length: 0
< 
* Connection #0 to host 10.29.1.2 left intact


(base) lv@lv:tmp$ curl 10.29.1.2:12345/objects/test6 > /tmp/output
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   97k    0   97k    0     0  2504k      0 --:--:-- --:--:-- --:--:-- 2504k

(base) lv@lv:tmp$ diff -s /tmp/file /tmp/output 
檔案 /tmp/file 和 /tmp/output 相同

(base) lv@lv:tmp$ curl 10.29.1.2:12345/objects/test6 -H "range: bytes=32000-" > /tmp/output2
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 68000    0 68000    0     0  5108k      0 --:--:-- --:--:-- --:--:-- 5108k

(base) lv@lv:tmp$ diff -s /tmp/second /tmp/output2 
檔案 /tmp/second 和 /tmp/output2 相同

```