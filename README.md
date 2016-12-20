# DroneServiceContainerWeirdness

Service containers only work in drone if the container explicity connects to localhost or the loopback. If a service container gets the ip based on the hostname (a la riak), connecting to localhost won't work because the service conatiner is bound to the internal docker hostname ip. 

As of December 20, 2016 using basho/riak-kv:latest image

