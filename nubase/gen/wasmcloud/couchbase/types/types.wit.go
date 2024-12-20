// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package types represents the imported interface "wasmcloud:couchbase/types@0.1.0-draft".
//
// General types that can be used throughout the couchbase interface
package types

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
)

// DocumentID represents the string "wasmcloud:couchbase/types@0.1.0-draft#document-id".
//
// # ID of a Document
//
// Note that Document IDs are expected to be:
// - Assigned at the application level
// - Be valid UTF-8 at the application level
// - Be no longer than 250 bytes (not characters)
//
//	type document-id = string
type DocumentID string

// SubdocumentPath represents the string "wasmcloud:couchbase/types@0.1.0-draft#subdocument-path".
//
// Path to a subdocument inside an existing document
//
//	type subdocument-path = string
type SubdocumentPath string

// RequestSpan represents the string "wasmcloud:couchbase/types@0.1.0-draft#request-span".
//
// The span of a request, normally used when performing/enabling tracing
//
//	type request-span = string
type RequestSpan string

// JSONString represents the string "wasmcloud:couchbase/types@0.1.0-draft#json-string".
//
// A string that is properly formatted JSON
//
//	type json-string = string
type JSONString string

// DocumentFieldName represents the string "wasmcloud:couchbase/types@0.1.0-draft#document-field-name".
//
// Document field name (ex. to include in a search)
//
//	type document-field-name = string
type DocumentFieldName string

// CollectionName represents the string "wasmcloud:couchbase/types@0.1.0-draft#collection-name".
//
// Collection name to include in a search
//
//	type collection-name = string
type CollectionName string

// BucketName represents the string "wasmcloud:couchbase/types@0.1.0-draft#bucket-name".
//
// Bucket name
//
//	type bucket-name = string
type BucketName string

// SearchIndexName represents the string "wasmcloud:couchbase/types@0.1.0-draft#search-index-name".
//
// Name of a search index
//
//	type search-index-name = string
type SearchIndexName string

// ReplicaReadLevel represents the enum "wasmcloud:couchbase/types@0.1.0-draft#replica-read-level".
//
// Whether to enable replica reads for the request
//
//	enum replica-read-level {
//		off,
//		on
//	}
type ReplicaReadLevel uint8

const (
	ReplicaReadLevelOff ReplicaReadLevel = iota
	ReplicaReadLevelOn
)

var stringsReplicaReadLevel = [2]string{
	"off",
	"on",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e ReplicaReadLevel) String() string {
	return stringsReplicaReadLevel[e]
}

// SortDirection represents the enum "wasmcloud:couchbase/types@0.1.0-draft#sort-direction".
//
// Direction in which to perform sorting
//
//	enum sort-direction {
//		asc,
//		desc
//	}
type SortDirection uint8

const (
	SortDirectionAsc SortDirection = iota
	SortDirectionDesc
)

var stringsSortDirection = [2]string{
	"asc",
	"desc",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e SortDirection) String() string {
	return stringsSortDirection[e]
}

// Time represents the record "wasmcloud:couchbase/types@0.1.0-draft#time".
//
// Port of https://pkg.go.dev/time#Time
//
//	record time {
//		offset: s8,
//		year: s32,
//		month: u8,
//		day: u8,
//		hour: u8,
//		minute: u8,
//		second: u8,
//		milliseconds: u32,
//		nanoseconds: u32,
//	}
type Time struct {
	_ cm.HostLayout
	// Offsets are assumed to be against the western hemisphere (GMT)
	// i.e. -14 -> 0 -> +14 UTC
	Offset int8
	Year   int32
	Month  uint8
	Day    uint8

	// 24hr time (i.e. 6PM is represented as `18`)
	Hour         uint8
	Minute       uint8
	Second       uint8
	Milliseconds uint32
	Nanoseconds  uint32
}

// Collection represents the record "wasmcloud:couchbase/types@0.1.0-draft#collection".
//
// A keyspace-path that identifies a collection
//
//	record collection {
//		bucket: bucket-name,
//		scope: option<string>,
//		name: collection-name,
//	}
type Collection struct {
	_ cm.HostLayout
	// Bucket the collection belongs to
	Bucket BucketName

	// Scope of the collection (ex. "_default" if not specified)
	Scope cm.Option[string]

	// Name of the collection (which may be a simple name, if only this value is specified)
	Name CollectionName
}

