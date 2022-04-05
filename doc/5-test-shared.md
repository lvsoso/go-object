```shell
(base) lv@lv:go-object$ echo -n "this object will be separate to 4+2 shards" | openssl dgst -sha256 -binary | base64
MBMxWHrPMsuOBaVYHkwScZQRyTRMQyiKp2oelpLZza8=

(base) lv@lv:go-object$ curl -v 10.29.1.2:12345/objects/test5 -XPUT -d "this object will be separate to 4+2 shards" -H "Digest:SHA-256=MBMxWHrPMsuOBaVYHkwScZQRyTRMQyiKp2oelpLZza8="
*   Trying 10.29.1.2:12345...
* TCP_NODELAY set
* Connected to 10.29.1.2 (10.29.1.2) port 12345 (#0)
> PUT /objects/test5 HTTP/1.1
> Host: 10.29.1.2:12345
> User-Agent: curl/7.65.3
> Accept: */*
> Digest:SHA-256=MBMxWHrPMsuOBaVYHkwScZQRyTRMQyiKp2oelpLZza8=
> Content-Length: 42
> Content-Type: application/x-www-form-urlencoded
> 
* upload completely sent off: 42 out of 42 bytes
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Tue, 05 Apr 2022 12:12:03 GMT
< Content-Length: 0
< 
* Connection #0 to host 10.29.1.2 left intact


(base) lv@lv:go-object$ curl 10.29.1.2:12345/objects/test5
this object will be separate to 4+2 shards

(base) lv@lv:go-object$ curl 10.29.1.2:12345/locate/MBMxWHrPMsuOBaVYHkwScZQRyTRMQyiKp2oelpLZza8=
{"0":"10.29.1.5:12345","1":"10.29.1.10:12345","2":"10.29.1.4:12345","3":"10.29.1.6:12345","4":"10.29.1.7:12345","5":"10.29.1.11:12345"}

# 损坏分片数据
(base) lv@lv:go-object$ docker exec -it deploy_n6_1 /bin/bash
root@24ff105fc9d4:/# ls /tmp/objects/
'MBMxWHrPMsuOBaVYHkwScZQRyTRMQyiKp2oelpLZza8=.4.i8xiyIwSO2cRJwnmkO4ieUV9v26H6e8tu5Y%2F3Op%2F4zE='
root@24ff105fc9d4:/# rm /tmp/objects/MBMxWHrPMsuOBaVYHkwScZQRyTRMQyiKp2oelpLZza8\=.4.i8xiyIwSO2cRJwnmkO4ieUV9v26H6e8tu5Y%2F3Op%2F4zE\= 
root@24ff105fc9d4:/# exit
exit

# 污染分片数据
(base) lv@lv:go-object$ docker exec -it deploy_n4_1 /bin/bash
root@f37af37f0d7a:/# echo some_garbage > /tmp/objects/MBMxWHrPMsuOBaVYHkwScZQRyTRMQyiKp2oelpLZza8\=.0.XVFHp5%2F5kZ89051XQo6UEkWW8OGzyXwLWS4Ln9f0Ncg\= 
root@f37af37f0d7a:/# exit
exit

(base) lv@lv:go-object$  curl 10.29.1.2:12345/objects/test5
this object will be separate to 4+2 shards

```