export PATH := /usr/local/go/bin:$(PATH) # for go

BUILD_DATE=`date  +%Y%m%d.%H%M%S`
LAST_COMMIT_SHA=`git show --pretty=format:%h --no-patch`

GOBIN=go

# $(GOBIN) env -w CGO_ENABLED=1 구문이 필요한 이유는?
#  - go-sqlite3 패키지를 사용하기 위해서
#  - CGO_ENABLED 값이 0인 경우에 발생하는 에러 메시지 : Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work.
#  - cgo 활성를 위해서는 gcc도 필요하기 때문에 /ndlpsdk2/bin 경로도 PATH에 추가해주어야함
all:
	$(GOBIN) env -w CGO_ENABLED=1
	GOOS=linux $(GOBIN) build -buildvcs=false -ldflags="-X 'main.version=v1.0.0-$(LAST_COMMIT_SHA)' -X 'main.bdate=$(BUILD_DATE)'" -o ./bin/wavynoted ./cmd/wavynoted;
	
clean:
	cd bin/; /bin/rm -rf wavynoted