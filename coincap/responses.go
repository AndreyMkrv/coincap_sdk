package coincap

import "fmt"

type assetResponse struct {
	Asset     Asset `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

type assetsResponse struct {
	Assets    []Asset `json:"data"`
	Timestamp int64   `json:"timestamp"`
}

type Asset struct {
	ID            string `json:"id"`
	Rank          string `json:"rank"`
	Symbol        string `json:"symbol"`
	Name          string `json:"name"`
	Supply        string `json:"supply"`
	MaxSupply     string `json:"maxsupply"`
	MarketCapUsd  string `json:"marketCupUsd"`
	VolumeUsd24Hr string `json:"volumeUsd24Hr"`
	PriceUsd      string `json:"priceUsd"`
}

func (d Asset) Info() string {
	return fmt.Sprintf("[ID] %s\n[RANK] %s\n[Symbol] %s\n[Name] %s\n[Supply] %s\n[MaxSupply] %s\n[MarketCupUsd] %s\n[VolumeUsd24Hr] %s\n[PriceUsd] %s\n\n ", d.ID, d.Rank, d.Symbol, d.Name, d.Supply, d.MaxSupply, d.MarketCapUsd, d.VolumeUsd24Hr, d.PriceUsd)
}
