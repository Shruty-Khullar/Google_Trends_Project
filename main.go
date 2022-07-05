package main

//Golang on its own do not understand xml or json. This is why we need external packages like encoding/json or encoding/xml to understand such lang. It only understand strings,structures,etc
//we have one to one mapping of golang var/structures with xml vars
import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//we will write the name of the xml field as it appeared in xml file so that golang will kno this is the xml field corresponding to the golang vars declared 
type RSS struct {
    XMLName    xml.Name             `xml:"rss"`         //datatype from xml package
	Channel    *Channel             `xml:"channel"`     //Pointer to the Channel
}
type Channel struct {
    Title       string              `xml:"title"`                
	ItemList    []Item              `xml:"item"`
}
type Item struct {
	Title      string               `xml:"title"`
	Link       string               `xml:"link"`
	Traffic    string               `xml:"approx_traffic"`
	NewsItems  []News               `xml:"news_item"`         //this is slice of News struct i.e. it can have instances of News struct
}
//using this xml encoding golang is converting xml fields into struct of go
//basically conversion of xml fields into golang struct is known as unmarshalling 

type News struct {
	Headline        string         `xml:"news_item_title"`
    HeadlineLink    string         `xml:"news_item_url"`
}

func main(){
	var r RSS
	data := readGoogleTrends()     //the data we get from getGoogleTrends and readGoogleTrends is in xml format and we need to convert it into format that the golang understand
    
	//unmarshall function is very intelligent. It is given by encoding/xml and it reads the code and understand the xml vars mapped to golang vars and find the relationship btw diff structs
    //we will pass the pointer to var and unmarshall will automatically convert xml to golang vars and after that we will use golang vars instead of xml ones
	err := xml.Unmarshal(data, &r)  //the data we received and the struct we want to store unmarshalled info in
	if err!=nil {
		fmt.Println("Error3: ",err)
	}
    fmt.Println("\n *****Google Search Trends for Today*****")
	fmt.Println("-----------------------------------------------------------------------------")
	for i := range r.Channel.ItemList {
		rank := (i + 1)
		fmt.Println("#", rank)
		fmt.Println("Search Term: ",r.Channel.ItemList[i].Title)
		fmt.Println("Link to the Trend: ",r.Channel.ItemList[i].Link)
		fmt.Println("***Related Headlines***")
		for j := range r.Channel.ItemList[i].NewsItems {
			fmt.Println(j,"th Headline: ", r.Channel.ItemList[i].NewsItems[j].Headline)
			//fmt.Println("Link to Related Article: ",r.Channel.ItemList[i].NewsItems[j].HeadlineLink)
		}

        fmt.Println("-----------------------------------------------")
	}
}
//Golang is so modular that for every task we need to import packages and nothing is in golang by default. This is one of the reason why golang is soo fast
func getGoogleTrends() *http.Response{      //will call url for RSS and will return the response it gets 
	resp, err := http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=IN")     //Get is used to make url req 
	if err!= nil{
		fmt.Println("Error1: ", err)
		os.Exit(1)
	}
	return resp
}
func readGoogleTrends() []byte {       //this func will return slice of bytes as read 
	resp := getGoogleTrends()         //we need to need the response as golang won't be able to read this response
	data, err := ioutil.ReadAll(resp.Body)  //whenever we make a url request we get a response and best way to handle this is by reading the response body which is done by this 
    if err!= nil {
		fmt.Println("Error2: ", err)
		os.Exit(1)
	}
    return data
}


