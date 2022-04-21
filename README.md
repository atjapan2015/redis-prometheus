```azure
docker build -t atjapan2015/redis-prometheus:latest .
```

```azure
docker run -d --name redis-prometheus -v /tmp/config:/u01/data/sd_configs/redis -p 9091:9091 atjapan2015/redis-prometheus:latest
```

```azure
curl 192.168.31.15:9091/prometheus/targets/add/192.168.1.3_6379
```
