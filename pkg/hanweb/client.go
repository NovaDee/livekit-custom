package hanweb

import (
	"bytes"
	"github.com/golang/protobuf/proto"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/livekit/livekit-server/pkg/config"
	"github.com/livekit/protocol/logger"
	"net/http"
	"strconv"
	"time"
)

var LH *LiveHook

type LiveHook struct {
	dataChannel chan *Data
	client      *retryablehttp.Client
	url         string
	init        bool
}

// NewWebhookModule 初始化WebhookModule实例
func NewWebhookModule(conf config.HanWebCustomConfig) *LiveHook {
	if !conf.CustomHook.Enabled {
		return nil
	}
	return &LiveHook{
		dataChannel: make(chan *Data, 1000),
		client:      newClient(),
		url:         conf.CustomHook.URL,
		init:        conf.CustomHook.Enabled,
		// 初始化数据通道
		// 其他初始化逻辑
	}
}

// NewClient 自定义retryHttpClient
func newClient() *retryablehttp.Client {
	client := retryablehttp.NewClient()
	// 设置 Request 拦截器，在每次请求时在请求头塞入请求时间
	client.RequestLogHook = func(logger retryablehttp.Logger, req *http.Request, retry int) {
		req.Header.Set("StartTime", strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10))
	}
	// 设置 Response 拦截器，在每次请求完成后记录结束时间，并计算耗时
	client.ResponseLogHook = func(logger retryablehttp.Logger, resp *http.Response) {
		startTimeStr := resp.Request.Header.Get("StartTime")
		parseInt, err := strconv.ParseInt(startTimeStr, 10, 64)
		if err == nil {
			responseTime := time.Now().UnixNano() / int64(time.Millisecond)
			i := responseTime - parseInt
			resp.Header.Set("CostTime", strconv.FormatInt(i, 10))
		}
	}
	return client
}

// Start 启动WebhookModule
func Start(conf *config.Config) *LiveHook {
	logger.Infow("hanweb module ", "enable", conf.Hanweb.CustomHook.Enabled, "url", conf.Hanweb.CustomHook.URL)
	module := NewWebhookModule(conf.Hanweb)
	if module == nil {
		return nil
	}
	LH = module
	LH.ConsumeData()
	return LH
}

// ConsumeData 数据消费逻辑，从数据通道中读取数据并执行消息推送
func (LH *LiveHook) ConsumeData() {
	var data *Data
	go func() {
		for {
			select {
			case data = <-LH.dataChannel:
				_ = LH.postMessage(data)
			}
		}
	}()
}

// SendData 将数据发送到数据通道
func SendData(data *Data) {
	if LH != nil && LH.init {
		LH.dataChannel <- data
	}
}

func (wm *LiveHook) postMessage(data *Data) error {
	// 使用 proto.Marshal 将 data 序列化为二进制数据
	encoded, err := proto.Marshal(data)
	if err != nil {
		return err
	}
	r, err := retryablehttp.NewRequest("POST", wm.url, bytes.NewReader(encoded))
	if err != nil {
		// ignore and continue
		return err
	}
	r.Header.Set(authHeader, "")
	r.Header.Set("content-type", "application/protobuf")
	res, err := wm.client.Do(r)
	if err != nil {
		return err
	}
	status := res.StatusCode
	if status == 200 && data.Logger {
		logger.Infow("send livehook url: ", "cost: ", res.Header.Get("CostTime")+"ms", "url: ", wm.url, "status: ", status, "event: ", data.Event)
	}
	_ = res.Body.Close()
	return nil
}
