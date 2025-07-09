#!/bin/sh

awslocal sqs send-message \
    --queue-url http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/ProductionQueueAppOrder \
    --message-body "{ "id": "0196c198-fe43-3abe-c786-9470e3e0b80e", "status": "PENDING"}" 