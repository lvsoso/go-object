创建索引
```json
{
"mappings":{
"objects":{
"properties":{
"name":{"type": "string","index":"not analyzed"},
"version":{"type":"integer"},
"size":{"type":"integer"},
"hash":{"type":"string"}
}
}
}
}
```

```shell
#获取对象当前最新版本的元数据
GET /metadata/_search?q=name:<object_name>&size=1&sort=version:desc

#添加新索引 
PUT /metadata/objects/<object_name>_<version>?op_type=create

#获取指定版本号的元数据
GET /metadata/objects/<object_name>_<version_id>/_source

# 获取全部对象版本列表，from和size用于分页显示
GET /metadata/_search?sort=name,version&from=<from>&size=<size>

# 获取指定对象版本列表，from和size用于分页显示
GET /metadata/_search?sort=name,version&from=<from>&size=<size>&q=name: <object_name>


```