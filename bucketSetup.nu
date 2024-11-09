# Create the given bucket, scope, collection and primary index if they don't exist.
export def dbSetup [bucket: string; scope: string; collection: string] {
    cb-env | print
    if (buckets| where name == $bucket | is-empty) {buckets create --replicas 1 $bucket 100; print $"Create Bucket ($bucket)"} else {print "Bucket already exist"}
    cb-env bucket $bucket
    if (scopes| where scope == $scope | is-empty) {scopes create $scope; print $"Create Scope ($scope)"} else {print "Scope already exist"}
    cb-env scope $scope
    if (collections| where collection == $collection | is-empty) {collections create $collection; print $"Create Collection ($collection)"} else {print "Collection already exist"}
    cb-env collection $collection
    query $"CREATE PRIMARY INDEX IF NOT EXISTS ON `default`:`($bucket)`.`($scope)`.`($collection)`"
    cb-env
}
