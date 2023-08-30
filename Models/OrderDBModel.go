package Models

import "one_test_case/DBConfig"

func GetAllOrders(orders *[]Order) (err error) {
	if err = DBConfig.DB.Find(&orders).Error; err != nil {
		return err
	}
	return nil
}

func CreateOrder(order *Order) (err error) {
	if err = DBConfig.DB.Create(order).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderById(order *Order, id int32) (err error) {
	if err = DBConfig.DB.Where("id = ?", id).First(order).Error; err != nil {
		return err
	}
	return nil
}
