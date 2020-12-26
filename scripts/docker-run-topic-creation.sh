#!/usr/bin/env bash
docker run --net=device-data-sample_default -it -v "$(pwd)/initial:/scripts" --rm --name kafka-cli-create-topics confluentinc/cp-enterprise-kafka:6.0.1 bash /scripts/create-topics.sh
