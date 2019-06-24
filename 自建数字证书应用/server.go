package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(
			            "   沁园春·雪\n"+
		               "   [现代] 毛泽东\n" +
			           "望长城内外，惟馀莽莽；\n" +
			           "大河上下，顿失滔滔。\n" +
			           "山舞银蛇，原驰蜡象，欲与天公试比高。\n" +
			           "须晴日，看红妆素裹，分外妖娆。\n" +
			           "江山如此多娇，\n" +
			           "引无数英雄竞折腰。\n" +
			           "惜秦皇汉武，略输文采；\n" +
			           "唐宗宋祖，稍逊风骚。\n" +
			           "一代天骄，成吉思汗，只识弯弓射大雕。\n" +
			           "俱往矣，数风流人物，还看今朝。\n"))
	})

	fmt.Println("服务器启动。。。")

	err := http.ListenAndServeTLS(":8082", "server.crt", "server.key", nil)
	if err != nil {
		fmt.Println("ListenAndServeTLS err ")
		return
	}
}
