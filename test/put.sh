#!/bin/bash

curl -v 127.0.0.1:12345/objects/test2 -XPUT -d"this is object test2"

curl 127.0.0.1:12345/locate/test2
echo
curl 10.29.2.1:12345/objects/test2
echo
