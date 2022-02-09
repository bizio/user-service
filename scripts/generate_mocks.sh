#!/bin/bash
PWD=$(pwd)

mockgen  -package=mock -self_package=github.com/bizio/user-service -destination="$PWD/test/mock/mock_datastore.go" \
    github.com/bizio/user-service/pkg/service/v1/data Datastore

mockgen  -package=mock -self_package=github.com/bizio/user-service -destination="$PWD/test/mock/mock_messagequeue.go" \
    github.com/bizio/user-service/pkg/service/v1/cloudpubsub Queue

mockgen  -package=mock -self_package=github.com/bizio/wa-srv-base -destination="$PWD/test/mock/mock_cache.go" \
    github.com/bizio/wa-srv-base/cache CacheService
