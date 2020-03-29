package tangle

import (
	"github.com/iotaledger/hive.go/marshalutil"
	"github.com/iotaledger/hive.go/objectstorage"
	"github.com/iotaledger/hive.go/stringify"

	"github.com/iotaledger/goshimmer/packages/binary/valuetransfer/payload"
	"github.com/iotaledger/goshimmer/packages/binary/valuetransfer/transaction"
)

// Attachment stores the information which transaction was attached by which payload. We need this to be able to perform
// reverse lookups from transactions to their corresponding payloads, that attach them.
type Attachment struct {
	objectstorage.StorableObjectFlags

	transactionId transaction.Id
	payloadId     payload.Id

	storageKey []byte
}

// NewAttachment creates an attachment object with the given information.
func NewAttachment(transactionId transaction.Id, payloadId payload.Id) *Attachment {
	return &Attachment{
		transactionId: transactionId,
		payloadId:     payloadId,

		storageKey: marshalutil.New(AttachmentLength).
			WriteBytes(transactionId.Bytes()).
			WriteBytes(payloadId.Bytes()).
			Bytes(),
	}
}

// AttachmentFromBytes unmarshals an Attachment from a sequence of bytes - it either creates a new object or fills the
// optionally provided one with the parsed information.
func AttachmentFromBytes(bytes []byte, optionalTargetObject ...*Attachment) (result *Attachment, err error, consumedBytes int) {
	// determine the target object that will hold the unmarshaled information
	switch len(optionalTargetObject) {
	case 0:
		result = &Attachment{}
	case 1:
		result = optionalTargetObject[0]
	default:
		panic("too many arguments in call to AttachmentFromBytes")
	}

	// parse the bytes
	marshalUtil := marshalutil.New(bytes)
	if result.transactionId, err = transaction.ParseId(marshalUtil); err != nil {
		return
	}
	if result.payloadId, err = payload.ParseId(marshalUtil); err != nil {
		return
	}
	result.storageKey = marshalutil.New(bytes[:AttachmentLength]).Bytes(true)
	consumedBytes = marshalUtil.ReadOffset()

	return
}

// Parse is a wrapper for simplified unmarshaling of Attachments from a byte stream using the marshalUtil package.
func ParseAttachment(marshalUtil *marshalutil.MarshalUtil) (*Attachment, error) {
	if attachment, err := marshalUtil.Parse(func(data []byte) (interface{}, error, int) { return AttachmentFromBytes(data) }); err != nil {
		return nil, err
	} else {
		return attachment.(*Attachment), nil
	}
}

// AttachmentFromStorageKey gets called when we restore an Attachment from the storage - it parses the key bytes and
// returns the new object.
func AttachmentFromStorageKey(keyBytes []byte) (objectstorage.StorableObject, error, int) {
	result, err, _ := AttachmentFromBytes(keyBytes)
	if err != nil {
		return nil, err, 0
	}

	return result, nil, 0
}

// TransactionId returns the transaction id of this Attachment.
func (attachment *Attachment) TransactionId() transaction.Id {
	return attachment.transactionId
}

// PayloadId returns the payload id of this Attachment.
func (attachment *Attachment) PayloadId() payload.Id {
	return attachment.payloadId
}

// Bytes marshals the Attachment into a sequence of bytes.
func (attachment *Attachment) Bytes() []byte {
	return attachment.ObjectStorageKey()
}

// String returns a human readable version of the Attachment.
func (attachment *Attachment) String() string {
	return stringify.Struct("Attachment",
		stringify.StructField("transactionId", attachment.TransactionId()),
		stringify.StructField("payloadId", attachment.PayloadId()),
	)
}

// ObjectStorageKey returns the key that is used to store the object in the database.
func (attachment *Attachment) ObjectStorageKey() []byte {
	return attachment.storageKey
}

// ObjectStorageValue marshals the "content part" of an Attachment to a sequence of bytes. Since all of the information
// for this object are stored in its key, this method does nothing and is only required to conform with the interface.
func (attachment *Attachment) ObjectStorageValue() (data []byte) {
	return
}

// UnmarshalObjectStorageValue unmarshals the "content part" of an Attachment from a sequence of bytes. Since all of the information
// for this object are stored in its key, this method does nothing and is only required to conform with the interface.
func (attachment *Attachment) UnmarshalObjectStorageValue(data []byte) (err error, consumedBytes int) {
	return
}

// Update is disabled - updates are supposed to happen through the setters (if existing).
func (attachment *Attachment) Update(other objectstorage.StorableObject) {
	panic("update forbidden")
}

// Interface contract: make compiler warn if the interface is not implemented correctly.
var _ objectstorage.StorableObject = &Attachment{}

// AttachmentLength holds the length of a marshaled Attachment in bytes.
const AttachmentLength = transaction.IdLength + payload.IdLength
