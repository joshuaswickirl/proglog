#!/bin/bash

# This can be turned into acceptance tests

curl -X POST localhost:8080 -d '{"record": {"value": "TGV0J3MgR28gIzEK"}}'
curl -X POST localhost:8080 -d '{"record": {"value": "TGV0J3MgR28gIzIK"}}'
curl -X POST localhost:8080 -d '{"record": {"value": "TGV0J3MgR28gIzMK"}}'

curl -X GET localhost:8080 -d '{"offset": 0}'
curl -X GET localhost:8080 -d '{"offset": 1}'
curl -X GET localhost:8080 -d '{"offset": 2}'


# {"offset":0}
# {"offset":1}
# {"offset":2}
# {"record":{"value":"TGV0J3MgR28gIzEK","offset":0}}
# {"record":{"value":"TGV0J3MgR28gIzIK","offset":1}}
# {"record":{"value":"TGV0J3MgR28gIzMK","offset":2}}
