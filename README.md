# slogger
Smart logging service, custom one, how to handle 1M requests / s.
Repository contains all microservices with terraform and code to scale the thing up!

# Architecture:
`N * [clog] -> [Aggregator] -> [Sender]`

For now we can have **N** clients of gRPC service to send logs. **Aggregator** is grabbing logs from all
the clients into one channel that is used in the sender. For sure, we need to create more advanced logic
for the **aggregator** (maybe spawn M Aggregators) which then are going to be read in the sender. There
are a lot of places that can be a bottleneck, and all the code needs to be well benchmark-tested.