// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// WalletABI is the input ABI used to generate the binding from.
const WalletABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"pendingTopupLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cancelWhitelistAddition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submittedWhitelistAddition\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingWhitelistRemoval\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submittedTopupLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"confirmWhitelistAddition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingWhitelistAddition\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"confirmWhitelistRemoval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cancelWhitelistRemoval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cancelSpendLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initializedSpendLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"initializeSpendLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"submitWhitelistRemoval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"topupAvailable\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"submitWhitelistAddition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"confirmTopupLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"initializeTopupLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cancelTopupLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"confirmSpendLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addController\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spendLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initializedWhitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_asset\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingSpendLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initializedTopupLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"submitTopupLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submittedSpendLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"topupLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"submitSpendLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spendAvailable\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submittedWhitelistRemoval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"topupGas\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"initializeWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeController\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_oracle\",\"type\":\"address\"},{\"name\":\"_controllers\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SetTopupLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SubmitTopupLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"CancelTopupLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TopupGas\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SetSpendLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SubmitSpendLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"CancelSpendLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"WhitelistAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"SubmitWhitelistAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"CancelWhitelistAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"WhitelistRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"SubmitWhitelistRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"CancelWhitelistRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"AddController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"RemoveController\",\"type\":\"event\"}]"

// WalletBin is the compiled bytecode used for deploying new contracts.
const WalletBin = `0x60806040523480156200001157600080fd5b506040516200328538038062003285833981018060405281019080805190602001909291908051906020019092919080518201929190505050828282600068056bc75e2d631000006007819055504260088190555060075460098190555083600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600b60026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600090505b815181101562000140576200013282828151811015156200011357fe5b9060200190602002015162000165640100000000026401000000009004565b8080600101915050620000f6565b505050506706f05b59d3b20000600c81905550600c54600e81905550505050620002c1565b6000808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151515620001be57600080fd5b60016000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600081548092919060010191905055507fec493e5f7bb2653b285a1e0af9af1c883375111370296bb332c8c5e3077689593382604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a150565b612fb480620002d16000396000f3006080604052600436106101ee576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630978f5a01461026557806315b9a8b81461029057806322401bde146102bb57806326d05ab2146102d2578063294f4025146103015780633671bb0d1461036d5780633af32abf1461039c5780633b98fe84146103f757806347b55a9d1461040e5780634b7304081461047a5780635204110c1461049157806354f599d7146104a85780635658eff0146104bf57806358453569146104ee5780636137d6701461051b5780636fc95b89146105565780637dc0d1d0146105815780637fd004fa146105d85780638da5cb5b146106135780638e112cf91461066a5780638ff3bf331461068157806395eccddb146106ae57806396136585146106c5578063a7fc7a07146106dc578063aceaf92d1461071f578063afa0fd9b1461074a578063b429afeb14610779578063beabacc8146107d4578063c8ecaddb14610841578063cd96895c1461086c578063d3a605861461089b578063d5666590146108c8578063d87d2a39146108f7578063d9ec301814610922578063dae37fac1461094f578063de212bf31461097a578063e3d670d7146109a9578063eca7179214610a00578063f4199bb814610a2d578063f6a74ed714610a68575b6000341115610263577fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c3334604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a15b005b34801561027157600080fd5b5061027a610aab565b6040518082815260200191505060405180910390f35b34801561029c57600080fd5b506102a5610ab1565b6040518082815260200191505060405180910390f35b3480156102c757600080fd5b506102d0610ab7565b005b3480156102de57600080fd5b506102e7610b9c565b604051808215151515815260200191505060405180910390f35b34801561030d57600080fd5b50610316610baf565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561035957808201518184015260208101905061033e565b505050509050019250505060405180910390f35b34801561037957600080fd5b50610382610c3d565b604051808215151515815260200191505060405180910390f35b3480156103a857600080fd5b506103dd600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c50565b604051808215151515815260200191505060405180910390f35b34801561040357600080fd5b5061040c610c70565b005b34801561041a57600080fd5b50610423610eb1565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561046657808201518184015260208101905061044b565b505050509050019250505060405180910390f35b34801561048657600080fd5b5061048f610f3f565b005b34801561049d57600080fd5b506104a6611180565b005b3480156104b457600080fd5b506104bd611265565b005b3480156104cb57600080fd5b506104d4611344565b604051808215151515815260200191505060405180910390f35b3480156104fa57600080fd5b5061051960048036038101908080359060200190929190505050611357565b005b34801561052757600080fd5b50610554600480360381019080803590602001908201803590602001919091929391929390505050611461565b005b34801561056257600080fd5b5061056b6115c7565b6040518082815260200191505060405180910390f35b34801561058d57600080fd5b506105966115eb565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156105e457600080fd5b50610611600480360381019080803590602001908201803590602001919091929391929390505050611611565b005b34801561061f57600080fd5b5061062861186b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561067657600080fd5b5061067f611891565b005b34801561068d57600080fd5b506106ac600480360381019080803590602001909291905050506119cb565b005b3480156106ba57600080fd5b506106c3611aff565b005b3480156106d157600080fd5b506106da611bde565b005b3480156106e857600080fd5b5061071d600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611ced565b005b34801561072b57600080fd5b50610734611d50565b6040518082815260200191505060405180910390f35b34801561075657600080fd5b5061075f611d56565b604051808215151515815260200191505060405180910390f35b34801561078557600080fd5b506107ba600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611d69565b604051808215151515815260200191505060405180910390f35b3480156107e057600080fd5b5061083f600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611d89565b005b34801561084d57600080fd5b5061085661219c565b6040518082815260200191505060405180910390f35b34801561087857600080fd5b506108816121a2565b604051808215151515815260200191505060405180910390f35b3480156108a757600080fd5b506108c6600480360381019080803590602001909291905050506121b5565b005b3480156108d457600080fd5b506108dd6122e5565b604051808215151515815260200191505060405180910390f35b34801561090357600080fd5b5061090c6122f8565b6040518082815260200191505060405180910390f35b34801561092e57600080fd5b5061094d600480360381019080803590602001909291905050506122fe565b005b34801561095b57600080fd5b50610964612404565b6040518082815260200191505060405180910390f35b34801561098657600080fd5b5061098f612428565b604051808215151515815260200191505060405180910390f35b3480156109b557600080fd5b506109ea600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061243b565b6040518082815260200191505060405180910390f35b348015610a0c57600080fd5b50610a2b60048036038101908080359060200190929190505050612559565b005b348015610a3957600080fd5b50610a66600480360381019080803590602001908201803590602001919091929391929390505050612786565b005b348015610a7457600080fd5b50610aa9600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612a59565b005b600f5481565b60015481565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515610b0e57600080fd5b60046000610b1c9190612e5f565b6000600660006101000a81548160ff0219169083151502179055507f558fc446b058402d71bbea9bab10c2c469b90898aeac942895a7251fc2b6847833604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1565b600660009054906101000a900460ff1681565b60606005805480602002602001604051908101604052809291908181526020018280548015610c3357602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610be9575b5050505050905090565b601060009054906101000a900460ff1681565b60036020528060005260406000206000915054906101000a900460ff1681565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515610cc957600080fd5b6000600480549050118015610cea5750600660009054906101000a900460ff165b1515610cf557600080fd5b600090505b600480549050811015610da557600160036000600484815481101515610d1c57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508080600101915050610cfa565b7ff0ff099f6211e3a70fce6ca4dd04c29600335b95f2d56b4e212d5ac69d6a0cfd336004604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001806020018281038252838181548152602001915080548015610e7657602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610e2c575b5050935050505060405180910390a160046000610e939190612e5f565b6000600660006101000a81548160ff02191690831515021790555050565b60606004805480602002602001604051908101604052809291908181526020018280548015610f3557602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610eeb575b5050505050905090565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515610f9857600080fd5b6000600580549050118015610fb95750600660019054906101000a900460ff165b1515610fc457600080fd5b600090505b60058054905081101561107457600060036000600584815481101515610feb57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508080600101915050610fc9565b7fce26ffefe745221a0fc931cb9394346a68c154ba02bc4c5e760bfe29a533f094336005604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818154815260200191508054801561114557602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116110fb575b5050935050505060405180910390a1600560006111629190612e5f565b6000600660016101000a81548160ff02191690831515021790555050565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156111d757600080fd5b600560006111e59190612e5f565b6000600660016101000a81548160ff0219169083151502179055507f8978ec16114c856a3a04b746619c1d83d552fbfd4c115c0b3cbe066a9a56e01d33604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156112bc57600080fd5b6000600a819055506000600b60006101000a81548160ff0219169083151502179055507f04df2cb40e86637ce3f74891f08d714137e4f2767a45e0495b9eed18578fdfa033604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1565b600b60019054906101000a900460ff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156113b357600080fd5b600b60019054906101000a900460ff161515156113cf57600080fd5b6113d881612abc565b6001600b60016101000a81548160ff0219169083151502179055507f068f112e5ec923d412be64779fe69e0fcbb6784c6617e94cccc8fd348f2e0f213382604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a150565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156114bd57600080fd5b818180806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050600181511015801561150157506014815111155b151561150c57600080fd5b600660019054906101000a900460ff161580156115365750600660009054906101000a900460ff16155b151561154157600080fd5b828260059190611552929190612e80565b506001600660016101000a81548160ff0219169083151502179055507f88c78a1f512216b3262ba17462fd5cfdffd48bc03c449e30aa77f71e4b50c1e68383604051808060200182810382528484828181526020019250602002808284378201915050935050505060405180910390a1505050565b600062015180600d54014211156115e257600c5490506115e8565b600e5490505b90565b600b60029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561166d57600080fd5b81818080602002602001604051908101604052809392919081815260200183836020028082843782019150505050505060018151101580156116b157506014815111155b15156116bc57600080fd5b82828080602002602001604051908101604052809392919081815260200183836020028082843782019150505050505060008090505b815181101561177c57600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16828281518110151561174257fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff161415151561176f57600080fd5b80806001019150506116f2565b600660009054906101000a900460ff161580156117a65750600660019054906101000a900460ff16155b15156117b157600080fd5b8484600491906117c2929190612e80565b506001600660006101000a81548160ff021916908315150217905550600660029054906101000a900460ff161515611810576001600660026101000a81548160ff0219169083151502179055505b7f3e8b4ac7dd097b817f72dee5feb5c43e5316cf4f03a6f8fb6936376b8962703a8585604051808060200182810382528484828181526020019250602002808284378201915050935050505060405180910390a15050505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156118e857600080fd5b601060009054906101000a900460ff16151561190357600080fd5b600f5466038d7ea4c680001115801561192657506706f05b59d3b20000600f5411155b151561192e57fe5b611939600f54612ae4565b7f542c9d9a32f2597d63c74984374742431933d68fe2fd019ee3592121b5e53e8133600f54604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a16000600f819055506000601060006101000a81548160ff021916908315150217905550565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611a2757600080fd5b601060019054906101000a900460ff16151515611a4357600080fd5b8066038d7ea4c6800011158015611a6257506706f05b59d3b200008111155b1515611a6d57600080fd5b611a7681612ae4565b6001601060016101000a81548160ff0219169083151502179055507f542c9d9a32f2597d63c74984374742431933d68fe2fd019ee3592121b5e53e813382604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a150565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515611b5657600080fd5b6000600f819055506000601060006101000a81548160ff0219169083151502179055507f3dfa16a690ee101e8342b8200b911ee7c0cb8e7e6cc596930a4e23a03a53941c33604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515611c3557600080fd5b600b60009054906101000a900460ff161515611c5057600080fd5b611c5b600a54612abc565b7f068f112e5ec923d412be64779fe69e0fcbb6784c6617e94cccc8fd348f2e0f2133600a54604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a16000600b60006101000a81548160ff0219169083151502179055506000600a81905550565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515611d4457600080fd5b611d4d81612b0c565b50565b60075481565b600660029054906101000a900460ff1681565b60006020528060005260406000206000915054906101000a900460ff1681565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611de757600080fd5b8160008114151515611df857600080fd5b600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515611fa057611e53612c67565b60008473ffffffffffffffffffffffffffffffffffffffff16141515611f7a57600b60029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166367c6e39c85856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015611f3857600080fd5b505af1158015611f4c573d6000803e3d6000fd5b505050506040513d6020811015611f6257600080fd5b81019080805190602001909291905050509150611f7e565b8291505b6009548211151515611f8f57600080fd5b816009600082825403925050819055505b60008473ffffffffffffffffffffffffffffffffffffffff161415156120ae578373ffffffffffffffffffffffffffffffffffffffff1663a9059cbb86856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561206357600080fd5b505af1158015612077573d6000803e3d6000fd5b505050506040513d602081101561208d57600080fd5b810190808051906020019092919050505015156120a957600080fd5b6120f6565b8473ffffffffffffffffffffffffffffffffffffffff166108fc849081150290604051600060405180830381858888f193505050501580156120f4573d6000803e3d6000fd5b505b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef858585604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a15050505050565b600a5481565b601060019054906101000a900460ff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561221157600080fd5b601060009054906101000a900460ff1615151561222d57600080fd5b8066038d7ea4c680001115801561224c57506706f05b59d3b200008111155b151561225757600080fd5b80600f819055506001601060006101000a81548160ff021916908315150217905550601060019054906101000a900460ff1615156122ab576001601060016101000a81548160ff0219169083151502179055505b7ffb68cbcb26831437ffeaeec06583fba34b716958f3363288069f862a909d9112816040518082815260200191505060405180910390a150565b600b60009054906101000a900460ff1681565b600c5481565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561235a57600080fd5b600b60009054906101000a900460ff1615151561237657600080fd5b80600a819055506001600b60006101000a81548160ff021916908315150217905550600b60019054906101000a900460ff1615156123ca576001600b60016101000a81548160ff0219169083151502179055505b7ff56c67a9b328e0778c4806a71b89e6e351bc301249815bd53d4110a5bc59a133816040518082815260200191505060405180910390a150565b6000620151806008540142111561241f576007549050612425565b60095490505b90565b600660019054906101000a900460ff1681565b6000808273ffffffffffffffffffffffffffffffffffffffff16141515612539578173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b1580156124f757600080fd5b505af115801561250b573d6000803e3d6000fd5b505050506040513d602081101561252157600080fd5b81019080805190602001909291905050509050612554565b3073ffffffffffffffffffffffffffffffffffffffff163190505b919050565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16806125ff5750600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b151561260a57600080fd5b816000811415151561261b57600080fd5b612623612caf565b6000600e5411151561263457600080fd5b829150600e5482111561264757600e5491505b81600e60008282540392505081905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f193505050501580156126bf573d6000803e3d6000fd5b507f11bb310b94280c15845698b8ce945817e14456a5d1582e387e6e4a01ef2c674232600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a1505050565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156127e457600080fd5b828280806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050600181511015801561282857506014815111155b151561283357600080fd5b83838080602002602001604051908101604052809392919081815260200183836020028082843782019150505050505060008090505b81518110156128f357600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1682828151811015156128b957fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff16141515156128e657600080fd5b8080600101915050612869565b600660029054906101000a900460ff1615151561290f57600080fd5b600093505b858590508410156129ae57600160036000888888818110151561293357fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508380600101945050612914565b6001600660026101000a81548160ff0219169083151502179055507ff0ff099f6211e3a70fce6ca4dd04c29600335b95f2d56b4e212d5ac69d6a0cfd338787604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001806020018281038252848482818152602001925060200280828437820191505094505050505060405180910390a1505050505050565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515612ab057600080fd5b612ab981612cf7565b50565b612ac4612c67565b806007819055506007546009541115612ae1576007546009819055505b50565b612aec612caf565b80600c81905550600c54600e541115612b0957600c54600e819055505b50565b6000808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151515612b6457600080fd5b60016000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600081548092919060010191905055507fec493e5f7bb2653b285a1e0af9af1c883375111370296bb332c8c5e3077689593382604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a150565b60006201518060085401421115612cac57620151806008544203811515612c8a57fe5b0490506201518081026008600082825401925050819055506007546009819055505b50565b600062015180600d5401421115612cf45762015180600d544203811515612cd257fe5b049050620151808102600d60008282540192505081905550600c54600e819055505b50565b6000808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff168015612d50575060018054115b1515612d5b57600080fd5b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600160008154809291906001900391905055507f98da1b1dd7d69af3ffee8826b8a31d3e98874a91a2e90e819fb6df0cfa91ca4d3382604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a150565b5080546000825590600052602060002090810190612e7d9190612f20565b50565b828054828255906000526020600020908101928215612f0f579160200282015b82811115612f0e57823573ffffffffffffffffffffffffffffffffffffffff168260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190612ea0565b5b509050612f1c9190612f45565b5090565b612f4291905b80821115612f3e576000816000905550600101612f26565b5090565b90565b612f8591905b80821115612f8157600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101612f4b565b5090565b905600a165627a7a72305820d3f14beeb90e989e717b360f398bc27f57fe04f0728341f170dcd59dc4c6f41a0029`

