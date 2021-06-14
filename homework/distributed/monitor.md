# 监控安装设置

## grafana

### install

```sh
yum install -y fontconfig freetype* urw-fonts
wget https://dl.grafana.com/oss/release/grafana-8.0.1-1.x86_64.rpm
rpm -ivh grafana-8.0.1-1.x86_64.rpm    
```

### import 763 dashboard

## prometheus

### prometheus-server

1. install

    ```sh
    wget https://github.com/prometheus/prometheus/releases/download/v2.27.1/prometheus-2.27.1.linux-amd64.tar.gz
    tar xzf prometheus-2.27.1.linux-amd64.tar.gz
    ```

2. 配置

    ```sh
    cd prometheus-2.27.1.linux-amd64
    vi prometheus.yml
    - job_name: redis_exporter
        static_configs:
        - targets:
            - redis://localhost:6379
            - redis://localhost:6380
        metrics_path: /scrape
        relabel_configs:
        - source_labels: [__address__]
            target_label: __param_target
        - source_labels: [__param_target]
            target_label: instance
        - target_label: __address__
            replacement: 127.0.0.1:9121      
    ```

3. 启动

    ```sh
    nohup ./prometheus --web.enable-lifecycle --storage.tsdb.retention=32d >/dev/null 2>&1 &
    #验证
    curl -v http://localhost:9090/targets
    ```

### redis_exporter

```sh
wget https://github.com/oliver006/redis_exporter/releases/download/v1.24.0/redis_exporter-v1.24.0.linux-amd64.tar.gz

tar xzf redis_exporter-v1.24.0.linux-amd64.tar.gz

cd redis_exporter-v1.24.0.linux-amd64

nohup ./redis_exporter >/dev/null 2>&1 &

curl -v http://localhost:9121/metrics
```
