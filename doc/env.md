```shell
cd deploy
docker-compose up -d

docker exec -it deploy_rabbitmq_1 rabbitmqctl set_permissions -p / test ".*" ".*" ".*"
wget localhost:15672/cli/rabbitmqadmin

rabbitmqadmin declare exchange --vhost=/ name=test type=direct -u test -p test
rabbitmqadmin declare exchange --vhost=/ name=test2 type=direct -u test -p test

rabbitmqadmin declare exchange --vhost=/ name=apiServers type=fanout -u test -p test
rabbitmqadmin declare exchange  --vhost=/  name=dataServers type=fanout -u test -p test

docker exec -it deploy_n3_1  /bin/bash -c "mkdir /tmp/objects"
docker exec -it deploy_n4_1  /bin/bash -c "mkdir /tmp/objects"
docker exec -it deploy_n5_1  /bin/bash -c "mkdir /tmp/objects"
docker exec -it deploy_n6_1  /bin/bash -c "mkdir /tmp/objects"

```

ES 

```shell
curl http://10.29.1.9:9200/metadata -XPUT -H 'Content-Type: application/json' -d'{"mappings":{"objects": {"properties": {"name":{"type":"keyword"},"version":{"type":"integer"}, "size":{"type":"integer"}, "hash":{"type":"keyword"}}}}}'
```

prepare node

```shell
cd ./dataServer
go build
cd -

cd ./apiServer
go build
cd -


docker cp ./apiServer/apiServer deploy_n1_1:/
docker cp ./apiServer/apiServer deploy_n2_1:/

docker cp ./dataServer/dataServer deploy_n3_1:/
docker cp ./dataServer/dataServer deploy_n4_1:/
docker cp ./dataServer/dataServer deploy_n5_1:/
docker cp ./dataServer/dataServer deploy_n6_1:/
```

start node

```shell
docker exec -it -e LISTEN_ADDRESS=10.29.1.2:12345  -e ES_SERVER=10.29.1.9:9200 deploy_n1_1 ./apiServer
docker exec -it -e LISTEN_ADDRESS=10.29.1.3:12345  -e ES_SERVER=10.29.1.9:9200 deploy_n2_1 ./apiServer

docker exec -it -e LISTEN_ADDRESS=10.29.1.4:12345 deploy_n3_1 ./dataServer
docker exec -it -e LISTEN_ADDRESS=10.29.1.5:12345 deploy_n4_1 ./dataServer
docker exec -it -e LISTEN_ADDRESS=10.29.1.6:12345 deploy_n5_1 ./dataServer
docker exec -it -e LISTEN_ADDRESS=10.29.1.7:12345 deploy_n6_1 ./dataServer

docker exec -it -e LISTEN_ADDRESS=10.29.1.2:12345 deploy_n1_1 ./apiServer
docker exec -it -e LISTEN_ADDRESS=10.29.1.3:12345 deploy_n2_1 ./apiServer
```
