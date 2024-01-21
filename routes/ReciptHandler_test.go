package routes

import (
	"receipt-processor-challenge/models"
	"receipt-processor-challenge/pkg"
	"testing"
)

func TestReciptHandler(t *testing.T) {
	rs := pkg.NewMemReciptStore()
	rh := NewReciptHandler(rs)
	r := models.Recipt{
		Retailer:     "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Total:        "9.00",
		Items: []models.Item{
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
		},
	}

	res := rh.reciptStore.AddNewRecipt(r)
	resGet := rh.reciptStore.GetRecipt(res)
	if resGet != 109 {
		t.Errorf("Expected %d points, but got %d", 109, resGet)
	}
}

func TestReciptHandler2(t *testing.T) {
	rs := pkg.NewMemReciptStore()
	rh := NewReciptHandler(rs)
	r := models.Recipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Total:        "35.35",
		Items: []models.Item{
			{
				ShortDescription: "Mountain Dew 12PK",
				Price:            "6.49",
			}, {
				ShortDescription: "Emils Cheese Pizza",
				Price:            "12.25",
			}, {
				ShortDescription: "Knorr Creamy Chicken",
				Price:            "1.26",
			}, {
				ShortDescription: "Doritos Nacho Cheese",
				Price:            "3.35",
			}, {
				ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
				Price:            "12.00",
			},
		},
	}

	res := rh.reciptStore.AddNewRecipt(r)
	resGet := rh.reciptStore.GetRecipt(res)
	if resGet != 28 {
		t.Errorf("Expected %d points, but got %d", 28, resGet)
	}
}
