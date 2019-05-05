package main

import (
	"github.com/eosforce/bus-service/force-relay/cfg"
	"github.com/eosforce/bus-service/force-relay/chainhandler"
	"github.com/eosforce/bus-service/force-relay/logger"
	"github.com/eosforce/bus-service/force-relay/relay"
	"github.com/eosforce/bus-service/force-relay/side"
	"github.com/fanyang1988/force-block-ev/blockdb"
	"github.com/fanyang1988/force-go/p2p"
	"github.com/pkg/errors"
)

func startSideService() {
	// frome side need to commit block to relay
	chainCfgs, _ := cfg.GetChainCfg("relay")

	_, p2ps := cfg.GetChainCfg("side")
	chainTyp := cfg.GetChainTyp("side")

	side.InitCommitWorker(chainCfgs, cfg.GetTransfers())

	// for p2p chain id
	info, err := relay.Client().GetInfoData()
	if err != nil {
		panic(errors.New("get info err"))
	}

	lastCommitted, err := side.GetLastCommittedBlock()
	if err != nil {
		logger.Errorf("err by %s", err.Error())
		panic(errors.New("GetLastCommittedBlock info err"))
	}

	logger.Debugf("get last committed block %v", lastCommitted)

	lastNum := lastCommitted.Num
	if lastNum > 3 {
		lastNum -= 2
	}

	var syncData *p2p.P2PSyncData
	if false {
		lastBlockData, err := relay.Client().GetBlockDataByNum(lastNum)
		if err != nil {
			panic(errors.Errorf("get block num %d err by %s", lastNum, err.Error()))
		}
		lastIrBlockData, err := relay.Client().GetBlockDataByNum(lastNum - 30)
		if err != nil {
			panic(errors.Errorf("get block num %d err by %s", lastNum-30, err.Error()))
		}
		syncData = &p2p.P2PSyncData{
			HeadBlockNum:             lastBlockData.BlockNum,
			HeadBlockID:              lastBlockData.ID,
			HeadBlockTime:            lastBlockData.Timestamp,
			LastIrreversibleBlockNum: lastIrBlockData.BlockNum,
			LastIrreversibleBlockID:  lastIrBlockData.ID,
		}
	}

	p2pPeers := p2p.NewP2PClient(chainTyp, p2p.P2PInitParams{
		Name:       "relay",
		ClientID:   info.ChainID.String(),
		StartBlock: syncData,
		Peers:      p2ps,
		Logger:     logger.Logger(),
	})

	p2pPeers.RegHandler(&handlerImp{
		verifier: blockdb.NewFastBlockVerifier(p2ps, 0, chainhandler.NewChainHandler(
			func(block *chainhandler.Block, actions []chainhandler.Action) {
				side.HandSideBlock(block, actions)
			}, chainTyp)),
	})
	p2pPeers.Start()
}