// DeployWallet deploys a new Ethereum contract, binding an instance of Wallet to it.
func DeployWallet(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _oracle common.Address, _controllers []common.Address) (common.Address, *types.Transaction, *Wallet, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WalletBin), backend, _owner, _oracle, _controllers)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Wallet{WalletCaller: WalletCaller{contract: contract}, WalletTransactor: WalletTransactor{contract: contract}, WalletFilterer: WalletFilterer{contract: contract}}, nil
}

// Wallet is an auto generated Go binding around an Ethereum contract.
type Wallet struct {
	WalletCaller     // Read-only binding to the contract
	WalletTransactor // Write-only binding to the contract
	WalletFilterer   // Log filterer for contract events
}

// WalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletSession struct {
	Contract     *Wallet           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletCallerSession struct {
	Contract *WalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletTransactorSession struct {
	Contract     *WalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletRaw struct {
	Contract *Wallet // Generic contract binding to access the raw methods on
}

// WalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletCallerRaw struct {
	Contract *WalletCaller // Generic read-only contract binding to access the raw methods on
}

// WalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletTransactorRaw struct {
	Contract *WalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWallet creates a new instance of Wallet, bound to a specific deployed contract.
func NewWallet(address common.Address, backend bind.ContractBackend) (*Wallet, error) {
	contract, err := bindWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wallet{WalletCaller: WalletCaller{contract: contract}, WalletTransactor: WalletTransactor{contract: contract}, WalletFilterer: WalletFilterer{contract: contract}}, nil
}

// NewWalletCaller creates a new read-only instance of Wallet, bound to a specific deployed contract.
func NewWalletCaller(address common.Address, caller bind.ContractCaller) (*WalletCaller, error) {
	contract, err := bindWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletCaller{contract: contract}, nil
}

// NewWalletTransactor creates a new write-only instance of Wallet, bound to a specific deployed contract.
func NewWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletTransactor, error) {
	contract, err := bindWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletTransactor{contract: contract}, nil
}

// NewWalletFilterer creates a new log filterer instance of Wallet, bound to a specific deployed contract.
func NewWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletFilterer, error) {
	contract, err := bindWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletFilterer{contract: contract}, nil
}

// bindWallet binds a generic wrapper to an already deployed contract.
func bindWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wallet *WalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Wallet.Contract.WalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wallet *WalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.Contract.WalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wallet *WalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wallet.Contract.WalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wallet *WalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Wallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wallet *WalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wallet *WalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wallet.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(_asset address) constant returns(uint256)
func (_Wallet *WalletCaller) Balance(opts *bind.CallOpts, _asset common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "balance", _asset)
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(_asset address) constant returns(uint256)
func (_Wallet *WalletSession) Balance(_asset common.Address) (*big.Int, error) {
	return _Wallet.Contract.Balance(&_Wallet.CallOpts, _asset)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(_asset address) constant returns(uint256)
func (_Wallet *WalletCallerSession) Balance(_asset common.Address) (*big.Int, error) {
	return _Wallet.Contract.Balance(&_Wallet.CallOpts, _asset)
}

// ControllerCount is a free data retrieval call binding the contract method 0x15b9a8b8.
//
// Solidity: function controllerCount() constant returns(uint256)
func (_Wallet *WalletCaller) ControllerCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "controllerCount")
	return *ret0, err
}

