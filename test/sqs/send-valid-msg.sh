#!/bin/bash

product_id="133257344"
skus=("133670391" "133257352")

json_object="{\"message\":\"test message Daiana soares\"}"


echo "$json_object"

 aws sqs send-message --endpoint-url http://localhost:4566 --queue-url http://localhost:4566/000000000000/testQueue --message-body "$json_object" --profile localstack
