package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	resp := `{"header":["單位","員編","名稱","分機","手機","信箱","啟用"],"datas":[["[無單位]","63","Ken Lee","","0963476927","ken.lee@softleader.com.tw","false"],["[無單位]","83","Joshua Lu","","0977420203","joshua.lu@softleader.com.tw","false"],["[無單位]","93","Winni Hsu","","0987302098","winni.hsu@softleader.com.tw","false"],["[無單位]","94","Ellie Liu","","0985430959","ellie.liu@softleader.com.tw","false"],["[無單位]","106","Benny Chen","","0932781642","benny.chen@softleader.com.tw","false"],["[無單位]","116","Van Lin","","0955089656","van.lin@softleader.com.tw","false"],["[無單位]","119","Zoe Liu","","0981073022","zoe.liu@softleader.com.tw","false"],["[無單位]","133","Tim Lin","","0988053957","tim.lin@softleader.com.tw","false"],["[無單位]","134","Willy Liu","","0917812112","willy.liu@softleader.com.tw","false"],["[無單位]","135","Hilda Chang","","0920009329","hilda.chang@softleader.com.tw","false"],["[無單位]","137","Dennis Shen","","0963007146","dennis.shen@softleader.com.tw","false"],["[無單位]","141","Irene Lee","","0979321114","irene.lee@softleader.com.tw","false"],["[無單位]","143","Leo Chen","","0919066873","Leo.chen@softleader.com.tw","false"],["[無單位]","144","Ryan Chang","","0975557797","ryan.chang@softleader.com.tw","false"],["[無單位]","151","Archer Chang","","0988768512","archer.chang@softleader.com.tw","false"],["[無單位]","153","Peggy Tsai","","0932464117","peggy.tsai@softleader.com.tw","false"],["[無單位]","157","Shawn Ma","","0928088615","shawn.ma@softleader.com.tw","false"],["[無單位]","161","Abby Chen","","0970712083","abby.chen@softleader.com.tw","false"],["[無單位]","163","Van Fan","","0987177885","van.fan@softleader.com.tw","false"],["[無單位]","166","Emerald Huang","","0970208593","emerald.huang@softleader.com.tw","false"],["[無單位]","169","Sarah Chiang","","0980333588","sarah.chiang@softleader.com.tw","false"]]}`
	var buf bytes.Buffer
	if err := print(&buf, []byte(resp), false); err != nil {
		t.Error(err)
	}
	if buf.String() == "No search results" {
		t.Error("excepted response not empty")
	}
	fmt.Println(buf.String())
}

func TestPrintEmpty(t *testing.T) {
	resp := `{"header":null,"datas":null}`
	var buf bytes.Buffer
	if err := print(&buf, []byte(resp), false); err != nil {
		t.Error(err)
	}
	if buf.String() != "No search results" {
		t.Error("excepted empty response")
	}
}
