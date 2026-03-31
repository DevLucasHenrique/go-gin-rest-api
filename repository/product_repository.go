package repository

import (
	"database/sql"
	"fmt"

	"github.com/DevLucasHenrique/go-gin-rest-api/model"
)

// Repository serve para fazer a conexao entre o banco de dados e o useCase que pede para o repository salvar algo
// Em outros casos ele vai ter a array mokada quando é só para testes

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) { /* 
		função entra como metodo de ProductRepository e retorna uma struct do model.product
	*/
	query := "SELECT id, product_name, price FROM products"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}


func (pr *ProductRepository) CreateProduct(product model.Product) (int, error)  { /* 
	função entra como metodo de ProductRepository pega como parametro um product que segue a estrutura do model.Product
	*/
	var id int
	query, err := pr.connection.Prepare("INSERT INTO products (product_name, price) VALUES ($1,$2) RETURNING id") 
	if(err!=nil) {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if(err!=nil) {
		fmt.Println(err)
		return 0, err  
	}

	query.Close()
	return id, nil
}

func (pr *ProductRepository) GetProductById(id_product uint) (model.Product, error) {
	// defer query.Close() Boas praticas para a query fechar sozinha no final

	query, err := pr.connection.Prepare("SELECT * FROM products WHERE id = $1") // prepara para a query
	if err!=nil {
		fmt.Println(err)
		return model.Product{}, err
	}

	var product model.Product

	
	err = query.QueryRow(id_product).Scan( // executa a query e retorna os dados passados no scan
		&product.ID,
		&product.Name,
		&product.Price,
	)

	if err!=nil {
		if(err == sql.ErrNoRows) { // se o erro for que o banco de dados não conseguiu achar o produto do id passsado
			return model.Product{}, nil // então o erro não é do server e sim a anta do cliente, então retorna nil
		}

		return model.Product{}, err // caso não seja então foi erro no server mesmo
	}

	query.Close() // NÃO ESQUECE DE FECHAR A QUERY
	return product, nil
}


func (pr *ProductRepository) UpdateProduct( product_id uint, product model.Product) (model.Product, error) {
	query, err := pr.connection.Prepare("UPDATE products SET product_name = $1, price = $2 WHERE id = $3")
	if err!=nil {
		return model.Product{}, err
	}
	defer query.Close()

	_, err = query.Exec(product.Name, product.Price, product_id)

	if err!=nil {
		fmt.Println("ERROR: ", err)
		return model.Product{}, err
	}

	return product, nil
}