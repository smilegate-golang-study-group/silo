# hello-world

- 환경: windows 10 / WSL2 (Ubuntu-20.04 zsh)
- go version: v1.18
- IDE: VScode

## go 초기 설치 방법
- 공식 문서: https://go.dev/doc/install 
- 하지만 버전 관리가 어려워보임

## golang 버전 관리

- v1.18 이전 버전을 사용하는 경우가 많음
- gvm (go version manager)로 버전 관리 쉽게 가능
- 참고: https://goax.tistory.com/7

```bash
# gvm 설치
❯ zsh < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
# 터미널 새 창으로 열기 or source 필요
❯ source ~/.gvm/scripts/gvm

# 기존 버전
❯ go version
go version go1.18 linux/amd64

# 설치 가능한 버전 확인 가능
❯ gvm listall

# 1.17 버전 설치 & 사용
❯ gvm install go1.17
❯ gvm use go1.17
Now using version go1.17

# 1.17 버전 반영 (단 새 터미널을 열면 default 버전으로 반영되니 확인 필요)
❯ go version
go version go1.17 linux/amd64
```

---

## 프로젝트 초기 설정

- 패키지명, 사용 모듈 초기화 필요
- main.go 초기 작성

```bash
# hello-world 라는 이름의 패키지 생성
❯ go mod init hello-world

# 사용하는 외부 모듈이 있는 경우 실행
❯ go mod tidy
# 외부 모듈이 있는 경우 - go.sum에 종속성 정보가 저장됩니다.
# hello-world의 경우 기본 패키지인 fmt만 사용하기에 생성되지 않음.
```
