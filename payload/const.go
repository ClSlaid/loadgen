package payload

const (
	normalOLTPSuiteName           = "normal-oltp"
	fullTableScanSuiteName        = "full-table-scan"
	fullIndexScanSuiteName        = "full-index-scan"
	fullIndexLookupSuiteName      = "full-index-lookup"
	randPointGetSuiteName         = "rand-point-get"
	randBatchPointGetSuiteName    = "rand-batch-point-get"
	fixPointGetSuiteName          = "fix-point-get"
	writeAutoIncSuiteName         = "write-auto-inc"
	pointGetForUpdateSuiteName    = "point-get-for-update"
	indexLookupForUpdateSuiteName = "index-lookup-for-update"
)

const (
	symbolSeparator  = ":"
	symbolAssignment = "="
	flagRows         = "rows"
	flagAgg          = "agg" // aggregation
	flagInsert       = "insert"
	flagUpdate       = "update"
	flagSelect       = "select"
	flagPointGet     = "point-get"
	flagIgnore       = "ignore" // ignore execute sql error
	flagBatchSize    = "batch-size"
	flagRowID        = "rowid"
	flagRandRowID    = "rand-rowid"
	flagColCnt       = "col-cnt"
)
