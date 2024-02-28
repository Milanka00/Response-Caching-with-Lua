# Response-Caching-with-Lua
Implement response caching in gateways using Lua filters

                        -- Lua filter configuration
                        -- Called on the response path.
                        function envoy_on_response(response_handle)
                          local response_body = response_handle:body()

                          local redis = require("resty.redis")
                          local red = redis:new()

                          red:set_timeout(1000)  -- 1 second timeout

                          -- Connecting to Redis
                          local ok, err = red:connect("redis", 6379)
                          if not ok then
                            ngx.log(ngx.ERR, "failed to connect to redis: ", err)
                            return
                          end

                          -- Saving the response to Redis
                          local res, err = red:set("response_key", response_body)
                          if not res then
                            ngx.log(ngx.ERR, "failed to save the response in redis: ", err)
                          end

                          -- Closing the Redis connection
                          red:close()
                        end
                         
                         ---------------------------
                        
                        -- Lua filter configuration
                          -- Called on the response path.
                          function envoy_on_response(response_handle)
                              print("inside the response path")
                              local response_body = response_handle:body()
                              print("extracted the body")
                              local redis = require("resty.redis")  -- error here
                              local red = redis:new()
                              print("new redis connection")

                              red:set_timeout(1000)  -- 1 second timeout

                              -- Connecting to Redis
                              local ok, err = red:connect("redis", 6379)
                              if not ok then
                                  print("failed to connect to redis: ", err)
                                  return
                              else
                                  print("connected to redis successfully")
                              end

                              -- Saving the response to Redis
                              local res, err = red:set("response_key", response_body)
                              if not res then
                                  print("failed to save the response in redis: ", err)
                              else
                                  print("response saved in redis successfully")
                              end

                              -- Closing the Redis connection
                              local ok, err = red:close()
                              if not ok then
                                  print("failed to close redis connection: ", err)
                              else
                                  print("redis connection closed successfully")
                              end
                          end
                envoy_redis-1     | inside the response path
envoy_redis-1     | extracted the body
envoy_redis-1     | [2024-02-28 10:54:18.039][26][error][lua] [source/extensions/filters/http/lua/lua_filter.cc:802] script log: /usr/local/share/lua/5.3/resty/redis.lua:6: attempt to index global 'ngx' (a nil value)