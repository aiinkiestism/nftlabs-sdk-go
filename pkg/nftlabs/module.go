package nftlabs

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	InterfaceIdErc165           = [4]byte{1, 255, 201, 167}  // 0x01ffc9a7
	InterfaceIdErc721           = [4]byte{128, 172, 88, 205} // 0x80ac58cd
	InterfaceIdErc721Metadata   = [4]byte{91, 94, 19, 159}   // 0x5b5e139f
	InterfaceIdErc721Enumerable = [4]byte{120, 14, 157, 99}  // 0x780e9d63
	InterfaceIdErc1155          = [4]byte{217, 182, 122, 38} // 0xd9b67a26
)

type SigningMethod = func(common.Address, *types.Transaction) (*types.Transaction, error)

// Will likely phase out of the codebase soon. Only used on unwrapped abi interfaces
type commonModule interface {
	SetPrivateKey(privateKey string) error
	getSignerAddress() common.Address
}
