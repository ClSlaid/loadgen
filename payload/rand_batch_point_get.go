package payload

import (
	"fmt"
	"github.com/crazycs520/loadgen/cmd"
	"github.com/crazycs520/loadgen/config"
	"github.com/spf13/cobra"
	"math/rand"
	"strconv"
	"strings"
)

type RandBatchPointGetSuite struct {
	*basicQuerySuite
	batchSize int
}

func (c *RandBatchPointGetSuite) Name() string {
	return randBatchPointGetSuiteName
}

func (c *RandBatchPointGetSuite) GenQuerySQL() string {
	vs := make([]string, 0, c.batchSize)
	for i := 0; i < c.batchSize; i++ {
		n := rand.Intn(c.rows)
		vs = append(vs, strconv.Itoa(n))
	}
	return fmt.Sprintf("select * from %v where a in (%v)", c.tblInfo.DBTableName(), strings.Join(vs, ","))
}

func NewRandBatchPointGetSuite(cfg *config.Config) cmd.CMDGenerater {
	suite := &RandBatchPointGetSuite{
		batchSize: 100,
	}
	basic := NewBasicQuerySuite(cfg, suite)
	suite.basicQuerySuite = basic
	return suite
}

func (c *RandBatchPointGetSuite) Cmd() *cobra.Command {
	cmd := c.basicQuerySuite.Cmd()
	cmd.Flags().IntVarP(&c.batchSize, flagBatchSize, "", 100, "the batch size of batch point get")
	return cmd
}

func (c *RandBatchPointGetSuite) ParseCmd(combinedCmd string) bool {
	return ParsePayloadCmd(combinedCmd, c.querySuite.Name(), func(flag, value string) error {
		switch flag {
		case flagRows:
			v, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			c.rows = v
		case flagAgg:
			v, err := strconv.ParseBool(value)
			if err != nil {
				return err
			}
			c.agg = v
		case flagBatchSize:
			v, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			c.batchSize = v
		default:
			return fmt.Errorf("unknow flag %v", flag)
		}
		return nil
	})
}
