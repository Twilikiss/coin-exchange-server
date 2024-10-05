// Package model
// @Author twilikiss 2024/8/10 1:24:24
package model

type BlockHashResult struct {
	Id     string `json:"id"`
	Error  string `json:"error"`
	Result string `json:"result"`
}

type BlockResult struct {
	Id     string      `json:"id"`
	Error  string      `json:"error"`
	Result BlockSimple `json:"result"`
}
type BlockSimple struct {
	Hash string   `json:"hash"`
	Tx   []string `json:"tx"`
	Time int64    `json:"time"`
}

type RawTransaction struct {
	TxId      string `json:"txid"`
	Hash      string `json:"hash"`
	Locktime  int64  `json:"locktime"`
	Version   int    `json:"version"`
	Size      int    `json:"size"`
	Vsize     int    `json:"vsize"`
	Weight    int    `json:"weight"`
	Vin       []Vin  `json:"vin"`
	Vout      []Vout `json:"vout"`
	Time      int64  `json:"time"`
	Hex       string `json:"hex"`
	Blocktime int64  `json:"blocktime"`
	Blockhash string `json:"blockhash"`
}

type RawTransactionResult struct {
	Id     string         `json:"id"`
	Error  string         `json:"error"`
	Result RawTransaction `json:"result"`
}

type Vin struct {
	Txid        string            `json:"txid"`
	Vout        int               `json:"vout"`
	Txinwitness []string          `json:"txinwitness"`
	Sequence    int64             `json:"sequence"`
	ScriptSig   map[string]string `json:"scriptSig"`
}

type Vout struct {
	Value        float64      `json:"value"`
	N            int          `json:"n"`
	ScriptPubKey ScriptPubKey `json:"scriptPubKey"`
}

type ScriptPubKey struct {
	Asm     string `json:"asm"`
	Desc    string `json:"desc"`
	Hex     string `json:"hex"`
	Address string `json:"address"`
	Type    string `json:"type"`
}
