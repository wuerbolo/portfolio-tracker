package main

import "github.com/gin-gonic/gin"
import "net/http"
import "fmt"

func PortfolioValues(c *gin.Context) {
	values := []float32{100.72, 105.2, 103.3, 104.72, 125.2, 117.3, 119.72, 125.2, 133.3}
	dates := []string{"2021-09-06", "2021-09-07", "2021-09-08", "2021-09-00", "2021-09-10", "2021-09-11", "2021-09-12", "2021-09-13", "2021-09-14"}
	chart_data := make(map[string]float32)
	for index, date := range dates {
		chart_data[date] = values[index]
	}
	c.JSON(http.StatusOK, gin.H{"list": chart_data})
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func ProcessTransactionFile() {
	transactions := ParseDegiroTransactionsFile("Transactions.csv")
	fmt.Println(transactions)

	// mongostuff
	client := ConnectMongo()

	for _, transaction := range transactions {
		InsertTransaction(client, transaction)
	}

	DisconnectMongo(client)
}

// func main() {
// 	r := gin.Default()
// 	r.Use(CORSMiddleware())
// 	// r.Use(static.Serve("/", static.LocalFile("./todo-vue/dist", false)))
// 	r.GET("/api/chart", PortfolioValues)
// 	r.Run()
// }

func main() {
	ProcessTransactionFile()
}