// ControllerCount is a free data retrieval call binding the contract method 0x15b9a8b8.
//
// Solidity: function controllerCount() constant returns(uint256)
func (_Wallet *WalletSession) ControllerCount() (*big.Int, error) {
	return _Wallet.Contract.ControllerCount(&_Wallet.CallOpts)
}

// ControllerCount is a free data retrieval call binding the contract method 0x15b9a8b8.
//
// Solidity: function controllerCount() constant returns(uint256)
func (_Wallet *WalletCallerSession) ControllerCount() (*big.Int, error) {
	return _Wallet.Contract.ControllerCount(&_Wallet.CallOpts)
}

// InitializedSpendLimit is a free data retrieval call binding the contract method 0x5658eff0.
//
// Solidity: function initializedSpendLimit() constant returns(bool)
func (_Wallet *WalletCaller) InitializedSpendLimit(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "initializedSpendLimit")
	return *ret0, err
}

// InitializedSpendLimit is a free data retrieval call binding the contract method 0x5658eff0.
//
// Solidity: function initializedSpendLimit() constant returns(bool)
func (_Wallet *WalletSession) InitializedSpendLimit() (bool, error) {
	return _Wallet.Contract.InitializedSpendLimit(&_Wallet.CallOpts)
}

// InitializedSpendLimit is a free data retrieval call binding the contract method 0x5658eff0.
//
// Solidity: function initializedSpendLimit() constant returns(bool)
func (_Wallet *WalletCallerSession) InitializedSpendLimit() (bool, error) {
	return _Wallet.Contract.InitializedSpendLimit(&_Wallet.CallOpts)
}

// InitializedTopupLimit is a free data retrieval call binding the contract method 0xcd96895c.
//
// Solidity: function initializedTopupLimit() constant returns(bool)
func (_Wallet *WalletCaller) InitializedTopupLimit(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "initializedTopupLimit")
	return *ret0, err
}

// InitializedTopupLimit is a free data retrieval call binding the contract method 0xcd96895c.
//
// Solidity: function initializedTopupLimit() constant returns(bool)
func (_Wallet *WalletSession) InitializedTopupLimit() (bool, error) {
	return _Wallet.Contract.InitializedTopupLimit(&_Wallet.CallOpts)
}

// InitializedTopupLimit is a free data retrieval call binding the contract method 0xcd96895c.
//
// Solidity: function initializedTopupLimit() constant returns(bool)
func (_Wallet *WalletCallerSession) InitializedTopupLimit() (bool, error) {
	return _Wallet.Contract.InitializedTopupLimit(&_Wallet.CallOpts)
}

// InitializedWhitelist is a free data retrieval call binding the contract method 0xafa0fd9b.
//
// Solidity: function initializedWhitelist() constant returns(bool)
func (_Wallet *WalletCaller) InitializedWhitelist(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "initializedWhitelist")
	return *ret0, err
}

// InitializedWhitelist is a free data retrieval call binding the contract method 0xafa0fd9b.
//
// Solidity: function initializedWhitelist() constant returns(bool)
func (_Wallet *WalletSession) InitializedWhitelist() (bool, error) {
	return _Wallet.Contract.InitializedWhitelist(&_Wallet.CallOpts)
}

// InitializedWhitelist is a free data retrieval call binding the contract method 0xafa0fd9b.
//
// Solidity: function initializedWhitelist() constant returns(bool)
func (_Wallet *WalletCallerSession) InitializedWhitelist() (bool, error) {
	return _Wallet.Contract.InitializedWhitelist(&_Wallet.CallOpts)
}

// IsController is a free data retrieval call binding the contract method 0xb429afeb.
//
// Solidity: function isController( address) constant returns(bool)
func (_Wallet *WalletCaller) IsController(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "isController", arg0)
	return *ret0, err
}

// IsController is a free data retrieval call binding the contract method 0xb429afeb.
//
// Solidity: function isController( address) constant returns(bool)
func (_Wallet *WalletSession) IsController(arg0 common.Address) (bool, error) {
	return _Wallet.Contract.IsController(&_Wallet.CallOpts, arg0)
}

// IsController is a free data retrieval call binding the contract method 0xb429afeb.
//
// Solidity: function isController( address) constant returns(bool)
func (_Wallet *WalletCallerSession) IsController(arg0 common.Address) (bool, error) {
	return _Wallet.Contract.IsController(&_Wallet.CallOpts, arg0)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted( address) constant returns(bool)
func (_Wallet *WalletCaller) IsWhitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "isWhitelisted", arg0)
	return *ret0, err
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted( address) constant returns(bool)
func (_Wallet *WalletSession) IsWhitelisted(arg0 common.Address) (bool, error) {
	return _Wallet.Contract.IsWhitelisted(&_Wallet.CallOpts, arg0)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted( address) constant returns(bool)
func (_Wallet *WalletCallerSession) IsWhitelisted(arg0 common.Address) (bool, error) {
	return _Wallet.Contract.IsWhitelisted(&_Wallet.CallOpts, arg0)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_Wallet *WalletCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "oracle")
	return *ret0, err
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_Wallet *WalletSession) Oracle() (common.Address, error) {
	return _Wallet.Contract.Oracle(&_Wallet.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_Wallet *WalletCallerSession) Oracle() (common.Address, error) {
	return _Wallet.Contract.Oracle(&_Wallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Wallet *WalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Wallet *WalletSession) Owner() (common.Address, error) {
	return _Wallet.Contract.Owner(&_Wallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Wallet *WalletCallerSession) Owner() (common.Address, error) {
	return _Wallet.Contract.Owner(&_Wallet.CallOpts)
}

// PendingSpendLimit is a free data retrieval call binding the contract method 0xc8ecaddb.
//
// Solidity: function pendingSpendLimit() constant returns(uint256)
func (_Wallet *WalletCaller) PendingSpendLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "pendingSpendLimit")
	return *ret0, err
}

// PendingSpendLimit is a free data retrieval call binding the contract method 0xc8ecaddb.
//
// Solidity: function pendingSpendLimit() constant returns(uint256)
func (_Wallet *WalletSession) PendingSpendLimit() (*big.Int, error) {
	return _Wallet.Contract.PendingSpendLimit(&_Wallet.CallOpts)
}

// PendingSpendLimit is a free data retrieval call binding the contract method 0xc8ecaddb.
//
// Solidity: function pendingSpendLimit() constant returns(uint256)
func (_Wallet *WalletCallerSession) PendingSpendLimit() (*big.Int, error) {
	return _Wallet.Contract.PendingSpendLimit(&_Wallet.CallOpts)
}

// PendingTopupLimit is a free data retrieval call binding the contract method 0x0978f5a0.
//
// Solidity: function pendingTopupLimit() constant returns(uint256)
func (_Wallet *WalletCaller) PendingTopupLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "pendingTopupLimit")
	return *ret0, err
}

// PendingTopupLimit is a free data retrieval call binding the contract method 0x0978f5a0.
//
// Solidity: function pendingTopupLimit() constant returns(uint256)
func (_Wallet *WalletSession) PendingTopupLimit() (*big.Int, error) {
	return _Wallet.Contract.PendingTopupLimit(&_Wallet.CallOpts)
}

// PendingTopupLimit is a free data retrieval call binding the contract method 0x0978f5a0.
//
// Solidity: function pendingTopupLimit() constant returns(uint256)
func (_Wallet *WalletCallerSession) PendingTopupLimit() (*big.Int, error) {
	return _Wallet.Contract.PendingTopupLimit(&_Wallet.CallOpts)
}

// PendingWhitelistAddition is a free data retrieval call binding the contract method 0x47b55a9d.
//
// Solidity: function pendingWhitelistAddition() constant returns(address[])
func (_Wallet *WalletCaller) PendingWhitelistAddition(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "pendingWhitelistAddition")
	return *ret0, err
}

// PendingWhitelistAddition is a free data retrieval call binding the contract method 0x47b55a9d.
//
// Solidity: function pendingWhitelistAddition() constant returns(address[])
func (_Wallet *WalletSession) PendingWhitelistAddition() ([]common.Address, error) {
	return _Wallet.Contract.PendingWhitelistAddition(&_Wallet.CallOpts)
}

// PendingWhitelistAddition is a free data retrieval call binding the contract method 0x47b55a9d.
//
// Solidity: function pendingWhitelistAddition() constant returns(address[])
func (_Wallet *WalletCallerSession) PendingWhitelistAddition() ([]common.Address, error) {
	return _Wallet.Contract.PendingWhitelistAddition(&_Wallet.CallOpts)
}

// PendingWhitelistRemoval is a free data retrieval call binding the contract method 0x294f4025.
//
// Solidity: function pendingWhitelistRemoval() constant returns(address[])
func (_Wallet *WalletCaller) PendingWhitelistRemoval(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "pendingWhitelistRemoval")
	return *ret0, err
}

// PendingWhitelistRemoval is a free data retrieval call binding the contract method 0x294f4025.
//
// Solidity: function pendingWhitelistRemoval() constant returns(address[])
func (_Wallet *WalletSession) PendingWhitelistRemoval() ([]common.Address, error) {
	return _Wallet.Contract.PendingWhitelistRemoval(&_Wallet.CallOpts)
}

// PendingWhitelistRemoval is a free data retrieval call binding the contract method 0x294f4025.
//
// Solidity: function pendingWhitelistRemoval() constant returns(address[])
func (_Wallet *WalletCallerSession) PendingWhitelistRemoval() ([]common.Address, error) {
	return _Wallet.Contract.PendingWhitelistRemoval(&_Wallet.CallOpts)
}

// SpendAvailable is a free data retrieval call binding the contract method 0xdae37fac.
//
// Solidity: function spendAvailable() constant returns(uint256)
func (_Wallet *WalletCaller) SpendAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "spendAvailable")
	return *ret0, err
}

// SpendAvailable is a free data retrieval call binding the contract method 0xdae37fac.
//
// Solidity: function spendAvailable() constant returns(uint256)
func (_Wallet *WalletSession) SpendAvailable() (*big.Int, error) {
	return _Wallet.Contract.SpendAvailable(&_Wallet.CallOpts)
}

