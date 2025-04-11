package flag

import (
	"os"
	"server/global"
	"strings"
)

func SQLImport(sqlPath string) (errs []error) {
	byteData, err := os.ReadFile(sqlPath)
	if err != nil {
		return append(errs, err)
	}
	sqlList := strings.Split(string(byteData), ";")
	for _, sql := range sqlList {
		sql = strings.TrimSpace(sql)
		if sql == "" {
			continue
		}
		err = global.DB.Exec(sql).Error
		if err != nil {
			errs = append(errs, err)
			continue
		}
	}
	return nil
}
