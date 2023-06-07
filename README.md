
# 什么是 Grafana 
Grafana 是由 Grafana Labs 开发的开源监控监控系统，你可以用它来监控线上系统，极大地简化运维和开发工作。

它不对数据源作假设，因此你可以用包括 Prometheus, MySQL 在内的任何（时序）数据库搭配使用。

# Getting Started

### 如何启动 Grafana

请 clone 本教程代码，然后确认本地已经安装 docker-compose 后，在本目录运行

``````
 docker-compose -f docker-compose.yml up -d
``````

notes:
   - docker-compose.yml中的networks，需要和 easynode环境在同一网络之内


### 登录管理页面

在你的浏览器中访问 `localhost:4000` 即可看到运行的 Grafana

notes：
 - Grafana 默认账号密码： admin/admin

## 配置

- 配置prometheus
  - 进入 Connections/Connect Data,Search Prometheus 插件并安装
  - HTTP 配置，URL: http://prometheus:4090,其他字段默认即可
  - 点击 Save&test

- 配置 Kafka 的dashboards
  ``````
  Grafana Dashboard ID: 7589
  name: Kafka Exporter Overview.
  ``````

　[dashboard](https://grafana.com/grafana/dashboards/7589-kafka-exporter-overview/)
  [exporter](https://github.com/danielqsj/kafka_exporter)

- 配置node 的dashboards

  ``````
  Grafana Dashboard ID:11074
  name:Node Exporter Full
  
  ``````
  [dashboard](https://grafana.com/grafana/dashboards/1860-node-exporter-full/)
 
- 配置 clickhouse 

  - 进入 Home/Connections/Connect data 目录
  - search ClickHouse,进入详情页并安装
  - 进入 clickhouse settings 标签
    - Server address: clickhouse 服务器IP (tcp 协议)
    - Server port：clickhouse 服务器端口，默认：9000
    - username:clickhouse 用户名，默认：test
    - password:clickhouse 密码， 默认：test
  - 点击 "save&test"
  - 进入 dashboards 标签 ，选择自己喜欢的风格，点击import，完成。建议选择 "Query Analysis"
  - 扩展配置自己视图
   








[grafana教程](https://grafana.com/docs/grafana/latest/getting-started/)

[dashboards](https://grafana.com/grafana/dashboards/?plcmt=footer)


