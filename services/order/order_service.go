package order

import (
	"bringeee-capstone/entities"
	"bringeee-capstone/entities/web"
	"mime/multipart"
)

type OrderService struct {
	
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

/*
 * Find All
 * -------------------------------
 * Mengambil data order berdasarkan filters dan sorts
 *
 * @var limit 	batas limit hasil query
 * @var offset 	offset hasil query
 * @var filters	query untuk penyaringan data, { field, operator, value }
 * @var sorts	pengurutan data, { field, value[bool] }
 * @return order	list order dalam bentuk entity domain
 * @return error	error
 */
func (service OrderService) FindAll(limit int, offset int, filters []map[string]string, sorts []map[string]interface{}) ([]entities.OrderResponse, error) {
	panic("implement me")
}
/*
 * Get Pagination
 * -------------------------------
 * Mengambil data pagination berdasarkan filters dan sorts
 *
 * @var limit 	batas limit hasil query
 * @var page 	halaman sekarang diakses
 * @var filters	query untuk penyaringan data, { field, operator, value }
 * @return order	response pagination
 * @return error	error
 */
func (service OrderService) GetPagination(limit int, page int, filters []map[string]string) (web.Pagination, error) {
	panic("implement me")
}
/*
 * Find
 * -------------------------------
 * Mengambil data order tunggal berdasarkan ID
 *
 * @var id 		id order
 * @return order	order tunggal dalam bentuk response
 * @return error	error
 */
func (service OrderService) Find(id int) (entities.Order, error) {
	panic("implement me")
}
/*
 * Customer Create order
 * -------------------------------
 * Membuat order baru berdasarkan user yang sedang login
 * @var orderRequest		request create order oleh customer
 * @var files				list file request untuk diteruskan ke validation dan upload
 * @return OrderResponse	order response 
 */
func (service OrderService) Create(orderRequest entities.CustomerCreateOrderRequest, files []*multipart.FileHeader) (entities.OrderResponse, error) {
	panic("implement me")
}
/*
 * Admin Set fixed price order
 * -------------------------------
 * Set fixed price order oleh admin untuk diteruskan kembali ke user agar di konfirmasi/cancel
 * @var orderRequest		request create order oleh customer
 * @return OrderResponse	order response 
 */
func (service OrderService) SetFixOrder(setPriceRequest entities.AdminSetPriceOrderRequest) error  {
	panic("implement me")
}
/*
 * Confirm Order
 * -------------------------------
 * Confirm order yang sudah dibuat
 * @var orderID				ID Order yang akan di cancel
 * @return OrderResponse	order response 
 */
func (service OrderService) ConfirmOrder(orderID int) error  {
	panic("implement me")
}
/*
 * Cancel Order
 * -------------------------------
 * Cancel order yang sudah dibuat
 * @var orderID				ID Order yang akan di cancel
 * @return OrderResponse	order response 
 */
func (service OrderService) CancelOrder(orderID int) error  {
	panic("implement me")
}
/*
 * Create Payment
 * -------------------------------
 * Buat pembayaran untuk order tertentu ke layanan pihak ketiga
 * @var orderID					ID Order yang akan di cancel
 * @var createPaymentRequest	request payment
 * @return PaymentResponse		response payment 
 */
func (service OrderService) CreatePayment(orderID int, createPaymentRequest entities.CreatePaymentRequest) (entities.PaymentResponse, error) {
	panic("implement me")
}
/*
 * Get Payment
 * -------------------------------
 * Mengambil data pembayaran yang sudah ada berdasarkan transaction_id yang sudah di set pada order
 * @var orderID					ID Order yang akan di cancel
 * @var createPaymentRequest	request payment
 * @return PaymentResponse		response payment 
 */
func (service OrderService) GetPayment(orderID int, createPaymentRequest entities.CreatePaymentRequest) (entities.PaymentResponse, error) {
	panic("implement me")
}
	
/*
 * Find All History
 * -------------------------------
 * Mengambil data order berdasarkan filters dan sorts
 *
 * @var limit 	batas limit hasil query
 * @var offset 	offset hasil query
 * @var filters	query untuk penyaringan data, { field, operator, value }
 * @var sorts	pengurutan data, { field, value[bool] }
 * @return order	list order dalam bentuk entity domain
 * @return error	error
 */
func (service OrderService) FindAllHistory(sorts []map[string]interface{}) ([]entities.OrderHistoryResponse, error) {
	panic("implement me")
}
/*
 * Webhook
 * -------------------------------
 * Payment Webhook notification, dikirimkan oleh layanan pihak ketiga
 * referensi: https://docs.midtrans.com/en/after-payment/http-notification
 *
 * @var limit 	batas limit hasil query
 * @var offset 	offset hasil query
 * @var filters	query untuk penyaringan data, { field, operator, value }
 * @var sorts	pengurutan data, { field, value[bool] }
 * @return order	list order dalam bentuk entity domain
 * @return error	error
 */
func (service OrderService) PaymentWebhook(orderID int) error {

	// if status settlement, set order to MANIFESTED
	// if status is deny, cancel, expire, set to CANCELLED
	
	panic("implement me")
}

/*
 * Take order for shipping
 * -------------------------------
 * Pengambilan order oleh driver untuk di set statusnya menjadi ON_PROCESS
 * @var orderID 	order id terkait
 * @var userID		authenticated user (role: driver)
 */
func (service OrderService) TakeOrder(orderID int, userID int) error {
	panic("implement me")
}

/*
 * Finish Order
 * -------------------------------
 * penyelesaian order oleh driver untuk di set statusnya menjadi DELIVERED
 * @var orderID 	order id terkait
 * @var userID		authenticated user (role: driver)
 */
func (service OrderService) FinishOrder(orderID int, userID int, files []*multipart.FileHeader) error {
	panic("implement me")
}