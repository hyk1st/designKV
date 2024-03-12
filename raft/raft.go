package raft

import (
	"fmt"
	"github.com/lni/dragonboat/v4"
	"github.com/lni/dragonboat/v4/config"
	"os"
	"path/filepath"
	"raftKV/stateMachine"
)

var Raft *dragonboat.NodeHost

func StartRaftNode(nodeID, shardId uint64, addr string, initMembers map[uint64]string, join bool, stop chan struct{}, msg chan string) {
	conf := config.Config{
		ReplicaID:          nodeID,
		ShardID:            shardId,
		ElectionRTT:        10,
		HeartbeatRTT:       1,
		CheckQuorum:        true,
		SnapshotEntries:    0,
		CompactionOverhead: 5,
	}
	walDir := filepath.Join("data", fmt.Sprintf("node%d", nodeID), "wal")
	dataDir := filepath.Join("data", fmt.Sprintf("node%d", nodeID), "data")
	nhc := config.NodeHostConfig{
		WALDir:         walDir,
		NodeHostDir:    dataDir,
		RTTMillisecond: 200,
		RaftAddress:    addr,
	}
	var err error
	Raft, err = dragonboat.NewNodeHost(nhc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "新建raft结点错误，%t\n", err)
		os.Exit(1)
	}
	err = Raft.StartOnDiskReplica(initMembers, join, stateMachine.NewDiskKV, conf)
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "启动raft结点错误，%t\n", err)
		os.Exit(1)
	}

}
