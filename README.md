
## Tcp-Proxy
simple tcp proxy for server

[![Build Status](https://travis-ci.org/kobehaha/tcp-proxy-system.svg?branch=master)](https://travis-ci.org/YOUR/PROJECT)
[![Go Report Card](https://goreportcard.com/badge/github.com/kobehaha/tcp-proxy-system)](https://goreportcard.com/report/github.com/kobehaha/tcp-proxy-system)
[![CircleCI Status](https://circleci.com/gh/kobehaha/tcp-proxy-system.svg?style=shield)](https://circleci.com/gh/kobehaha/tcp-proxy-system)

## Overview

Tcp-proxy is a high-performance proxy for tcp layer by Go, Just like other tcp proxy such as nginx, haproxy. you can use it to split your traffic to a few backends. Tcp-proxy aims to simplify the tcp-proxy solution of Proxy.

## Feature

### 1 Basic Function
- Support limiting the max count of connections to proxy
- Support custom load balance algorithm
- Support dynamic checking backend by heartbeat service

### 2. Automatic Function
- Dynamic add/remove backend without restarting or chaning config of proxy [doing]
- Client's ip ACL control [doing]


## Install

```
    1 Install Go
    2 git clone https://github.com/kobehaha/tcp-proxy-system.git
    3 cd tcp-proxy-system
    4 make
    5 cd bin
    6 proxyd

```

## Usege


* **default config file**

	*default config PATH*

		$(tcp-proxy-system)/config/default.json

	*start*

		tcp-proxy-system/bin/proxyd


* **special config file**

	*start*

		proxyd -conf=/etc/config.json

	*log*

	```
	2017/10/21 23:31:12 success load config file: /etc/default.json
	2017/10/21 23:31:12 server.go:69: start server and server address is :1001
	2017/10/21 23:31:12 server.go:76: proxy server start ok
	2017/10/21 23:31:17 proxy_tcp.go:151: clean backendUrl which is not available ----->  192.168.33.19:8000
	```






## Details of Tcp-proxy

- design (doing)

## License

Tcp-proxy is under the GPL license. See the [LICENSE](./LICENSE) directory for details.
