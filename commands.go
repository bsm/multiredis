package multiredis

import (
	"time"

	"gopkg.in/redis.v3"
)

// Commands exposes all supported commands
type Commands interface {
	Append(key, value string) *redis.IntCmd
	BitCount(key string, bitCount *redis.BitCount) *redis.IntCmd
	BitOpAnd(destKey string, keys ...string) *redis.IntCmd
	BitOpNot(destKey string, key string) *redis.IntCmd
	BitOpOr(destKey string, keys ...string) *redis.IntCmd
	BitOpXor(destKey string, keys ...string) *redis.IntCmd
	BitPos(key string, bit int64, pos ...int64) *redis.IntCmd
	ClientKill(ipPort string) *redis.StatusCmd
	ClientList() *redis.StringCmd
	ClientPause(dur time.Duration) *redis.BoolCmd
	ClientSetName(name string) *redis.BoolCmd
	ClusterAddSlots(slots ...int) *redis.StatusCmd
	ClusterAddSlotsRange(min, max int) *redis.StatusCmd
	ClusterCountFailureReports(nodeID string) *redis.IntCmd
	ClusterCountKeysInSlot(slot int) *redis.IntCmd
	ClusterDelSlots(slots ...int) *redis.StatusCmd
	ClusterDelSlotsRange(min, max int) *redis.StatusCmd
	ClusterFailover() *redis.StatusCmd
	ClusterForget(nodeID string) *redis.StatusCmd
	ClusterInfo() *redis.StringCmd
	ClusterKeySlot(key string) *redis.IntCmd
	ClusterMeet(host, port string) *redis.StatusCmd
	ClusterNodes() *redis.StringCmd
	ClusterReplicate(nodeID string) *redis.StatusCmd
	ClusterResetHard() *redis.StatusCmd
	ClusterResetSoft() *redis.StatusCmd
	ClusterSaveConfig() *redis.StatusCmd
	ClusterSlaves(nodeID string) *redis.StringSliceCmd
	ClusterSlots() *redis.ClusterSlotCmd
	ConfigGet(parameter string) *redis.SliceCmd
	ConfigResetStat() *redis.StatusCmd
	ConfigSet(parameter, value string) *redis.StatusCmd
	DbSize() *redis.IntCmd
	DebugObject(key string) *redis.StringCmd
	Decr(key string) *redis.IntCmd
	DecrBy(key string, decrement int64) *redis.IntCmd
	Del(keys ...string) *redis.IntCmd
	Dump(key string) *redis.StringCmd
	Echo(message string) *redis.StringCmd
	Eval(script string, keys []string, args []string) *redis.Cmd
	EvalSha(sha1 string, keys []string, args []string) *redis.Cmd
	Exists(key string) *redis.BoolCmd
	Expire(key string, expiration time.Duration) *redis.BoolCmd
	ExpireAt(key string, tm time.Time) *redis.BoolCmd
	FlushAll() *redis.StatusCmd
	FlushDb() *redis.StatusCmd
	GeoAdd(key string, geoLocation ...*redis.GeoLocation) *redis.IntCmd
	GeoDist(key string, member1, member2, unit string) *redis.FloatCmd
	GeoHash(key string, members ...string) *redis.StringSliceCmd
	GeoRadius(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd
	GeoRadiusByMember(key, member string, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd
	Get(key string) *redis.StringCmd
	GetBit(key string, offset int64) *redis.IntCmd
	GetRange(key string, start, end int64) *redis.StringCmd
	GetSet(key string, value interface{}) *redis.StringCmd
	HDel(key string, fields ...string) *redis.IntCmd
	HExists(key, field string) *redis.BoolCmd
	HGet(key, field string) *redis.StringCmd
	HGetAll(key string) *redis.StringSliceCmd
	HGetAllMap(key string) *redis.StringStringMapCmd
	HIncrBy(key, field string, incr int64) *redis.IntCmd
	HIncrByFloat(key, field string, incr float64) *redis.FloatCmd
	HKeys(key string) *redis.StringSliceCmd
	HLen(key string) *redis.IntCmd
	HMGet(key string, fields ...string) *redis.SliceCmd
	HMSet(key, field, value string, pairs ...string) *redis.StatusCmd
	HScan(key string, cursor int64, match string, count int64) *redis.ScanCmd
	HSet(key, field, value string) *redis.BoolCmd
	HSetNX(key, field, value string) *redis.BoolCmd
	HVals(key string) *redis.StringSliceCmd
	Incr(key string) *redis.IntCmd
	IncrBy(key string, value int64) *redis.IntCmd
	IncrByFloat(key string, value float64) *redis.FloatCmd
	Info(section ...string) *redis.StringCmd
	Keys(pattern string) *redis.StringSliceCmd
	LIndex(key string, index int64) *redis.StringCmd
	LInsert(key, op, pivot, value string) *redis.IntCmd
	LLen(key string) *redis.IntCmd
	LPop(key string) *redis.StringCmd
	LPush(key string, values ...string) *redis.IntCmd
	LPushX(key, value interface{}) *redis.IntCmd
	LRange(key string, start, stop int64) *redis.StringSliceCmd
	LRem(key string, count int64, value interface{}) *redis.IntCmd
	LSet(key string, index int64, value interface{}) *redis.StatusCmd
	LTrim(key string, start, stop int64) *redis.StatusCmd
	LastSave() *redis.IntCmd
	MGet(keys ...string) *redis.SliceCmd
	MSet(pairs ...string) *redis.StatusCmd
	MSetNX(pairs ...string) *redis.BoolCmd
	Migrate(host, port, key string, db int64, timeout time.Duration) *redis.StatusCmd
	Move(key string, db int64) *redis.BoolCmd
	ObjectEncoding(keys ...string) *redis.StringCmd
	ObjectIdleTime(keys ...string) *redis.DurationCmd
	ObjectRefCount(keys ...string) *redis.IntCmd
	PExpire(key string, expiration time.Duration) *redis.BoolCmd
	PExpireAt(key string, tm time.Time) *redis.BoolCmd
	PFAdd(key string, fields ...string) *redis.IntCmd
	PFCount(keys ...string) *redis.IntCmd
	PFMerge(dest string, keys ...string) *redis.StatusCmd
	PTTL(key string) *redis.DurationCmd
	Persist(key string) *redis.BoolCmd
	Ping() *redis.StatusCmd
	Process(cmd redis.Cmder)
	Quit() *redis.StatusCmd
	RPop(key string) *redis.StringCmd
	RPopLPush(source, destination string) *redis.StringCmd
	RPush(key string, values ...string) *redis.IntCmd
	RPushX(key string, value interface{}) *redis.IntCmd
	RandomKey() *redis.StringCmd
	ReadWrite() *redis.StatusCmd
	Readonly() *redis.StatusCmd
	Rename(key, newkey string) *redis.StatusCmd
	RenameNX(key, newkey string) *redis.BoolCmd
	Restore(key string, ttl time.Duration, value string) *redis.StatusCmd
	RestoreReplace(key string, ttl time.Duration, value string) *redis.StatusCmd
	SAdd(key string, members ...string) *redis.IntCmd
	SCard(key string) *redis.IntCmd
	SDiff(keys ...string) *redis.StringSliceCmd
	SDiffStore(destination string, keys ...string) *redis.IntCmd
	SInter(keys ...string) *redis.StringSliceCmd
	SInterStore(destination string, keys ...string) *redis.IntCmd
	SIsMember(key string, member interface{}) *redis.BoolCmd
	SMembers(key string) *redis.StringSliceCmd
	SMove(source, destination string, member interface{}) *redis.BoolCmd
	SPop(key string) *redis.StringCmd
	SRandMember(key string) *redis.StringCmd
	SRandMemberN(key string, count int64) *redis.StringSliceCmd
	SRem(key string, members ...string) *redis.IntCmd
	SScan(key string, cursor int64, match string, count int64) *redis.ScanCmd
	SUnion(keys ...string) *redis.StringSliceCmd
	SUnionStore(destination string, keys ...string) *redis.IntCmd
	Save() *redis.StatusCmd
	Scan(cursor int64, match string, count int64) *redis.ScanCmd
	ScriptExists(scripts ...string) *redis.BoolSliceCmd
	ScriptFlush() *redis.StatusCmd
	ScriptKill() *redis.StatusCmd
	ScriptLoad(script string) *redis.StringCmd
	Select(index int64) *redis.StatusCmd
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	SetBit(key string, offset int64, value int) *redis.IntCmd
	SetNX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd
	SetRange(key string, offset int64, value string) *redis.IntCmd
	Shutdown() *redis.StatusCmd
	ShutdownNoSave() *redis.StatusCmd
	ShutdownSave() *redis.StatusCmd
	SlaveOf(host, port string) *redis.StatusCmd
	SlowLog()
	Sort(key string, sort redis.Sort) *redis.StringSliceCmd
	SortInterfaces(key string, sort redis.Sort) *redis.SliceCmd
	StrLen(key string) *redis.IntCmd
	TTL(key string) *redis.DurationCmd
	Time() *redis.StringSliceCmd
	Type(key string) *redis.StatusCmd
	ZAdd(key string, members ...redis.Z) *redis.IntCmd
	ZAddCh(key string, members ...redis.Z) *redis.IntCmd
	ZAddNX(key string, members ...redis.Z) *redis.IntCmd
	ZAddNXCh(key string, members ...redis.Z) *redis.IntCmd
	ZAddXX(key string, members ...redis.Z) *redis.IntCmd
	ZAddXXCh(key string, members ...redis.Z) *redis.IntCmd
	ZCard(key string) *redis.IntCmd
	ZCount(key, min, max string) *redis.IntCmd
	ZIncr(key string, member redis.Z) *redis.FloatCmd
	ZIncrBy(key string, increment float64, member string) *redis.FloatCmd
	ZIncrNX(key string, member redis.Z) *redis.FloatCmd
	ZIncrXX(key string, member redis.Z) *redis.FloatCmd
	ZInterStore(destination string, store redis.ZStore, keys ...string) *redis.IntCmd
	ZRange(key string, start, stop int64) *redis.StringSliceCmd
	ZRangeByLex(key string, opt redis.ZRangeByScore) *redis.StringSliceCmd
	ZRangeByScore(key string, opt redis.ZRangeByScore) *redis.StringSliceCmd
	ZRangeByScoreWithScores(key string, opt redis.ZRangeByScore) *redis.ZSliceCmd
	ZRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd
	ZRank(key, member string) *redis.IntCmd
	ZRem(key string, members ...string) *redis.IntCmd
	ZRemRangeByRank(key string, start, stop int64) *redis.IntCmd
	ZRemRangeByScore(key, min, max string) *redis.IntCmd
	ZRevRange(key string, start, stop int64) *redis.StringSliceCmd
	ZRevRangeByLex(key string, opt redis.ZRangeByScore) *redis.StringSliceCmd
	ZRevRangeByScore(key string, opt redis.ZRangeByScore) *redis.StringSliceCmd
	ZRevRangeByScoreWithScores(key string, opt redis.ZRangeByScore) *redis.ZSliceCmd
	ZRevRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd
	ZRevRank(key, member string) *redis.IntCmd
	ZScan(key string, cursor int64, match string, count int64) *redis.ScanCmd
	ZScore(key, member string) *redis.FloatCmd
	ZUnionStore(dest string, store redis.ZStore, keys ...string) *redis.IntCmd
}
