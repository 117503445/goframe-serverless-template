package misc

import (
	"context"
	"database/sql"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"

	_ "github.com/go-sql-driver/mysql"
)

func ExecSQL(ctx context.Context) error {
	_, err := g.Cfg().Get(ctx, "database.link")
	if err != nil {
		return err
	}

	sqlDB, err := sql.Open("mysql", "root:MYSQL_PASSWORD@tcp(db:3306)/goframe_serverless_template?multiStatements=true")
	defer sqlDB.Close()

	if err != nil {
		return err
	}

	sql := gfile.GetContents("../manifest/sql/create.sql")

	if _, err = sqlDB.Exec(sql); err != nil {
		return err
	}

	return nil
}
