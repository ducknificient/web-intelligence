package datastore

import (
	"context"
	"errors"
	"fmt"

	"github.com/ducknificient/web-intelligence/go/config"
	"github.com/ducknificient/web-intelligence/go/logger"
	"github.com/jackc/pgx/v4/pgxpool"
)

func GetPgDataSource(host string, port string, user string, pwd string, dbname string, sslmode string, poolmaxcon string) string {
	return fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s pool_max_conns=%s",
		host, port, user, pwd, dbname, sslmode, poolmaxcon)
}

func NewPostgreSQLDB(ctx context.Context, logger logger.Logger) (pgdb *PostgresDB) {
	return &PostgresDB{
		Ctx:    ctx,
		Logger: logger,
	}
}

type PostgresDB struct {
	Ctx    context.Context
	Conn   *pgxpool.Pool
	Logger logger.Logger
	// Option     OracleOption
	PgInfo config.PgInfo
}

func GetListPgDatabase(logger logger.Logger) (listPostgresDB []PostgresDB) {

	// fmt.Printf("dblist : %#v\n\n", *general.Conf.DBList)
	for _, info := range *config.Conf.PgList {
		cicDB := PostgresDB{Logger: logger, PgInfo: info}
		listPostgresDB = append(listPostgresDB, cicDB)
	}

	return listPostgresDB
}

func (db *PostgresDB) Connect() (err error) {

	/*Connect to Postgresql*/
	db.Conn, err = pgxpool.Connect(db.Ctx, GetPgDataSource(
		*db.PgInfo.Host,
		*db.PgInfo.Port,
		*db.PgInfo.Username,
		*db.PgInfo.Password,
		*db.PgInfo.DbName,
		*db.PgInfo.SSLMode,
		*db.PgInfo.PoolMaxConn,
	))
	if err != nil {
		db.Logger.Info(fmt.Sprintf("Unable connect main postgresql database:" + err.Error()))
	}

	db.Logger.Info("connect main postgresql database success")

	return err
}

func (db *PostgresDB) StoreD(pagesource string, link string, task string) (err error) {

	// Prepare SQL statement to insert data
	sqlStatement := `INSERT INTO webintelligence.crawlpage (pagesource, link, task) VALUES ($1, $2, $3)`
	_, err = db.Conn.Exec(db.Ctx, sqlStatement, pagesource, link, task)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to insert into webintelligence.crawlpage. q: %v, param: %v,%v,%v . err: %#v", sqlStatement, pagesource, link, task, err.Error())
		err = errors.New(errMsg)
		return err
	}

	return err

}

func (db *PostgresDB) StorePdf(filename string, link string, task string, pdf []byte) (err error) {

	// Prepare SQL statement to insert data
	sqlStatement := `INSERT INTO webintelligence.crawlpage (pagesource, link, task, document) VALUES ($1, $2, $3, $4)`
	_, err = db.Conn.Exec(db.Ctx, sqlStatement, filename, link, task, pdf)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to insert into webintelligence.crawlpage. q: %v, param: %v,%v,%v,%v . err: %#v", sqlStatement, filename, link, task, pdf, err.Error())
		err = errors.New(errMsg)
		return err
	}

	return err

}

func (db *PostgresDB) StoreE(link string, href string, task string) (err error) {

	// Prepare SQL statement to check if data exists
	sqlStatement := "SELECT COUNT(1) FROM webintelligence.crawlhref WHERE link = $1 AND href = $2"
	row := db.Conn.QueryRow(db.Ctx, sqlStatement, link, href)
	var count int
	err = row.Scan(&count)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to select from webintelligence.crawlhref. q: %v, param: %v,%v . err: %#v", sqlStatement, link, href, err.Error())
		err = errors.New(errMsg)
		return err
	}

	// If data does not exist, insert into tableE
	if count <= 0 {
		sqlStatement = `INSERT INTO webintelligence.crawlhref (link, href, task) VALUES ($1, $2, $3)`
		_, err := db.Conn.Exec(db.Ctx, sqlStatement, link, href, task)
		if err != nil {
			errMsg := fmt.Sprintf("Unable to insert into webintelligence.crawlhref. q: %v, param: %v,%v . err: %#v", sqlStatement, link, href, err.Error())
			err = errors.New(errMsg)
			return err
		}
	}

	return err
}

func (db *PostgresDB) ContainsD(link string) (contains bool, err error) {

	// Prepare SQL statement to check if data exists in tableD
	row := db.Conn.QueryRow(db.Ctx, "SELECT COUNT(*) FROM webintelligence.crawlpage WHERE link = $1", link)
	var count int
	err = row.Scan(&count)
	if err != nil {
		return false, err
	}

	// If count > 0, data exists in tableD
	if count > 0 {
		return true, nil
	}

	return false, err
}
