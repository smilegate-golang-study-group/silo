# 1주차 Study

## 환경 구성
- go1.18 v 설치(windows 10)


## Go Module
- 패키지 버전 묶음
- go1.11 버전에서 패키지 버전 및 의존성 관리를 지원하는 go mod가 도입됨.

### go.mod
- 패키지 실행에 필요한 **모듈 정보와 버전을 관리**하는 파일
- 현재 모듈 정보와 코드에서 사용하는 외부 패키지 의존성 정보를 포함
- 루트 경로에 존재

### go.sum
- dependency **모듈의 버전 및 해시** 저장
- go command 실행 시 go.mod에 정의된 모듈을 다운로드할 때 모듈의 checksum이 저장됨.
- 로컬에 설치된 모듈의 해시와 go.sum에 저장된 해시를 비교하여 **유효성을 검증**

```
# 모듈 생성
go mod init <module-name>

# source 파일에 필요한 dependency를 다운로드하고 go.mod 파일 업데이트
go mod tidy

# go.sum 파일에 기록된 checksum과 캐시된 모듈 복사본의 일치 여부 확인
go mod verify

# 새로운 모듈을 버전을 지정하여 추가
go get <module-path>@<module-query>
```
