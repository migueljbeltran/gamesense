#!/bin/sh
set -e

awslocal s3 mb s3://gamesense-dev || true

awslocal s3api put-bucket-cors \
  --bucket gamesense-dev \
  --cors-configuration '{
    "CORSRules": [
      {
        "AllowedOrigins": ["http://localhost:3000"],
        "AllowedMethods": ["GET", "PUT", "POST", "DELETE", "HEAD"],
        "AllowedHeaders": ["*"],
        "ExposeHeaders": ["ETag"],
        "MaxAgeSeconds": 3000
      }
    ]
  }'

DLQ_URL=$(awslocal sqs create-queue --queue-name gamesense-jobs-dlq --query QueueUrl --output text)
DLQ_ARN=$(awslocal sqs get-queue-attributes \
  --queue-url "$DLQ_URL" \
  --attribute-names QueueArn \
  --query 'Attributes.QueueArn' \
  --output text)

awslocal sqs create-queue \
  --queue-name gamesense-jobs \
  --attributes "RedrivePolicy={\"deadLetterTargetArn\":\"$DLQ_ARN\",\"maxReceiveCount\":\"3\"},VisibilityTimeout=600" \
  || true
