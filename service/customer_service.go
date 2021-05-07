package service

import (
	"github.com/pulpfree/shts-api/model"
)

func (svc *Service) CreateCustomer(input *model.CreateCustomer) (*model.Customer, error) {
	return svc.repo.CreateCustomer(input)
}

func (svc *Service) GetCustomer(id string) (*model.Customer, error) {
	return svc.repo.GetOne(id)
}

func (svc *Service) ListCustomers() ([]*model.Customer, error) {
	return svc.repo.GetAll()
}

// Adds an Observer to the CustomerCreationStream. The returned object is holding a personal channel
func (svc *Service) SubscribeCustomerCreation() *CustomerServiceObserver {
	svc.mutex.Lock()
	deliveryChannel := make(chan *model.Customer)
	observer := &CustomerServiceObserver{deliveryChannel}
	svc.customerCreationObservers = append(svc.customerCreationObservers, observer)
	svc.mutex.Unlock()
	return observer
}

// Remove an Observer from the CustomerCreationStream
func (svc *Service) UnsubscribeCustomerCreation(observer *CustomerServiceObserver) {
	svc.mutex.Lock()
	close(observer.CreationStream)
	j := 0
	for _, entry := range svc.customerCreationObservers {
		if entry == observer {
			svc.customerCreationObservers[j] = entry
		}
	}
	svc.customerCreationObservers = svc.customerCreationObservers[:j]
	svc.mutex.Unlock()
}

// This multiplexer routes incoming Customers from the the customerRepository to every active subscriber.
func (svc *Service) customerCreationStreamMultiplexer() {
	incoming := svc.repo.GetCreationStream()
	for {
		customer := <-incoming
		svc.mutex.Lock()
		for _, entry := range svc.customerCreationObservers {
			entry.CreationStream <- customer
		}
		svc.mutex.Unlock()
	}
}
