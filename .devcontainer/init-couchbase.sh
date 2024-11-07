#!/bin/bash

# Wait for Couchbase to be up and running
until curl -s http://couchbase:8091/pools > /dev/null; do
  echo "Waiting for Couchbase to be available..."
  sleep 1
done

CLI=/opt/couchbase/bin/couchbase-cli

# Initialize the cluster with the specified username and password
if $CLI server-list -c couchbase:8091 -u $COUCHBASE_ADMINISTRATOR_USERNAME -p $COUCHBASE_ADMINISTRATOR_PASSWORD > /dev/null; then
  echo "Cluster already initialized"
else
echo "Initializing cluster..."
$CLI cluster-init -c couchbase:8091 --cluster-username $COUCHBASE_ADMINISTRATOR_USERNAME --cluster-password $COUCHBASE_ADMINISTRATOR_PASSWORD --cluster-ramsize 512 --services data,index,query,fts
fi

sleep 5

# Create the bucket
if $CLI bucket-list -c couchbase:8091 -u $COUCHBASE_ADMINISTRATOR_USERNAME -p $COUCHBASE_ADMINISTRATOR_PASSWORD | grep $COUCHBASE_BUCKET > /dev/null; then
  echo "Bucket already created"
  exit 0
else 
echo "Creating bucket..."
$CLI bucket-create -c couchbase:8091 --username $COUCHBASE_ADMINISTRATOR_USERNAME --password $COUCHBASE_ADMINISTRATOR_PASSWORD --bucket $COUCHBASE_BUCKET --bucket-type couchbase --bucket-ramsize 256 --enable-flush 1
fi
