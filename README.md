# 说明

基于[phonedata](https://github.com/xluohome/phonedata)的golang的手机归属地查询系统，特点：
1. docker / docker swarm / k8s 部署
2. rest api接口进行分词

## Docker & Swarm/Composer
```bash
docker pull hetao29/mobilequery:latest
```

## 编译
```bash
make build
```

## 运行

```bash
make start
```

## 测试

```bash
make test
```
结果
```json
curl "http://127.0.0.1:8040/query?mobile="
{"message":"pong","record":{"PhoneNum":"","Province":"","City":"","ZipCode":"","AreaZone":"","CardType":""}}
```