// SpendAvailable is a free data retrieval call binding the contract method 0xdae37fac.
//
// Solidity: function spendAvailable() constant returns(uint256)
func (_Wallet *WalletCallerSession) SpendAvailable() (*big.Int, error) {
	return _Wallet.Contract.SpendAvailable(&_Wallet.CallOpts)
}

// SpendLimit is a free data retrieval call binding the contract method 0xaceaf92d.
//
// Solidity: function spendLimit() constant returns(uint256)
func (_Wallet *WalletCaller) SpendLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "spendLimit")
	return *ret0, err
}

// SpendLimit is a free data retrieval call binding the contract method 0xaceaf92d.
//
// Solidity: function spendLimit() constant returns(uint256)
func (_Wallet *WalletSession) SpendLimit() (*big.Int, error) {
	return _Wallet.Contract.SpendLimit(&_Wallet.CallOpts)
}

// SpendLimit is a free data retrieval call binding the contract method 0xaceaf92d.
//
// Solidity: function spendLimit() constant returns(uint256)
func (_Wallet *WalletCallerSession) SpendLimit() (*big.Int, error) {
	return _Wallet.Contract.SpendLimit(&_Wallet.CallOpts)
}

// SubmittedSpendLimit is a free data retrieval call binding the contract method 0xd5666590.
//
// Solidity: function submittedSpendLimit() constant returns(bool)
func (_Wallet *WalletCaller) SubmittedSpendLimit(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "submittedSpendLimit")
	return *ret0, err
}

// SubmittedSpendLimit is a free data retrieval call binding the contract method 0xd5666590.
//
// Solidity: function submittedSpendLimit() constant returns(bool)
func (_Wallet *WalletSession) SubmittedSpendLimit() (bool, error) {
	return _Wallet.Contract.SubmittedSpendLimit(&_Wallet.CallOpts)
}

// SubmittedSpendLimit is a free data retrieval call binding the contract method 0xd5666590.
//
// Solidity: function submittedSpendLimit() constant returns(bool)
func (_Wallet *WalletCallerSession) SubmittedSpendLimit() (bool, error) {
	return _Wallet.Contract.SubmittedSpendLimit(&_Wallet.CallOpts)
}

// SubmittedTopupLimit is a free data retrieval call binding the contract method 0x3671bb0d.
//
// Solidity: function submittedTopupLimit() constant returns(bool)
func (_Wallet *WalletCaller) SubmittedTopupLimit(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "submittedTopupLimit")
	return *ret0, err
}

// SubmittedTopupLimit is a free data retrieval call binding the contract method 0x3671bb0d.
//
// Solidity: function submittedTopupLimit() constant returns(bool)
func (_Wallet *WalletSession) SubmittedTopupLimit() (bool, error) {
	return _Wallet.Contract.SubmittedTopupLimit(&_Wallet.CallOpts)
}

// SubmittedTopupLimit is a free data retrieval call binding the contract method 0x3671bb0d.
//
// Solidity: function submittedTopupLimit() constant returns(bool)
func (_Wallet *WalletCallerSession) SubmittedTopupLimit() (bool, error) {
	return _Wallet.Contract.SubmittedTopupLimit(&_Wallet.CallOpts)
}

// SubmittedWhitelistAddition is a free data retrieval call binding the contract method 0x26d05ab2.
//
// Solidity: function submittedWhitelistAddition() constant returns(bool)
func (_Wallet *WalletCaller) SubmittedWhitelistAddition(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "submittedWhitelistAddition")
	return *ret0, err
}

// SubmittedWhitelistAddition is a free data retrieval call binding the contract method 0x26d05ab2.
//
// Solidity: function submittedWhitelistAddition() constant returns(bool)
func (_Wallet *WalletSession) SubmittedWhitelistAddition() (bool, error) {
	return _Wallet.Contract.SubmittedWhitelistAddition(&_Wallet.CallOpts)
}

// SubmittedWhitelistAddition is a free data retrieval call binding the contract method 0x26d05ab2.
//
// Solidity: function submittedWhitelistAddition() constant returns(bool)
func (_Wallet *WalletCallerSession) SubmittedWhitelistAddition() (bool, error) {
	return _Wallet.Contract.SubmittedWhitelistAddition(&_Wallet.CallOpts)
}

// SubmittedWhitelistRemoval is a free data retrieval call binding the contract method 0xde212bf3.
//
// Solidity: function submittedWhitelistRemoval() constant returns(bool)
func (_Wallet *WalletCaller) SubmittedWhitelistRemoval(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "submittedWhitelistRemoval")
	return *ret0, err
}

// SubmittedWhitelistRemoval is a free data retrieval call binding the contract method 0xde212bf3.
//
// Solidity: function submittedWhitelistRemoval() constant returns(bool)
func (_Wallet *WalletSession) SubmittedWhitelistRemoval() (bool, error) {
	return _Wallet.Contract.SubmittedWhitelistRemoval(&_Wallet.CallOpts)
}

// SubmittedWhitelistRemoval is a free data retrieval call binding the contract method 0xde212bf3.
//
// Solidity: function submittedWhitelistRemoval() constant returns(bool)
func (_Wallet *WalletCallerSession) SubmittedWhitelistRemoval() (bool, error) {
	return _Wallet.Contract.SubmittedWhitelistRemoval(&_Wallet.CallOpts)
}

// TopupAvailable is a free data retrieval call binding the contract method 0x6fc95b89.
//
// Solidity: function topupAvailable() constant returns(uint256)
func (_Wallet *WalletCaller) TopupAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "topupAvailable")
	return *ret0, err
}

// TopupAvailable is a free data retrieval call binding the contract method 0x6fc95b89.
//
// Solidity: function topupAvailable() constant returns(uint256)
func (_Wallet *WalletSession) TopupAvailable() (*big.Int, error) {
	return _Wallet.Contract.TopupAvailable(&_Wallet.CallOpts)
}

// TopupAvailable is a free data retrieval call binding the contract method 0x6fc95b89.
//
// Solidity: function topupAvailable() constant returns(uint256)
func (_Wallet *WalletCallerSession) TopupAvailable() (*big.Int, error) {
	return _Wallet.Contract.TopupAvailable(&_Wallet.CallOpts)
}

// TopupLimit is a free data retrieval call binding the contract method 0xd87d2a39.
//
// Solidity: function topupLimit() constant returns(uint256)
func (_Wallet *WalletCaller) TopupLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "topupLimit")
	return *ret0, err
}

// TopupLimit is a free data retrieval call binding the contract method 0xd87d2a39.
//
// Solidity: function topupLimit() constant returns(uint256)
func (_Wallet *WalletSession) TopupLimit() (*big.Int, error) {
	return _Wallet.Contract.TopupLimit(&_Wallet.CallOpts)
}

// TopupLimit is a free data retrieval call binding the contract method 0xd87d2a39.
//
// Solidity: function topupLimit() constant returns(uint256)
func (_Wallet *WalletCallerSession) TopupLimit() (*big.Int, error) {
	return _Wallet.Contract.TopupLimit(&_Wallet.CallOpts)
}

// AddController is a paid mutator transaction binding the contract method 0xa7fc7a07.
//
// Solidity: function addController(_account address) returns()
func (_Wallet *WalletTransactor) AddController(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "addController", _account)
}

// AddController is a paid mutator transaction binding the contract method 0xa7fc7a07.
//
// Solidity: function addController(_account address) returns()
func (_Wallet *WalletSession) AddController(_account common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.AddController(&_Wallet.TransactOpts, _account)
}

// AddController is a paid mutator transaction binding the contract method 0xa7fc7a07.
//
// Solidity: function addController(_account address) returns()
func (_Wallet *WalletTransactorSession) AddController(_account common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.AddController(&_Wallet.TransactOpts, _account)
}

// CancelSpendLimit is a paid mutator transaction binding the contract method 0x54f599d7.
//
// Solidity: function cancelSpendLimit() returns()
func (_Wallet *WalletTransactor) CancelSpendLimit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "cancelSpendLimit")
}

// CancelSpendLimit is a paid mutator transaction binding the contract method 0x54f599d7.
//
// Solidity: function cancelSpendLimit() returns()
func (_Wallet *WalletSession) CancelSpendLimit() (*types.Transaction, error) {
	return _Wallet.Contract.CancelSpendLimit(&_Wallet.TransactOpts)
}

// CancelSpendLimit is a paid mutator transaction binding the contract method 0x54f599d7.
//
// Solidity: function cancelSpendLimit() returns()
func (_Wallet *WalletTransactorSession) CancelSpendLimit() (*types.Transaction, error) {
	return _Wallet.Contract.CancelSpendLimit(&_Wallet.TransactOpts)
}

// CancelTopupLimit is a paid mutator transaction binding the contract method 0x95eccddb.
//
// Solidity: function cancelTopupLimit() returns()
func (_Wallet *WalletTransactor) CancelTopupLimit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "cancelTopupLimit")
}

// CancelTopupLimit is a paid mutator transaction binding the contract method 0x95eccddb.
//
// Solidity: function cancelTopupLimit() returns()
func (_Wallet *WalletSession) CancelTopupLimit() (*types.Transaction, error) {
	return _Wallet.Contract.CancelTopupLimit(&_Wallet.TransactOpts)
}

// CancelTopupLimit is a paid mutator transaction binding the contract method 0x95eccddb.
//
// Solidity: function cancelTopupLimit() returns()
func (_Wallet *WalletTransactorSession) CancelTopupLimit() (*types.Transaction, error) {
	return _Wallet.Contract.CancelTopupLimit(&_Wallet.TransactOpts)
}

// CancelWhitelistAddition is a paid mutator transaction binding the contract method 0x22401bde.
//
// Solidity: function cancelWhitelistAddition() returns()
func (_Wallet *WalletTransactor) CancelWhitelistAddition(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "cancelWhitelistAddition")
}

// CancelWhitelistAddition is a paid mutator transaction binding the contract method 0x22401bde.
//
// Solidity: function cancelWhitelistAddition() returns()
func (_Wallet *WalletSession) CancelWhitelistAddition() (*types.Transaction, error) {
	return _Wallet.Contract.CancelWhitelistAddition(&_Wallet.TransactOpts)
}

