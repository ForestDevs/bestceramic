package utils

import (
	"bestceramic-parser/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/takuoki/clmconv"

	"github.com/xuri/excelize/v2"
)

func productsSheet(f *excelize.File, collections []models.Collection, keys map[string]int) {
	index := f.NewSheet("Товары")
	f.SetActiveSheet(index)
	for k, v := range keys {
		f.SetCellValue("Товары", clmconv.Itoa(v-1)+"1", k)
	}
	i := 2
	for _, c := range collections {
		for _, p := range c.Products {
			if p.Name == "" {
				continue
			}
			f.SetCellValue("Товары", clmconv.Itoa(keys["Имя"]-1)+strconv.Itoa(i), p.Name)
			f.SetCellValue("Товары", clmconv.Itoa(keys["Цена"]-1)+strconv.Itoa(i), p.Price)
			f.SetCellValue("Товары", clmconv.Itoa(keys["Картинки"]-1)+strconv.Itoa(i), strings.Join(p.Images, ";"))
			for k, v := range p.Features {
				f.SetCellValue("Товары", clmconv.Itoa(keys[k]-1)+strconv.Itoa(i), v)
			}
			i++
		}
	}
	f.SetActiveSheet(index)
}

func collectionsSheet(f *excelize.File, collections []models.Collection, keys map[string]int) {
	index := f.NewSheet("Коллекция")
	f.SetActiveSheet(index)
	i := 2
	for k, v := range keys {
		f.SetCellValue("Коллекция", clmconv.Itoa(v-1)+"1", k)
	}
	for _, c := range collections {
		f.SetCellValue("Коллекция", clmconv.Itoa(keys["Имя"]-1)+strconv.Itoa(i), c.Name)
		f.SetCellValue("Коллекция", clmconv.Itoa(keys["Бренд"]-1)+strconv.Itoa(i), c.Brand)
		f.SetCellValue("Коллекция", clmconv.Itoa(keys["Картинки"]-1)+strconv.Itoa(i), c.Image)
		f.SetCellValue("Коллекция", clmconv.Itoa(keys["Цена"]-1)+strconv.Itoa(i), c.Price)
		i++
	}
	f.SetActiveSheet(index)
}

func writeUniqueKeysCollection() {
	var settingsKeysForCollections map[string]int = map[string]int{
		"Имя":      1,
		"Цена":     2,
		"Бренд":    3,
		"Картинки": 4,
	}

	file, err := json.Marshal(settingsKeysForCollections)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	_ = ioutil.WriteFile("collectionsKeys.json", file, 755)
}

func writeUniqueKeysProduct(collections []models.Collection) {
	var settingsKeysForProducts map[string]int = map[string]int{
		"Имя":      1,
		"Цена":     2,
		"Картинки": 3,
	}
	i := len(settingsKeysForProducts)
	for _, c := range collections {
		for _, p := range c.Products {
			for t, _ := range p.Features {
				_, ok := settingsKeysForProducts[t]
				if !ok {
					i++
					settingsKeysForProducts[t] = i
				}
			}
		}
	}

	file, err := json.Marshal(settingsKeysForProducts)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	_ = ioutil.WriteFile("productsKeys.json", file, 755)
}

func ExcelWriteMultipleData(collections []models.Collection, collectionsKeys map[string]int, productsKeys map[string]int, updateKeys bool) {
	if updateKeys {
		writeUniqueKeysCollection()
		writeUniqueKeysProduct(collections)
	}
	f := excelize.NewFile()
	collectionsSheet(f, collections, collectionsKeys)
	productsSheet(f, collections, productsKeys)
	if _, err := os.Stat("./data/"); os.IsNotExist(err) {
		if err := os.Mkdir("data", os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
	if err := f.SaveAs("./data/" + collections[1].Brand + ".xlsx"); err != nil {
		fmt.Println(err)
	}
}
