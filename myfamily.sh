#!/bin/bash
export HERO_ID
curl https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json  | jq --argjson HERO_ID "$HERO_ID"  ' .[] | select( .id== $HERO_ID ) | .connections.relatives' | sed s/^.// | sed  s/.$//
