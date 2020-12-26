# Sample project to test Streams


## Initialize the topics

The scripts folder has a script which will run a container which will execute a command similar to the below.  This will create the topic(s) needed for this sample.  Kafka should be up and running

```bash
# net=(network for bridge created by docker compose)
docker network ls 
docker run -it \
  --net=device-data-sample_default \
  -v "$(pwd)/scripts/initial:/scripts" \
  --rm --name kafka-cli-create-topics \
  confluentinc/cp-enterprise-kafka:6.0.1 \
  bash /scripts/create-topics.sh
```

