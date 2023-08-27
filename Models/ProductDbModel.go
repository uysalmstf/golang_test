package Models

import "one_test_case/DBConfig"

func GetAllProducts(products *[]Product) (err error) {
	if err = DBConfig.DB.Find(products).Error; err != nil {
		return err
	}
	return nil
}

func CreateProduct(product *Product) (err error) {
	if err = DBConfig.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}
