# slogger
Smart logging service, custom one, how to handle 1M requests / s.
Repository contains all microservices with terraform and code to scale the thing up!

# Architecture:
`(N * [clog] -> [Aggregator] -> [Sender]) -> cLB-> ([serialize])`

For now we can have **N** clients of gRPC service to send logs. **Aggregator** is grabbing logs from all
the clients into one channel that is used in the sender. For sure, we need to create more advanced logic
for the **aggregator** (maybe spawn M Aggregators) which then are going to be read in the sender. There
are a lot of places that can be a bottleneck, and all the code needs to be well benchmark-tested.

# LoadBalancer
Between sender and serialize we are planning to have load balancer. After studying it for a time, the best option
is going to be the **client load balancer** where client is `sender`.

### +Pros
 - High performance because elimination of extra hop

### -Cons
 - Complex client
 - Client keeps track of server load and health
 - Client implements load balancing algorithm
 - Client needs to be trusted, or the trust boundary needs to be handled by a lookaside LB

As we can see, the high performance is kind of selling point here. But we need to answer those cons.
Let's start with complex client. Even if the client is going to be complex, who cares? Everything is part
of our logger context, so implementation details are not intresting for users. Rest of the points can be
answered by the same thing, the implementation detail belongs to our app, no one needs to use this balancer
anywhere else.