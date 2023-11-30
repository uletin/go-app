package helpers

type CustomerFunc func(*CustomerOptions)

type CustomerOptions struct {
	fullname, address string
	age               int
}

type Customer struct {
	CustomerOptions
}

func defaultCustomerOptions() CustomerOptions {
	return CustomerOptions{
		fullname: "John Doe",
		address:  "Universitas Mulia",
		age:      0,
	}
}

func WithFullname(fullname string) CustomerFunc {
	return func(co *CustomerOptions) { co.fullname = fullname }
}

func NewCustomer(customerOptions ...CustomerFunc) *Customer {
	nc := defaultCustomerOptions()

	for _, fn := range customerOptions {
		fn(&nc)
	}

	return &Customer{CustomerOptions: nc}
}
