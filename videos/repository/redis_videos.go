package repository

import (
	"bytes"
	"encoding/gob"
	"errors"
	models "github.com/dbond762/my-player/videos"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

const (
	sleepTime   = 250 * time.Millisecond
	maxIdle     = 3
	idleTimeout = 240 * time.Second
	ttl         = 24 * 60 * 60
)

var (
	VideosNotInCache = errors.New("videos not in cache")
)

type RedisVideosRepository interface {
	GetVideosFromCache(query string) ([]models.Video, error)
	AddVideosToCache(query string, videos []models.Video) error
}

type redisVideosRepository struct {
	Pool *redis.Pool
}

func NewRedisVideosRepository(addr string) *redisVideosRepository {
	return &redisVideosRepository{
		Pool: &redis.Pool{
			MaxIdle:     maxIdle,
			IdleTimeout: idleTimeout,
			Dial: func() (redis.Conn, error) {
				return redis.DialURL(addr)
			},
		},
	}
}

func (r *redisVideosRepository) GetVideosFromCache(query string) ([]models.Video, error) {
	conn := r.Pool.Get()
	defer conn.Close()

	reply, err := redis.Int(conn.Do("EXISTS", "lock_query_"+query))
	if err != nil && err != redis.ErrNil {
		log.Print(err)
		return []models.Video{}, err
	}
	if reply == 1 {
		time.Sleep(sleepTime)
	}

	res, err := redis.Bytes(conn.Do("GET", "query_"+query))
	if err == redis.ErrNil {
		return []models.Video{}, VideosNotInCache
	} else if err != nil {
		log.Print(err)
		return []models.Video{}, err
	}

	var videos []models.Video

	decoder := gob.NewDecoder(bytes.NewBuffer(res))
	err = decoder.Decode(videos)
	if err != nil {
		log.Print(err)
		return []models.Video{}, err
	}

	log.Print("Redis: Search from cache")

	return videos, nil
}

func (r *redisVideosRepository) AddVideosToCache(query string, videos []models.Video) error {
	conn := r.Pool.Get()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Print(err)
		}
	}()

	_, err := conn.Do("SET", "lock_query_"+query, 1)
	if err != nil {
		return err
	}
	defer func() {
		if _, err := conn.Do("DEL", "lock_query_"+query); err != nil {
			log.Print(err)
		}
	}()

	var buf bytes.Buffer

	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(videos); err != nil {
		return err
	}

	if _, err := conn.Do("SET", "query_"+query, buf.Bytes(), "EX", ttl); err != nil {
		return err
	}

	log.Print("Redis: Search from api")

	return nil
}
