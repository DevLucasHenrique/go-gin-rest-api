package repository

import (
	"database/sql"
	"fmt"

	"github.com/DevLucasHenrique/go-gin-rest-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

/*
 * @param MyParam 
 
 */

func (pr *ProductRepository) GetProducts() ([]model.Product, error) { /* 
		função entra como metodo de ProductRepository e retorna uma struct do model.product
	*/
	query := "SELECT id, product_name, price FROM product"
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
	query, err := pr.connection.Prepare("INSERT INTO product (product_name, price) VALUES ($1,$2) RETURNING id") 
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
