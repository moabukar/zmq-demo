# Zero MQ testing

- [Zero MQ testing](#zero-mq-testing)
  - [Using Python](#using-python)
    - [Method 1: Req Reply pattern](#method-1-req-reply-pattern)
    - [Method 2: Pub Sub pattern](#method-2-pub-sub-pattern)
    - [Method 3: Push Pull (Pipeline) pattern](#method-3-push-pull-pipeline-pattern)
    - [Method 4: Pair pattern](#method-4-pair-pattern)
  - [Using Go](#using-go)
      - [Method 1: Pub Sub pattern](#method-1-pub-sub-pattern)
      - [Method 2: Req Reply pattern](#method-2-req-reply-pattern)

## Using Python

- Run the server and client in separate terminals 

- `pip3 install zmq`

### Method 1: Req Reply pattern

```
python3 py/reqreply/server.py (run the server)

python3 py/reqreply/client.py (send messages to server)

```

### Method 2: Pub Sub pattern

```

python3 py/pubsub/pub.py (run the publisher)

python3 py/pubsub/sub.py (run the subscriber) - gets messages from publisher

```

### Method 3: Push Pull (Pipeline) pattern

```

python3 py/pipeline/manager.py (run the manager)

python3 py/pipeline/worker.py (run the worker) - gets messages from manager

```

### Method 4: Pair pattern

```

python3 py/pair/bind.py (run the bind)

python3 py/pair/connect.py (run the connect) - gets messages from bind

```

## Using Go


#### Method 1: Pub Sub pattern

```

go run go/pubsub/pub.go (run the publisher)

go run go/pubsub/sub.go (run the subscriber) - gets messages from publisher

```

#### Method 2: Req Reply pattern

```

go run go/reqreply/server.go (run the server)

go run go/reqreply/worker.go (run the worker) - gets messages from server

```
