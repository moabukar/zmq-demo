# Zero MQ testing


## Using Python

```

## Run the server and client in separate terminals 

pip3 install zmq

# Method 1: Req Reply pattern

python3 py/reqreply/server.py (run the server)

python3 py/reqreply/client.py (send messages to server)

# Method 2: Pub Sub pattern

python3 py/pubsub/pub.py (run the publisher)

python3 py/pubsub/sub.py (run the subscriber) - gets messages from publisher

# Method 3: Push Pull (Pipeline) pattern

python3 py/pipeline/manager.py (run the manager)

python3 py/pipeline/worker.py (run the worker) - gets messages from manager

```

## Using Go
