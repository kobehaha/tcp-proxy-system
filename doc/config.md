
# Config parameters #

| **Parameter Name** | **Required?** | **Default Value** | **Example** |
|:-------------------|:--------------|:--------------------|:------------------|:------------|
|[service](#service)|Yes           |Default Service: tcp proxy  |"service" : "http proxy for tomcat"|
|[host](#host)        |No            |127.0.0.1                   |“host” : "192.168.1.3"|
|[port](#port)      |No            |80                          |"port" : 3306   |
|[strategy](#strategy)|No          |iphash                      |"strategy" : "iphash" |
|[keepalive](#keepalive)|No           |20                          |"keepalive" : 20|
|[protocol](#protocol)|No          |tcp                         |"protocal" : tcp |
|[maxprocessor](#maxprocessor)|No          |2                         |"maxprocessor" :2 |
|[logfile](#logfile)|No          |/tmp/logs                         |"logfile" : "/tmps/logs" |
|[backends](#backends)|Yes          |-                         |"backends" :   { "host":"192.168.33.19", "port": 8000 }"|
