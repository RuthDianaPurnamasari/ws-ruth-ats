package controller

import (
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	cek "github.com/RuthDianaPurnamasari/package_ats/module"

)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetAllPenelitian(c *fiber.Ctx) error {
	ps := cek.GetAllPenelitian()
	return c.JSON(ps)
	}