package mysql

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type Client struct {
	Host     string
	DB       string
	Username string
	Port     string
	Password string
}

var mysqlClient *Client
var mysqlClientOnce sync.Once

func GetClientClient() *Client {
	mysqlClientOnce.Do(func() {
		mysqlClient = &Client{
			Host:     os.Getenv("MYSQL_HOST"),
			DB:       os.Getenv("MYSQL_DATABASE"),
			Username: os.Getenv("MYSQL_USER"),
			Port:     "3306",
			Password: os.Getenv("MYSQL_PASSWORD"),
		}
	})
	return mysqlClient
}

func (c *Client) getConnectionString(user, pass, host, port, db string) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&multiStatements=True", user, pass, host, port, db,
	)
}

func (c *Client) getConnection() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", c.getConnectionString(c.Username, c.Password, c.Host, c.Port, string(c.DB)))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get mysql connection")
	}
	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "failed to ping mysql server")
	}
	return db, nil
}

func (c *Client) ExecuteQuery(query string) ([][]interface{}, []string, error) {
	db, err := c.getConnection()
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to establish connection to mysql")
	}
	rows, err := db.Query(query)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to execute query")
	}
	defer rows.Close()
	defer db.Close()

	columns, _ := rows.Columns()
	count := len(columns)
	var results [][]interface{}

	for rows.Next() {
		values := make([]interface{}, count)
		valuePtrs := make([]interface{}, count)
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		rows.Scan(valuePtrs...)
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to scan rows")
		}
		results = append(results, values)
	}
	err = rows.Err()
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to scan rows")
	}
	return results, columns, nil
}
