#!/bin/sh

# create song
curl \
  -X POST \
  -d '{"title": "Raised by Wolves", "artist": "Falling in Reverse"}' \
  -H "Content-Type: application/json" \
  http://localhost:8080/song

# get song
curl \
  -X GET \
  -H "Content-Type: application/json" \
  http://localhost:8080/song/38
