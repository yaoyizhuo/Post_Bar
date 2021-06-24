package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"time"
)

var node *snowflake.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-1-2", startTime)
	if err != nil {
		return
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)

	return
}

func GetID() int64  {
	return node.Generate().Int64()
}

//func main() {
//	if err:=Init("2021-4-8",1);err!=nil{
//		fmt.Println(err)
//		return
//	}
//	id := GetID()
//	fmt.Println(id)
//}