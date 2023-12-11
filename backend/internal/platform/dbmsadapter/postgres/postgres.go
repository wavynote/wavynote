package postgres

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/lib/pq"

	"github.com/wavynote/internal/platform/dbmsadapter"
)

type dsn struct {
	host     string
	port     int
	user     string
	password string
	dbName   string
	sslMode  string
	appName  string
}

type DbmsadapterService struct {
	db   *sql.DB
	conn dsn
}

func NewService(host string, port int, user, password, dbname, sslmode, appname string) *DbmsadapterService {
	dsn := dsn{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbName:   dbname,
		sslMode:  sslmode,
		appName:  appname,
	}

	service := &DbmsadapterService{conn: dsn}
	return service
}

func (d *DbmsadapterService) Open() error {
	var err error
	if d.db != nil {
		if err = d.db.Ping(); err != nil {
			return err
		}

		return nil
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s application_name=%s",
		d.conn.host,
		d.conn.port,
		d.conn.user,
		d.conn.password,
		d.conn.dbName,
		d.conn.sslMode,
		d.conn.appName)

	if d.db, err = sql.Open("postgres", dsn); err != nil {
		return err
	}

	if err = d.db.Ping(); err != nil {
		return err
	}

	return nil
}

func (d *DbmsadapterService) Close() error {
	if d.db != nil {
		err := d.db.Close()
		if err != nil {
			return err
		}
		d.db = nil
	}
	return nil
}

func (d *DbmsadapterService) ExecuteQuery(query string) (int64, error) {
	res, err := d.db.Exec(query)
	if err != nil {
		return -1, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}

	return count, nil
}

func (d *DbmsadapterService) SelectQuery(query string) (dbmsadapter.SelectQueryResultType, error) {
	rows, err := d.db.Query(query)
	if err != nil {
		return nil, err
	}

	cols, err := rows.Columns()
	if err != nil {
		fmt.Printf("db columns fail: %s\n", err)
		return nil, err
	}

	var result dbmsadapter.SelectQueryResultType

	for rows.Next() {
		values := make([]interface{}, len(cols))
		pointers := make([]interface{}, len(cols))

		for i := range values {
			pointers[i] = &values[i]
		}

		if err = rows.Scan(pointers...); err != nil {
			fmt.Printf("db scan fail: %s\n", err)
			return nil, err
		}

		row := make(map[string]interface{})
		for i, val := range values {
			row[strings.ToLower(cols[i])] = val
		}

		result = append(result, row)
	}

	return result, nil
}

func (d *DbmsadapterService) BeginTx() (*sql.Tx, error) {
	tx, err := d.db.Begin()
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (d *DbmsadapterService) QueryTx(tx *sql.Tx, query string) (dbmsadapter.SelectQueryResultType, error) {
	rows, err := tx.Query(query)
	if err != nil {
		return nil, err
	}

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var result dbmsadapter.SelectQueryResultType

	for rows.Next() {
		values := make([]interface{}, len(cols))
		pointers := make([]interface{}, len(cols))

		for i := range values {
			pointers[i] = &values[i]
		}

		if err = rows.Scan(pointers...); err != nil {
			return nil, err
		}

		row := make(map[string]interface{})
		for i, val := range values {
			row[strings.ToLower(cols[i])] = val
		}

		result = append(result, row)
	}

	return result, nil
}

func (d *DbmsadapterService) ExecTx(tx *sql.Tx, query string) (int64, error) {
	res, err := tx.Exec(query)
	if err != nil {
		return -1, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}

	return count, nil
}

func (d *DbmsadapterService) CommitTx(tx *sql.Tx) error {
	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (d *DbmsadapterService) RollbackTx(tx *sql.Tx) error {
	err := tx.Rollback()
	if err != nil {
		return err
	}
	return nil
}

func (d *DbmsadapterService) ExecuteTransaction(query []string) (string, error) {
	tx, err := d.db.Begin()
	if err != nil {
		return "", err
	}

	for _, q := range query {
		_, err = tx.Exec(q)
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				return q, rollbackErr
			}
			return q, nil
		}
	}

	err = tx.Commit()
	if err != nil {
		return "", err
	}

	return "", err
}

func (d *DbmsadapterService) GetString(col interface{}) string {
	if col == nil {
		return ""
	}

	result := fmt.Sprintf("%v", col)

	if strings.Contains(result, "'") == true {
		result = strings.Replace(result, "'", "''", -1)
	}

	return result
}

func (d *DbmsadapterService) GetInteger(col interface{}) int {
	val, err := strconv.Atoi(fmt.Sprintf("%v", col))
	if err != nil {
		return 0
	}
	return val
}

func (d *DbmsadapterService) GetBigInteger(col interface{}) int64 {
	val, err := strconv.ParseInt(fmt.Sprintf("%v", col), 10, 64)
	if err != nil {
		return 0
	}
	return val
}

func (d *DbmsadapterService) GetBytes(col interface{}) []byte {
	buf, ok := col.([]byte)
	if !ok {
		return nil
	}
	return buf
}

func (d *DbmsadapterService) GetUUID(col interface{}) string {
	if col == nil {
		return ""
	}

	result := fmt.Sprintf("%s", col)
	return result
}

func (d *DbmsadapterService) GetBoolean(col interface{}) bool {
	buf, ok := col.(bool)
	if !ok {
		return false
	}

	return buf
}

func (d *DbmsadapterService) CheckDBIntegrity() bool {
	return true
}
