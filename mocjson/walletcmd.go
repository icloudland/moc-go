package mocjson

type GetCurNodeBlockHeightCmd struct {
}

type GetBlockInfoByBlockHeightDealWithDataCmd struct {
	BlockHeight int64 `qstring:"blockHeight"`
}

func NewGetBlockInfoByBlockHeightDealWithDataCmd(blockHeight int64) *GetBlockInfoByBlockHeightDealWithDataCmd {
	return &GetBlockInfoByBlockHeightDealWithDataCmd{
		BlockHeight: blockHeight,
	}
}

type GetTransactionByHashCmd struct {
	Hash string `qstring:"hash"`
}

func NewGetTransactionByHashCmd(hash string) *GetTransactionByHashCmd {
	return &GetTransactionByHashCmd{
		Hash: hash,
	}
}

type TransferAccountsCmd struct {
	TransFrom     string `json:"transFrom" qstring:"transFrom"`
	TransTo       string `json:"transTo" qstring:"transTo"`
	TransValue    string `json:"transValue" qstring:"transValue"`
	Fee           string `json:"fee" qstring:"fee"`
	PassWord      string `json:"passWord" qstring:"passWord"`
	TransToPubkey string `json:"transToPubkey" qstring:"transToPubkey"`
	Remark        string `json:"remark" qstring:"remark"`
	TokenName     string `json:"tokenName" qstring:"tokenName"`
}

func NewTransferAccountsCmd(transFrom, transTo, transValue, fee,
	passWord, transToPubkey, remark, tokenName string) *TransferAccountsCmd {

	return &TransferAccountsCmd{
		TransFrom:     transFrom,
		TransTo:       transTo,
		TransValue:    transValue,
		Fee:           fee,
		PassWord:      passWord,
		TransToPubkey: transToPubkey,
		Remark:        remark,
		TokenName:     tokenName,
	}
}

type GetAccountBalanceCmd struct {
	Address   string `qstring:"address"`
	TokenName string `qstring:"tokenName"`
}

func NewGetAccountBalanceCmd(address, tokenName string) *GetAccountBalanceCmd {
	return &GetAccountBalanceCmd{
		Address:   address,
		TokenName: tokenName,
	}
}

func init() {
	// No special flags for commands in this file.
	flags := UsageFlag(0)
	MustRegisterCmd("WalletController/getCurNodeBlockHeight:get", (*GetCurNodeBlockHeightCmd)(nil), flags)
	MustRegisterCmd("WalletController/getBlockInfoByBlockHeightDealWithData:get", (*GetBlockInfoByBlockHeightDealWithDataCmd)(nil), flags)
	MustRegisterCmd("WalletController/getTransactionByHash:post", (*GetTransactionByHashCmd)(nil), flags)
	MustRegisterCmd("WalletController/transferAccounts:post", (*TransferAccountsCmd)(nil), flags)
	MustRegisterCmd("WalletController/getAccountBalance:get", (*GetAccountBalanceCmd)(nil), flags)

}
