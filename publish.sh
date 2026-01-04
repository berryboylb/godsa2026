#!/bin/bash

if [ $# -eq 0 ]; then
  echo "Usage: ./publish.sh \"Gist Title\""
  exit 1
fi

gh gist create solution.go problem.md notes.md \
  -d "$1" \
  -p
echo "Gist published successfully."