package side

import (
	"fmt"

	"github.com/eosforce/bus-service/force-relay/chainhandler"

	eos "github.com/eosforce/goeosforce"
	"github.com/fanyang1988/force-go"
	"github.com/pkg/errors"
)

// client client to force relay chain
var client *force.Client

// CreateClient create client to force relay chain
func CreateClient(configPath string) {
	var err error
	client, err = force.NewClientFromFile(configPath)
	if err != nil {
		fmt.Println("create client error  ", err.Error())
		panic(err)
		return
	}
}

type lastCommitBlockInfo struct {
	Chain eos.Name           `json:"chain"`
	Last  chainhandler.Block `json:"last"`
}

// GetLastCommittedBlock get last committed block to relay chain
func GetLastCommittedBlock() (*chainhandler.Block, error) {
	req := eos.GetTableRowsRequest{
		Code:  "force.relay",
		Scope: string(cfg.Chain),
		Table: "relaystat",
	}

	res, err := client.GetTableRows(req)
	if err != nil {
		return nil, err
	}

	rspBlock := make([]lastCommitBlockInfo, 0, 32)
	err = res.BinaryToStructs(&rspBlock)
	if err != nil {
		return nil, err
	}

	if len(rspBlock) == 0 {
		return nil, errors.New("rsp block info no find")
	}

	return &rspBlock[0].Last, nil
}
