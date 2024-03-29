package api

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
	"net"
	"math/rand"
	"time"
	"github.com/olekukonko/tablewriter"
	"os"
)

func httpSend(p *Platform)  string{
	urlParam := buildUrl(p.Body)
	var req *http.Request

	if p.Method == "POST"{
		req, _ = http.NewRequest(p.Method,p.Url,strings.NewReader(urlParam))
	}else {
		req, _ = http.NewRequest(p.Method, p.Url + "?" + urlParam, nil)
	}
	// headers
	for k,v := range p.Headers{
		req.Header.Set(k,v)
	}
	resp, err := (&http.Client{}).Do(req)
	defer resp.Body.Close()

	if err != nil{
		fmt.Println("http err: ",err)
	}
	body,err:= ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("read err: ",err)
	}
	return string(body)
}

func buildUrl(requestBody body)  string{
	i := 0
	var str string
	for k,v := range requestBody{
		if i == 0 {
			str = fmt.Sprintf("%v=%v",k,v)
		}else{
			str += fmt.Sprintf("&%v=%v",k,v)
		}
		i++
	}
	return str
}

func Long2ip(ip int) net.IP {
	a := byte((ip >> 24) & 0xFF)
	b := byte((ip >> 16) & 0xFF)
	c := byte((ip >> 8) & 0xFF)
	d := byte(ip & 0xFF)
	return net.IPv4(a, b, c, d)
}

func GenerateRangeNum(min,max int) int{
	rand.Seed(time.Now().Unix())
	num := rand.Intn(max-min) + min
	return num
}

func RenderTable(songsArr []Song)  {
	var data [][]string
	for _,s := range songsArr{
		arts := strings.Join(s.Artist,",")
		data = append(data,[]string{
			s.Id,
			s.Name,
			arts,
			s.Album,
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "名称", "歌手","专辑"})

	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor})

	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor})

	table.AppendBulk(data)
	table.Render()
}
