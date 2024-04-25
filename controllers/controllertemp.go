package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gabriel/loja/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaProduto()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		if err != nil {
			log.Println("Erro na conversão do preco:", err)
		}
		quantidade, err := strconv.Atoi(r.FormValue("quantidade"))
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.NovoProduto(nome, descricao, preco, int64(quantidade))
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idProduto)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := models.EditarProduto(idProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "POST" {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			log.Println("Erro na conversão do ID:", err)
		}
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		if err != nil {
			log.Println("Erro na conversão do PREÇO:", err)
		} 
		quantidade, err := strconv.Atoi(r.FormValue("quantidade"))
		if err != nil {
			log.Println("Erro na conversão da QUANTIDADE:", err)
		}

		models.AtualizaProduto(nome, descricao, preco, int64(quantidade), int64(id))
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}