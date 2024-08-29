#!/bin/bash

# ChatStream test
grpcurl -proto proto/aibro/v1/aibro.proto -plaintext -d '{"message":"Hello"}' \
  localhost:8080 aibro.v1.AIBroService/ChatStream
