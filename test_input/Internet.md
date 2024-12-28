2024/08/24 17:03
Status: #idea
Tags: #academic #technology #concept 
# What is the Internet?
The internet is a collection of computers and networks that transmit data using the standard Internet Protocol (IP). In general, a computer on its own cannot get onto the internet, it communicates with a sort of "middleman" to have access to the internet, the "middleman" is reffered to as a proxy server or gateway. It is the computer that the computer must communicate with to get access to the internet.
___
# Computers and the Internet
- A shared medium, language, and protocol is required to communicate with others on the internet; medium is generally a wired or wireless link
- Internet Service Provider may be a phone, cable, or satellite company
- [[Computers|Computer]] communicates with an ISP computer which relays information to and from the internet on your behalf
- Computers on the internet are given a unique IP address
- Communication format (language) and protocol used is called TCP/IP (Transmission Control Protocol/Internet Protocol)
- The internet is an asynchronous, point-to-point communication system however the speed of electronic communications may cause it to appear synchronous

# TCP/IP (Transmission Control Protocol/Internet Protocol)
TCP/IP (Transmission Control Protocol/Internet Protocol) is the structure (language) and protocol used for communication between computers on the internet. 
## How TCP/IP Works:
- Information is broken into sequence of small fixed-size units called IP packets. Each packet has space for the unit of data, the source and destination IP addresses, and a sequence number
- Packets are sent over the internet one at a time using whatever route is available
- Because each packet can take a different route, congestion and service interruptions do not delay transmissions
- Receiver reassembles packets using their sequence numbers
## Moving Packets
- The way packets are transmitted between computers on the internet is dependent on the medium of transmission
- Internet backbone (defined by the principal data routes between large, strategically interconnected computer networks and core routers of the Internet) has largest capacity (bandwidth) and is optical fiber
- Smaller bandwidth connections may use copper fiber or cable to connect machines to hubs and routers
- End users may use phone lines, cable lines, or even fiber optic connections as their first hop (connection) to the internet
# The Internet's Layers
The internet is divided into layers so that its complexity can be dealt with more efficiently. Layering is a technique for dealing with complex systems in which each layer: provides services for the layers above and uses services for the layers below.
## Five Internet Layers:
- Application- supports messages between programs
- Transport- process-to-process data transfer
- Network (internet)- send packets from source to destination
- Link- data transfer between neighbors
- Physical- encoding of bits on medium
![[Internet Layers Visual|700]]
The above visual is the [[Encapsulation|encapsulation]] of application data descending through the layers. 
![[Communication of Computers|700]]
The above shows two internet hosts connected via two routers and the corresponding layers used at each hop.

# Communication Systems
Communication is the act of sending information from one party to another. There is a sender and a receiver. During broadcast communication (multicast), there is a single sender and many receivers, while with point to point, there is only a single sender and a single receiver.
<u>Synchronous communication:</u> sender and receiver are active at the same time e.g. telephone call, instant messaging
<u>Asynchronous communication:</u> sending and receiving occur at different times e.g. email

# IP Addresses
An IP or internet protocol address is a series of numbers that identifies any device on a network (used to differentiate different devices connected to the internet). Computers use IP addresses to communicate with each other both over the internet as well as on other networks. There are two main IP address versions, IPv4 and IPv6. IP addresses are unique but because they are so long, we require domain names- text names for computers that are easier to remember.
### IPv4
- Consists of four numbers (each of the four sections is referred to as an octet) from the range of 0 to 255
- Seperated by dots
	-  Example: 142.231.95.1
### IPv6
- Because there are an ever-increasing number of computers and devices being added to the internet, there is an ongoing transfer to IP version 6 addresses which is 128 bits -> eight 16 bit numbers separated by colons
- Example: `2002:CE57:25A2:0000:0000:0000:CE57:25A2`
- Note: When a two-byte chunk is comprised of all zeroes, you can simply omit the zeroes
	- `2001:0db8:c9d2:0012:0000:0000:0000:0051` is the same as `2001:db8:c9d2:12::51`
