package controllers

import (
	"fmt"
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

		heightTotal := float64(0)
		for _, character := range characters {
			height:=character.Height
			heightTotal += height
		}
		//1 cm = 0.032808 ft
		//1 cm = 0.3937 in
		//For instance, 170cm makes 5ft and 6.93 inches.
		//6.93 inches = 17.6000664 cm
		ft := heightTotal * 0.032808
		feetToString := fmt.Sprintf("%f", ft)
		var a []string
		if strings.Contains(feetToString,  ".") {
			a = strings.Split(feetToString,".")
		}else{
			a = []string{feetToString, "0"}
		}
		stringToFloat:= a[0]
		floatNum, err := strconv.ParseFloat(stringToFloat, 64)
		if err != nil {
			return
		}
		value:= 170-(floatNum*30.48)
		newConv := fmt.Sprintf("%f %s %s %s %f %s", heightTotal, "is equal to", a[0], "feet and", value/2.54, "inches")

		fmt.Println(newConv)
		c.JSON(http.StatusOK, gin.H{
			"status":   http.StatusOK,
			"response": gin.H{
				"matched characters": len(characters),
				"converted height": newConv,
			},
		})
	}
}
