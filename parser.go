package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("urlS.txt") //читаем файл, данные прихоядт массивом
	if err != nil {                          //если функция не возвращает nil то выводим ошибку
		fmt.Println(err)
	}
	var urlAdr string = string(data)
	//var obrezok = strings.Trim(urlAdr, " https://")
	//fmt.Println(urlAdr)

	arr := strings.Split(urlAdr, " ")

	for i := 0; i < len(arr); i++ {
		resp, err := http.Get(arr[i]) //получаем http через переменную из блока маин
		var obrezok = strings.Trim(arr[i], " https://") + ".txt"
		fmt.Println(obrezok)
		if err != nil { //если функция не возвращает nil то выводим ошибку
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		var giveHtp string
		for true {

			bs := make([]byte, 1014)     //создаем массив
			n, err := resp.Body.Read(bs) //строки
			giveHtp = string(bs[:n])
			fmt.Println(giveHtp)
			if n == 0 || err != nil {
				break
			}
		}

		file, err := os.Create(obrezok) // создаем файл
		if err != nil {                 // если возникла ошибка
			fmt.Println("Unable to create file:", err)
			os.Exit(1) // выходим из программы
		}
		defer file.Close() // закрываем файл
		file.WriteString(giveHtp)
	}

}
