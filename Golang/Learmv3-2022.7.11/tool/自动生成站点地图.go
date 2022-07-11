package main

import (
	"fmt"
	"github.com/douyacun/gositemap"
	"time"
)

func main() {
	st := gositemap.NewSiteMap()
	st.SetPretty(true)

	url := gositemap.NewUrl()

	url.SetLoc("https://www.douyacun.com/")
	url.SetLastmod(time.Now())
	url.SetPriority(1)
	st.AppendUrl(url)
	bt, err := st.ToXml()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("%s", bt)
}
