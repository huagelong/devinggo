// Package cache
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE
package cache

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
)

type TagCache struct {
	client               *gredis.Redis
	setScript            string
	deleteScript         string
	invalidateTagsScript string
}

func NewTagCache(ctx context.Context, client *gredis.Redis) (*TagCache, error) {
	const setScriptSrc = `
local function split(input, sep)
    if sep == nil then
        sep = "%s"
    end
    local t = {}
    for str in string.gmatch(input, "([^"..sep.."]+)") do
        t[#t+1] = str
    end
    return t
end
local key = KEYS[1]
local value = ARGV[1]
local expire = tonumber(ARGV[2])
local tags = ARGV[3]
local tagsArray = {}
if tags ~= "" then
    tagsArray = split(tags, ",")
end
redis.call("SET", key, value, "EX", expire)
for i=1, #tagsArray do
    redis.call("SADD", "tags:" .. tagsArray[i], key)
    redis.call("SADD", "item_tags:" .. key, tagsArray[i])
end
`

	const deleteScriptSrc = `
local key = KEYS[1]
local tagsCmd = redis.call("SMEMBERS", "item_tags:" .. key)
for i=1, #tagsCmd do
    redis.call("SREM", "tags:" .. tagsCmd[i], key)
end
redis.call("DEL", key)
redis.call("DEL", "item_tags:" .. key)
`

	const invalidateTagsScriptSrc = `
local tags = ARGV
for i=1, #tags do
    local keys = redis.call("SMEMBERS", "tags:" .. tags[i])
    for j=1, #keys do
        redis.call("DEL", keys[j])
		redis.call("SREM", "item_tags:" .. keys[j], tags[i])
    end
    redis.call("DEL", "tags:" .. tags[i])
end
`
	setScript, err := client.ScriptLoad(ctx, setScriptSrc)
	if err != nil {
		return nil, err
	}
	deleteScript, err := client.ScriptLoad(ctx, deleteScriptSrc)
	if err != nil {
		return nil, err
	}
	invalidateTagsScript, err := client.ScriptLoad(ctx, invalidateTagsScriptSrc)
	if err != nil {
		return nil, err
	}
	return &TagCache{client: client, setScript: setScript, deleteScript: deleteScript, invalidateTagsScript: invalidateTagsScript}, nil
}

func (c *TagCache) GetClient() *gredis.Redis {
	return c.client
}

func (c *TagCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration, tags []string) error {
	keys := []string{key}
	numkey := gconv.Int64(len(keys))
	tagsStr := strings.Join(tags, ",")
	var interfaceSlice []interface{}
	interfaceSlice = append(interfaceSlice, value)
	interfaceSlice = append(interfaceSlice, gconv.Int(expiration.Seconds()))
	interfaceSlice = append(interfaceSlice, tagsStr)
	_, err := c.client.EvalSha(ctx, c.setScript, numkey, keys, interfaceSlice)
	return err
}

func (c *TagCache) Get(ctx context.Context, key string) (*gvar.Var, error) {
	res, err := c.client.Get(ctx, key)
	if err != nil {
		if err == redis.Nil {
			err = c.cleanUpItemTags(ctx, key)
			if err != nil {
				return nil, err
			}
		}
		return nil, err
	}

	if g.IsEmpty(res) {
		err = c.cleanUpItemTags(ctx, key)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (c *TagCache) Delete(ctx context.Context, key string) error {
	keys := []string{key}
	numkeys := gconv.Int64(len(keys))
	_, err := c.client.EvalSha(ctx, c.deleteScript, numkeys, keys, nil)
	return err
}

func (c *TagCache) InvalidateTags(ctx context.Context, tags []string) error {
	var interfaceSlice []interface{}
	for i := 0; i < len(tags); i++ {
		interfaceSlice = append(interfaceSlice, tags[i])
	}
	_, err := c.client.EvalSha(ctx, c.invalidateTagsScript, 0, []string{}, interfaceSlice)
	return err
}

func (c *TagCache) cleanUpItemTags(ctx context.Context, key string) error {
	tags, err := c.client.SMembers(ctx, "item_tags:"+key)
	if err != nil && err != redis.Nil {
		return err
	}
	if !g.IsEmpty(tags) {
		for _, tag := range tags.Strings() {
			_, err = c.client.SRem(ctx, "tags:"+tag, key)
			if err != nil {
				return err
			}
		}
	}
	_, err = c.client.Del(ctx, "item_tags:"+key)
	if err != nil {
		return err
	}
	return nil
}
