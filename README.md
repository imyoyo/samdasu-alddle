MSA 교육중 만든 프로젝트를 잘 정리해본다. Go, gRPC, docker 를 쓸 때 참고할꼬얌

#### 개발 환경 설정
- $GOPATH/src 에 클론을 받음. Go 의 법칙이다. 예) ~/go/src/samdasu-alddle 
- [발표자료](https://docs.google.com/presentation/d/1iXq9_UP0MCULyLkZuWbQvrpcBuj42JE1NvMfl7nelMQ/edit#slide=id.g3c6e259a04_0_11) 참고.

#### 로컬에서 프로세스로 실행하기
1. 의존성 다운로드 : resolve-dependencies.sh
1. protobuf 생성 : protoc.sh
1. 각 티어 빌드 : (server,client,scheduler)/build.sh
1. 빌드된 실행파일을 실행.

#### 로컬에서 도커로 실행하기
1. 빌드 : docker-build.sh
2. 실행 : docker-run.sh
