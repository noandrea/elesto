#!/bin/bash

echo "Create a DID doc for the regulator (by the regulator account)"
elestod tx did create-did regulator \
 --from regulator \
 --chain-id elesto -y --broadcast-mode block

elestod query did did did:cosmos:net:elesto:regulator --output json | jq

echo "Create a DID doc for the EMTi (by the validator)"
elestod tx did create-did emti --from validator \
 --chain-id elesto -y --broadcast-mode block

elestod query did did did:cosmos:net:elesto:emti --output json | jq

echo "Add the EMTi account verification method to the the EMTi DID doc (by the validator account)"
elestod tx did add-verification-method emti $(elestod keys show emti -p) \
 --from validator \
 --chain-id elesto -y --broadcast-mode block

elestod query did did did:cosmos:net:elesto:emti --output json | jq

echo "Querying dids"
elestod query did dids --output json | jq

echo "Add a service to the EMTi DID doc (by the EMTi account)"
elestod tx did add-service emti emti-agent DIDComm "https://agents.elesto.app.beta.starport.cloud/emti" \
--from emti \
--chain-id elesto -y --broadcast-mode block

elestod query did did did:cosmos:net:elesto:emti --output json | jq

echo "Adding a verification relationship from decentralized did for validator"
elestod tx did set-verification-relationship emti $(elestod keys show validator -a) --relationship assertionMethod --relationship capabilityInvocation \
--from emti \
--chain-id elesto -y --broadcast-mode block

elestod query did did did:cosmos:net:elesto:emti --output json | jq

echo "Revoking verification method from decentralized did for user: validator"
elestod tx did revoke-verification-method emti $(elestod keys show validator -a) \
--from emti \
--chain-id elesto -y --broadcast-mode block

echo "Querying dids"
elestod query did did did:cosmos:net:elesto:emti --output json | jq

echo "Deleting service from EMTi did document (by EMTi user)"
elestod tx did delete-service emti emti-agent \
--from emti \
--chain-id elesto -y --broadcast-mode block


echo "Add a controller to the EMTi did document (by EMTi user)"
elestod tx did add-controller emti $(elestod keys show alice -a) \
--from emti \
--chain-id elesto -y --broadcast-mode block

echo "Querying dids"
elestod query did did did:cosmos:net:elesto:emti --output json | jq

echo "Remove a controller from the EMTi did document (by alice user)"
elestod tx did delete-controller emti $(elestod keys show alice -a) \
--from alice \
--chain-id elesto -y --broadcast-mode block

echo "Querying dids"
elestod query did did did:cosmos:net:elesto:emti --output json | jq