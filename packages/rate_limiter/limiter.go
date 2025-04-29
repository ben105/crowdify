package rate_limiter

import (
    "context"
    "fmt"
    "time"
    "github.com/redis/go-redis/v9"
    "github.com/google/uuid"
)

type RateLimiter interface {
	func AllowRequest() bool
}

type RedisRateLimiter struct {
	r *redis.Client
	key string
	ctx context.Context
	window time.Duration
	limit int
}

func NewRateLimiter(key string) RateLimiter {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &RedisRateLimiter{
		r: rdb,
		key: key,
		ctx: context.Background(),
		window: time.Second * 30,
		limit: 100,
	}
}

func (rl *RedisRateLimiter) AllowRequest() bool {
	now := time.Now().UnixMilli()
	member := uuid.New().String()

	script := redis.NewScript(`
		-- KEYS[1]: redis key
		-- ARGV[1]: current timestamp in milliseconds
		-- ARGV[2]: window size in ms
		-- ARGV[3]: max requests
		-- ARGV[4]: member ID (UUID )

		redis.call("ZREMRANGEBYSCORE", KEYS[1], "-inf", ARGV[1] - ARGV[2])

		local count = redis.call("ZCARD", KEYS[1])

		if tonumber(count) < tonumber(ARGV[3]) then
			redis.call("ZADD", KEYS[1], ARGV[1], ARGV[4])
			redis.call("PEXPIRE", KEYS[1], ARGV[2])
			return 1
		else
			return 0
		end
	`)
	res, err := script.Run(ctx, rl.r, []string{rl.key}, now, rl.window.Milliseconds(), rl.limit, member).Int()
    if err != nil {
        return false, err
    }

    return res == 1, nil
}