// DocumentError represents the variant "wasmcloud:couchbase/types@0.1.0-draft#document-error".
//
// Errors that can occur during operations on documents
//
//	variant document-error {
//		not-found,
//		cas-mismatch,
//		locked,
//		not-locked,
//		unretrievable,
//		already-exists,
//		not-json,
//		path-not-found,
//		path-invalid,
//		path-too-deep,
//		invalid-value,
//		subdocument-delta-invalid,
//	}
type DocumentError uint8

const (
	// Document could not be found
	DocumentErrorNotFound DocumentError = iota

	// CAS revision submitted to operation does not match value on server
	DocumentErrorCasMismatch

	// Document is locked
	DocumentErrorLocked

	// Unlock con a document that is not locked
	DocumentErrorNotLocked

	// Document cannot be retrieved (from any replica)
	DocumentErrorUnretrievable

	// Document already exists (i.e. a duplicate)
	DocumentErrorAlreadyExists

	// Document is not JSON
	DocumentErrorNotJSON

	// Document does not contain the specified sub-document path
	DocumentErrorPathNotFound

	// Sub-document path is invalid (could not be parsed)
	DocumentErrorPathInvalid

	// Sub-document path is too deeply nested
	DocumentErrorPathTooDeep

	// Sub-document operation could not insert
	DocumentErrorInvalidValue

	// An invalid delta is used when performing a sub-document operation
	DocumentErrorSubdocumentDeltaInvalid
)

var stringsDocumentError = [12]string{
	"not-found",
	"cas-mismatch",
	"locked",
	"not-locked",
	"unretrievable",
	"already-exists",
	"not-json",
	"path-not-found",
	"path-invalid",
	"path-too-deep",
	"invalid-value",
	"subdocument-delta-invalid",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e DocumentError) String() string {
	return stringsDocumentError[e]
}

// DocumentValueCreateError represents the variant "wasmcloud:couchbase/types@0.1.0-draft#document-value-create-error".
//
// Errors that occur when building/using document values
//
//	variant document-value-create-error {
//		invalid-json(string),
//	}
type DocumentValueCreateError cm.Variant[uint8, string, string]

// DocumentValueCreateErrorInvalidJSON returns a [DocumentValueCreateError] of case "invalid-json".
//
// JSON used to create the document value was invalid
func DocumentValueCreateErrorInvalidJSON(data string) DocumentValueCreateError {
	return cm.New[DocumentValueCreateError](0, data)
}

// InvalidJSON returns a non-nil *[string] if [DocumentValueCreateError] represents the variant case "invalid-json".
func (self *DocumentValueCreateError) InvalidJSON() *string {
	return cm.Case[string](self, 0)
}

var stringsDocumentValueCreateError = [1]string{
	"invalid-json",
}

// String implements [fmt.Stringer], returning the variant case name of v.
func (v DocumentValueCreateError) String() string {
	return stringsDocumentValueCreateError[v.Tag()]
}

// DocumentValue represents the imported resource "wasmcloud:couchbase/types@0.1.0-draft#document-value".
//
// An implementer specific (efficient) implementation of a Document value
//
//	resource document-value
type DocumentValue cm.Resource

// ResourceDrop represents the imported resource-drop for resource "document-value".
//
// Drops a resource handle.
//
//go:nosplit
func (self DocumentValue) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_DocumentValueResourceDrop((uint32)(self0))
	return
}

// NewDocumentValue represents the imported constructor for resource "document-value".
//
// Construct an empty document value (a dictionary, by default)
//
//	constructor()
//
//go:nosplit
func NewDocumentValue() (result DocumentValue) {
	result0 := wasmimport_NewDocumentValue()
	result = cm.Reinterpret[DocumentValue]((uint32)(result0))
	return
}

// DocumentValueFromJSON represents the imported static function "from-json".
//
// Build a document-value from a stringified JSON value
//
//	from-json: static func(json: string) -> result<document-value, document-value-create-error>
//
//go:nosplit
func DocumentValueFromJSON(json string) (result cm.Result[DocumentValueCreateErrorShape, DocumentValue, DocumentValueCreateError]) {
	json0, json1 := cm.LowerString(json)
	wasmimport_DocumentValueFromJSON((*uint8)(json0), (uint32)(json1), &result)
	return
}

