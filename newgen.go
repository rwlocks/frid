package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

var (
	CheckSumBits int8 = 4
	NewNodeBits  int8 = 8
	NewStepBits  int8 = 10

	//the timestamp of 2010-07-02 09:30:00
	MyTimeStamp int64 = 1278063000000
)

//snowflake mode
//+--------------------------------------------------------------------------+
//| 1 Bit Unused | 41 Bit Timestamp |  10 Bit NodeID  |   12 Bit Sequence ID |
//+--------------------------------------------------------------------------+
//
//new mode
//+--------------------------------------------------------------------------------------------+
//| 1 Bit Unused | 41 Bit Timestamp |  8 Bit NodeID  |   10 Bit Sequence ID |  4 Bit CheckSum  |
//+--------------------------------------------------------------------------------------------+

type IDService struct {
	Node int64
	Step int64

	SfNode *snowflake.Node
}

func (is *IDService)Init(n int64) error  {

	snowflake.Epoch = MyTimeStamp

	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(n)
	if err != nil {
		fmt.Println(err)
		return err
	}

	is.SfNode = node

	return nil
}


func (is *IDService)GenerateNewID() (int64, error){
	// Generate a snowflake ID.
	ID := is.SfNode.Generate()
	id := ID.Int64()

	timeShift := snowflake.NodeBits + snowflake.StepBits

	nodeNum := ID.Node()
	stepNum := ID.Step()

	cs := CheckSum(id)

	id = (id & (-1 << timeShift)) | nodeNum << (NewStepBits + CheckSumBits) | stepNum << CheckSumBits | cs

	return id, nil
}

func CheckSum(id int64) int64 {

	// 4 bit checksum
	return (id >> CheckSumBits) % 16
}

func Valid(cs int64, id int64) bool {
	return cs == CheckSum(id)
}



