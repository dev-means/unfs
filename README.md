# unfs
minio object storage

```shell
#!/usr/bin/env bash

kill $(cat pid)

export MINIO_ACCESS_KEY=minio
export MINIO_SECRET_KEY=minioAdmin

nohup \
./bin/minio server --address :9000 \
./data/minio >> nohup.log & printf $! > pid
sleep 0.1s
```