// ToString represents the imported method "to-string".
//
// Convert this JSON value into a string
//
//	to-string: func() -> string
//
//go:nosplit
func (self DocumentValue) ToString() (result string) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_DocumentValueToString((uint32)(self0), &result)
	return
}

// Document represents the imported variant "wasmcloud:couchbase/types@0.1.0-draft#document".
//
// WIT cannot currently support passing recursive value types, which means JSON cannot
// be properly
// represetned as a WIT level type.
//
// What we use instead is a variant with two types:
// - a valid JSON string (which may or may not be an object)
// - an implementer-controlled value which is a valid JSON object (as defined by the
// implementer)
//
// The implementer-controlled value can be any implementation (usually a more efficient
// one)
// that can be understood by the implementer. One requirement on this value is that
// it must be
// possible to convert it to/from a raw JSON string at any time.
//
// An opaque resource that represents a JSON value, and can be efficiently
// manipulated by implementers.
//
//	variant document {
//		raw(json-string),
//		%resource(document-value),
//	}
type Document cm.Variant[uint8, string, JSONString]

// DocumentRaw returns a [Document] of case "raw".
//
// A stringified JSON value represented by a resource, often used by callers
// that cannot or do not want to construct an efficient implementer specific JSON
// value
func DocumentRaw(data JSONString) Document {
	return cm.New[Document](0, data)
}

// Raw returns a non-nil *[JSONString] if [Document] represents the variant case "raw".
func (self *Document) Raw() *JSONString {
	return cm.Case[JSONString](self, 0)
}

// DocumentResource returns a [Document] of case "resource".
//
// A JSON value represented by a more efficient but opaque implementer-specific representation
func DocumentResource(data DocumentValue) Document {
	return cm.New[Document](1, data)
}

// Resource returns a non-nil *[DocumentValue] if [Document] represents the variant case "resource".
func (self *Document) Resource() *DocumentValue {
	return cm.Case[DocumentValue](self, 1)
}

var stringsDocument = [2]string{
	"raw",
	"resource",
}

// String implements [fmt.Stringer], returning the variant case name of v.
func (v Document) String() string {
	return stringsDocument[v.Tag()]
}

// BucketError represents the variant "wasmcloud:couchbase/types@0.1.0-draft#bucket-error".
//
// Errors related to bucket usage
//
//	variant bucket-error {
//		unexpected(string),
//	}
type BucketError cm.Variant[uint8, string, string]

// BucketErrorUnexpected returns a [BucketError] of case "unexpected".
//
// A completely unexpected error
func BucketErrorUnexpected(data string) BucketError {
	return cm.New[BucketError](0, data)
}

// Unexpected returns a non-nil *[string] if [BucketError] represents the variant case "unexpected".
func (self *BucketError) Unexpected() *string {
	return cm.Case[string](self, 0)
}

var stringsBucketError = [1]string{
	"unexpected",
}

// String implements [fmt.Stringer], returning the variant case name of v.
func (v BucketError) String() string {
	return stringsBucketError[v.Tag()]
}

// MutationMetadata represents the record "wasmcloud:couchbase/types@0.1.0-draft#mutation-metadata".
//
// Metadata related to any mutation on a Couchbase collection (ex. CRUD operations)
//
//	record mutation-metadata {
//		cas: u64,
//		bucket: string,
//		partition-id: u64,
//		partition-uuid: u64,
//		seq: u64,
//	}
type MutationMetadata struct {
	_ cm.HostLayout
	// CAS revision of the document
	Cas uint64

	// The bucket on which the mutation was performed
	Bucket string

	// The ID of the vbucket (partition) that the operation was performed on
	PartitionID uint64

	// The UUID of the vbucket (partition) that the operation was performed on
	PartitionUUID uint64

	// The sequence number of the operation performed on the vbucket (partition)
	Seq uint64
}

// DurabilityLevel represents the enum "wasmcloud:couchbase/types@0.1.0-draft#durability-level".
//
// Durability level that should be used for operations
//
//	enum durability-level {
//		unknown,
//		none,
//		replicate-majority,
//		replicate-majority-persist-master,
//		persist-majority
//	}
type DurabilityLevel uint8

const (
	DurabilityLevelUnknown DurabilityLevel = iota
	DurabilityLevelNone

	// Replicate (hold in memory) to a majority of nodes
	DurabilityLevelReplicateMajority

	// Replicate (hold in memory) to a majority of nodes, persist (write to disk) to master
	DurabilityLevelReplicateMajorityPersistMaster

	// Persist to a majority of nodes
	DurabilityLevelPersistMajority
)

