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

func GetProductByCode(product *Product, code string) (err error) {
	if err = DBConfig.DB.Where("code = ?", code).First(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProductById(product *Product, id int32) (err error) {
	if err = DBConfig.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}
