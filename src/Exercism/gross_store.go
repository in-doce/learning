package exercism

// Units stores the Gross Store unit measurements.
func Units() (units map[string]int) {
	units = map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return
}

// exists checks if a given key exists in a map.
func exists(m map[string]int, key string) bool {
	_, exists := m[key]
	return exists
}

// NewBill creates a new bill.
func NewBill() (bill map[string]int) {
	bill = make(map[string]int)
	return
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if !exists(units, unit) {
		return false
	}
	bill[item] += units[unit]
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	newQuantity := bill[item] - units[unit]
	switch {
	case !exists(bill, item):
		return false
	case !exists(units, unit):
		return false
	case newQuantity < 0:
		return false
	case newQuantity == 0:
		delete(bill, item)
		return true
	default:
		bill[item] = newQuantity
		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if !exists(bill, item) {
		return 0, false
	}
	return bill[item], true
}
