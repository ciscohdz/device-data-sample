echo "Waiting for Kafka..."

cub kafka-ready -b kafka-1:9092 1 20

# create the topic which will receive device signals
kafka-topics \
  --bootstrap-server kafka-1:9092 \
  --topic raw-device-signals \
  --replication-factor 1 \
  --partitions 3 \
  --create
