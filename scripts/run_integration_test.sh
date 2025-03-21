#!/bin/bash

if ./scripts/integration_test.sh; then
  echo "SUCCESS! YEAH!"
  exit 0
else
  echo "Something is not well :("
  exit 1
fi