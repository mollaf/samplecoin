#!/bin/bash

samplecli tx nameservice buy-name jack.id 5nametoken --from jack -y
samplecli tx nameservice set-name jack.id 8.8.8.8 --from jack -y
samplecli query nameservice resolve jack.id

