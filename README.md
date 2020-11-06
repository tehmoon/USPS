# Ultimate Secure Proxy Server

USPS is a new project that aims to replace the complicated and old Squid forward proxy for filtering outgoing traffic in compliance and secure environments with a modern language and container ready simple executable.

Here's a list of various features that would be great to implement:
- Simple executable written in Go
- Allow list based on protocols:
  - http host header
  - tls SNI
  - SSH
- Prometheus ready
- Formatted audit logs to stdout
- Health checks
- Transparent proxy with protocol bypass
- Transparent proxy with protocol interception
- SOCKS proxy server endpoint
- HTTP proxy server endpoint
- Service discovery
  - docker
  - kubernetes
- IP client allow list
- IP/DNS target allow list
- Configuration datastore
  - etcd
  - file
- Automatic protocol detection
- Scalable, small executable, minimal resources 

# Rollout phases

Phase one:
- Config datastore file with yaml
- allow list with http/tls protocol support
- Formatted audit logs to stdout
- HTTP proxy server endpoint
Phase two:
- Health checks
- Prometheus metrics endpoint

# Architecture

TBD

```
connection established -> protocol detection -> host:port destination parsing -> allow/block list check -> proxy
```
