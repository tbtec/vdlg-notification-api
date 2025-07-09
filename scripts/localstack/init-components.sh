#!/bin/sh

awslocal sns create-topic \
    --name OrderTopic

awslocal sns create-topic \
    --name ProductionTopic

awslocal sqs create-queue \
    --queue-name ProductionQueueAppOrder \

awslocal sqs create-queue \
    --queue-name ProductionQueueAppNotification \

awslocal sqs create-queue \
    --queue-name OrderQueueAppProduction \

awslocal sqs create-queue \
    --queue-name OrderQueueAppNotification \
