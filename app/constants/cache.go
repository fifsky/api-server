package constants

// CacheKey 缓存key
const (
	CachedStaffKey string = "cached_model:staff:%s"
	CachedRoleKey  string = "cached_model:role:%s"

	CacheMainStaffInfoKey       string = "cached_model:corp:%s:dept:%s:offset:%d:limit:%d"
	CacheMainStaffInfoKeyPrefix string = "cached_model:corp:%s*"

	CacheCustomerSummaryKey string = "cached_model:corp:%s:staff:%s"

	StaffIDConverterKey string = "staff_id_converter"
)

const (
	DelCacheMainStaffInfoKeyScripts string = `
	local cursor = '0'
	repeat
    local result = redis.call(KEYS[1], cursor，'MATCH', ARGV[1], 'COUNT', 100)
    cursor = result[1]
	local keys = result[2]
	for _, key in ipairs(keys) do
		redis.call('DEL', key)
	end
	until cursor == '0'
	`
)
