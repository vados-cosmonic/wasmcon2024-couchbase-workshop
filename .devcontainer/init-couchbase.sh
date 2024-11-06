#!/bin/bash

# Wait for Couchbase to be up and running
until curl -s http://couchbase:8091/pools > /dev/null; do
  echo "Waiting for Couchbase to be available..."
  sleep 1
done

# Initialize the cluster with the specified username and password
if /opt/couchbase/bin/couchbase-cli server-list -c couchbase:8091 -u Administrator -p password > /dev/null; then
  echo "Cluster already initialized"
else
echo "Initializing cluster..."
/opt/couchbase/bin/couchbase-cli cluster-init -c couchbase:8091 --cluster-username Administrator --cluster-password password --cluster-ramsize 512 --services data,index,query,fts
fi

sleep 5
# Create the bucket
if /opt/couchbase/bin/couchbase-cli bucket-list -c couchbase:8091 -u Administrator -p password | grep test > /dev/null; then
  echo "Bucket already created"
  exit 0
else 
echo "Creating bucket..."
/opt/couchbase/bin/couchbase-cli bucket-create -c couchbase:8091 --username Administrator --password password --bucket test --bucket-type couchbase --bucket-ramsize 256 --enable-flush 1
fi
