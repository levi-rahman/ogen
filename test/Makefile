test_sync:
	go test -tags sync_test -c integration/sync_test.go
	./sync.test
	make clean

test_backup:
	go test -tags backup_test -c integration/backup_test.go
	./backup.test
	make clean

test_rpc:
	go test -tags rpc_test -c functional/rpc_test.go
	./rpc.test
	make clean

test_chain:
	go test -tags chain_test -c functional/chain_test.go
	./chain.test
	make clean

test_mempools:
	go test -tags mempools_test -c functional/mempools_test.go
	./mempools.test
	make clean
	
clean:
	rm -rf data_*	
	rm -rf *.test