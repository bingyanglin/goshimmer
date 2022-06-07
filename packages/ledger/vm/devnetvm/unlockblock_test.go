package devnetvm

import (
	"testing"

	"github.com/iotaledger/hive.go/crypto/ed25519"
	"github.com/stretchr/testify/assert"
)

func TestNewRISC0UnlockBlock(t *testing.T) {
	message := CreateTestMesage()

	// test a valid set of UnlockBlocks
	{
		unlockBlocks := UnlockBlocks{
			NewRISC0UnlockBlock(*message),
		}
		marshaledUnlockBlocks := unlockBlocks.Bytes()
		parsedUnlockBlocks, consumedBytes, err := UnlockBlocksFromBytes(marshaledUnlockBlocks)
		assert.NoError(t, err)
		assert.Equal(t, len(marshaledUnlockBlocks), consumedBytes)
		assert.Equal(t, unlockBlocks.String(), parsedUnlockBlocks.String())
		commit := unlockBlocks[0].(*RISC0UnlockBlock).Verify()
		t.Log("Commit: ", GetCommitString(commit))
	}
}

func TestUnlockBlockFromBytes(t *testing.T) {
	keyPair := ed25519.GenerateKeyPair()

	// test a valid set of UnlockBlocks
	{
		unlockBlocks := UnlockBlocks{
			NewSignatureUnlockBlock(NewED25519Signature(keyPair.PublicKey, keyPair.PrivateKey.Sign([]byte("testdata")))),
			NewReferenceUnlockBlock(0),
		}
		marshaledUnlockBlocks := unlockBlocks.Bytes()
		parsedUnlockBlocks, consumedBytes, err := UnlockBlocksFromBytes(marshaledUnlockBlocks)
		assert.NoError(t, err)
		assert.Equal(t, len(marshaledUnlockBlocks), consumedBytes)
		assert.Equal(t, unlockBlocks, parsedUnlockBlocks)
	}

	// test an invalid set of UnlockBlocks
	// TODO: this should be enabled again once duplicate validation is enabled
	// {
	// 	unlockBlocks := UnlockBlocks{
	// 		NewSignatureUnlockBlock(NewED25519Signature(keyPair.PublicKey, keyPair.PrivateKey.Sign([]byte("testdata")))),
	// 		NewSignatureUnlockBlock(NewED25519Signature(keyPair.PublicKey, keyPair.PrivateKey.Sign([]byte("testdata")))),
	// 	}
	// 	_, _, err := UnlockBlocksFromBytes(unlockBlocks.Bytes())
	// 	assert.Error(t, err)
	// }
}
