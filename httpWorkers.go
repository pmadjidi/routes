package main

import "log"

func (s *system) httpWorkers(workerId int) {
	log.Printf("worker (%d) starting....\n ", workerId)
loop:
	for {
		select {
		case req := <-s.apiRequest:
			url := s.apiUrl + string(req.route.src) + ";" + string(req.route.dst) + "?overview=false"
			log.Printf("worker (%d) calling url (%s)\n ", workerId, url)
			thisBody, err := s.apiCallHttp(url)
			if err != nil {
				req.inError = true
				req.msg = err.Error()
				req.resp <- req
			} else {
				req.body = thisBody
				req.resp <- req
				thisCacheEntrie := newCacheEntry(thisBody)
				cacheRequest := s.newCacheRequest(req.route.src, req.route.dst)
				cacheRequest.val = thisCacheEntrie
				s.setCache[cacheRequest.bucket] <- cacheRequest
			}
		case <-s.shutDown:
			break loop
		}
	}
	log.Printf("worker (%d) exiting....\n ", workerId)

}
