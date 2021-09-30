package mock

import (
	"github.com/ElrondNetwork/elrond-sdk-erdgo/core"
	"github.com/ElrondNetwork/elrond-sdk-erdgo/data"
)

// ElrondProxyStub -
type ElrondProxyStub struct {
	GetNetworkConfigCalled              func() (*data.NetworkConfig, error)
	SendTransactionCalled               func(transaction *data.Transaction) (string, error)
	GetTransactionInfoWithResultsCalled func(hash string) (*data.TransactionInfo, error)
	ExecuteVMQueryCalled                func(vmRequest *data.VmValueRequest) (*data.VmValuesResponseData, error)
	GetAccountCalled                    func(address core.AddressHandler) (*data.Account, error)
}

// GetNetworkConfig -
func (eps *ElrondProxyStub) GetNetworkConfig() (*data.NetworkConfig, error) {
	if eps.GetNetworkConfigCalled != nil {
		return eps.GetNetworkConfigCalled()
	}

	return &data.NetworkConfig{}, nil
}

// SendTransaction -
func (eps *ElrondProxyStub) SendTransaction(transaction *data.Transaction) (string, error) {
	if eps.SendTransactionCalled != nil {
		return eps.SendTransactionCalled(transaction)
	}

	return "", nil
}

// GetTransactionInfoWithResults -
func (eps *ElrondProxyStub) GetTransactionInfoWithResults(hash string) (*data.TransactionInfo, error) {
	if eps.GetTransactionInfoWithResultsCalled != nil {
		return eps.GetTransactionInfoWithResultsCalled(hash)
	}

	return &data.TransactionInfo{}, nil
}

// ExecuteVMQuery -
func (eps *ElrondProxyStub) ExecuteVMQuery(vmRequest *data.VmValueRequest) (*data.VmValuesResponseData, error) {
	if eps.ExecuteVMQueryCalled != nil {
		return eps.ExecuteVMQueryCalled(vmRequest)
	}

	return &data.VmValuesResponseData{}, nil
}

// GetAccount -
func (eps *ElrondProxyStub) GetAccount(address core.AddressHandler) (*data.Account, error) {
	if eps.GetAccountCalled != nil {
		return eps.GetAccountCalled(address)
	}

	return &data.Account{}, nil
}

// IsInterfaceNil -
func (eps *ElrondProxyStub) IsInterfaceNil() bool {
	return eps == nil
}
