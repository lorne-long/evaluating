package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"net"
)

// App struct
type App struct {
	ctx context.Context
}
type Result struct {
	code int
	msg  string
	data interface{}
}
type topic struct {
	name    string
	options string
	answer  string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// tss := {"code":200,"msg":"成功","data":[{"answerList":["从不","很少","偶尔","经常","总是"],"report":["从不","很少","偶尔","经常","总是"],"type":1,"topicList":["朗读时一个字一个字地读，速度很慢","对前后左右上下等方位很难辨别","钻洞、攀爬类游戏困难，布置如何控制摆放身体","根据样图照样子涂色，会涂错位置","很难找到或发现掉落在地上的物品"]},{"answerList":["从不","很少","偶尔","经常","总是"],"report":["从不","很少","偶尔","经常","总是"],"type":2,"topicList":["跑步姿势怪异，动作不协调","平时好动，但是，运动技巧却不佳","容易打翻东西，弄脏或损坏衣物或作业本","不会描画或者所画线条偏离目标线条明显","运笔能力弱，握笔姿势怪异或错误"]},{"answerList":["从不","很少","偶尔","经常","总是"],"report":["从不","很少","偶尔","经常","总是"],"type":3,"topicList":["由于忽视细节或粗心，在学校作业或其他活动中常犯些粗心的错误","很难在某件事情或活动上专注较长时间","对于读音相近的字母或拼音分不清，比如“兔子”和“肚子”、“t”和l","打断或搅扰别人，比如打断别人的谈话或游戏","发音不准确，尤其是跟读时发音错误"]},{"answerList":["从不","很少","偶尔","经常","总是"],"report":["从不","很少","偶尔","经常","总是"],"type":4,"topicList":["很会背诵，但是背诵后却一问三不知","很难复述刚刚交代的话","没有按照指示做事没做作业或指定的事情（不是由于故意反抗或不理解指示）","对妈妈或老师说的话仿佛没听到，爱问“刚刚说什么了？","别人对他说话时，他好象没注意听"]},{"answerList":["从不","很少","偶尔","经常","总是"],"report":["从不","很少","偶尔","经常","总是"],"type":5,"topicList":["在别人还没说完问题前，就脱口乱答","对于听说结合的游戏缺乏兴趣，不能参与其中","只顾自己活动，无视教师或家长的指令","说话时言语缺乏组织，句型过分简单，逻辑性不足","很会朗读，但是对内容却一知半解"]},{"answerList":["从不","很少","偶尔","经常","总是"],"report":["从不","很少","偶尔","经常","总是"],"type":6,"topicList":["朗读时增字、漏字，或忽略句号、逗号","写字时，偏旁部首经张冠李戴、左右颠倒","写字时笔顺颠倒，笔画不正确","对数学应用题感到十分困难，不解题意","不能找到两个图片不同的地方"]},{"answerList":["从不","很少","偶尔","经常","总是"],"report":["从不","很少","偶尔","经常","总是"],"type":7,"topicList":["写作业速度很慢，时间拖得很长","做算术题时，忘记计算过程中的进位或错位","不能选出刚刚看到的图片上的相应物品","提笔忘字，对文字的记忆能力差","丢失上课或活动时的必需品，如作业本铅笔课本玩具或工具"]},{"answerList":["从不","很少","偶尔","经常","总是"],"report":["从不","很少","偶尔","经常","总是"],"type":8,"topicList":["阅读时，需要用手指帮助知识文字方向","写字时，看一眼写一笔","看动画的时候，只注重绘本上最突出的一点，不能对整个画面观察就进入下一页","写字太小，不易分辨","画画涂色只在纸张很小的一块区域"]}]}

// 启动是在应用程序启动时调用的。上下文已保存
// 这样我们就可以调用运行时方法
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet返回对给定名称的问候
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Greet1(name string) string {
	return fmt.Sprintf("Hello1 %s, It's show time!", name)
}

func (a *App) VerifyCode(code string) string {
	const publicKeyStr = `-----BEGIN RSA PUBLIC KEY-----
	MIIBCgKCAQEAxgAO0IjskGoSR/3Ndm9TdEuF/8wGAAOgtKNMIFZXWd+2ruo/wtA5
	ELT8dlZQPI6ZZ9WlsVg2BR7IyopH3NI3KgGggcrv+qyPzOsLOWHboT9aTEbAsdy8
	+VEfgp3HsXPIOcVFQ7N7PvLqU7VfUaNMPc9e5RI29ZpQws/o7/2VIwSASfsnkmdF
	EpxGx/hS3GsTxYqcsNZ/gK/VNd5ijy7RUsI4f9Z/FmoulhrES0qDpBiOAaR5xEoV
	X1CbUzEQdsdZrQS+7pZB/ab0atAkH3AZ0SlCmaNIAampFC6ALgk1+xG8SPg2a3+7
	85fYdniTl0U3qYfO9FMuR5QLe95lQNAeNwIDAQAB
	-----END RSA PUBLIC KEY-----`
	var reader = rand.Reader
	// privateKey, err := rsa.GenerateKey(reader, 2048)
	// if err != nil {
	// 	return "1"
	// }
	key, _ := hex.DecodeString(publicKeyStr)
	publicKey, _ := x509.ParsePKCS1PublicKey(key)

	//The public key is a part of the *rsa.PrivateKey struct
	// publicKey := privateKey.PublicKey
	encryptedBytes, err := rsa.EncryptPKCS1v15(
		reader,
		publicKey,
		[]byte(GetMac()[0]),
	)
	if err != nil {
		return string(encryptedBytes)
	}
	return "ets"

}
func GetMac() (macAddrs []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("fail to get net interfaces: %v\n", err)
		return macAddrs
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}
		macAddrs = append(macAddrs, macAddr)
	}
	return macAddrs
}