// CancelWhitelistAddition is a paid mutator transaction binding the contract method 0x22401bde.
//
// Solidity: function cancelWhitelistAddition() returns()
func (_Wallet *WalletTransactorSession) CancelWhitelistAddition() (*types.Transaction, error) {
	return _Wallet.Contract.CancelWhitelistAddition(&_Wallet.TransactOpts)
}

// CancelWhitelistRemoval is a paid mutator transaction binding the contract method 0x5204110c.
//
// Solidity: function cancelWhitelistRemoval() returns()
func (_Wallet *WalletTransactor) CancelWhitelistRemoval(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "cancelWhitelistRemoval")
}

// CancelWhitelistRemoval is a paid mutator transaction binding the contract method 0x5204110c.
//
// Solidity: function cancelWhitelistRemoval() returns()
func (_Wallet *WalletSession) CancelWhitelistRemoval() (*types.Transaction, error) {
	return _Wallet.Contract.CancelWhitelistRemoval(&_Wallet.TransactOpts)
}

// CancelWhitelistRemoval is a paid mutator transaction binding the contract method 0x5204110c.
//
// Solidity: function cancelWhitelistRemoval() returns()
func (_Wallet *WalletTransactorSession) CancelWhitelistRemoval() (*types.Transaction, error) {
	return _Wallet.Contract.CancelWhitelistRemoval(&_Wallet.TransactOpts)
}

// ConfirmSpendLimit is a paid mutator transaction binding the contract method 0x96136585.
//
// Solidity: function confirmSpendLimit() returns()
func (_Wallet *WalletTransactor) ConfirmSpendLimit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "confirmSpendLimit")
}

// ConfirmSpendLimit is a paid mutator transaction binding the contract method 0x96136585.
//
// Solidity: function confirmSpendLimit() returns()
func (_Wallet *WalletSession) ConfirmSpendLimit() (*types.Transaction, error) {
	return _Wallet.Contract.ConfirmSpendLimit(&_Wallet.TransactOpts)
}

// ConfirmSpendLimit is a paid mutator transaction binding the contract method 0x96136585.
//
// Solidity: function confirmSpendLimit() returns()
func (_Wallet *WalletTransactorSession) ConfirmSpendLimit() (*types.Transaction, error) {
	return _Wallet.Contract.ConfirmSpendLimit(&_Wallet.TransactOpts)
}

// ConfirmTopupLimit is a paid mutator transaction binding the contract method 0x8e112cf9.
//
// Solidity: function confirmTopupLimit() returns()
func (_Wallet *WalletTransactor) ConfirmTopupLimit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "confirmTopupLimit")
}

// ConfirmTopupLimit is a paid mutator transaction binding the contract method 0x8e112cf9.
//
// Solidity: function confirmTopupLimit() returns()
func (_Wallet *WalletSession) ConfirmTopupLimit() (*types.Transaction, error) {
	return _Wallet.Contract.ConfirmTopupLimit(&_Wallet.TransactOpts)
}

// ConfirmTopupLimit is a paid mutator transaction binding the contract method 0x8e112cf9.
//
// Solidity: function confirmTopupLimit() returns()
func (_Wallet *WalletTransactorSession) ConfirmTopupLimit() (*types.Transaction, error) {
	return _Wallet.Contract.ConfirmTopupLimit(&_Wallet.TransactOpts)
}

// ConfirmWhitelistAddition is a paid mutator transaction binding the contract method 0x3b98fe84.
//
// Solidity: function confirmWhitelistAddition() returns()
func (_Wallet *WalletTransactor) ConfirmWhitelistAddition(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "confirmWhitelistAddition")
}

// ConfirmWhitelistAddition is a paid mutator transaction binding the contract method 0x3b98fe84.
//
// Solidity: function confirmWhitelistAddition() returns()
func (_Wallet *WalletSession) ConfirmWhitelistAddition() (*types.Transaction, error) {
	return _Wallet.Contract.ConfirmWhitelistAddition(&_Wallet.TransactOpts)
}

// ConfirmWhitelistAddition is a paid mutator transaction binding the contract method 0x3b98fe84.
//
// Solidity: function confirmWhitelistAddition() returns()
func (_Wallet *WalletTransactorSession) ConfirmWhitelistAddition() (*types.Transaction, error) {
	return _Wallet.Contract.ConfirmWhitelistAddition(&_Wallet.TransactOpts)
}

// ConfirmWhitelistRemoval is a paid mutator transaction binding the contract method 0x4b730408.
//
// Solidity: function confirmWhitelistRemoval() returns()
func (_Wallet *WalletTransactor) ConfirmWhitelistRemoval(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "confirmWhitelistRemoval")
}

// ConfirmWhitelistRemoval is a paid mutator transaction binding the contract method 0x4b730408.
//
// Solidity: function confirmWhitelistRemoval() returns()
func (_Wallet *WalletSession) ConfirmWhitelistRemoval() (*types.Transaction, error) {
	return _Wallet.Contract.ConfirmWhitelistRemoval(&_Wallet.TransactOpts)
}

// ConfirmWhitelistRemoval is a paid mutator transaction binding the contract method 0x4b730408.
//
// Solidity: function confirmWhitelistRemoval() returns()
func (_Wallet *WalletTransactorSession) ConfirmWhitelistRemoval() (*types.Transaction, error) {
	return _Wallet.Contract.ConfirmWhitelistRemoval(&_Wallet.TransactOpts)
}

// InitializeSpendLimit is a paid mutator transaction binding the contract method 0x58453569.
//
// Solidity: function initializeSpendLimit(_amount uint256) returns()
func (_Wallet *WalletTransactor) InitializeSpendLimit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "initializeSpendLimit", _amount)
}

// InitializeSpendLimit is a paid mutator transaction binding the contract method 0x58453569.
//
// Solidity: function initializeSpendLimit(_amount uint256) returns()
func (_Wallet *WalletSession) InitializeSpendLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.InitializeSpendLimit(&_Wallet.TransactOpts, _amount)
}

// InitializeSpendLimit is a paid mutator transaction binding the contract method 0x58453569.
//
// Solidity: function initializeSpendLimit(_amount uint256) returns()
func (_Wallet *WalletTransactorSession) InitializeSpendLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.InitializeSpendLimit(&_Wallet.TransactOpts, _amount)
}

// InitializeTopupLimit is a paid mutator transaction binding the contract method 0x8ff3bf33.
//
// Solidity: function initializeTopupLimit(_amount uint256) returns()
func (_Wallet *WalletTransactor) InitializeTopupLimit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "initializeTopupLimit", _amount)
}

// InitializeTopupLimit is a paid mutator transaction binding the contract method 0x8ff3bf33.
//
// Solidity: function initializeTopupLimit(_amount uint256) returns()
func (_Wallet *WalletSession) InitializeTopupLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.InitializeTopupLimit(&_Wallet.TransactOpts, _amount)
}

// InitializeTopupLimit is a paid mutator transaction binding the contract method 0x8ff3bf33.
//
// Solidity: function initializeTopupLimit(_amount uint256) returns()
func (_Wallet *WalletTransactorSession) InitializeTopupLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.InitializeTopupLimit(&_Wallet.TransactOpts, _amount)
}

// InitializeWhitelist is a paid mutator transaction binding the contract method 0xf4199bb8.
//
// Solidity: function initializeWhitelist(_addresses address[]) returns()
func (_Wallet *WalletTransactor) InitializeWhitelist(opts *bind.TransactOpts, _addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "initializeWhitelist", _addresses)
}

// InitializeWhitelist is a paid mutator transaction binding the contract method 0xf4199bb8.
//
// Solidity: function initializeWhitelist(_addresses address[]) returns()
func (_Wallet *WalletSession) InitializeWhitelist(_addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.InitializeWhitelist(&_Wallet.TransactOpts, _addresses)
}

// InitializeWhitelist is a paid mutator transaction binding the contract method 0xf4199bb8.
//
// Solidity: function initializeWhitelist(_addresses address[]) returns()
func (_Wallet *WalletTransactorSession) InitializeWhitelist(_addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.InitializeWhitelist(&_Wallet.TransactOpts, _addresses)
}

// RemoveController is a paid mutator transaction binding the contract method 0xf6a74ed7.
//
// Solidity: function removeController(_account address) returns()
func (_Wallet *WalletTransactor) RemoveController(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "removeController", _account)
}

// RemoveController is a paid mutator transaction binding the contract method 0xf6a74ed7.
//
// Solidity: function removeController(_account address) returns()
func (_Wallet *WalletSession) RemoveController(_account common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.RemoveController(&_Wallet.TransactOpts, _account)
}

// RemoveController is a paid mutator transaction binding the contract method 0xf6a74ed7.
//
// Solidity: function removeController(_account address) returns()
func (_Wallet *WalletTransactorSession) RemoveController(_account common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.RemoveController(&_Wallet.TransactOpts, _account)
}

// SubmitSpendLimit is a paid mutator transaction binding the contract method 0xd9ec3018.
//
// Solidity: function submitSpendLimit(_amount uint256) returns()
func (_Wallet *WalletTransactor) SubmitSpendLimit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "submitSpendLimit", _amount)
}

// SubmitSpendLimit is a paid mutator transaction binding the contract method 0xd9ec3018.
//
// Solidity: function submitSpendLimit(_amount uint256) returns()
func (_Wallet *WalletSession) SubmitSpendLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.SubmitSpendLimit(&_Wallet.TransactOpts, _amount)
}

// SubmitSpendLimit is a paid mutator transaction binding the contract method 0xd9ec3018.
//
// Solidity: function submitSpendLimit(_amount uint256) returns()
func (_Wallet *WalletTransactorSession) SubmitSpendLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.SubmitSpendLimit(&_Wallet.TransactOpts, _amount)
}

// SubmitTopupLimit is a paid mutator transaction binding the contract method 0xd3a60586.
//
// Solidity: function submitTopupLimit(_amount uint256) returns()
func (_Wallet *WalletTransactor) SubmitTopupLimit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "submitTopupLimit", _amount)
}

// SubmitTopupLimit is a paid mutator transaction binding the contract method 0xd3a60586.
//
// Solidity: function submitTopupLimit(_amount uint256) returns()
func (_Wallet *WalletSession) SubmitTopupLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.SubmitTopupLimit(&_Wallet.TransactOpts, _amount)
}

