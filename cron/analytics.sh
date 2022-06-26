#!/bin/bash

URL=$1
MEASUREMENT_ID=$2
API_SECRET=$3

DATA=`curl -X GET ${URL} | jq ".icestats.source"`
LISTENERS=`jq .listeners <<< $DATA`
LISTENER_PEAK=`jq .listener_peak <<< $DATA`
SINCE=`jq .stream_start <<< $DATA`

echo -e "listeners: $LISTENERS\tlistener_peak: $LISTENER_PEAK\tsince: $SINCE"

# Google Analytics
# Send the data to Google Analytics
curl -X POST "https://www.google-analytics.com/mp/collect?measurement_id=${MEASUREMENT_ID}&api_secret=${API_SECRET}" -d "{\"client_id\":\"raspberry_pi\",\"events\":[{\"name\":\"listeners\",\"params\":{\"Now\":\"$LISTENERS\",\"Peak\":\"$LISTENER_PEAK\",\"Since\":$SINCE}}]}"
