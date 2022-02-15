package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/toluwase1/busha-assessment/models"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
)
func (server *Server) GetCharacterList() gin.HandlerFunc {
	return func(c *gin.Context) {
		//clears previous error if any
		errList = map[string]string{}
		sortParameter := c.Query("sort")
		filterParameter := strings.TrimSpace(c.Query("filter"))
		orderParameter := c.Query("order")
		id, err := strconv.Atoi(c.Param("movie_id"))
		if err != nil {
			errList["Id"] = "movie Id invalid"
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  errList,
			})
			return
		}

		characters := server.Cache.GetCharacters("movie" + strconv.Itoa(id) + "char")
		if characters == nil {
			url, err := models.GetCharacterListById(c)
			if err != nil {
				errList["Id"] = "movie Id invalid"
				log.Println(err)
				c.JSON(http.StatusBadRequest, gin.H{
					"status": http.StatusBadRequest,
					"error":  errList,
				})
				return
			}

			for _, link := range *url {
				info, _ := models.GetCharInformation(link)
				characters = append(characters, info)
			}
			server.Cache.SetCharacters("movie_"+strconv.Itoa(id)+"_characters", characters)
		}

		if orderParameter == "descending" {
			switch sortParameter {
			case "name":
				sort.Slice(characters, func(i, j int) bool {
					return characters[i].Name > characters[j].Name
				})
			case "height":
				sort.Slice(characters, func(i, j int) bool {
					return characters[i].Height > characters[j].Height
				})
			case "gender":
				sort.Slice(characters, func(i, j int) bool {
					return characters[i].Gender > characters[j].Gender
				})
			}
		} else {
			switch sortParameter {
			case "name":
				sort.Slice(characters, func(i, j int) bool {
					return characters[i].Name < characters[j].Name
				})
			case "height":
				sort.Slice(characters, func(i, j int) bool {
					return characters[i].Height < characters[j].Height
				})
			case "gender":
				sort.Slice(characters, func(i, j int) bool {
					return characters[i].Gender < characters[j].Gender
				})
			}
		}

		if filterParameter == "male" || filterParameter == "female" || filterParameter == "n/a" || filterParameter == "hermaphrodite" {
			filteredList := []models.CharacterList{}
			for _, character := range characters {
				if character.Gender == filterParameter {
					filteredList = append(filteredList, character)
				}
			}
			characters = filteredList
		}

		//heightTotal := float64(0)
		//for _, character := range characters {
		//	height:=character.Height
		//	heightTotal += height
		//}
		//ft := heightTotal / 34
		//inches := heightTotal / 24.53
		//heightString := fmt.Sprintf("%.2fcm || %.2fft || %.2finches", heightTotal, ft, inches)
		c.JSON(http.StatusOK, gin.H{
			"status":   http.StatusOK,
			"response": characters,
		})
		//c.JSON(http.StatusOK, gin.H{"message": "user info retrieved successfully", "data": characters,
		//	"metadata": gin.H{"matching_characters": len(characters), "total_height": heightString}})
	}
}
