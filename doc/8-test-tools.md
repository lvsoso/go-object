```shell

(base) lv@lv:go-object$ echo -n "this is object test8 version 1" | openssl dgst -sha256 -binary | base64
2IJQkIth94IVsnPQMrsNxz1oqfrsPo0E2ZmZfJLDZnE=

(base) lv@lv:go-object$ curl 10.29.1.2:12345/objects/test8 -XPUT -d"this is object test8 version 1" -H "Digest:SHA-256=2IJQkIth94IVsnPQMrsNxz1oqfrsPo0E2ZmZfJLDZnE="

(base) lv@lv:go-object$ echo -n "this is object test8 version 2-6" | openssl dgst -sha256 -binary | base64
66WuRH0s0albWDZ9nTmjFo9JIqTTXmB6EiRkhTh1zeA=

(base) lv@lv:go-object$ curl 10.29.1.2:12345/objects/test8 -XPUT -d"this is object test8 version 2-6" -H "Digest:SHA-256=66WuRH0s0albWDZ9nTmjFo9JIqTTXmB6EiRkhTh1zeA="
(base) lv@lv:go-object$ curl 10.29.1.2:12345/objects/test8 -XPUT -d"this is object test8 version 2-6" -H "Digest:SHA-256=66WuRH0s0albWDZ9nTmjFo9JIqTTXmB6EiRkhTh1zeA="
(base) lv@lv:go-object$ curl 10.29.1.2:12345/objects/test8 -XPUT -d"this is object test8 version 2-6" -H "Digest:SHA-256=66WuRH0s0albWDZ9nTmjFo9JIqTTXmB6EiRkhTh1zeA="
(base) lv@lv:go-object$ curl 10.29.1.2:12345/objects/test8 -XPUT -d"this is object test8 version 2-6" -H "Digest:SHA-256=66WuRH0s0albWDZ9nTmjFo9JIqTTXmB6EiRkhTh1zeA="
(base) lv@lv:go-object$ curl 10.29.1.2:12345/objects/test8 -XPUT -d"this is object test8 version 2-6" -H "Digest:SHA-256=66WuRH0s0albWDZ9nTmjFo9JIqTTXmB6EiRkhTh1zeA="

(base) lv@lv:go-object$ curl 10.29.1.2:12345/versions/test8
{"Name":"test8","Version":1,"Size":104857600,"Hash":"IEkqTQ2E+L6xdn9mFiKfhdRMKCe2S9v7Jg7hL6EQng4="}
{"Name":"test8","Version":2,"Size":30,"Hash":"2IJQkIth94IVsnPQMrsNxz1oqfrsPo0E2ZmZfJLDZnE="}
{"Name":"test8","Version":3,"Size":32,"Hash":"66WuRH0s0albWDZ9nTmjFo9JIqTTXmB6EiRkhTh1zeA="}
{"Name":"test8","Version":4,"Size":32,"Hash":"66WuRH0s0albWDZ9nTmjFo9JIqTTXmB6EiRkhTh1zeA="}
{"Name":"test8","Version":5,"Size":32,"Hash":"66WuRH0s0albWDZ9nTmjFo9JIqTTXmB6EiRkhTh1zeA="}
{"Name":"test8","Version":6,"Size":32,"Hash":"66WuRH0s0albWDZ9nTmjFo9JIqTTXmB6EiRkhTh1zeA="}
{"Name":"test8","Version":7,"Size":32,"Hash":"66WuRH0s0albWDZ9nTmjFo9JIqTTXmB6EiRkhTh1zeA="}

```