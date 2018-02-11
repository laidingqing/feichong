package pholcus_lib

// 基础包
import (
	//必需

	"log"
	"strconv"

	"github.com/henrylee2cn/pholcus/app/downloader/request"
	. "github.com/henrylee2cn/pholcus/app/spider"
	"github.com/henrylee2cn/pholcus/common/goquery"
)

func init() {
	ZhiPin.Register()
}

var page = 1
var baseURL = "https://www.zhipin.com"

// ZhiPin search

var ZhiPin = &Spider{
	Name:         "Boss直聘",
	Description:  "https://www.zhipin.com/c101230100-p150301/h_101230100/?page=xx", //150301为职位代码，循环替换
	Keyin:        KEYIN,
	Limit:        LIMIT,
	EnableCookie: false,
	RuleTree: &RuleTree{
		Root: func(ctx *Context) {
			ctx.Aid(map[string]interface{}{"loop": [2]int{0, 1}, "Rule": "生成请求"}, "生成请求")
		},
		Trunk: map[string]*Rule{
			"生成请求": {
				AidFunc: func(ctx *Context, aid map[string]interface{}) interface{} {
					for loop := aid["loop"].([2]int); loop[0] < loop[1]; loop[0]++ {
						ctx.AddQueue(
							&request.Request{
								Url:  "https://www.zhipin.com/c101230100-p150301/h_101230100/?page=" + strconv.Itoa(loop[0]+1),
								Rule: aid["Rule"].(string),
							},
						)
					}
					return nil
				},
				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()
					next := query.Find(".page a[ka='page-next']")
					href, _ := next.Attr("href")
					if href == "javascript:;" {
						log.Println("最后一页了")
						ctx.SetLimit(page)
					} else {
						log.Println("next:" + href)
						page++
						ctx.Aid(map[string]interface{}{"loop": [2]int{1, ctx.GetLimit()}, "Rule": "搜索结果"})
					}
					ctx.Parse("搜索结果")
				},
			},
			"搜索结果": {
				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()
					query.Find(".job-primary").Each(func(i int, s *goquery.Selection) {
						ctx.SetTemp("html", s)
						ctx.Parse("职位列表")
					})
				},
			},
			"职位列表": {
				ParseFunc: func(ctx *Context) {
					var selector = ctx.GetTemp("html", &goquery.Selection{}).(*goquery.Selection)
					detailURL, _ := selector.Find(".info-primary .name a").First().Attr("href")
					jobTitle := selector.Find(".info-primary .job-title").First().Text()
					jobSalary := selector.Find(".info-primary .red").First().Text()
					companyName := selector.Find(".info-company .company-text a").First().Text()
					ctx.AddQueue(&request.Request{
						Url:  baseURL + detailURL,
						Rule: "职位详情",
						Temp: ctx.CreatItem(map[int]interface{}{
							0: companyName,
							1: jobTitle,
							2: jobSalary,
						}, "职位详情"),
						Priority: 1,
					})
				},
			},
			"职位详情": {
				ItemFields: []string{
					"company_name",
					"job_title",
					"salary",
					"job_desc",
					"company_addr",
					"spider_from",
				},
				ParseFunc: func(ctx *Context) {
					r := ctx.CopyTemps()
					query := ctx.GetDom()
					desc := query.Find(".job-sec .text").First().Html
					addr := query.Find(".job-location .location-address")
					r["3"] = desc
					r["4"] = addr
					r["5"] = "zhipin"
					ctx.Output(r)
				},
			},
		},
	},
}
