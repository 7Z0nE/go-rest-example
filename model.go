package main

type Vendor struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type Product struct {
	Id     int    `json:"id" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Vendor Vendor `json:"vendor" binding:"required"`
}
