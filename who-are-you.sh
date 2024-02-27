#!/bin/bash
curl -s https://learn.zone01kisumu.ke/assets/superhero/all.json | jq -r '.[] | select(.id == 70) | .name' | \
while read -r name; do
    echo "\"$name\""
done