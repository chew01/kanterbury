package proxy

import (
	"bytes"
	"encoding/json"
	"github.com/chew01/kanterbury/utils"
	"github.com/elazarl/goproxy"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func (p *Proxy) handleHttps(host string, _ *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
	return goproxy.MitmConnect, host
}

var hostRegex = regexp.MustCompile(`gc-openapi-zinny3.kakaogames.com`)

func (p *Proxy) handleReq(req *http.Request, _ *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	// Return if not game traffic
	if !hostRegex.MatchString(req.URL.Host) {
		return req, nil
	}

	body, err := ioutil.ReadAll(req.Body)
	utils.Must(err)
	req.Body = ioutil.NopCloser(bytes.NewReader(body))

	opStr := strings.Trim(req.URL.Path, "/service/v3/")
	opArgs := strings.Split(opStr, "/")

	if opArgs[0] == "log" {
		switch opArgs[1] {
		case "writeActionLog":
			var player PlayerData
			utils.Must(json.Unmarshal(body, &player))
			p.State.Player = &player
		case "writeRoundLog":
			var activity ActivityData
			utils.Must(json.Unmarshal(body, &activity))
			if activity.EndTime == 0 {
				p.State.Activity = &activity
			} else {
				p.State.Activity = &ActivityData{}
			}
		default:
			var character CharacterData
			utils.Must(json.Unmarshal(body, &character))
			p.State.Character = &character
		}
	}

	return req, nil
}

func (p *Proxy) handleRes(res *http.Response, _ *goproxy.ProxyCtx) *http.Response {
	return res
}
