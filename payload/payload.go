package payload

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/crazycs520/loadgen/cmd"
	"github.com/crazycs520/loadgen/config"
	"github.com/crazycs520/loadgen/util"
	"strings"
)

func init() {
	cmd.RegisterCaseCmd(NewFullTableScanSuite)
	cmd.RegisterCaseCmd(NewFullIndexScanSuite)
	cmd.RegisterCaseCmd(NewFullIndexLookUpSuite)
	cmd.RegisterCaseCmd(NewRandPointGetSuite)
	cmd.RegisterCaseCmd(NewRandBatchPointGetSuite)
	cmd.RegisterCaseCmd(NewFixPointGetSuite)
	cmd.RegisterCaseCmd(NewWriteAutoIncSuite)

	cmd.RegisterCaseCmd(NewGenStmtSuite)
	cmd.RegisterCaseCmd(NewPointGetForUpdateGetSuite)
	cmd.RegisterCaseCmd(NewIndexLookupForUpdateSuite)
	cmd.RegisterCaseCmd(NewWriteHotSuite)
	cmd.RegisterCaseCmd(NewNormalOLTPSuite)
	cmd.RegisterCaseCmd(NewWriteConflictSuite)
}

// ParsePayloadCmd return true if the combined cmd is valid, otherwise, return false.
func ParsePayloadCmd(combinedCmd string, payloadName string, fn func(flag, value string) error) bool {
	cmds := strings.Split(combinedCmd, symbolSeparator)
	if len(cmds) == 0 || cmds[0] != payloadName {
		return false
	}
	for i := 1; i < len(cmds); i++ {
		fields := strings.Split(cmds[i], symbolAssignment)
		if len(fields) != 2 {
			fmt.Printf("cmd %v is invalid, the valid format is like this `rows=1000`\n", cmds[i])
			return false
		}
		err := fn(fields[0], fields[1])
		if err != nil {
			fmt.Printf("parse cmd %v for payload %v, please check the flag of payload %v, err: %v\n",
				cmds[i], payloadName, payloadName, err)
			return false
		}
	}
	return true
}

func execSQLLoop(ctx context.Context, cfg *config.Config, genSQL func() string) error {
	db := util.GetSQLCli(cfg)
	defer func() {
		db.Close()
	}()
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}
		sql := genSQL()
		err := execSQL(db, sql)
		if err != nil {
			return err
		}
	}
}

func execSQL(db *sql.DB, sql string) error {
	if strings.HasPrefix(strings.ToLower(sql), "select") {
		rows, err := db.Query(sql)
		if err != nil {
			return err
		}
		for rows.Next() {
		}
		return rows.Close()
	}

	_, err := db.Exec(sql)
	return err
}
