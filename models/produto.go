package models

import (
	"github.com/gabriel/loja/db"
)

type Produto struct {
	Id              int64
	Nome, Descricao string
	Preco           float64
	Quantidade      int64
}

func BuscaProduto() []Produto {
	db := db.ConnectionDB()

	selectProdutos, err := db.Query("SELECT * FROM produtos ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var id, quantidade int64
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func NovoProduto(nome, descricao string, preco float64, quantidade int64) {
	db := db.ConnectionDB()

	insert, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.ConnectionDB()

	delete, err := db.Prepare("DELETE FROM produtos WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)

	defer db.Close()
}

func EditarProduto(idProduto string) Produto {
	db := db.ConnectionDB()
	
	dbProduto, err := db.Query("SELECT * FROM produtos WHERE id=$1", idProduto)
	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}
	for dbProduto.Next() {
		var nome, descricao string
		var preco float64
		var quantidade, id int64

		err = dbProduto.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade
	}
	defer db.Close()
	return produto
}

func AtualizaProduto(nome, descricao string, preco float64, quantidade, id int64) {
	db := db.ConnectionDB()

	update, err := db.Prepare("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}

	update.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()
}