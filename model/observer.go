package model

import "time"

type ObserverResponse struct {
	ChainID string `json:"chain_id"`
	Type    string `json:"type"`
	Data    Data   `json:"data"`
}
type Header struct {
	Height string `json:"height"`
}
type Block struct {
	Header Header `json:"header"`
}
type NativeToken struct {
	Denom string `json:"denom"`
}
type Info struct {
	NativeToken NativeToken `json:"native_token"`
}
type OfferAsset struct {
	Amount string `json:"amount"`
	Info   Info   `json:"info"`
}
type Swap struct {
	BeliefPrice string     `json:"belief_price"`
	MaxSpread   string     `json:"max_spread"`
	OfferAsset  OfferAsset `json:"offer_asset"`
}
type ExecuteMsg struct {
	Swap Swap `json:"swap"`
}
type Coins struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}
type Messages struct {
	Type       string     `json:"@type"`
	Sender     string     `json:"sender"`
	Contract   string     `json:"contract"`
	ExecuteMsg ExecuteMsg `json:"execute_msg"`
	Coins      []Coins    `json:"coins"`
}
type Body struct {
	Messages                    []Messages    `json:"messages"`
	Memo                        string        `json:"memo"`
	TimeoutHeight               string        `json:"timeout_height"`
	ExtensionOptions            []interface{} `json:"extension_options"`
	NonCriticalExtensionOptions []interface{} `json:"non_critical_extension_options"`
}
type PublicKey struct {
	Type string `json:"@type"`
	Key  string `json:"key"`
}
type Single struct {
	Mode string `json:"mode"`
}
type ModeInfo struct {
	Single Single `json:"single"`
}
type SignerInfos struct {
	PublicKey PublicKey `json:"public_key"`
	ModeInfo  ModeInfo  `json:"mode_info"`
	Sequence  string    `json:"sequence"`
}
type Amount struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}
type Fee struct {
	Amount   []Amount `json:"amount"`
	GasLimit string   `json:"gas_limit"`
	Payer    string   `json:"payer"`
	Granter  string   `json:"granter"`
}
type AuthInfo struct {
	SignerInfos []SignerInfos `json:"signer_infos"`
	Fee         Fee           `json:"fee"`
}
type Attributes struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type Events struct {
	Type       string       `json:"type"`
	Attributes []Attributes `json:"attributes"`
}
type Logs struct {
	MsgIndex int      `json:"msg_index"`
	Log      string   `json:"log"`
	Events   []Events `json:"events"`
}
type Tx struct {
	Type       string   `json:"@type"`
	Body       Body     `json:"body"`
	AuthInfo   AuthInfo `json:"auth_info"`
	Signatures []string `json:"signatures"`
}
type Txs struct {
	Body       Body      `json:"body"`
	AuthInfo   AuthInfo  `json:"auth_info"`
	Signatures []string  `json:"signatures"`
	Height     string    `json:"height"`
	Txhash     string    `json:"txhash"`
	Codespace  string    `json:"codespace"`
	Code       int       `json:"code"`
	Data       string    `json:"data"`
	RawLog     string    `json:"raw_log"`
	Logs       []Logs    `json:"logs"`
	Info       string    `json:"info"`
	GasWanted  string    `json:"gas_wanted"`
	GasUsed    string    `json:"gas_used"`
	Tx         Tx        `json:"tx"`
	Timestamp  time.Time `json:"timestamp"`
	Events     []Events  `json:"events"`
}
type Data struct {
	Block Block `json:"block"`
	Txs   []Txs `json:"txs"`
}
