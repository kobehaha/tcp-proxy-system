
## Tcp-Proxy
simple tcp proxy for server

[![Build Status](https://travis-ci.org/kobehaha/tcp-proxy-system.svg?branch=master)](https://travis-ci.org/YOUR/PROJECT)
[![Go Report Card](https://goreportcard.com/badge/github.com/kobehaha/tcp-proxy-system)](https://goreportcard.com/report/github.com/kobehaha/tcp-proxy-system)

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


## Details of Tcp-proxy

- doing 

## License

Tcp-proxy is under the GPL license. See the [LICENSE](./LICENSE) directory for details.



