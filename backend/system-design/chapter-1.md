# chapter-1
- [chapter-1](#chapter-1)
  - [IP](#ip)
  - [TCP and UDP](#tcp-and-udp)
    - [TCP](#tcp)
    - [UDP](#udp)
    - [TCP vs UDP](#tcp-vs-udp)
  - [Domain Name System (DNS)](#domain-name-system-dns)
    - [How DNS works](#how-dns-works)
  - [Caching](#caching)
  - [Storage](#storage)
    - [RAID](#raid)
    - [](#)

## IP
An IP address is a unique address that identifies a device on the internet or a local network. IP stands for *Internet Protocol*  

## TCP and UDP
### TCP
Transmission Control Protocol (TCP) is connection-oriented, meaning once a connection has been established, data can be transmitted in both directions

TCP has built-in systems to check for errors and to guarantee data will be delivered in the order it was sent

![Img](./FILES/chapter-1.md/2c5e74d2.png)

### UDP
User Datagram Protocol (UDP) is a simplier, connectionless internet protocol in which error-checking and recovery services are not required

![Img](https://raw.githubusercontent.com/karanpratapsingh/portfolio/master/public/static/courses/system-design/chapter-I/tcp-and-udp/udp.png)


### TCP vs UDP

| Feature | TCP | UDP |
| -- | -- | -- |
| Connection | Requires an established connection | Connectionless protocol |
| Guaranteed delivery | Can guarantee delivery of data | Can not guarantee delivery of data |
| Re-transmission | Re-transmission of lost packets is possible | No re-transmission of lost packets |
| Speed | Slower than UDP | Faster than TCP |
| Broadcasting | Does not support broadcasting | Supports broadcasting |
| Use cases | HTTPS / HTTP / SMTP / POP / FTP , etc | Video streaming, DNS, etc |

## Domain Name System (DNS)
Domain Name System (DNS) is a hierachical and decentralized naming system used for translating human-readable domain names to IP addresses

### How DNS works
![Img](https://raw.githubusercontent.com/karanpratapsingh/portfolio/master/public/static/courses/system-design/chapter-I/domain-name-system/how-dns-works.png)

DNS lookup involves the following 8 steps:
1. A client types example.com into a web browser, the query travels to the internet and is received by a DNS resolver
2. The resolver then recursively queries a DNS root nameserver
3. The root server responds to the resolver with the address of a Top Level Domain (TLD)
4. The resolver then makes a request to the `.com` TLD
5. The TLD server then responds with the IP address of the domain's nameserver, `example.com`
6. Lastly, the recursive resolver sends a query to the domain's nameserver
7. The IP address for `example.com` is then returned to the resolver from the nameserver
8. The DNS resolver then responds to the web browser with the IP address of the domain requested initially

## Caching

![Img](https://raw.githubusercontent.com/karanpratapsingh/portfolio/master/public/static/courses/system-design/chapter-I/caching/caching.png)


## Storage
Storage is a mechanism that enables a system to retain data, either temporarily or permanently


### RAID

### 