// SubmitTopupLimit is a paid mutator transaction binding the contract method 0xd3a60586.
//
// Solidity: function submitTopupLimit(_amount uint256) returns()
func (_Wallet *WalletTransactorSession) SubmitTopupLimit(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.SubmitTopupLimit(&_Wallet.TransactOpts, _amount)
}

// SubmitWhitelistAddition is a paid mutator transaction binding the contract method 0x7fd004fa.
//
// Solidity: function submitWhitelistAddition(_addresses address[]) returns()
func (_Wallet *WalletTransactor) SubmitWhitelistAddition(opts *bind.TransactOpts, _addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "submitWhitelistAddition", _addresses)
}

// SubmitWhitelistAddition is a paid mutator transaction binding the contract method 0x7fd004fa.
//
// Solidity: function submitWhitelistAddition(_addresses address[]) returns()
func (_Wallet *WalletSession) SubmitWhitelistAddition(_addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.SubmitWhitelistAddition(&_Wallet.TransactOpts, _addresses)
}

// SubmitWhitelistAddition is a paid mutator transaction binding the contract method 0x7fd004fa.
//
// Solidity: function submitWhitelistAddition(_addresses address[]) returns()
func (_Wallet *WalletTransactorSession) SubmitWhitelistAddition(_addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.SubmitWhitelistAddition(&_Wallet.TransactOpts, _addresses)
}

// SubmitWhitelistRemoval is a paid mutator transaction binding the contract method 0x6137d670.
//
// Solidity: function submitWhitelistRemoval(_addresses address[]) returns()
func (_Wallet *WalletTransactor) SubmitWhitelistRemoval(opts *bind.TransactOpts, _addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "submitWhitelistRemoval", _addresses)
}

// SubmitWhitelistRemoval is a paid mutator transaction binding the contract method 0x6137d670.
//
// Solidity: function submitWhitelistRemoval(_addresses address[]) returns()
func (_Wallet *WalletSession) SubmitWhitelistRemoval(_addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.SubmitWhitelistRemoval(&_Wallet.TransactOpts, _addresses)
}

// SubmitWhitelistRemoval is a paid mutator transaction binding the contract method 0x6137d670.
//
// Solidity: function submitWhitelistRemoval(_addresses address[]) returns()
func (_Wallet *WalletTransactorSession) SubmitWhitelistRemoval(_addresses []common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.SubmitWhitelistRemoval(&_Wallet.TransactOpts, _addresses)
}

// TopupGas is a paid mutator transaction binding the contract method 0xeca71792.
//
// Solidity: function topupGas(_amount uint256) returns()
func (_Wallet *WalletTransactor) TopupGas(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "topupGas", _amount)
}

// TopupGas is a paid mutator transaction binding the contract method 0xeca71792.
//
// Solidity: function topupGas(_amount uint256) returns()
func (_Wallet *WalletSession) TopupGas(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.TopupGas(&_Wallet.TransactOpts, _amount)
}