var stringsDurabilityLevel = [5]string{
	"unknown",
	"none",
	"replicate-majority",
	"replicate-majority-persist-master",
	"persist-majority",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e DurabilityLevel) String() string {
	return stringsDurabilityLevel[e]
}

// RetryStrategy represents the variant "wasmcloud:couchbase/types@0.1.0-draft#retry-strategy".
//
// As functions cannot be represented as part of types in WIT,
// we represent static retry strategies
//
//	variant retry-strategy {
//		interval-times-ms(tuple<u64, u64>),
//	}
type RetryStrategy cm.Variant[uint8, [2]uint64, [2]uint64]

// RetryStrategyIntervalTimesMs returns a [RetryStrategy] of case "interval-times-ms".
//
// Retry a certain number of times with a given interval between each retry)
func RetryStrategyIntervalTimesMs(data [2]uint64) RetryStrategy {
	return cm.New[RetryStrategy](0, data)
}

// IntervalTimesMs returns a non-nil *[[2]uint64] if [RetryStrategy] represents the variant case "interval-times-ms".
func (self *RetryStrategy) IntervalTimesMs() *[2]uint64 {
	return cm.Case[[2]uint64](self, 0)
}

var stringsRetryStrategy = [1]string{
	"interval-times-ms",
}

// String implements [fmt.Stringer], returning the variant case name of v.
func (v RetryStrategy) String() string {
	return stringsRetryStrategy[v.Tag()]
}

// QueryScanConsistency represents the enum "wasmcloud:couchbase/types@0.1.0-draft#query-scan-consistency".
//
// Level of data consistency required for a query
//
//	enum query-scan-consistency {
//		not-bounded,
//		request-plus
//	}
type QueryScanConsistency uint8

const (
	// No data consistency required
	QueryScanConsistencyNotBounded QueryScanConsistency = iota

	// Request-level data consistency is required
	QueryScanConsistencyRequestPlus
)

var stringsQueryScanConsistency = [2]string{
	"not-bounded",
	"request-plus",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e QueryScanConsistency) String() string {
	return stringsQueryScanConsistency[e]
}

// QueryProfileMode represents the enum "wasmcloud:couchbase/types@0.1.0-draft#query-profile-mode".
//
// Profiling mode to use during a query
//
//	enum query-profile-mode {
//		none,
//		phases,
//		timings
//	}
type QueryProfileMode uint8

const (
	// Disable query profiling
	QueryProfileModeNone QueryProfileMode = iota

	// Include phase-boudned profiling information in the response
	QueryProfileModePhases

	// Include timing profiling informatoin in the query response
	QueryProfileModeTimings
)

var stringsQueryProfileMode = [3]string{
	"none",
	"phases",
	"timings",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e QueryProfileMode) String() string {
	return stringsQueryProfileMode[e]
}

// SearchScanConsistency represents the enum "wasmcloud:couchbase/types@0.1.0-draft#search-scan-consistency".
//
// Level of data consistency required for a search
//
//	enum search-scan-consistency {
//		not-bounded
//	}
type SearchScanConsistency uint8

const (
	// No data consistency required
	SearchScanConsistencyNotBounded SearchScanConsistency = iota
)

var stringsSearchScanConsistency = [1]string{
	"not-bounded",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e SearchScanConsistency) String() string {
	return stringsSearchScanConsistency[e]
}

// MutationToken represents the record "wasmcloud:couchbase/types@0.1.0-draft#mutation-token".
//
// Individual mutation token
//
//	record mutation-token {
//		bucket-name: string,
//		partition-uuid: u64,
//		partition-id: u64,
//		sequence-number: u64,
//	}
type MutationToken struct {
	_              cm.HostLayout
	BucketName     string
	PartitionUUID  uint64
	PartitionID    uint64
	SequenceNumber uint64
}

// MutationState represents the record "wasmcloud:couchbase/types@0.1.0-draft#mutation-state".
//
// Collection of mutation tokens that represent a single state
//
//	record mutation-state {
//		tokens: list<mutation-token>,
//	}
type MutationState struct {
	_      cm.HostLayout
	Tokens cm.List[MutationToken]
}
