#!/bin/bash

if [ -z "$TRAVIS_TAG" ]; then
  echo 'No tag, no relase'
  exit 0
fi

go get github.com/inconshreveable/mousetrap
curl -sL https://git.io/goreleaser | bash