- The address `::1` is the _loopback address_. It always means “this machine I’m running on now”. In IPv4, the loopback address is `127.0.0.1`.
- Finally, there’s an IPv4-compatibility mode for IPv6 addresses that you might come across. If you want, for example, to represent the IPv4 address `192.0.2.33` as an IPv6 address, you use the following notation: “`::ffff:192.0.2.33`”.
## Manipulating IP Addresses
Suppose you have a `struct sockaddr_in ina` and you have an IP address “`10.12.110.57`” or “`2001:db8:63b3:1::3490`” that you want to store into it, the function you want to use, `inet_pton()`, converts an IP address in numbers-and-dots into either a `struct in_addr` or `struct in6_addr` depending on whether you specify `AF_INET` or `AF_INET6`. `pton` stands for "presentation to network" or "printable to network" as some like to call it. The conversion can be made as follows:
```C
struct sockaddr_in sa; // IPv4
struct sockaddr_in6 sa6; // IPv6  

inet_pton(AF_INET, "10.12.110.57", &(sa.sin_addr)); // IPv4
inet_pton(AF_INET6, "2001:db8:63b3:1::3490", &(sa6.sin6_addr)); // IPv6
```
**NOTE**: The old way of doing things involved using a function called `inet_addr()` or `inet_aton()` which are now obsolete and don't work with IPv6.
`inet_pton()` returns -1 on error or 0 if the address is messed up so when calling the function, it's important to have error handling for those cases (check that result is greater than 0).
To convert from the binary representation of IP to presentation format (numbers-and-dots for IPv4 and hex-and-colons for IPv6), you can use `inet_ntop()` ("ntop" means network to presentation/printable):
```C
//IPv4:
char ip4[INET_ADDRSTRLEN]; // space to hold the IPv4 string
struct sockaddr_in sa; // assume this is loaded with something

inet_ntop(AF_INET, &(sa.sin_addr), ip4, INET_ADDRSTRLEN);
printf("The IPv4 address is: %s\n", ip4);

//IPv6:
char ip6[INET6_ADDRSTRLEN]; // space to hold the IPv6 string
struct sockaddr_in6 sa6; // assume this as well is loaded

inet6_ntop(AF_INET6, &(sa6.sin6_addr), ip6, INET6_ADDRSTRLEN);
printf("The IPv6 address is: %s\n", ip6);
```
When you call it, you pass the address type, the address, a pointer to a string to hold the result, and the maximum length of that string. (Two macros conveniently hold the size of the string you'll need to hold the largest IPv4 or IPv6 address: `INET_ADDRSTRLEN` and `INET6_ADDRSTRLEN`.)
**NOTE**: The old way of doing this "network to presentation" conversion (which did not work for IPv6 btw) was `inet_ntoa()` which is now obsolete.
Furthermore, these functions only work with numeric IP addresses and cannot deal with nameserver DNS lookup on a hostname - "`www.example.com`". `getaddrinfo()` will be used to do that.
# From IPv4 to IPv6
1. Try to use `getaddrinfo()` to get all the `struct sockaddr` info instead of packing the structs by hand. This will keep you IP version-agnostic and eliminate many of the subsequent steps
2. Any place where you're hard-coding anything related to the IP version, try to wrap up in a helper function
3. Change `AF_INET` to `AF_INET6`
4. Change `PF_INET` to `PF_INET6`
5. Change `INADDR_ANY` assignments to `in6addr_any` assignments, which are slightly different:
```C
struct sockaddr_in sa;
struct sockaddr_in6 sa6;

sa.sin_addr.s_addr = INADDR_ANY; // use my IPv4 address
sa6.sin6_addr = in6addr_any; // use my IPv6 address
```
Also, the value `IN6ADDR_ANY_INIT` can be used as an initializer when the `struct in6_addr` is declared, like so:
```C
struct in6_addr ia6 = IN6ADDR_ANY_INIT;
```
6. Instead of `struct sockaddr_in` use `sockaddr_in6`, being sure to add "6" to the fields as appropriate. There is no `sin6_zero` field.
7. Instead of `struct in_addr` use `struct in6_addr`, being sure to add "6" to the fields as appropriate.
8. Instead of `inet_aton()` or `inet_addr()`, use `inet_pton()`
9. Instead of `inet_ntoa()`, use `inet_ntop()`
10. Instead of `gethostbyaddr()`, use the superior `getnameinfo()` (although `gethostbyaddr()` can still work for IPv6)
11. `INADDR_BROADCAST` no longer works. Use IPv6 multicast instead
# Subnets
- For organizational reasons, it is sometimes useful to declare that "this first part of this IP address up through this bit is the network portion of the IP address, and the remainder is the host portion"
	- For instance, with IPv4, if we have the IP `192.0.2.12`, we could say that the first three bytes are the network and the last byte is the host `192.0.2.0` (we zero out the byte that was the host)
	- We would, thus, be talking about host 12 on network 
- In the old days, there even used to exist classes of subnets where the first one, two, or three bytes of the address was the network part
	- The classes:
		- **Class A**: One byte for the network and three for the host -> allowing for 24-bits worth of hosts on your network (~16 million)
		- **Class C**: Three bytes of network and one byte of host (256 hosts minus a few reserved ones)![](Pasted%20image%2020240802172851.png)
	- Thus, there were only a few class As, a huge pile of class Cs, and some class Bs somewhere in the middle
	- The network portion of the IP address is described by something called the *netmask*, which you bitwise-AND with the IP address to get the network number out of it
		- The netmask usually looks something like `255.255.255.0`
		- (E.g. with that netmask, if your IP is `192.0.2.12`, then your network is `192.0.2.12` AND `255.255.255.0` which gives `192.0.2.0`.)
	- Unfortunately, this wasn't fine-grained enough for the eventual needs of the internet and we are running out of Class C networks quite quickly
		- To remedy this, the Internet gods allowed for the netmask to be an arbritrary number of bits, not just 8, 16, or 24
		- So you might have a netmask of, say `255.255.255.252`, which is 30 bits of network, and 2 bits of host allowing for four hosts on the network. (Note that the netmask is _ALWAYS_ a bunch of 1-bits followed by a bunch of 0-bits.)
	- However, it’s a bit unusual to use a big string of numbers like `255.192.0.0` as a netmask
		- First of all, people don’t have an intuitive idea of how many bits that is, and secondly, it’s really not compact
		- In the new style, you just put a slash after the IP address, and then follow that by the number of network bits in decimal. Like this: `192.0.2.12/30`.
		- Or, for IPv6, something like this: `2001:db8::/32` or `2001:db8:5413:4028::9db9/64`.
# Port Numbers
- Besides an IP address (used by the IP layer), there is another address that is used by TCP (stream sockets), and also by UDP (datagram sockets) - the *port number*
	- It's a 16 bit number that's like the local address for the connection
- Think of IP address as the street address of a hotel and the port number as the room number 
- If a computer is to handle incoming mail and web services, how can you differentiate between the two on a computer with a single IP address?
	- Different services on the internet have different well-known port numbers
		- You can see them all in [the Big IANA Port List](https://www.iana.org/assignments/port-numbers) or, if you’re on a Unix box, in your `/etc/services` file
	- HTTP (the web) is port 80, telnet is port 23, SMTP is port 25, the game DOOM used port 666 (smh my head)
	- Ports under 1024 are often considered special and require special OS privileges to use
# Domains
 A domain is a related group of networked computers. Domain names are organized hierarchically. The most general part of the hierarchy at the end of the name. A domain hierarchy or a DNS hierarchy is a system used to sort the parts of a domain according to their importance. Top-level domains appear in the last part of the domain name; some examples consist of: `com`, `edu`, `org`, `net`, `mil`, `gov`, and `ca`. Top-level domains are controlled by ICANN
![[Domain Heirarchy|700]]

# DNS Servers
DNS stands for Domain Name System and translates the human-readable names into IP addresses. A DNS server has the same function as a phonebook.
- Each computer knows the IP address of its nearest DNS server
- When you use a domain name in a request, computer asks the DNS server to look up the IP address
- If closest DNS server does not recognize the IP address, it asks a root name server which keeps a master list of name-to-address relationships
- There are 13 root name servers (with multiple mirrored instances) distributed across the globe
- <u>Dynamic Host Configuration Protocol (DHCP)</u> automatically assigns an IP address to any device that connects to a network
	- IP will change when you connect to a different access point
	- Static IP address for servers does not change
	- DHCP also assigns local DNS server addresses

# The World Wide Web (WWW)
- The world wide web is an application built on top of the internet that allows for the display and transmission of documents called web pages 
- Web page is a document that contains mark-up that allows it to be displayed graphically by a web browser. The page may also contain hyperlinks to link to related web pages.
- A web server is a computer on the internet with the task of storing web pages and repsonding to client's requests for them
- The World Wide Web (WWW) is the web servers and the files they store.
## Clients and Servers:
When you click a link in your browser, your computer becomes the client requesting the appropriate web page from the server that stores that page (web server). Once the web page is sent to you, the client-server interaction is complete.
<u>Server:</u> a computer that stores information such as a web page, email, database, etc.
<u>Client:</u> a client is the computer that is requesting information stored at a server
## Requesting a Webpage:
<u>Web page is requested by the user by either:</u>
- Typing in a URL (Universal Resource Locator) into the web browser's address field OR
- clicking on a hyperlink in a document that contains a URL
<u>A request for a URL has three parts:</u>
- Protocol: `http://` Hypertext Transfer Protocol
- Server computer's domain name or IP address
- Page's path and file name: tells the server which file (page) is requested and where to find it
![[Pasted image 20221002183112.png]]
## Hypertext
The HTTP protocol can transmit any file type, not just documents but it is most commonly used to transmit documents written in HTML. HTML describes the layout of a document including fonts, text style, image placement, and hypertext links. Hypertext links provides a way to jump from point to point in documents (non-linear). Links may jump within a document, between documents on a server, and to documents on other servers.
## Firewalls
A firewall is a network device that is installed on the edge of a network to prevent unauthorized network traffic from entering a local network.
- Uses information in the packets (IP addresses, ports) to determine good traffic from bad traffic
- Admins can restrict access to certain sites or applications using a firewall
# Ethernet Protocol
Ethernet is a commonly used protocol for communicating with computers on a local area network (LAN). It is a very simple and decentralized protocol. To send:
- A computer listens to the channel and if its quiet, then its free
- The computer starts sending on the channel
- While sending, the computer listens to make sure it is the only one sending. If not, it stops for a random amount of time and then continues again
# How Search Engines Work
All popular search engines have two basic parts:
- <u>Crawler</u>: visits sites on the internet discovering web pages and building an index to the web's content; search engines have crawlers running continuously to refresh and update its index database of web pages; when a crawler visits the page, it identifies the terms on the page and processes any outgoing links
- <u>Query processor</u>: looks up user-submitted keywords in the index and reports back which web pages the crawler has found containing those word; The query processor does not search the Internet – it only returns answers previously found by the crawler; the ranking algorithm to identify important pages is critical to success of the search engine. Google uses the PageRank algorithm

___
# References
- Brian “Beej Jorgensen” Hall. "Beej's Guide to Network Programming Using Internet Sockets." v3.1.12, July 17, 2024, https://beej.us/guide/bgnet/html//index.html#connect