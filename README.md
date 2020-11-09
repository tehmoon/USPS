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

# Configuration

```
modules:
  http_proxy:
    port: 3128
    intercept_tls: false

# list of host: ip that you can use in conjuction of service discovery
# Use the FQDN if you want to add a regex domain
static_hosts:
  web1: 172.17.0.5
  web2: 172.17.0.6
sources:
  marketing:
  - 172.16.59/24
destinations:
  facebook_443:
  - hostnames:
    - $facebook\.com^
    - $faceblick\.com^
    ports:
    - 443

# A source has the following types:
# - cidr (default): simply CIDR that is checked against the tcp source packet
# - hostname: regex translation of the ip address to a name. Useful for auto discovery

# A destination is either:
# - cidr (default when string)
# - a regex hostname
# - an object Destination:
#   - hosts: array of regex hostnames
#   - ports: array of ports
allow:
- sources: marketing
  destinations: facebook_443
  modules:
  - http_proxy
- sources:
  - 127.0.0.1/32
  - web1
  - web2
  destinations:
  - hostnames:
    - $google\.com^
    ports:
    - 443
# By default, all modules are matching
  modules:
  - http_proxy

# Flow:
# - check allow list
#   - in: proceed
#   - not in: fail
# - check deny list
#   - in: fail
#   - no in: proceed
deny:
- destinations:
  - $facebook\.com^
  - 127.0.0.1/32
```
