-- Define a Lua table to store cached responses
local cache = {}

-- Function to handle request
function handle_request(request_handle)
    local key = request_handle:headers():get(":path")
    
    -- Check if there is a cached response available for this request
    if cache[key] ~= nil then
        -- Check if the cached response is still fresh
        local cached_time = cache[key]["cached_time"]
        local current_time = os.time()
        local cache_max_age = cache[key]["cache_max_age"]
        if current_time - cached_time <= cache_max_age then
            -- Set cached response headers and body
            request_handle:headers():add("cache-control", "public, max-age=" .. cache_max_age)
            request_handle:headers():copyFrom(cache[key]["headers"])
            request_handle:streamInfo():setDynamicMetadata("envoy.filters.http.lua", "response_cached", true)
            
            -- Log and return cached response
            request_handle:logInfo("Cached response served")
            return
        end
    end
    
    -- Cache the request
    request_handle:logInfo("No cached response found, forwarding request to backend")
    request_handle:streamInfo():setDynamicMetadata("envoy.filters.http.lua", "cache_enabled", true)
end

-- Function to handle response
function handle_response(response_handle)
    local cache_enabled = response_handle:streamInfo():dynamicMetadata():get("envoy.filters.http.lua")["cache_enabled"]
    
    -- Check if caching is enabled for the request
    if cache_enabled == true then
        local key = response_handle:streamInfo():requestHeaders():get(":path")
        
        -- Cache the response
        local response_headers = {}
        response_handle:headers():iterate(function(name, value)
            response_headers[name] = value
        end)
        local response_body = response_handle:body():getBytes(0, response_handle:body():length())
        local cache_max_age = tonumber(response_handle:streamInfo():dynamicMetadata():get("envoy.filters.http.lua")["cache_max_age"])
        
        cache[key] = {
            headers = response_headers,
            body = response_body,
            cached_time = os.time(),
            cache_max_age = cache_max_age
        }
        
        -- Set cache-control header in response
        response_handle:headers():add("cache-control", "public, max-age=" .. cache_max_age)
        
        -- Log and cache the response
        response_handle:logInfo("Caching enabled for the response")
        response_handle:streamInfo():setDynamicMetadata("envoy.filters.http.lua", "response_cached", true)
    end
end
