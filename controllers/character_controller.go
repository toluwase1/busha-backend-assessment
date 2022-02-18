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

// @Summary Get characters
// @Description Get all characters for a movie by movie id use the sort parameter to sort the results by name or height or gender, and the order parameter to order in assending or desending order eg /api/v1/movies/{movie_id}/characters?sort_by=height&filter_by=male&order=descending
// @Produce  json
// @Param movie_id path int true "Movie ID"
// @Param sort_by query string false "Sort by height or name or gender"
// @Param order query string false "ascending or descending order"
// @Param filter_by query string false "Filter by male or female or n/a or hermaphrodite"
// @Success 200 {object} []models.CharacterList
// @Failure 404 {object} models.ApiError
// @Failure 500 {object} models.ApiError
// @Router /api/v1/movies/{movie_id}/characters [get]
func (server *Server) GetCharacterList() gin.HandlerFunc {
	return func(c *gin.Context) {
		//clears previous error if any
		errList = map[string]string{}
		sortParameter := c.Query("sort_by")
		filterParameter := strings.TrimSpace(c.Query("filter_by"))
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

		characters := server.Cache.GetCharactersFromCache("movie_" + strconv.Itoa(id) + "_characters")
		if characters == nil {
			url, err := models.GetCharacterListById(id)
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
			server.Cache.SetCharToCache("movie_"+strconv.Itoa(id)+"_characters", characters)
		}

		if orderParameter == "descending" {
			switch sortParameter {
			case "name":
				sort.Slice(characters, func(i, j int) bool {
					return characters[i].Name > characters[j].Name
				})
			case "height":
				sort.Slice(characters, func(i, j int) bool {
					heightI, _ := strconv.ParseFloat(characters[i].Height, 64)
					heightJ, _ := strconv.ParseFloat(characters[j].Height, 64)
					return heightJ < heightI
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
					heightI, _ := strconv.ParseFloat(characters[i].Height, 64)
					heightJ, _ := strconv.ParseFloat(characters[j].Height, 64)
					return heightJ < heightI
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
			height, err := strconv.ParseFloat(character.Height, 64)
			if err != nil {
				log.Println(err)
				continue
			}
			heightTotal += height
		}
		//1 cm = 0.032808 ft
		//1 cm = 0.3937 in
		//For instance, 170cm makes 5ft and 6.93 inches.
		//6.93 inches = 17.6000664 cm
		ft := heightTotal * 0.032808
		feetToString := fmt.Sprintf("%f", ft)
		var a []string
		if strings.Contains(feetToString, ".") {
			a = strings.Split(feetToString, ".")
		} else {
			a = []string{feetToString, "0"}
		}
		stringToFloat := a[0]
		floatNum, err := strconv.ParseFloat(stringToFloat, 64)
		if err != nil {
			return
		}
		value := 170 - (floatNum * 30.48)
		newConv := fmt.Sprintf("%f %s %s %s %f %s", heightTotal, "is equal to", a[0], "feet and", value/2.54, "inches")

		fmt.Println(newConv)
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"response": gin.H{
				"matched characters": len(characters),
				"converted height":   newConv,
				"characters":         characters,
			},
		})
	}
}
