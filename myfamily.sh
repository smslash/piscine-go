#!/bin/bash

curl https://01.alem.school/assets/superhero/all.json | jq --argjson id "$HERO_ID" '. [] | select(.id == $id) | .connections.relatives' | sed 's/^"\(.*\)"$/\1/'
