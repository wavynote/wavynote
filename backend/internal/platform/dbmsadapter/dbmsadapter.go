package dbmsadapter

import "database/sql"

type SelectQueryResultType []map[string]interface{} // SELECT 쿼리 결과를 동적으로 처리하기 위해 사용하는 반환값 타입

type DbmsadapterService interface {
	// DB 세션 관리 함수 제공
	Open() error
	Close() error

	ExecuteQuery(string) (int64, error)                // (SELECT를 제외한)쿼리 수행을 위한 함수 제공
	SelectQuery(string) (SelectQueryResultType, error) // SELECT 쿼리 결과를 동적으로 처리하기 위한 함수 제공

	// Transaction 처리를 위한 함수 제공
	BeginTx() (*sql.Tx, error)
	ExecTx(*sql.Tx, string) (int64, error)                  // Transaction내에서 (SELECT를 제외한)쿼리 수행을 위한 함수
	QueryTx(*sql.Tx, string) (SelectQueryResultType, error) // Transaction내에서 SELECT 쿼리 결과를 처리하기 위한 함수
	CommitTx(*sql.Tx) error
	RollbackTx(*sql.Tx) error
	ExecuteTransaction([]string) (string, error)

	// 타입에 따른 칼럼 값 반환 함수 제공
	GetString(interface{}) string
	GetInteger(interface{}) int
	GetBigInteger(interface{}) int64
	GetBytes(interface{}) []byte
	GetUUID(interface{}) string
	GetArray(interface{}) []string

	// 무결성 체크 함수 제공
	CheckDBIntegrity() bool
}
