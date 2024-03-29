package cg

import (
	"cgss/ipc"
	"encoding/json"
	"errors"
)

type CenterClient struct {
	*ipc.IpcClient
}

func (client *CenterClient) AddPlayer(player *Player) error {
	b, err := json.Marshal(*player)
	if err != nil {
		return err
	}

	params := string(b)
	resp, err := client.Call("addplayer", params)
	if err == nil && resp.Code == "200" {
		return nil
	}
	return err
}

func (client *CenterClient) RemovePlayer(name string) error {
	ret, err := client.Call("removeplayer", name)
	if err != nil {
		return err
	}
	if ret.Code == "200" {
		return nil
	}
	return errors.New(ret.Code)
}

func (client *CenterClient) ListPlayer(params string) (ps []*Player, err error) {
	resp, _ := client.Call("listplayer", params)
	if resp.Code != "200" {
		return nil, errors.New(resp.Code)
	}
	err = json.Unmarshal([]byte(resp.Body), &ps)
	return
}

func (client *CenterClient) Broadcast(message string) error {
	m := &Message{Content: message}
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	resp, _ := client.Call("broadcast", string(b))
	if resp.Code == "200" {
		return nil
	}
	return errors.New(resp.Code)
}
