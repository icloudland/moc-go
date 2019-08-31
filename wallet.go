package moc

import (
	"encoding/json"
	"github.com/icloudland/moc-go/mocjson"
	"github.com/tidwall/gjson"
)

type FutureGetCurNodeBlockHeight chan *response

func (r FutureGetCurNodeBlockHeight) Receive() (int64, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return 0, err
	}

	var height int64
	err = json.Unmarshal(res, &height)
	if err != nil {
		return height, err
	}

	return height, nil
}

func (c *Client) GetCurNodeBlockHeightAsync() FutureGetCurNodeBlockHeight {

	cmd := &mocjson.GetCurNodeBlockHeightCmd{}
	return c.sendCmd(cmd)
}

func (c *Client) GetCurNodeBlockHeight() (int64, error) {
	return c.GetCurNodeBlockHeightAsync().Receive()
}

type FutureGetBlockInfoByBlockHeightDealWithData chan *response

func (r FutureGetBlockInfoByBlockHeightDealWithData) Receive() (
	*mocjson.GetBlockInfoByBlockHeightDealWithDataResult, error) {

	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	block := gjson.GetBytes(res, "block").String()
	if block == "" {
		return nil, ErrUndefined
	}

	var resObj mocjson.GetBlockInfoByBlockHeightDealWithDataResult
	err = json.Unmarshal([]byte(block), &resObj)
	if err != nil {
		return nil, err
	}

	return &resObj, nil
}

func (c *Client) GetBlockInfoByBlockHeightDealWithDataAsync(
	height int64) FutureGetBlockInfoByBlockHeightDealWithData {

	cmd := mocjson.NewGetBlockInfoByBlockHeightDealWithDataCmd(height)
	return c.sendCmd(cmd)
}

func (c *Client) GetBlockInfoByBlockHeightDealWithData(height int64) (
	*mocjson.GetBlockInfoByBlockHeightDealWithDataResult, error) {
	return c.GetBlockInfoByBlockHeightDealWithDataAsync(height).Receive()
}

type FutureGetTransactionByHash chan *response

func (r FutureGetTransactionByHash) Receive() ([]mocjson.GetTransactionByHashResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	transactions := gjson.GetBytes(res, "transactions").String()
	if transactions == "" {
		return nil, ErrUndefined
	}

	var resObj []mocjson.GetTransactionByHashResult
	err = json.Unmarshal([]byte(transactions), &resObj)
	if err != nil {
		return nil, err
	}

	return resObj, nil
}

func (c *Client) GetTransactionByHashAsync(hash string) FutureGetTransactionByHash {

	cmd := mocjson.NewGetTransactionByHashCmd(hash)
	return c.sendCmd(cmd)
}

func (c *Client) GetTransactionByHash(hash string) ([]mocjson.GetTransactionByHashResult, error) {
	return c.GetTransactionByHashAsync(hash).Receive()
}

type FutureTransferAccounts chan *response

func (r FutureTransferAccounts) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	hash := gjson.GetBytes(res, "transHash").String()

	return hash, nil
}

func (c *Client) TransferAccountsAsync(transFrom, transTo, transValue, fee,
	passWord, transToPubkey, remark, tokenName string) FutureTransferAccounts {

	cmd := mocjson.NewTransferAccountsCmd(
		transFrom, transTo, transValue, fee, passWord, transToPubkey, remark, tokenName)
	return c.sendCmd(cmd)
}

func (c *Client) TransferAccounts(transFrom, transTo, transValue, fee,
	passWord, transToPubkey, remark, tokenName string) (string, error) {
	return c.TransferAccountsAsync(
		transFrom, transTo, transValue, fee, passWord, transToPubkey, remark, tokenName).Receive()
}

type FutureGetAccountBalance chan *response

func (r FutureGetAccountBalance) Receive() (*mocjson.GetAccountBalanceResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	restString := gjson.GetBytes(res, "accountMap").String()
	if restString == "" {
		return nil, ErrUndefined
	}

	var resObj mocjson.GetAccountBalanceResult
	err = json.Unmarshal([]byte(restString), &resObj)
	if err != nil {
		return nil, err
	}

	return &resObj, nil
}

func (c *Client) GetAccountBalanceAsync(address, tokenName string) FutureGetAccountBalance {

	cmd := mocjson.NewGetAccountBalanceCmd(address, tokenName)
	return c.sendCmd(cmd)
}

func (c *Client) GetAccountBalance(address, tokeName string) (*mocjson.GetAccountBalanceResult, error) {
	return c.GetAccountBalanceAsync(address, tokeName).Receive()
}
