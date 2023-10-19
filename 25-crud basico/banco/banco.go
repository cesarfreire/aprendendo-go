package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver de conexao
)

func Conectar() (*sql.DB, error) {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {
		return nil, erro
	}

	// garante que a conexão foi feita com sucesso
	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	// retorna a conexão
	return db, nil
}
