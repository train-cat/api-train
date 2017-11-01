package repositories

import "fmt"

// ValueExist return true if value 'v' exist in 's' table for field 'f'
func ValueExist(s interface{}, f string, v string) (bool, error) {
	count := 0

	err := db.Model(s).Where(fmt.Sprintf("%s = ?", f), v).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// IDExist return true if ID exist in table
func IDExist(table string, id uint64) (bool, error) {
	count := 0

	err := db.Table(table).Where("id = ?", id).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
