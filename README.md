```azure
docker build -t atjapan2015/redis-prometheus:latest .
docker push atjapan2015/redis-prometheus:latest
```

```azure
docker run -d --name redis-prometheus -v /tmp/config:/u01/data/sd_configs/redis -p 9091:9091 atjapan2015/redis-prometheus:latest
```

```azure
curl 127.0.0.1:9091/prometheus/targets/add/127.0.0.1_9121
curl 127.0.0.1:9091/prometheus/targets/add/127.0.0.1_9121
```
