!/bin/bash

aws --version
export AWS_ACCESS_KEY_ID="test"
export AWS_SECRET_ACCESS_KEY="test"
export AWS_DEFAULT_REGION="us-east-1"
aws configure --profile localstack
aws configure list


# cria a fila local

aws sqs create-queue --endpoint-url http://localhost:4566 --queue-name testQueue --profile localstack
