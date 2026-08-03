### REST
Representational State Transfer

1. Client-server decoupling (clients could be anything ... server would still respond to them)
2. Uniform interface: all requests for the same resource should look the same, no matter where the request comes from.
e.g.,   GET /v1/users/feed
       HTTP METHOD, RESOURCE, ACTION/RESPONSE
        GET /v1/posts/:postId/comments/:commentId
        There is a hierarchy. 

Go HTTP Handlerfunc()

Create     HTTP             POST (PUT)
Read       ----->           GET
Update                      PATCH/PUT
Delete                      DELETE

3. Statelessness 

------------


- industry standard for building web services 
- All about client-server communication 



+ Why? 
- Standardize -> less decision making 
- Stateless -> Scalable & easy to modify
- Cacheability -> performance

--------------




TCP: a "reliable" connection protocol that runs on top of an unreliable (the client may or may not receive the message in order or if it receives at all) protocol: IP (short for Internet Protocol)




Reliable Delivery 
                                    Acknowledgement (3-way handshake) 
                        
Program (server) <> --------------   Computer (client)

E.g. When we go to google and type in a url such as "https://hello.com", the browser (client) will do a DNS lookup to look for the ip address of the server so that when we connect to the server, there's going to be a TCP connection/handshake, then HTTP will take over. Then ther's going to the HTTP responses from the clients back to the server and then you have your rendred webpage.

















# References
1. [12factor](https://12factor.net/) inspired by Martin Fowler
2. [Roy Fielding's original research](https://ics.uci.edu/~fielding/pubs/dissertation/rest_arch_style.htm)