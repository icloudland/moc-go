package mocjson

type GetBlockInfoByBlockHeightDealWithDataResult struct {
	Date              string `json:"date"`
	FoundryPublicKey  string `json:"foundryPublicKey"`
	BlockTransactions []struct {
		Date         string `json:"date"`
		Confirm      int    `json:"confirm"`
		TransType    int    `json:"transType"`
		ContractType int    `json:"contractType"`
		Fee          string `json:"fee"`
		TokenName    string `json:"tokenName"`
		From         string `json:"from"`
		Remark       string `json:"remark"`
		To           string `json:"to"`
		Value        string `json:"value"`
		TxHash       string `json:"txHash"`
	} `json:"blockTransactions"`
	BlockHeight    int     `json:"blockHeight"`
	TotalFee       float64 `json:"totalFee"`
	HashPrevBlock  string  `json:"hashPrevBlock"`
	HashBlock      string  `json:"hashBlock"`
	BlockSignature string  `json:"blockSignature"`
	BlockSize      int     `json:"blockSize"`
}

type GetTransactionByHashResult struct {
	CoinType     string  `json:"coinType"`
	Date         string  `json:"date"`
	Amount       float64 `json:"amount"`
	ContractType int     `json:"contractType"`
	Fee          float64 `json:"fee"`
	TokenName    string  `json:"tokenName"`
	Remark       string  `json:"remark"`
	Confirm      int     `json:"confirm"`
	BlockHeight  int     `json:"blockHeight"`
	TransType    int     `json:"transType"`
	From         string  `json:"from"`
	To           string  `json:"to"`
	TxHash       string  `json:"txHash"`
	Value        float64 `json:"value"`
}

type TransferAccountsResult struct {
	TransHash string `json:"transHash"`
}

type GetAccountBalanceResult struct {
	TotalEffectiveIncome string `json:"totalEffectiveIncome"`
	TotalExpenditure     string `json:"totalExpenditure"`
	TotalIncome          string `json:"totalIncome"`
	Address              string `json:"address"`
	Balance              string `json:"balance"`
	PubKey               string `json:"pubKey"`
}
