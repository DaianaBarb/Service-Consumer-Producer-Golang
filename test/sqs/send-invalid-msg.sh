#!/bin/bash
aws sqs send-message --endpoint-url http://localhost:4566 --queue-url http://localhost:4566/000000000000/migration-queue --message-body "dado-invalidao xD" --profile localstack
