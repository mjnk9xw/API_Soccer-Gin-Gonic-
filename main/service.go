package main

import (
	"Soccer-oauth2/err"
	"Soccer-oauth2/models"
	"Soccer-oauth2/server"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func main() {
	//write log
	gin.DisableConsoleColor()
	f, _ := os.Create("logService.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	//
	router := gin.Default()
	v1 := router.Group("api/v1/")
	{
		// Tour
		tour := v1.Group("Tour/")
		{
			tour.GET("", func(c *gin.Context) {
				data, err := server.GetAllTour()
				c.JSON(200, gin.H{
					"Code":        errSC.MapErrorCode[err],
					"Description": errSC.MapDescription[err],
					"Data":        data})
			})
			tour.GET(":idTour", func(c *gin.Context) {
				id := c.Param("idTour")
				data, err := server.GetTourByID(id)
				c.JSON(200, gin.H{
					"Code":        errSC.MapErrorCode[err],
					"Description": errSC.MapDescription[err],
					"Data":        data})
			})
			tour.POST("", func(c *gin.Context) {
				tour := &models.Tour{}
				err := c.ShouldBindBodyWith(&tour, binding.JSON)
				if err == nil {
					tour.SetTourID()
					err := server.AddTour(tour)
					c.JSON(200, gin.H{
						"Code":        errSC.MapErrorCode[err],
						"Description": errSC.MapDescription[err],
						"Data":        tour})
				} else {
					c.JSON(404, gin.H{
						"Code":        errSC.MapErrorCode[errSC.DataPost],
						"Description": errSC.MapDescription[errSC.DataPost],
						"Data":        nil})
				}
			})
			tour.PUT("", func(c *gin.Context) {
				tour := &models.Tour{}
				err := c.ShouldBindBodyWith(&tour, binding.JSON)
				if err == nil {
					err := server.UpdateTour(tour)
					c.JSON(200, gin.H{
						"Code":        errSC.MapErrorCode[err],
						"Description": errSC.MapDescription[err],
						"Data":        tour})
				} else {
					c.JSON(404, gin.H{
						"Code":        errSC.MapErrorCode[errSC.DataPost],
						"Description": errSC.MapDescription[errSC.DataPost],
						"Data":        nil})
				}
			})
			tour.DELETE(":idTour", func(c *gin.Context) {
				id := c.Param("idTour")
				err := server.DeleteTour(id)
				c.JSON(200, gin.H{
					"Code":        errSC.MapErrorCode[err],
					"Description": errSC.MapDescription[err],
					"Data":        id})
			})
		}
		// Round
		round := v1.Group("Tour/:idTour/Round/")
		{
			round.GET("", func(c *gin.Context) {
				idT := c.Param("idTour")
				data, err := server.GetAllRound(idT)
				c.JSON(200, gin.H{
					"Code":        errSC.MapErrorCode[err],
					"Description": errSC.MapDescription[err],
					"Data":        data})
			})
			round.GET(":idRound", func(c *gin.Context) {
				idT := c.Param("idTour")
				idR := c.Param("idRound")
				data, err := server.GetRoundByID(idT, idR)
				c.JSON(200, gin.H{
					"Code":        errSC.MapErrorCode[err],
					"Description": errSC.MapDescription[err],
					"Data":        data})
			})
			round.POST("", func(c *gin.Context) {
				idTour := c.Param("idTour")
				round := &models.Round{}
				err := c.ShouldBindBodyWith(&round, binding.JSON)
				if err == nil {
					round.SetRoundID()
					err := server.AddRound(idTour, round)
					c.JSON(200, gin.H{
						"Code":        errSC.MapErrorCode[err],
						"Description": errSC.MapDescription[err],
						"Data":        round})
				} else {
					c.JSON(404, gin.H{
						"Code":        errSC.MapErrorCode[errSC.DataPost],
						"Description": errSC.MapDescription[errSC.DataPost],
						"Data":        nil})
				}
			})
			round.PUT("", func(c *gin.Context) {
				idTour := c.Param("idTour")
				round := &models.Round{}
				err := c.ShouldBindBodyWith(&round, binding.JSON)
				if err == nil {
					err := server.UpdateRound(idTour, round)
					c.JSON(200, gin.H{
						"Code":        errSC.MapErrorCode[err],
						"Description": errSC.MapDescription[err],
						"Data":        round})
				} else {
					c.JSON(404, gin.H{
						"Code":        errSC.MapErrorCode[errSC.DataPost],
						"Description": errSC.MapDescription[errSC.DataPost],
						"Data":        nil})
				}
			})
			round.DELETE(":idRound", func(c *gin.Context) {
				idTour := c.Param("idTour")
				idRound := c.Param("idRound")
				err := server.DeleteRound(idTour, idRound)
				c.JSON(200, gin.H{
					"Code":        errSC.MapErrorCode[err],
					"Description": errSC.MapDescription[err],
					"Data":        idRound})
			})
		}
		//Table
		table := v1.Group("Tour/:idTour/Round/:idRound/Table/")
		{
			table.GET("", func(c *gin.Context) {
				idTour := c.Param("idTour")
				idRound := c.Param("idRound")
				data, err := server.GetAllTable(idTour, idRound)
				c.JSON(200, gin.H{
					"Code":        errSC.MapErrorCode[err],
					"Description": errSC.MapDescription[err],
					"Data":        data})
			})
			table.GET(":idTable", func(c *gin.Context) {
				idTour := c.Param("idTour")
				idRound := c.Param("idRound")
				idTable := c.Param("idTable")
				data, err := server.GetTableByID(idTour, idRound, idTable)
				c.JSON(200, gin.H{
					"Code":        errSC.MapErrorCode[err],
					"Description": errSC.MapDescription[err],
					"Data":        data})
			})
			table.POST("", func(c *gin.Context) {
				idT := c.Param("idTour")
				idR := c.Param("idRound")
				table := &models.Table{}
				err := c.ShouldBindBodyWith(&table, binding.JSON)
				if err == nil {
					table.SetTableID()
					err := server.AddTable(idT, idR, table)
					c.JSON(200, gin.H{
						"Code":        errSC.MapErrorCode[err],
						"Description": errSC.MapDescription[err],
						"Data":        table})
				} else {
					c.JSON(404, gin.H{
						"Code":        errSC.MapErrorCode[errSC.DataPost],
						"Description": errSC.MapDescription[errSC.DataPost],
						"Data":        nil})
				}
			})
			table.PUT("", func(c *gin.Context) {
				idTour := c.Param("idTour")
				idRound := c.Param("idRound")
				table := &models.Table{}
				err := c.ShouldBindBodyWith(&table, binding.JSON)
				if err == nil {
					err := server.UpdateTable(idTour, idRound, table)
					c.JSON(200, gin.H{
						"Code":        errSC.MapErrorCode[err],
						"Description": errSC.MapDescription[err],
						"Data":        table})
				} else {
					c.JSON(404, gin.H{
						"Code":        errSC.MapErrorCode[errSC.DataPost],
						"Description": errSC.MapDescription[errSC.DataPost],
						"Data":        nil})
				}
			})
			table.DELETE(":idTable", func(c *gin.Context) {
				idTour := c.Param("idTour")
				idRound := c.Param("idRound")
				idTable := c.Param("idTable")
				err := server.DeleteTable(idTour, idRound, idTable)
				c.JSON(200, gin.H{
					"Code":        errSC.MapErrorCode[err],
					"Description": errSC.MapDescription[err],
					"Data":        idTable})
			})
		}
		// .. get all data player in tour
		v1.GET("Tour/:idTour/Player", func(c *gin.Context) {
			idTour := c.Param("idTour")
			data, err := server.GetPlayerAll(idTour)
			c.JSON(200, gin.H{
				"Code":        errSC.MapErrorCode[err],
				"Description": errSC.MapDescription[err],
				"Data":        data})
		})
		// .. get player in tour
		v1.GET("Tour/:idTour/Player/:idPlayer", func(c *gin.Context) {
			idTour := c.Param("idTour")
			idPlayer := c.Param("idPlayer")
			data, err := server.GetPlayerInTour(idTour, idPlayer)
			c.JSON(200, gin.H{
				"Code":        errSC.MapErrorCode[err],
				"Description": errSC.MapDescription[err],
				"Data":        data})
		})
		// Player
		// v1.GET("Tour/:idTour/Round/:idRound/Table/:idTable/Player", func(c *gin.Context) {
		// 	idTour := c.Param("idTour")
		// 	idRound := c.Param("idRound")
		// 	idTable := c.Param("idTable")
		// 	data, err := server.GetAllPlayerInTable(idTour, idRound, idTable)
		// 	c.JSON(200, gin.H{
		// 		"Code":        errSC.MapErrorCode[err],
		// 		"Description": errSC.MapDescription[err],
		// 		"Data":        data})
		// })
		player := v1.Group("Tour/:idTour/Round/:idRound/Table/:idTable/Player/")
		{
			player.POST("", func(c *gin.Context) {
				idTour := c.Param("idTour")
				idRound := c.Param("idRound")
				idTable := c.Param("idTable")
				player := &models.Player{}
				err := c.ShouldBindBodyWith(&player, binding.JSON)
				if err == nil {
					player.SetPlayerID()
					err := server.AddPlayer(idTour, idRound, idTable, player)
					c.JSON(200, gin.H{
						"Code":        errSC.MapErrorCode[err],
						"Description": errSC.MapDescription[err],
						"Data":        player})
				} else {
					c.JSON(404, gin.H{
						"Code":        errSC.MapErrorCode[errSC.DataPost],
						"Description": errSC.MapDescription[errSC.DataPost],
						"Data":        nil})
				}
			})
			player.PUT("", func(c *gin.Context) {
				idTour := c.Param("idTour")
				idRound := c.Param("idRound")
				idTable := c.Param("idTable")
				player := &models.Player{}
				err := c.ShouldBindBodyWith(&player, binding.JSON)
				if err == nil {
					err := server.UpdatePlayer(idTour, idRound, idTable, player)
					c.JSON(200, gin.H{
						"Code":        errSC.MapErrorCode[err],
						"Description": errSC.MapDescription[err],
						"Data":        player})
				} else {
					c.JSON(404, gin.H{
						"Code":        errSC.MapErrorCode[errSC.DataPost],
						"Description": errSC.MapDescription[errSC.DataPost],
						"Data":        nil})
				}
			})
			player.DELETE(":idPlayer", func(c *gin.Context) {
				idTour := c.Param("idTour")
				idRound := c.Param("idRound")
				idTable := c.Param("idTable")
				idPlayer := c.Param("idPlayer")
				err := server.DeletePlayer(idTour, idRound, idTable, idPlayer)
				c.JSON(200, gin.H{
					"Code":        errSC.MapErrorCode[err],
					"Description": errSC.MapDescription[err],
					"Data":        idTable})
			})
		}
		// Match
		// v1.GET("Tour/:idTour/Round/:idRound/Table/:idTable/Match", func(c *gin.Context) {
		// 	idTour := c.Param("idTour")
		// 	idRound := c.Param("idRound")
		// 	idTable := c.Param("idTable")
		// 	data, err := server.GetAllMatchInTable(idTour, idRound, idTable)
		// 	c.JSON(200, gin.H{
		// 		"Code":        errSC.MapErrorCode[err],
		// 		"Description": errSC.MapDescription[err],
		// 		"Data":        data})
		// })
		match := v1.Group("Tour/:idTour/Round/:idRound/Table/:idTable/Match/")
		{
			match.GET(":idMatch", func(c *gin.Context) {
				idTour := c.Param("idTour")
				idRound := c.Param("idRound")
				idTable := c.Param("idTable")
				idMatch := c.Param("idMatch")
				data, err := server.GetMatchByID(idTour, idRound, idTable, idMatch)
				c.JSON(200, gin.H{
					"Code":        errSC.MapErrorCode[err],
					"Description": errSC.MapDescription[err],
					"Data":        data})
			})
			match.POST("", func(c *gin.Context) {
				idTour := c.Param("idTour")
				idRound := c.Param("idRound")
				idTable := c.Param("idTable")
				match := &models.Match{}
				err := c.ShouldBindBodyWith(&match, binding.JSON)
				if err == nil {
					match.SetMatchID()
					err := server.AddMatch(idTour, idRound, idTable, match)
					c.JSON(200, gin.H{
						"Code":        errSC.MapErrorCode[err],
						"Description": errSC.MapDescription[err],
						"Data":        match})
				} else {
					c.JSON(404, gin.H{
						"Code":        errSC.MapErrorCode[errSC.DataPost],
						"Description": errSC.MapDescription[errSC.DataPost],
						"Data":        nil})
				}
			})
			match.PUT("", func(c *gin.Context) {
				idTour := c.Param("idTour")
				idRound := c.Param("idRound")
				idTable := c.Param("idTable")
				match := &models.Match{}
				err := c.ShouldBindBodyWith(&match, binding.JSON)
				if err == nil {
					err := server.UpdateMatch(idTour, idRound, idTable, match)
					c.JSON(200, gin.H{
						"Code":        errSC.MapErrorCode[err],
						"Description": errSC.MapDescription[err],
						"Data":        match})
				} else {
					c.JSON(404, gin.H{
						"Code":        errSC.MapErrorCode[errSC.DataPost],
						"Description": errSC.MapDescription[errSC.DataPost],
						"Data":        nil})
				}
			})
			match.DELETE(":idMatch", func(c *gin.Context) {
				idTour := c.Param("idTour")
				idRound := c.Param("idRound")
				idTable := c.Param("idTable")
				idMatch := c.Param("idMatch")
				err := server.DeleteMatch(idTour, idRound, idTable, idMatch)
				c.JSON(200, gin.H{
					"Code":        errSC.MapErrorCode[err],
					"Description": errSC.MapDescription[err],
					"Data":        idTable})
			})
		}
	}
	router.Run(":8080")
}
