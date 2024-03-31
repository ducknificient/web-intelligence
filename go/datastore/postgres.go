package datastore

import (
	"context"
	"errors"
	"fmt"

	configpackage "github.com/ducknificient/web-intelligence/go/config"
	"github.com/ducknificient/web-intelligence/go/logger"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresDB struct {
	ctx    context.Context
	config configpackage.Configuration
	logger logger.Logger

	pginfo configpackage.PgInfo
	conn   *pgxpool.Pool
}

func GetListPgModel(ctx context.Context, config configpackage.Configuration, logger logger.Logger) (listPostgresModel []*PostgresDB) {

	for _, pginfo := range *config.GetConfiguration().PgList {
		// db := configpackage.PgInfo{Logger: logger, PgInfo: info}

		db := &PostgresDB{
			ctx:    ctx,
			config: config,
			logger: logger,
			pginfo: pginfo,
		}

		listPostgresModel = append(listPostgresModel, db)
	}

	return listPostgresModel
}

func GetPgDataSource(host string, port string, user string, pwd string, dbname string, sslmode string, poolmaxcon string) string {
	return fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s pool_max_conns=%s",
		host, port, user, pwd, dbname, sslmode, poolmaxcon)
}

func NewPostgresModelList(ctx context.Context, config configpackage.Configuration, logger logger.Logger) (mapPgModel map[string]*PostgresDB, err error) {

	listPgModel := GetListPgModel(ctx, config, logger)
	mapPgModel = make(map[string]*PostgresDB)

	for _, postgresModel := range listPgModel {

		err := postgresModel.Connect()
		if err != nil {
			logger.Error(fmt.Sprintf("unable to connect : %v", err.Error()))
			// panic(err.Error())
			return mapPgModel, err
		}

		err = postgresModel.conn.Ping(ctx)
		if err != nil {
			logger.Error(fmt.Sprintf("unable to ping : %v", err.Error()))
			// panic(err.Error())
			return mapPgModel, err
		}

		fmt.Printf("ping success : %#v\n", postgresModel.pginfo.DbName)

		mapPgModel[*postgresModel.pginfo.DbName] = postgresModel

	}

	return mapPgModel, err
}

func (m *PostgresDB) Connect() (err error) {

	// pginfo := m.config.GetConfiguration().Pg

	/*Connect to Postgresql*/
	m.conn, err = pgxpool.Connect(m.ctx, GetPgDataSource(
		*m.pginfo.Host,
		*m.pginfo.Port,
		*m.pginfo.Username,
		*m.pginfo.Password,
		*m.pginfo.DbName,
		*m.pginfo.SSLMode,
		*m.pginfo.PoolMaxConn,
	))
	if err != nil {
		// db.Logger.Info(fmt.Sprintf("Unable connect main postgresql database:" + err.Error()))
		m.logger.Info(fmt.Sprintf("Unable connect main postgresql database:" + err.Error()))
	}

	m.logger.Info("connect main postgresql database success")

	return err
}

func (d *PostgresDB) StoreD(pagesource string, link string, task string, documenttype string, mimetype string) (err error) {

	// Prepare SQL statement to insert data
	sqlStatement := `INSERT INTO webintelligence.crawlpage (pagesource, link, task, documenttype, mimetype) VALUES ($1, $2, $3, $4, $5)`
	_, err = d.conn.Exec(d.ctx, sqlStatement, pagesource, link, task, documenttype, mimetype)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to insert into webintelligence.crawlpage. q: %v, param: %v,%v,%v,%v,%v . err: %#v", sqlStatement, pagesource, link, task, documenttype, mimetype, err.Error())
		err = errors.New(errMsg)
		return err
	}

	return err

}

func (d *PostgresDB) StoreDocument(link string, task string, documenttype string, document []byte, documentcontenttype string) (err error) {

	// Prepare SQL statement to insert data
	sqlStatement := `INSERT INTO webintelligence.crawlpage (link, task,documenttype,document, mimetype) VALUES ($1, $2, $3, $4, $5)`
	_, err = d.conn.Exec(d.ctx, sqlStatement, link, task, documenttype, document, documentcontenttype)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to insert into webintelligence.crawlpage. q: %v, param: %v,%v,%v,%v,%v . err: %#v", sqlStatement, link, task, documenttype, document, documentcontenttype, err.Error())
		err = errors.New(errMsg)
		return err
	}

	return err

}

func (d *PostgresDB) StoreE(link string, href string, task string) (err error) {

	// Prepare SQL statement to check if data exists
	sqlStatement := "SELECT COUNT(1) FROM webintelligence.crawlhref WHERE link = $1 AND href = $2"
	row := d.conn.QueryRow(d.ctx, sqlStatement, link, href)
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
		_, err := d.conn.Exec(d.ctx, sqlStatement, link, href, task)
		if err != nil {
			errMsg := fmt.Sprintf("Unable to insert into webintelligence.crawlhref. q: %v, param: %v,%v,%v . err: %#v", sqlStatement, link, href, task, err.Error())
			err = errors.New(errMsg)
			return err
		}
	}

	return err
}

func (d *PostgresDB) ContainsD(link string) (contains bool, err error) {

	// Prepare SQL statement to check if data exists in tableD
	row := d.conn.QueryRow(d.ctx, "SELECT COUNT(*) FROM webintelligence.crawlpage WHERE link = $1", link)
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
