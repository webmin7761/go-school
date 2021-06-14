# benchmark-command

jemalloc与libc就只访问端口有差别，其它一样

## jemalloc

```sh
./redis-cli -p 6379 flushall
./redis-benchmark -q -d 10 -r 1000000 -n 500000 -t set -p 6379

./redis-cli -p 6379 flushall
./redis-benchmark -q -d 20 -r 1000000 -n 500000 -t set -p 6379

./redis-cli -p 6379 flushall
./redis-benchmark -q -d 50 -r 1000000 -n 500000 -t set -p 6379

./redis-cli -p 6379 flushall
./redis-benchmark -q -d 100 -r 1000000 -n 500000 -t set -p 6379

./redis-cli -p 6379 flushall
./redis-benchmark -q -d 200 -r 1000000 -n 500000 -t set -p 6379

./redis-cli -p 6379 flushall
./redis-benchmark -q -d 500 -r 1000000 -n 500000 -t set -p 6379

./redis-cli -p 6379 flushall
./redis-benchmark -q -d 1024 -r 1000000 -n 500000 -t set -p 6379

./redis-cli -p 6379 flushall
./redis-cli -p 6380 flushall
./redis-benchmark -q -d 5120 -r 1000000 -n 500000 -t set -p 6379
```

## libc

```sh
./redis-cli -p 6380 flushall
./redis-benchmark -q -d 10  -r 1000000 -n 500000 -t set -p 6380

./redis-cli -p 6380 flushall
./redis-benchmark -q -d 20 -r 1000000 -n 500000 -t set -p 6380

./redis-cli -p 6380 flushall
./redis-benchmark -q -d 50 -r 1000000 -n 500000 -t set -p 6380

./redis-cli -p 6380 flushall
./redis-benchmark -q -d 100 -r 1000000 -n 500000 -t set -p 6380

./redis-cli -p 6380 flushall
./redis-benchmark -q -d 200 -r 1000000 -n 500000 -t set -p 6380

./redis-cli -p 6380 flushall
./redis-benchmark -q -d 500 -r 1000000 -n 500000 -t set -p 6380

./redis-cli -p 6380 flushall
./redis-benchmark -q -d 1024 -r 1000000 -n 500000 -t set -p 6380

./redis-cli -p 6379 flushall
./redis-cli -p 6380 flushall
./redis-benchmark -q -d 5120 -r 1000000 -n 500000 -t set -p 6380
```
