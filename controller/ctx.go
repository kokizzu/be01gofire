package controller

import "github.com/gin-gonic/gin"

type Ctx struct {
	*Server
	*gin.Context
}

type Handler func(*Ctx) 
