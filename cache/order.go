package cache

//
//func getStock(eventID int) (int, error) {
//	key := fmt.Sprintf("ticket_stock:%d", eventID)
//	stock, err := RedisClient.Get(key).Int()
//	if err != nil {
//		return 0, err
//	}
//	return stock, nil
//}
//
//func decrementStock(eventID int) (bool, error) {
//	key := fmt.Sprintf("ticket_stock:%d", eventID)
//
//	luaScript := redis.NewScript(`
//        if (redis.call('exists', KEYS[1]) == 1) then
//            local stock = tonumber(redis.call('get', KEYS[1]))
//            if (stock <= 0) then
//                return 0
//            else
//                redis.call('decr', KEYS[1])
//                return stock - 1
//            end
//        else
//            return 0
//        end
//    `)
//
//	result, err := luaScript.Run(RedisClient, []string{key}).Int()
//	if err != nil {
//		return false, err
//	}
//
//	return result >= 0, nil
//}
