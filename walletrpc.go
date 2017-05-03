package walletrpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"net"
	"github.com/walletrpc/wallet"
)

// CoinDefaultHost   default coin address
const CoinDefaultHost string = "localhost"

//CoinDefaultPort   default coin port
const CoinDefaultPort int = 8332

//CoinDefaultProto  default coin protp
const CoinDefaultProto string = "http"

//RPCTimeOut default timeout(second)
const RPCTimeOut = 150

//RpcClient RPC struct
type RpcClient struct {
	// Configuration options
	username    string
	password    string
	proto       string
	host        string
	port        int
	certificate string
	// Information and debugging
	Status    int
	LastError error
	// rawResponse string
	responseData []byte
	id           int
	client       *http.Client
}

//NewCoin   create a new RPC instance
func NewClient(coinUser, coinPasswd, coinHost string, coinPort int) (client *RpcClient, err error) {
	client = &RpcClient{
		username: coinUser,
		password: coinPasswd,
		host:     coinHost,
		port:     coinPort,
		proto:    CoinDefaultProto,
	}
	if len(coinHost) == 0 {
		client.host = CoinDefaultHost
	}
	if coinPort < 0 || coinPort > 65535 {
		client.port = CoinDefaultPort
	}
	client.client = &http.Client{}
	//client.client.Timeout = time.Duration(RPCTimeOut) * time.Second
	client.client.Transport = &http.Transport{
		Dial: func(netw, addr string) (net.Conn, error) {
			deadline := time.Now().Add(30 * time.Second)
			c, err := net.DialTimeout(netw, addr, time.Second*30)
			if err != nil {
				return nil, err
			}
			c.SetDeadline(deadline)
			return c, nil
		},
	}
	//first access
	if _, err = client.Call("getinfo"); err != nil {
		return nil, err
	}
	if client.Status != http.StatusOK || client.LastError != nil {
		return nil, client.LastError
	}
	return client, nil
}

//SetSSL    setup certificate
func (client *RpcClient) SetSSL(certificate string) {
	client.proto = "https"
	client.certificate = certificate
}

func (client *RpcClient) access(data map[string]interface{}) (err error) {
	if len(data) != 2 {
		err = errors.New("params count error")
		return
	}
	if client.client == nil {
		err = errors.New("http client error")
		return
	}
	client.id++
	data["id"] = client.id
	client.LastError = nil
	client.responseData = nil
	client.Status = http.StatusOK
	var (
		jbuf []byte
		req  *http.Request
		resp *http.Response
	)
	if jbuf, err = json.Marshal(data); err != nil {
		return
	}

	addr := client.proto + "://" + client.host + ":" + strconv.Itoa(client.port)
	if req, err = http.NewRequest("POST", addr, bytes.NewReader(jbuf)); err != nil {
		client.LastError = err
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(client.username, client.password)
	//todo: setup ssl
	if resp, err = client.client.Do(req); err != nil {
		client.LastError = err
		return
	}
	client.Status = resp.StatusCode
	if client.Status != http.StatusOK {
		err = errors.New(resp.Status)
		return
	}
	defer resp.Body.Close()
	if client.responseData, err = ioutil.ReadAll(resp.Body); err != nil {
		client.LastError = err
		return
	}
	if len(client.responseData) == 0 {
		err = errors.New("response data is empty")
		return
	}
	return
}

//Call run RPC command
func (client *RpcClient) Call(method string, args ...interface{}) (data []byte, err error) {
	if method == "" {
		err = errors.New("method is not set")
		return
	}
	requestData := make(map[string]interface{})
	requestData["method"] = method
	requestData["params"] = args
	if err = client.access(requestData); err == nil {
		data = client.responseData
	}
	return
}

func (client *RpcClient) ListTransactions(account string, count, offset int64) (t wallet.TransactionAccountRecords, err error) {
	if b, err := client.Call("listtransactions", account, count, offset); err != nil {
		return t, err
	} else {
		var trans wallet.Transactions
		if err := json.Unmarshal(b, &trans); err != nil {
			return t, err
		}
		return trans.Result, nil
	}
}

func (client *RpcClient) GetBalance() (float64) {
	if b, err := client.Call("getbalance"); err != nil {
		return 0
	} else {
		var balance wallet.Balance
		if err := json.Unmarshal(b, &balance); err != nil {
			return 0
		}
		return balance.Result
	}
}

func (client *RpcClient) SendToaddress(address string, amount float64) (txid string, err error) {
	if b, err := client.Call("sendtoaddress", address, amount); err != nil {
		return "", err
	} else {
		var txid wallet.SendTxid
		if err := json.Unmarshal(b, &txid); err != nil {
			return "", err
		}
		return txid.Result, nil
	}
}
