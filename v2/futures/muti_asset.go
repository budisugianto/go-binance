package futures

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// MultiAssetService create order
type MultiAssetService struct {
	c                 *Client
	multiAssetsMargin bool
}

// ReduceOnly set reduceOnly
func (s *MultiAssetService) MultiAssetsMargin(multiAsset bool) *MultiAssetService {
	s.multiAssetsMargin = multiAsset
	return s
}

func (s *MultiAssetService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, header *http.Header, err error) {

	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"multiAssetsMargin": strconv.FormatBool(s.multiAssetsMargin),
	}

	r.setFormParams(m)
	data, header, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	return data, header, nil
}

// Do send request
func (s *MultiAssetService) Do(ctx context.Context, opts ...RequestOption) (res *MultiAssetResponse, err error) {
	data, header, err := s.createOrder(ctx, "/fapi/v1/multiAssetsMargin", opts...)
	if err != nil {
		return nil, err
	}
	res = new(MultiAssetResponse)
	err = json.Unmarshal(data, res)
	res.RateLimitOrder10s = header.Get("X-Mbx-Order-Count-10s")
	res.RateLimitOrder1m = header.Get("X-Mbx-Order-Count-1m")

	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateOrderResponse define create order response
type MultiAssetResponse struct {
	Code              int64  `json:"code"`
	Msg               string `json:"msg"`
	RateLimitOrder10s string `json:"rateLimitOrder10s,omitempty"`
	RateLimitOrder1m  string `json:"rateLimitOrder1m,omitempty"`
}
