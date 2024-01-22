package pkg

import (
	"math"
	"receipt-processor-challenge/models"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)

type ReciptStore struct {
	mu       sync.RWMutex
	receipts map[string]models.Recipt
}

func NewMemReciptStore() *ReciptStore {
	return &ReciptStore{
		receipts: make(map[string]models.Recipt),
	}
}

func (rs *ReciptStore) AddNewRecipt(r models.Recipt) string {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	id := uuid.New().String()
	r.ID = id
	rs.receipts[r.ID] = r
	return r.ID
}

func (rs *ReciptStore) GetRecipt(ID string) int {
	rs.mu.RLock()
	defer rs.mu.RUnlock()
	receipt, ok := rs.receipts[ID]
	if ok {
		return calculatePoints(receipt)
	}
	return -1
}

func calculatePoints(r models.Recipt) int {
	// totalPoints := 0
	retailNamePoint := 0
	roundAmountPoint := 0
	totalMultipleOf25Point := 0
	everyTwoItemsPoint := 0
	trimmedDescriptionPoint := 0
	purchaseDateOddDayPoint := 0
	purchaseTimeBetween2and4Point := 0

	retailNamePoint = len(nonAlphanumericRegex.ReplaceAllString(r.Retailer, ""))

	if strings.HasSuffix(r.Total, ".00") || !strings.Contains(r.Total, ".") {
		roundAmountPoint = 50
	}

	if total, err := strconv.ParseFloat(r.Total, 64); err == nil && math.Mod(total, 0.25) == 0 {
		totalMultipleOf25Point = 25
	}

	everyTwoItemsPoint = ((len(r.Items) / 2) * 5)

	for i := 0; i < len(r.Items); i++ {
		if len(strings.Trim(r.Items[i].ShortDescription, " "))%3 == 0 {
			price, err := strconv.ParseFloat(r.Items[i].Price, 64)
			if err == nil {
				roundedValue := int(math.Ceil(price * 0.2))
				trimmedDescriptionPoint += roundedValue
			}
		}
	}

	if date, err := time.Parse("2006-01-02", r.PurchaseDate); err == nil && date.Day()%2 != 0 {
		purchaseDateOddDayPoint = 6
	}

	if purchaseTime, err := time.Parse("15:04", r.PurchaseTime); err == nil && purchaseTime.Hour() >= 14 && purchaseTime.Hour() <= 16 {
		purchaseTimeBetween2and4Point = 10
	}

	return retailNamePoint + roundAmountPoint + totalMultipleOf25Point + everyTwoItemsPoint + trimmedDescriptionPoint + purchaseDateOddDayPoint + purchaseTimeBetween2and4Point
}
