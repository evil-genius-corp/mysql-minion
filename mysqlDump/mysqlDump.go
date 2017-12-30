package mysqlDump

import "mysql-minion/utils"

// MysqlDump wraps the mysqdump cmd
func MysqlDump(arg string) string {
	mysqlDumpBin := utils.Which("mysqldump")
	return utils.RunCmd(mysqlDumpBin, arg)
}
