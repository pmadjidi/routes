package main

import (
	"log"
	"time"
)

type CacheRequest struct {
	src    latlong
	dst    latlong
	bucket int
	val    *CacheEntry
	resp   chan []byte
}

type CacheEntry struct {
	stamp time.Time
	value []byte
}

func (s *system) newCacheRequest(src, dst latlong) *CacheRequest {
	bucket := s.hash(getKeyFromLatLong(src, dst))
	return &CacheRequest{src, dst, bucket, nil, make(chan []byte, 1)}
}

func newCacheEntry(val []byte) *CacheEntry {
	return &CacheEntry{time.Now(), val}
}

func (s *system) cacheProcessor(workerId int) {
	ticker := time.NewTicker(CACHEPRUNEFREQUENCY * time.Second)
	log.Printf("cache processor (%d) starting....\n ", workerId)
	localCache := make(map[string]CacheEntry)
	expirationTime, err := time.ParseDuration(CACHEEXPIRATION)
	if err != nil {
		expirationTime = 60 * time.Second
	}
	log.Printf("cache processor (%d) setting cache expiration time to (%s)....\n ", workerId, expirationTime)

loop:
	for {
		select {
		case req := <-s.cacheRequest[workerId]:
			log.Printf("cache processor(%d) looking for src,dst (%s,%s)\n ", workerId, req.src, req.dst)
			key := getKeyFromLatLong(req.src, req.dst)
			hit, ok := localCache[key]
			if ok {
				log.Printf("cache processor(%d) hit for key (%s)\n ", workerId, key)
				req.resp <- hit.value
			} else {
				log.Printf("cache processor(%d) miss for key (%s)\n ", workerId, key)
				req.resp <- nil
			}
		case <-ticker.C:
			now := time.Now()
			log.Printf("cache processor (%d) pruneing local cache...", workerId)
			for key, val := range localCache {
				if now.Sub(val.stamp) > expirationTime {
					log.Printf("cache processor (%d): Key (%s) is expired, removing...", workerId, key)
					delete(localCache, key)
				}
			}
			log.Printf("cache processor (%d) pruneing done...", workerId)

		case req := <-s.setCache[workerId]:
			key := getKeyFromLatLong(req.src, req.dst)
			log.Printf("cache processor(%d) setting key to val: (%s,%d)\n ", workerId, key, len(req.val.value))
			localCache[key] = *req.val

		case <-s.shutDown:
			break loop

		}
	}
	log.Printf("cache processor (%d) exiting....\n ", workerId)

}