// TopupGas is a paid mutator transaction binding the contract method 0xeca71792.
//
// Solidity: function topupGas(_amount uint256) returns()
func (_Wallet *WalletTransactorSession) TopupGas(_amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.TopupGas(&_Wallet.TransactOpts, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(_to address, _asset address, _amount uint256) returns()
func (_Wallet *WalletTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "transfer", _to, _asset, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(_to address, _asset address, _amount uint256) returns()
func (_Wallet *WalletSession) Transfer(_to common.Address, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.Transfer(&_Wallet.TransactOpts, _to, _asset, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(_to address, _asset address, _amount uint256) returns()
func (_Wallet *WalletTransactorSession) Transfer(_to common.Address, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.Transfer(&_Wallet.TransactOpts, _to, _asset, _amount)
}

// WalletAddControllerIterator is returned from FilterAddController and is used to iterate over the raw logs and unpacked data for AddController events raised by the Wallet contract.
type WalletAddControllerIterator struct {
	Event *WalletAddController // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WalletAddControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletAddController)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WalletAddController)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WalletAddControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletAddControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletAddController represents a AddController event raised by the Wallet contract.
type WalletAddController struct {
	Sender  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddController is a free log retrieval operation binding the contract event 0xec493e5f7bb2653b285a1e0af9af1c883375111370296bb332c8c5e307768959.
//
// Solidity: e AddController(_sender address, _account address)
func (_Wallet *WalletFilterer) FilterAddController(opts *bind.FilterOpts) (*WalletAddControllerIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "AddController")
	if err != nil {
		return nil, err
	}
	return &WalletAddControllerIterator{contract: _Wallet.contract, event: "AddController", logs: logs, sub: sub}, nil
}

// WatchAddController is a free log subscription operation binding the contract event 0xec493e5f7bb2653b285a1e0af9af1c883375111370296bb332c8c5e307768959.
//
// Solidity: e AddController(_sender address, _account address)
func (_Wallet *WalletFilterer) WatchAddController(opts *bind.WatchOpts, sink chan<- *WalletAddController) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "AddController")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletAddController)
				if err := _Wallet.contract.UnpackLog(event, "AddController", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// WalletCancelSpendLimitIterator is returned from FilterCancelSpendLimit and is used to iterate over the raw logs and unpacked data for CancelSpendLimit events raised by the Wallet contract.
type WalletCancelSpendLimitIterator struct {
	Event *WalletCancelSpendLimit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WalletCancelSpendLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletCancelSpendLimit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WalletCancelSpendLimit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WalletCancelSpendLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletCancelSpendLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletCancelSpendLimit represents a CancelSpendLimit event raised by the Wallet contract.
type WalletCancelSpendLimit struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCancelSpendLimit is a free log retrieval operation binding the contract event 0x04df2cb40e86637ce3f74891f08d714137e4f2767a45e0495b9eed18578fdfa0.
//
// Solidity: e CancelSpendLimit(_sender address)
func (_Wallet *WalletFilterer) FilterCancelSpendLimit(opts *bind.FilterOpts) (*WalletCancelSpendLimitIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "CancelSpendLimit")
	if err != nil {
		return nil, err
	}
	return &WalletCancelSpendLimitIterator{contract: _Wallet.contract, event: "CancelSpendLimit", logs: logs, sub: sub}, nil
}

// WatchCancelSpendLimit is a free log subscription operation binding the contract event 0x04df2cb40e86637ce3f74891f08d714137e4f2767a45e0495b9eed18578fdfa0.
//
// Solidity: e CancelSpendLimit(_sender address)
func (_Wallet *WalletFilterer) WatchCancelSpendLimit(opts *bind.WatchOpts, sink chan<- *WalletCancelSpendLimit) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "CancelSpendLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletCancelSpendLimit)
				if err := _Wallet.contract.UnpackLog(event, "CancelSpendLimit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// WalletCancelTopupLimitIterator is returned from FilterCancelTopupLimit and is used to iterate over the raw logs and unpacked data for CancelTopupLimit events raised by the Wallet contract.
type WalletCancelTopupLimitIterator struct {
	Event *WalletCancelTopupLimit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WalletCancelTopupLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletCancelTopupLimit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WalletCancelTopupLimit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WalletCancelTopupLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletCancelTopupLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletCancelTopupLimit represents a CancelTopupLimit event raised by the Wallet contract.
type WalletCancelTopupLimit struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCancelTopupLimit is a free log retrieval operation binding the contract event 0x3dfa16a690ee101e8342b8200b911ee7c0cb8e7e6cc596930a4e23a03a53941c.
//
// Solidity: e CancelTopupLimit(_sender address)
func (_Wallet *WalletFilterer) FilterCancelTopupLimit(opts *bind.FilterOpts) (*WalletCancelTopupLimitIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "CancelTopupLimit")
	if err != nil {
		return nil, err
	}
	return &WalletCancelTopupLimitIterator{contract: _Wallet.contract, event: "CancelTopupLimit", logs: logs, sub: sub}, nil
}

// WatchCancelTopupLimit is a free log subscription operation binding the contract event 0x3dfa16a690ee101e8342b8200b911ee7c0cb8e7e6cc596930a4e23a03a53941c.
//
// Solidity: e CancelTopupLimit(_sender address)
func (_Wallet *WalletFilterer) WatchCancelTopupLimit(opts *bind.WatchOpts, sink chan<- *WalletCancelTopupLimit) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "CancelTopupLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletCancelTopupLimit)
				if err := _Wallet.contract.UnpackLog(event, "CancelTopupLimit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// WalletCancelWhitelistAdditionIterator is returned from FilterCancelWhitelistAddition and is used to iterate over the raw logs and unpacked data for CancelWhitelistAddition events raised by the Wallet contract.
type WalletCancelWhitelistAdditionIterator struct {
	Event *WalletCancelWhitelistAddition // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WalletCancelWhitelistAdditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletCancelWhitelistAddition)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WalletCancelWhitelistAddition)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WalletCancelWhitelistAdditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletCancelWhitelistAdditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletCancelWhitelistAddition represents a CancelWhitelistAddition event raised by the Wallet contract.
type WalletCancelWhitelistAddition struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCancelWhitelistAddition is a free log retrieval operation binding the contract event 0x558fc446b058402d71bbea9bab10c2c469b90898aeac942895a7251fc2b68478.
//
// Solidity: e CancelWhitelistAddition(_sender address)
func (_Wallet *WalletFilterer) FilterCancelWhitelistAddition(opts *bind.FilterOpts) (*WalletCancelWhitelistAdditionIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "CancelWhitelistAddition")
	if err != nil {
		return nil, err
	}
	return &WalletCancelWhitelistAdditionIterator{contract: _Wallet.contract, event: "CancelWhitelistAddition", logs: logs, sub: sub}, nil
}

// WatchCancelWhitelistAddition is a free log subscription operation binding the contract event 0x558fc446b058402d71bbea9bab10c2c469b90898aeac942895a7251fc2b68478.
//
// Solidity: e CancelWhitelistAddition(_sender address)
func (_Wallet *WalletFilterer) WatchCancelWhitelistAddition(opts *bind.WatchOpts, sink chan<- *WalletCancelWhitelistAddition) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "CancelWhitelistAddition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletCancelWhitelistAddition)
				if err := _Wallet.contract.UnpackLog(event, "CancelWhitelistAddition", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// WalletCancelWhitelistRemovalIterator is returned from FilterCancelWhitelistRemoval and is used to iterate over the raw logs and unpacked data for CancelWhitelistRemoval events raised by the Wallet contract.
type WalletCancelWhitelistRemovalIterator struct {
	Event *WalletCancelWhitelistRemoval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WalletCancelWhitelistRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletCancelWhitelistRemoval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WalletCancelWhitelistRemoval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WalletCancelWhitelistRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletCancelWhitelistRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletCancelWhitelistRemoval represents a CancelWhitelistRemoval event raised by the Wallet contract.
type WalletCancelWhitelistRemoval struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCancelWhitelistRemoval is a free log retrieval operation binding the contract event 0x8978ec16114c856a3a04b746619c1d83d552fbfd4c115c0b3cbe066a9a56e01d.
//
// Solidity: e CancelWhitelistRemoval(_sender address)
func (_Wallet *WalletFilterer) FilterCancelWhitelistRemoval(opts *bind.FilterOpts) (*WalletCancelWhitelistRemovalIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "CancelWhitelistRemoval")
	if err != nil {
		return nil, err
	}
	return &WalletCancelWhitelistRemovalIterator{contract: _Wallet.contract, event: "CancelWhitelistRemoval", logs: logs, sub: sub}, nil
}

// WatchCancelWhitelistRemoval is a free log subscription operation binding the contract event 0x8978ec16114c856a3a04b746619c1d83d552fbfd4c115c0b3cbe066a9a56e01d.
//
// Solidity: e CancelWhitelistRemoval(_sender address)
func (_Wallet *WalletFilterer) WatchCancelWhitelistRemoval(opts *bind.WatchOpts, sink chan<- *WalletCancelWhitelistRemoval) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "CancelWhitelistRemoval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletCancelWhitelistRemoval)
				if err := _Wallet.contract.UnpackLog(event, "CancelWhitelistRemoval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// WalletDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Wallet contract.
type WalletDepositIterator struct {
	Event *WalletDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WalletDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WalletDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WalletDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletDeposit represents a Deposit event raised by the Wallet contract.
type WalletDeposit struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: e Deposit(_from address, _amount uint256)
func (_Wallet *WalletFilterer) FilterDeposit(opts *bind.FilterOpts) (*WalletDepositIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &WalletDepositIterator{contract: _Wallet.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: e Deposit(_from address, _amount uint256)
func (_Wallet *WalletFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *WalletDeposit) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletDeposit)
				if err := _Wallet.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// WalletRemoveControllerIterator is returned from FilterRemoveController and is used to iterate over the raw logs and unpacked data for RemoveController events raised by the Wallet contract.
type WalletRemoveControllerIterator struct {
	Event *WalletRemoveController // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WalletRemoveControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletRemoveController)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WalletRemoveController)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WalletRemoveControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletRemoveControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletRemoveController represents a RemoveController event raised by the Wallet contract.
type WalletRemoveController struct {
	Sender  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemoveController is a free log retrieval operation binding the contract event 0x98da1b1dd7d69af3ffee8826b8a31d3e98874a91a2e90e819fb6df0cfa91ca4d.
//
// Solidity: e RemoveController(_sender address, _account address)
func (_Wallet *WalletFilterer) FilterRemoveController(opts *bind.FilterOpts) (*WalletRemoveControllerIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "RemoveController")
	if err != nil {
		return nil, err
	}
	return &WalletRemoveControllerIterator{contract: _Wallet.contract, event: "RemoveController", logs: logs, sub: sub}, nil
}

// WatchRemoveController is a free log subscription operation binding the contract event 0x98da1b1dd7d69af3ffee8826b8a31d3e98874a91a2e90e819fb6df0cfa91ca4d.
//
// Solidity: e RemoveController(_sender address, _account address)
func (_Wallet *WalletFilterer) WatchRemoveController(opts *bind.WatchOpts, sink chan<- *WalletRemoveController) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "RemoveController")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletRemoveController)
				if err := _Wallet.contract.UnpackLog(event, "RemoveController", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// WalletSetSpendLimitIterator is returned from FilterSetSpendLimit and is used to iterate over the raw logs and unpacked data for SetSpendLimit events raised by the Wallet contract.
type WalletSetSpendLimitIterator struct {
	Event *WalletSetSpendLimit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WalletSetSpendLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSetSpendLimit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WalletSetSpendLimit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WalletSetSpendLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSetSpendLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSetSpendLimit represents a SetSpendLimit event raised by the Wallet contract.
type WalletSetSpendLimit struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetSpendLimit is a free log retrieval operation binding the contract event 0x068f112e5ec923d412be64779fe69e0fcbb6784c6617e94cccc8fd348f2e0f21.
//
// Solidity: e SetSpendLimit(_sender address, _amount uint256)
func (_Wallet *WalletFilterer) FilterSetSpendLimit(opts *bind.FilterOpts) (*WalletSetSpendLimitIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "SetSpendLimit")
	if err != nil {
		return nil, err
	}
	return &WalletSetSpendLimitIterator{contract: _Wallet.contract, event: "SetSpendLimit", logs: logs, sub: sub}, nil
}

// WatchSetSpendLimit is a free log subscription operation binding the contract event 0x068f112e5ec923d412be64779fe69e0fcbb6784c6617e94cccc8fd348f2e0f21.
//
// Solidity: e SetSpendLimit(_sender address, _amount uint256)
func (_Wallet *WalletFilterer) WatchSetSpendLimit(opts *bind.WatchOpts, sink chan<- *WalletSetSpendLimit) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "SetSpendLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSetSpendLimit)
				if err := _Wallet.contract.UnpackLog(event, "SetSpendLimit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// WalletSetTopupLimitIterator is returned from FilterSetTopupLimit and is used to iterate over the raw logs and unpacked data for SetTopupLimit events raised by the Wallet contract.
type WalletSetTopupLimitIterator struct {
	Event *WalletSetTopupLimit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WalletSetTopupLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSetTopupLimit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WalletSetTopupLimit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WalletSetTopupLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSetTopupLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSetTopupLimit represents a SetTopupLimit event raised by the Wallet contract.
type WalletSetTopupLimit struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetTopupLimit is a free log retrieval operation binding the contract event 0x542c9d9a32f2597d63c74984374742431933d68fe2fd019ee3592121b5e53e81.
//
// Solidity: e SetTopupLimit(_sender address, _amount uint256)
func (_Wallet *WalletFilterer) FilterSetTopupLimit(opts *bind.FilterOpts) (*WalletSetTopupLimitIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "SetTopupLimit")
	if err != nil {
		return nil, err
	}
	return &WalletSetTopupLimitIterator{contract: _Wallet.contract, event: "SetTopupLimit", logs: logs, sub: sub}, nil
}

// WatchSetTopupLimit is a free log subscription operation binding the contract event 0x542c9d9a32f2597d63c74984374742431933d68fe2fd019ee3592121b5e53e81.
//
// Solidity: e SetTopupLimit(_sender address, _amount uint256)
func (_Wallet *WalletFilterer) WatchSetTopupLimit(opts *bind.WatchOpts, sink chan<- *WalletSetTopupLimit) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "SetTopupLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSetTopupLimit)
				if err := _Wallet.contract.UnpackLog(event, "SetTopupLimit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// WalletSubmitSpendLimitIterator is returned from FilterSubmitSpendLimit and is used to iterate over the raw logs and unpacked data for SubmitSpendLimit events raised by the Wallet contract.
type WalletSubmitSpendLimitIterator struct {
	Event *WalletSubmitSpendLimit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WalletSubmitSpendLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSubmitSpendLimit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WalletSubmitSpendLimit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WalletSubmitSpendLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSubmitSpendLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSubmitSpendLimit represents a SubmitSpendLimit event raised by the Wallet contract.
type WalletSubmitSpendLimit struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubmitSpendLimit is a free log retrieval operation binding the contract event 0xf56c67a9b328e0778c4806a71b89e6e351bc301249815bd53d4110a5bc59a133.
//
// Solidity: e SubmitSpendLimit(_amount uint256)
func (_Wallet *WalletFilterer) FilterSubmitSpendLimit(opts *bind.FilterOpts) (*WalletSubmitSpendLimitIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "SubmitSpendLimit")
	if err != nil {
		return nil, err
	}
	return &WalletSubmitSpendLimitIterator{contract: _Wallet.contract, event: "SubmitSpendLimit", logs: logs, sub: sub}, nil
}

// WatchSubmitSpendLimit is a free log subscription operation binding the contract event 0xf56c67a9b328e0778c4806a71b89e6e351bc301249815bd53d4110a5bc59a133.
//
// Solidity: e SubmitSpendLimit(_amount uint256)
func (_Wallet *WalletFilterer) WatchSubmitSpendLimit(opts *bind.WatchOpts, sink chan<- *WalletSubmitSpendLimit) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "SubmitSpendLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSubmitSpendLimit)
				if err := _Wallet.contract.UnpackLog(event, "SubmitSpendLimit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// WalletSubmitTopupLimitIterator is returned from FilterSubmitTopupLimit and is used to iterate over the raw logs and unpacked data for SubmitTopupLimit events raised by the Wallet contract.
type WalletSubmitTopupLimitIterator struct {
	Event *WalletSubmitTopupLimit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WalletSubmitTopupLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSubmitTopupLimit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WalletSubmitTopupLimit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WalletSubmitTopupLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSubmitTopupLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSubmitTopupLimit represents a SubmitTopupLimit event raised by the Wallet contract.
type WalletSubmitTopupLimit struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubmitTopupLimit is a free log retrieval operation binding the contract event 0xfb68cbcb26831437ffeaeec06583fba34b716958f3363288069f862a909d9112.
//
// Solidity: e SubmitTopupLimit(_amount uint256)
func (_Wallet *WalletFilterer) FilterSubmitTopupLimit(opts *bind.FilterOpts) (*WalletSubmitTopupLimitIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "SubmitTopupLimit")
	if err != nil {
		return nil, err
	}
	return &WalletSubmitTopupLimitIterator{contract: _Wallet.contract, event: "SubmitTopupLimit", logs: logs, sub: sub}, nil
}

// WatchSubmitTopupLimit is a free log subscription operation binding the contract event 0xfb68cbcb26831437ffeaeec06583fba34b716958f3363288069f862a909d9112.
//
// Solidity: e SubmitTopupLimit(_amount uint256)
func (_Wallet *WalletFilterer) WatchSubmitTopupLimit(opts *bind.WatchOpts, sink chan<- *WalletSubmitTopupLimit) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "SubmitTopupLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSubmitTopupLimit)
				if err := _Wallet.contract.UnpackLog(event, "SubmitTopupLimit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// WalletSubmitWhitelistAdditionIterator is returned from FilterSubmitWhitelistAddition and is used to iterate over the raw logs and unpacked data for SubmitWhitelistAddition events raised by the Wallet contract.
type WalletSubmitWhitelistAdditionIterator struct {
	Event *WalletSubmitWhitelistAddition // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WalletSubmitWhitelistAdditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSubmitWhitelistAddition)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WalletSubmitWhitelistAddition)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WalletSubmitWhitelistAdditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSubmitWhitelistAdditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSubmitWhitelistAddition represents a SubmitWhitelistAddition event raised by the Wallet contract.
type WalletSubmitWhitelistAddition struct {
	Addresses []common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSubmitWhitelistAddition is a free log retrieval operation binding the contract event 0x3e8b4ac7dd097b817f72dee5feb5c43e5316cf4f03a6f8fb6936376b8962703a.
//
// Solidity: e SubmitWhitelistAddition(_addresses address[])
func (_Wallet *WalletFilterer) FilterSubmitWhitelistAddition(opts *bind.FilterOpts) (*WalletSubmitWhitelistAdditionIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "SubmitWhitelistAddition")
	if err != nil {
		return nil, err
	}
	return &WalletSubmitWhitelistAdditionIterator{contract: _Wallet.contract, event: "SubmitWhitelistAddition", logs: logs, sub: sub}, nil
}

// WatchSubmitWhitelistAddition is a free log subscription operation binding the contract event 0x3e8b4ac7dd097b817f72dee5feb5c43e5316cf4f03a6f8fb6936376b8962703a.
//
// Solidity: e SubmitWhitelistAddition(_addresses address[])
func (_Wallet *WalletFilterer) WatchSubmitWhitelistAddition(opts *bind.WatchOpts, sink chan<- *WalletSubmitWhitelistAddition) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "SubmitWhitelistAddition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSubmitWhitelistAddition)
				if err := _Wallet.contract.UnpackLog(event, "SubmitWhitelistAddition", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// WalletSubmitWhitelistRemovalIterator is returned from FilterSubmitWhitelistRemoval and is used to iterate over the raw logs and unpacked data for SubmitWhitelistRemoval events raised by the Wallet contract.
type WalletSubmitWhitelistRemovalIterator struct {
	Event *WalletSubmitWhitelistRemoval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WalletSubmitWhitelistRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSubmitWhitelistRemoval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WalletSubmitWhitelistRemoval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WalletSubmitWhitelistRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSubmitWhitelistRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSubmitWhitelistRemoval represents a SubmitWhitelistRemoval event raised by the Wallet contract.
type WalletSubmitWhitelistRemoval struct {
	Addresses []common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSubmitWhitelistRemoval is a free log retrieval operation binding the contract event 0x88c78a1f512216b3262ba17462fd5cfdffd48bc03c449e30aa77f71e4b50c1e6.
//
// Solidity: e SubmitWhitelistRemoval(_addresses address[])
func (_Wallet *WalletFilterer) FilterSubmitWhitelistRemoval(opts *bind.FilterOpts) (*WalletSubmitWhitelistRemovalIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "SubmitWhitelistRemoval")
	if err != nil {
		return nil, err
	}
	return &WalletSubmitWhitelistRemovalIterator{contract: _Wallet.contract, event: "SubmitWhitelistRemoval", logs: logs, sub: sub}, nil
}

// WatchSubmitWhitelistRemoval is a free log subscription operation binding the contract event 0x88c78a1f512216b3262ba17462fd5cfdffd48bc03c449e30aa77f71e4b50c1e6.
//
// Solidity: e SubmitWhitelistRemoval(_addresses address[])
func (_Wallet *WalletFilterer) WatchSubmitWhitelistRemoval(opts *bind.WatchOpts, sink chan<- *WalletSubmitWhitelistRemoval) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "SubmitWhitelistRemoval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSubmitWhitelistRemoval)
				if err := _Wallet.contract.UnpackLog(event, "SubmitWhitelistRemoval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// WalletTopupGasIterator is returned from FilterTopupGas and is used to iterate over the raw logs and unpacked data for TopupGas events raised by the Wallet contract.
type WalletTopupGasIterator struct {
	Event *WalletTopupGas // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WalletTopupGasIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletTopupGas)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WalletTopupGas)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WalletTopupGasIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletTopupGasIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletTopupGas represents a TopupGas event raised by the Wallet contract.
type WalletTopupGas struct {
	Sender common.Address
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTopupGas is a free log retrieval operation binding the contract event 0x11bb310b94280c15845698b8ce945817e14456a5d1582e387e6e4a01ef2c6742.
//
// Solidity: e TopupGas(_sender address, _owner address, _amount uint256)
func (_Wallet *WalletFilterer) FilterTopupGas(opts *bind.FilterOpts) (*WalletTopupGasIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "TopupGas")
	if err != nil {
		return nil, err
	}
	return &WalletTopupGasIterator{contract: _Wallet.contract, event: "TopupGas", logs: logs, sub: sub}, nil
}

// WatchTopupGas is a free log subscription operation binding the contract event 0x11bb310b94280c15845698b8ce945817e14456a5d1582e387e6e4a01ef2c6742.
//
// Solidity: e TopupGas(_sender address, _owner address, _amount uint256)
func (_Wallet *WalletFilterer) WatchTopupGas(opts *bind.WatchOpts, sink chan<- *WalletTopupGas) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "TopupGas")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletTopupGas)
				if err := _Wallet.contract.UnpackLog(event, "TopupGas", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// WalletTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Wallet contract.
type WalletTransferIterator struct {
	Event *WalletTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WalletTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WalletTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WalletTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletTransfer represents a Transfer event raised by the Wallet contract.
type WalletTransfer struct {
	To     common.Address
	Asset  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_to address, _asset address, _amount uint256)
func (_Wallet *WalletFilterer) FilterTransfer(opts *bind.FilterOpts) (*WalletTransferIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &WalletTransferIterator{contract: _Wallet.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_to address, _asset address, _amount uint256)
func (_Wallet *WalletFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WalletTransfer) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletTransfer)
				if err := _Wallet.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// WalletWhitelistAdditionIterator is returned from FilterWhitelistAddition and is used to iterate over the raw logs and unpacked data for WhitelistAddition events raised by the Wallet contract.
type WalletWhitelistAdditionIterator struct {
	Event *WalletWhitelistAddition // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WalletWhitelistAdditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletWhitelistAddition)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WalletWhitelistAddition)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WalletWhitelistAdditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletWhitelistAdditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletWhitelistAddition represents a WhitelistAddition event raised by the Wallet contract.
type WalletWhitelistAddition struct {
	Sender    common.Address
	Addresses []common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAddition is a free log retrieval operation binding the contract event 0xf0ff099f6211e3a70fce6ca4dd04c29600335b95f2d56b4e212d5ac69d6a0cfd.
//
// Solidity: e WhitelistAddition(_sender address, _addresses address[])
func (_Wallet *WalletFilterer) FilterWhitelistAddition(opts *bind.FilterOpts) (*WalletWhitelistAdditionIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "WhitelistAddition")
	if err != nil {
		return nil, err
	}
	return &WalletWhitelistAdditionIterator{contract: _Wallet.contract, event: "WhitelistAddition", logs: logs, sub: sub}, nil
}

// WatchWhitelistAddition is a free log subscription operation binding the contract event 0xf0ff099f6211e3a70fce6ca4dd04c29600335b95f2d56b4e212d5ac69d6a0cfd.
//
// Solidity: e WhitelistAddition(_sender address, _addresses address[])
func (_Wallet *WalletFilterer) WatchWhitelistAddition(opts *bind.WatchOpts, sink chan<- *WalletWhitelistAddition) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "WhitelistAddition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletWhitelistAddition)
				if err := _Wallet.contract.UnpackLog(event, "WhitelistAddition", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// WalletWhitelistRemovalIterator is returned from FilterWhitelistRemoval and is used to iterate over the raw logs and unpacked data for WhitelistRemoval events raised by the Wallet contract.
type WalletWhitelistRemovalIterator struct {
	Event *WalletWhitelistRemoval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WalletWhitelistRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletWhitelistRemoval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WalletWhitelistRemoval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WalletWhitelistRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletWhitelistRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletWhitelistRemoval represents a WhitelistRemoval event raised by the Wallet contract.
type WalletWhitelistRemoval struct {
	Sender    common.Address
	Addresses []common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWhitelistRemoval is a free log retrieval operation binding the contract event 0xce26ffefe745221a0fc931cb9394346a68c154ba02bc4c5e760bfe29a533f094.
//
// Solidity: e WhitelistRemoval(_sender address, _addresses address[])
func (_Wallet *WalletFilterer) FilterWhitelistRemoval(opts *bind.FilterOpts) (*WalletWhitelistRemovalIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "WhitelistRemoval")
	if err != nil {
		return nil, err
	}
	return &WalletWhitelistRemovalIterator{contract: _Wallet.contract, event: "WhitelistRemoval", logs: logs, sub: sub}, nil
}

// WatchWhitelistRemoval is a free log subscription operation binding the contract event 0xce26ffefe745221a0fc931cb9394346a68c154ba02bc4c5e760bfe29a533f094.
//
// Solidity: e WhitelistRemoval(_sender address, _addresses address[])
func (_Wallet *WalletFilterer) WatchWhitelistRemoval(opts *bind.WatchOpts, sink chan<- *WalletWhitelistRemoval) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "WhitelistRemoval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletWhitelistRemoval)
				if err := _Wallet.contract.UnpackLog(event, "WhitelistRemoval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
