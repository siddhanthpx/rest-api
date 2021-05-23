package rediscache

import (
	"bytes"
	"context"
	"encoding/gob"
	"os"
	"rest-api/data"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type Client struct {
	client *redis.Client
}

func NewClient() (*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:        os.Getenv("REDIS"),
		DB:          0,
		DialTimeout: 100 * time.Millisecond,
		ReadTimeout: 100 * time.Millisecond,
	})

	if _, err := client.Ping(context.TODO()).Result(); err != nil {
		return nil, err
	}

	return &Client{
		client: client,
	}, nil
}

func (c *Client) GetGame(ctx context.Context, id int) (*data.Game, error) {

	cmd := c.client.Get(ctx, strconv.Itoa(id))

	cmdb, err := cmd.Bytes()
	if err != nil {
		return nil, err
	}

	b := bytes.NewReader(cmdb)

	var game data.Game

	if err := gob.NewDecoder(b).Decode(&game); err != nil {
		return nil, err
	}

	return &game, nil

}

func (c *Client) SetGame(ctx context.Context, g data.Game) error {

	var b bytes.Buffer

	if err := gob.NewEncoder(&b).Encode(g); err != nil {
		return err
	}

	return c.client.Set(ctx, strconv.Itoa(g.ID), b.Bytes(), 25*time.Second).Err()

}
