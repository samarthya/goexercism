package gross

import "log"

// Units stores the Gross Store unit measurements.
func Units() (m map[string]int) {
	m = make(map[string]int)
	m["quarter_of_a_dozen"] = 3
	m["half_of_a_dozen"] = 6
	m["dozen"] = 12
	m["small_gross"] = 120
	m["gross"] = 144
	m["great_gross"] = 1728
	return
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if bill != nil && units != nil {
		if val, ok := units[unit]; ok {
			bill[item] = val
			return true
		}
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {

	if val1, ok := GetItem(bill, item); ok {
		log.Println(" Item ", item, " found ", val1)
		if val2, ok := GetUnits(units, unit); ok {
			log.Println(" Unit ", unit, " found ", val2)
			if val1 == val2 {
				delete(bill, item)
				return true
			} else if val1 > val2 {
				bill[item] -= val2
				return true
			}
		} else {
			return false
		}
	}

	return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	log.Println(" Looking for ", item, " in ", bill)
	if v, ok := bill[item]; ok {
		log.Println(" found ", v)
		return v, true
	}
	log.Println(" not found ")
	return 0, false
}

// GetUnits Get units
func GetUnits(units map[string]int, unit string) (int, bool) {
	log.Println(" Looking for ", unit, " in ", units)
	return GetItem(units, unit)
}
