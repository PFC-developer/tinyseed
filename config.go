package main

import (
	"fmt"
	"github.com/Entrio/subenv"
	"path/filepath"
)

// Config defines the configuration format
type Config struct {
	ListenAddress       string `toml:"laddr" comment:"Address to listen for incoming connections"`
	ChainID             string `toml:"chain_id" comment:"network identifier (todo move to cli flag argument? keeps the config network agnostic)"`
	NodeKeyFile         string `toml:"node_key_file" comment:"path to node_key (relative to tendermint-seed home directory or an absolute path)"`
	AddrBookFile        string `toml:"addr_book_file" comment:"path to address book (relative to tendermint-seed home directory or an absolute path)"`
	AddrBookStrict      bool   `toml:"addr_book_strict" comment:"Set true for strict routability rules\n Set false for private or local networks"`
	MaxNumInboundPeers  int    `toml:"max_num_inbound_peers" comment:"maximum number of inbound connections"`
	MaxNumOutboundPeers int    `toml:"max_num_outbound_peers" comment:"maximum number of outbound connections"`
	Seeds               string `toml:"seeds" comment:"seed nodes we can use to discover peers"`
}

// DefaultConfig returns a seed config initialized with default values
func DefaultConfig(basePath string) *Config {
	return &Config{
		ListenAddress:       fmt.Sprintf("tcp://0.0.0.0:%d", subenv.EnvI("LISTEN_PORT", 6969)),
		ChainID:             subenv.Env("CHAIN_ID", "osmosis-1"),
		NodeKeyFile:         filepath.Join(basePath, "config/node_key.json"),
		AddrBookFile:        filepath.Join(basePath, "data/addrbook.json"),
		AddrBookStrict:      subenv.EnvB("ADDR_STRICT", true),
		MaxNumInboundPeers:  subenv.EnvI("MAX_INBOUND", 1000),
		MaxNumOutboundPeers: subenv.EnvI("MAX_OUTBOUND", 1000),
		Seeds:               subenv.Env("SEEDS", "1b077d96ceeba7ef503fb048f343a538b2dcdf1b@136.243.218.244:26656,2308bed9e096a8b96d2aa343acc1147813c59ed2@3.225.38.25:26656,085f62d67bbf9c501e8ac84d4533440a1eef6c45@95.217.196.54:26656,f515a8599b40f0e84dfad935ba414674ab11a668@osmosis.blockpane.com:26656"),
	}
}
