package handler

import (
	"log"
	"net/http"
	"rumah-sakit/model/dto"
	usacase "rumah-sakit/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler interface {
	GetAllOrder(c *gin.Context)
	GetOrder(c *gin.Context)
	CreateOrder(c *gin.Context)
	UpdateOrder(c *gin.Context)
	DeleteOrder(c *gin.Context)
}

func CreateOrderHandler(usecase usacase.OrderUseCase) OrderHandler {
	return &orderHandler{
		orderUseCase: usecase,
	}
}

type orderHandler struct {
	orderUseCase usacase.OrderUseCase
}

func (order *orderHandler) GetAllOrder(c *gin.Context) {

	// Memanggil use case untuk mendapatkan semua order
	orders, err := order.orderUseCase.GetAllOrder()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan data order dalam format JSON
	c.JSON(200, map[string]any{
		"Data":        orders,
		"Status Code": 200,
		"Message":     "Success get all order",
	})
}

func (order *orderHandler) GetOrder(c *gin.Context) {

	id := c.Param("id")

	orderID, err := strconv.Atoi(id)

	if err != nil {
		log.Println("Invalid order ID:", err)
	}

	// Memanggil use case untuk mendapatkan order
	get_order, err := order.orderUseCase.GetOrder(orderID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan data order dalam format JSON
	c.JSON(200, map[string]any{
		"Data":        get_order,
		"Status Code": 200,
		"Message":     "Success get order",
	})
}

func (order *orderHandler) CreateOrder(c *gin.Context) {

	var request_order dto.Order_request

	if err := c.ShouldBind(&request_order); err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid input",
		})
		return
	}

	// Memanggil use case untuk mendapatkan order
	create_order, err := order.orderUseCase.CreateOrder(request_order)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan data order dalam format JSON
	c.JSON(200, map[string]any{
		"Data":        create_order,
		"Status Code": 201,
		"Message":     "Success create order",
	})
}

func (order *orderHandler) UpdateOrder(c *gin.Context) {

	var request_order dto.Order_request

	if err := c.ShouldBind(&request_order); err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid input",
		})
		return
	}

	// Memanggil use case untuk mendapatkan order
	update_order, err := order.orderUseCase.UpdateOrder(request_order)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan data order dalam format JSON
	c.JSON(200, map[string]any{
		"Data":        update_order,
		"Status Code": 200,
		"Message":     "Success update order",
	})
}

func (order *orderHandler) DeleteOrder(c *gin.Context) {

	id := c.Param("id")

	orderID, err := strconv.Atoi(id)

	if err != nil {
		log.Println("Invalid order ID:", err)
	}

	// Memanggil use case untuk mendapatkan order
	err = order.orderUseCase.DeleteOrder(orderID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan data order dalam format JSON
	c.JSON(200, map[string]any{

		"Status Code": 200,
		"Message":     "Success delete order",
	})
}
