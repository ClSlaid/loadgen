package payload

import (
	"github.com/crazycs520/loadgen/cmd"
	"github.com/crazycs520/loadgen/config"
	"github.com/crazycs520/loadgen/data"
)

type WriteAutoIncSuite struct {
	*basicWriteSuite
}

func (c *WriteAutoIncSuite) Name() string {
	return writeAutoIncSuiteName
}

func (c *WriteAutoIncSuite) UpdateTableDef(_ *data.TableInfo) {
}

func NewWriteAutoIncSuite(cfg *config.Config) cmd.CMDGenerater {
	suite := &WriteAutoIncSuite{}
	basic := NewBasicWriteSuite(cfg, suite)
	suite.basicWriteSuite = basic
	return suite
}
