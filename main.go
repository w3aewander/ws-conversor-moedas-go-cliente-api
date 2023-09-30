/*************************************************************
*  Aplicação didática
*  Consumir API para calculo do valor do cambio entre moedas
*  @author Wanderlei Silva do Carmo <wander.silva@gmail.com>
*
**************************************************************
package main

	import (
		
		"fmt"
		"io/ioutil"
		"net/http"
		"strconv"
		//"os"
		//"bufio"
		//"strings"
	   )


const url =	 "https://www.wmomodavix.com.br/api/public/?"


func getMoneyValue(moeda string) float64 {
   
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", url, moeda), nil)
	
	if err != nil {
	 fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "text/plain")
	req.Header.Add("Content-Type", "text/plain")
	
	resp, err := client.Do(req)
   
	defer resp.Body.Close()
	
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
	 fmt.Print(err.Error())
	}
   
	//fmt.Print("VAlor do dolar:" + string(bodyBytes) )
	//var responseObject Response
	//json.Unmarshal(bodyBytes, &responseObject)
	//fmt.Printf("API Response as struct %+v\n", responseObject)

	toFloat, _ := strconv.ParseFloat(string(bodyBytes),8)

	return toFloat;

   }


func main(){

	fmt.Println("====================")
	fmt.Println("Chamando a API em Go")
	fmt.Println("====================")

	valor_dolar := getMoneyValue("dolar")
	valor_euro := getMoneyValue("euro")
	valor_bitcoin := getMoneyValue("bitcoin")

	

	fmt.Printf("\nValor do dolar: %.3f\n", valor_dolar)
	fmt.Printf("Valor do euro: %.3f\n", valor_euro)
	fmt.Printf("Valor do bitcoin: %.3f\n", valor_bitcoin)


	fmt.Println("Digite o valor em dólares:")

	var m float64

	fmt.Scanf("%f", &m)

	fmt.Printf("\nDolar convertido em reais  : %.3f", m * valor_dolar )
	fmt.Printf("\nEuro convertido em reais   : %.3f", m * valor_euro )
	fmt.Printf("\nBitcoin convertido em reais: %.3f", m * valor_bitcoin )

  
}
