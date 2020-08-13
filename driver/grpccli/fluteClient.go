package grpccli

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"hcc/harp/action/grpc/rpcflute"
	"hcc/harp/lib/config"
	"hcc/harp/lib/logger"
	"strconv"
	"time"
)

var fluteConn *grpc.ClientConn

func initFlute() error {
	var err error

	addr := config.Flute.ServerAddress + ":" + strconv.FormatInt(config.Flute.ServerPort, 10)
	logger.Logger.Println("Trying to connect to flute module (" + addr + ")")
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(config.Flute.ConnectionTimeOutMs)*time.Millisecond)

	for i := 0; i < int(config.Flute.ConnectionRetryCount); i++ {
		fluteConn, err = grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			logger.Logger.Println("Failed to connect flute module ("+addr+"): %v", err)
			continue
		}

		RC.flute = rpcflute.NewFluteClient(fluteConn)
		logger.Logger.Println("gRPC client connected to flute module")

		return nil
	}

	return errors.New("retry count exceeded to connect flute module")
}

func cleanFlute() {
	_ = fluteConn.Close()
}

// GetNode : Get infos of the node
func (rc *RPCClient) GetNode(uuid string) (*rpcflute.Node, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	node, err := rc.flute.GetNode(ctx, &rpcflute.ReqGetNode{UUID: uuid})
	if err != nil {
		return nil, err
	}

	return node.Node, nil
}

// GetNodeList : Get the list of nodes by server UUID.
func (rc *RPCClient) GetNodeList(serverUUID string) ([]rpcflute.Node, error) {
	var nodeList []rpcflute.Node

	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	pnodeList, err := rc.flute.GetNodeList(ctx, &rpcflute.ReqGetNodeList{Node: &rpcflute.Node{ServerUUID: serverUUID}})
	if err != nil {
		return nil, err
	}

	for _, pnode := range pnodeList.Node {
		nodeList = append(nodeList, rpcflute.Node{
			UUID:        pnode.UUID,
			ServerUUID:  pnode.ServerUUID,
			BmcMacAddr:  pnode.BmcMacAddr,
			BmcIP:       pnode.BmcIP,
			PXEMacAddr:  pnode.PXEMacAddr,
			Status:      pnode.Status,
			CPUCores:    pnode.CPUCores,
			Memory:      pnode.Memory,
			Description: pnode.Description,
			Active:      pnode.Active,
			CreatedAt:   pnode.CreatedAt,
		})
	}

	return nodeList, nil
}
