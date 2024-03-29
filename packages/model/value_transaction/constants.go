package value_transaction

import (
	"strings"

	"github.com/iotaledger/goshimmer/packages/ternary"
)

const (
	ADDRESS_OFFSET                    = 0
	VALUE_OFFSET                      = ADDRESS_END
	TIMESTAMP_OFFSET                  = VALUE_END
	NONCE_OFFSET                      = TIMESTAMP_END
	SIGNATURE_MESSAGE_FRAGMENT_OFFSET = NONCE_END

	ADDRESS_SIZE                    = 243
	VALUE_SIZE                      = 81
	TIMESTAMP_SIZE                  = 27
	NONCE_SIZE                      = 81
	SIGNATURE_MESSAGE_FRAGMENT_SIZE = 6561
	BUNDLE_ESSENCE_SIZE             = ADDRESS_SIZE + VALUE_SIZE + SIGNATURE_MESSAGE_FRAGMENT_SIZE

	ADDRESS_END                    = ADDRESS_OFFSET + ADDRESS_SIZE
	VALUE_END                      = VALUE_OFFSET + VALUE_SIZE
	TIMESTAMP_END                  = TIMESTAMP_OFFSET + TIMESTAMP_SIZE
	NONCE_END                      = NONCE_OFFSET + NONCE_SIZE
	SIGNATURE_MESSAGE_FRAGMENT_END = SIGNATURE_MESSAGE_FRAGMENT_OFFSET + SIGNATURE_MESSAGE_FRAGMENT_SIZE

	TOTAL_SIZE = SIGNATURE_MESSAGE_FRAGMENT_END
)

var (
	EMPTY_SIGNATURE = ternary.Trytes(strings.Repeat("9", SIGNATURE_MESSAGE_FRAGMENT_SIZE/ternary.NUMBER_OF_TRITS_IN_A_TRYTE))
)
