<H1>wavynote backend</H1>

<H2>목차</H2>

- [wavynote REST API](#wavynote-rest-api)
- [Flowchart 기반 REST API 리스트업](#flowchart-기반-rest-api-리스트업)
- [요구사항](#요구사항)
- [요청](#요청)
- [응답](#응답)
- [Main 페이지](#main-페이지)
  - [존재하는 모든 폴더 목록 조회](#존재하는-모든-폴더-목록-조회)
  - [특정 폴더에 존재하는 모든 노트 조회](#특정-폴더에-존재하는-모든-노트-조회)
  - [특정 폴더 이름 변경](#특정-폴더-이름-변경)
  - [특정 폴더 삭제](#특정-폴더-삭제)
  - [내가 쓴 특정 노트 삭제](#내가-쓴-특정-노트-삭제)
  - [전체 폴더를 대상으로 노트 내용 검색](#전체-폴더를-대상으로-노트-내용-검색)
  - [특정 폴더를 대상으로 글 내용 검색](#특정-폴더를-대상으로-글-내용-검색)
- [Write 페이지](#write-페이지)
  - [내가 쓴 노트 저장](#내가-쓴-노트-저장)
  - [내가 쓴 노트를 특정 대상에게 보내기](#내가-쓴-노트를-특정-대상에게-보내기)
  - [내가 쓴 노트를 오픈 노트에 공유하기](#내가-쓴-노트를-오픈-노트에-공유하기)
  - [내가 쓴 노트를 랜덤 매칭을 통해 (비슷한 관심 주제를 갖는)임의의 대상의 노트로 보내기](#내가-쓴-노트를-랜덤-매칭을-통해-비슷한-관심-주제를-갖는임의의-대상의-노트로-보내기)
  - [내가 쓴 노트 조회](#내가-쓴-노트-조회)
- [Box 페이지](#box-페이지)
  - [노트를 주고 받는 대화방 목록 조회(최대 3개)](#노트를-주고-받는-대화방-목록-조회최대-3개)
  - [특정 친구와 주고받은 노트 목록](#특정-친구와-주고받은-노트-목록)
  - [특정 노트 조회](#특정-노트-조회)
  - [특정 대화방 삭제](#특정-대화방-삭제)
- [Profile 페이지](#profile-페이지)


## wavynote REST API
Wavy Note에서 제공할 백엔드 API를 설계 내용을 포함한다.

## Flowchart 기반 REST API 리스트업
https://www.figma.com/file/7fY7Jos8W7BYVs141IJsT1/wavy-note?type=design&node-id=76-152&mode=design&t=xxOKgOeHLWJe8P0d-0

## 요구사항
- Wavy note에서 제공하는 REST API는 SSL(Secure Sockets Layer)이 적용된 HTTPS 프로토콜로만 호출이 가능함

## 요청
- **메서드(Method)**: API 호출 시 사용해야할 HTTP 요청 메서드
- **IP 주소**: Wavy Note 백엔드 모듈이 구동 중인 서버의 IP 주소
- **포트 번호**: Wavy Note 백엔드 모듈이 열고 있는 REST API용 HTTPS 서버 포트
- **URL**: API를 통해 제공되는 리소스마다 지정된 요청 경로로 IP주소:PORT와 함께 각 API의 엔드포인트를 구성함. 반드시 Base URL(/wavynote/v1.0)을 추가해주어야 함
- **헤더(Header)**: API 호출 시 필요한 인증 정보(현재는 Basic Authentication 방식만을 제공함)를 전달하는데 사용 (예시: Authorization: Basic {ID:PW Base64 인코딩 데이터})
- **경로 변수(Path variable)**: API 호출 시 사용자가 전달한 값을 포함해 URL을 구성할 때 사용

## 응답
- REST API 항목 별로 응답 값을 별도로 정의함

## Main 페이지
### 존재하는 모든 폴더 목록 조회
- 메서드: GET
- URL: https://{ip}:{port}/wavynote/v1.0/main/folderlist
- 요청
  - 헤더
    - Authorization: Basic {ID:PW Base64 인코딩 데이터}
    - Content-Type: application/x-www-form-urlencoded
- 응답
  - 본문
    - `folder_id`: 폴더의 고유 id 값
    - `folder_name`: 폴더 이름
    - `note_count`: 폴더에 존재하는 노트의 총 개수
  - 응답 본문 예시
    ```json
    {
        "data": [
        {
          "folder_id": "a3106a0c-5ce7-40f6-81f4-ff9b8ebb240b",
          "folder_name": "생각정리",
          "note_count": "5"
        },
        {
          "folder_id": "980e71ba-0395-49aa-833e-3ebc76b3ec88",
          "folder_name": "나의웨이비노트",
          "note_count": "3"
        }
      ]
    }
    ```

### 특정 폴더에 존재하는 모든 노트 조회
*노트의 수가 많아지는 경우 페이징 처리는 어떻게 해야하는가?*

- 메서드: GET
- URL: https://{ip}:{port}/wavynote/v1.0/main/notelist?fid={folder_id}
- 요청
  - 헤더
    - Authorization: Basic {ID:PW Base64 인코딩 데이터}
    - Content-Type: application/x-www-form-urlencoded
  - 본문
    - `fid`: 폴더의 고유 id 값
    - ~~`pn`: 페이지 번호~~
    - ~~`mc`: 페이지별로 보여줄 최대 노트 개수~~
- 응답
  - 본문
    - `note_id`: 해당 폴더에 존재하는 노트의 고유 id 값
    - `title`: 해당 폴더에 존재하는 노트의 제목
    - `preview`: 해당 폴더에 존재하는 노트의 본문 미리보기(글자수 제한)
  - 응답 본문 예시
    ```json
    {
        "data": [
            {
                "note_id": "09d05df1-2958-4a3d-b910-3b4fb079327b",
                "title": "나의첫번째노트",
                "preview": "나의 첫 웨이비노트 본문 내용입니다.",
            },
            {
                "note_id": "1a092b35-dc9e-472e-be39-7391ca176040",
                "title": "나의두번째노트",
                "preview": "나의 두번째 웨이비노트 본문 내용입니다."
            }
        ]
    }
    ```

### 특정 폴더 이름 변경
- 메서드: POST
- URL: https://{ip}:{port}/wavynote/v1.0/main/folder
- 요청
  - 헤더
    - Authorization: Basic {ID:PW Base64 인코딩 데이터}
    - Content-Type: application/json
  - 본문
    - `folder_id`: 이름을 변경하고자 하는 폴더의 고유 id 값
    - `folder_name`: 변경할 폴더 이름
  - 요청 본문 예시
    ```json
    {
        "folder_id": "a3106a0c-5ce7-40f6-81f4-ff9b8ebb240b",
    	"folder_name": "생각정리",
    }
    ```
- 응답
  - 본문
    - `result`: 특정 폴더 이름 변경에 대한 성공/실패 여부(true/false)
    - `msg`: 실패 시 반환하는 에러 메시지
  - 응답 본문 예시
    ```json
    {
        "result": "true",
        "msg": ""
    }
    ```

### 특정 폴더 삭제
- 메서드: DELETE
- URL: https://{ip}:{port}/wavynote/v1.0/main/folder
- 요청
  - 헤더
    - Authorization: Basic {ID:PW Base64 인코딩 데이터}
    - Content-Type: application/json
  - 본문
    - `data`: 삭제할 폴더 정보를 배열 형태로 입력
      - `folder_id`: 삭제할 폴더의 고유 id 값
      - `folder_name`: 삭제할 폴더의 이름
  - 요청 본문 예시
    ```json
    {
      "data": [
        {
          "folder_id": "a3106a0c-5ce7-40f6-81f4-ff9b8ebb240b",
          "folder_name": "생각정리"
        },
        {
          "folder_id": "980e71ba-0395-49aa-833e-3ebc76b3ec88",
          "folder_name": "나의웨이비노트"
        }
      ]
    }
    ```
- 응답
  - 본문
    - `result`: 특정 폴더 삭제에 대한 성공/실패 여부(true/false)
    - `msg`: 실패 시 반환하는 에러 메시지
  - 응답 본문 예시
    ```json
    {
        "result": "true",
        "msg": ""
    }
    ```


### 내가 쓴 특정 노트 삭제
- 메서드: DELETE
- URL: https://{ip}:{port}/wavynote/v1.0/main/note
- 요청
  - 헤더
    - Authorization: Basic {ID:PW Base64 인코딩 데이터}
    - Content-Type: application/json
  - 본문
    - `data`: 삭제할 노트 정보를 배열 형태로 입력
      - `folder_id`: 삭제할 노트가 포함되어 있는 폴더의 고유 id 값
      - `note_id`: 삭제할 노트의 고유한 id 값
  - 요청 본문 예시
    ```json
    {
      "data": [
        {
          "folder_id": "a3106a0c-5ce7-40f6-81f4-ff9b8ebb240b",
          "note_id": "ebec3eaf-b2f3-45f9-9d7a-4488492f7bfc"
        },
        {
          "folder_id": "980e71ba-0395-49aa-833e-3ebc76b3ec88",
          "note_id": "09d05df1-2958-4a3d-b910-3b4fb079327b"
        }
      ]
    }
    ```
- 응답
  - 본문
    - `result`: 내가 쓴 특정 노트 삭제에 대한 성공/실패 여부(true/false)
    - `msg`: 실패 시 반환하는 에러 메시지
  - 응답 본문 예시
    ```json
    {
        "result": "true",
        "msg": ""
    }
    ```

### 전체 폴더를 대상으로 노트 내용 검색
- 메서드: GET
- URL: https://{ip}:{port}/wavynote/v1.0/search/top?query={검색어}
- 요청
  - 헤더
    - Authorization: Basic {ID:PW Base64 인코딩 데이터}
    - Content-Type: application/x-www-form-urlencoded
  - 본문
    - `query`: 검색어(노트의 제목과 본문에서 검색을 수행함)
- 응답
  - 본문
    - `data`: 검색된 노트 정보를 배열 형태로 전달
      - `note_id`: 검색어가 포함된 노트의 고유한 id 값
      - `title`: 검색어가 포함된 노트의 제목
      - `preview`: 검색어가 포함된 노트의 본문 미리보기
  - 응답 본문 예시
    ```json
    {
        "data": [
            {
                "note_id": "09d05df1-2958-4a3d-b910-3b4fb079327b",
                "title": "나의첫번째노트",
                "preview": "나의 첫 웨이비노트 본문 내용입니다.",
            },
            {
                "note_id": "1a092b35-dc9e-472e-be39-7391ca176040",
                "title": "나의두번째노트",
                "preview": "나의 두번째 웨이비노트 본문 내용입니다."
            }
        ]
    }
    ```

### 특정 폴더를 대상으로 글 내용 검색
- 메서드: GET
- URL: https://{ip}:{port}/wavynote/v1.0/search/folder?id={folderid}&query={검색어}
- 요청
  - 헤더
    - Authorization: Basic {ID:PW Base64 인코딩 데이터}
    - Content-Type: application/x-www-form-urlencoded
  - 본문
    - `id`: 검색을 수행할 폴더의 고유 id 값
    - `query`: 검색어(노트의 제목과 본문에서 검색을 수행함)
- 응답
  - 본문
    - `data`: 검색된 노트 정보를 배열 형태로 전달
      - `note_id`: 검색어가 포함된 노트의 고유한 id 값
      - `title`: 검색어가 포함된 노트의 제목
      - `preview`: 검색어가 포함된 노트의 본문 미리보기
  - 응답 본문 예시
    ```json
    {
        "data": [
            {
                "note_id": "09d05df1-2958-4a3d-b910-3b4fb079327b",
                "title": "나의첫번째노트",
                "preview": "나의 첫 웨이비노트 본문 내용입니다.",
            },
            {
                "note_id": "1a092b35-dc9e-472e-be39-7391ca176040",
                "title": "나의두번째노트",
                "preview": "나의 두번째 웨이비노트 본문 내용입니다."
            }
        ]
    }
    ```

## Write 페이지
### 내가 쓴 노트 저장
*대화 상대가 지정된 경우가 아님*<br>
*노트의 고유한 id 값은 저장 시점에 자동으로 생성(uuid)*<br>
*노트 별로 카테고리를 설정할 수 있는가?*<br>

- 메서드: POST
- URL: https://{ip}:{port}/wavynote/v1.0/write/save
- 요청
  - 헤더
    - Authorization: Basic {ID:PW Base64 인코딩 데이터}
    - Content-Type: application/json
  - 본문
    - `folder_id`: 내가 쓴 노트가 포함되어 있는 폴더의 고유 id 값
    - `from`: 작성자의 id
    - `save_at`: 노트 저장 시점의 timestamp 정보
    - `title`: 내가 쓴 노트의 제목
    - `content`: 내가 쓴 노트의 본문 내용
    - `keyword`: 내가 쓴 노트의 키워드
  - 요청 본문 예시
    ```json
    {
        "folder_id": "980e71ba-0395-49aa-833e-3ebc76b3ec88",
        "from": "wavynoteadmin",
        "save_at": "2023-11-01 21:00:00",
        "title": "나의첫번째노트",
        "content": "나의 첫 웨이비노트 본문 내용입니다.",
        "keyword": "일상"
    }
    ```
- 응답
  - 본문
    - `result`: 내가 쓴 노트 저장에 대한 성공/실패 여부(true/false)
    - `msg`: 실패 시 반환하는 에러 메시지
  - 응답 본문 예시
    ```json
    {
        "result": "true",
        "msg": ""
    }
    ```

### 내가 쓴 노트를 특정 대상에게 보내기
- 메서드: POST
- URL: https://{ip}:{port}/wavynote/v1.0/write/send
- 요청
  - 헤더
    - Authorization: Basic {ID:PW Base64 인코딩 데이터}
    - Content-Type: application/json
  - 본문
    - `note_id`: 내가 쓴 노트의 고유 id 값
    - `from`: 작성자의 id
    - `to`: 내가 쓴 노트를 보내는 대상의 id
    - `conversation_id`: 대화방의 고유 id 값
    - `send_at`: 보낸 시간
  - 요청 본문 예시
    ```json
    {
        "note_id": "09d05df1-2958-4a3d-b910-3b4fb079327b",
        "from": "wavynoteadmin",
        "to": "somebody",
        "conversation_id": "1afc571d-61bf-4cef-95ce-ab791f999297",
        "send_at": "2023-11-01 23:20:12"
    }
    ```
- 응답
  - 본문
    - `result`: 내가 쓴 노트 보내기에 대한 성공/실패 여부(true/false)
    - `msg`: 실패 시 반환하는 에러 메시지
  - 응답 본문 예시
    ```json
    {
        "result": "true",
        "msg": ""
    }
    ```

### 내가 쓴 노트를 오픈 노트에 공유하기
*대화방에 부여되는 고유한 id 값은 저장 시점에 자동으로 생성(uuid)*

- 메서드: POST
- URL: https://{ip}:{port}/wavynote/v1.0/write/opennote
- 요청
  - 헤더
    - Authorization: Basic {ID:PW Base64 인코딩 데이터}
    - Content-Type: application/json
  - 본문
    - `note_id`: 오픈 노트에 공유할 노트의 고유 id 값
    - `host_id`: 공유자의 id
  - 요청 본문 예시
    ```json
    {
        "note_id": "09d05df1-2958-4a3d-b910-3b4fb079327b",
        "host_id": "wavynoteadmin"
    }
    ```
- 응답
  - 본문
    - `result`: 내가 쓴 노트 (오픈 노트에)공유하기에 대한 성공/실패 여부(true/false)
    - `msg`: 실패 시 반환하는 에러 메시지
  - 응답 본문 예시
    ```json
    {
        "result": "true",
        "msg": ""
    }
    ```

### 내가 쓴 노트를 랜덤 매칭을 통해 (비슷한 관심 주제를 갖는)임의의 대상의 노트로 보내기
*랜덤 매칭의 대상은 오픈 노트에 등록된 노트로 한정하는것이 맞는가?*<br>
*응답값에 매칭된 대화방에 대한 정보가 필요한가?*

- 메서드: POST
- URL: https://{ip}:{port}/wavynote/v1.0/write/random
- 요청
  - 헤더
    - Authorization: Basic {ID:PW Base64 인코딩 데이터}
    - Content-Type: application/json
  - 본문
    - `note_id`: 랜덤 매칭을 통해 임의의 대상에게 보낼 노트의 고유 id 값
    - `from`: 노트를 보내는 사용자의 id
  - 요청 본문 예시
    ```json
    {
        "note_id": "09d05df1-2958-4a3d-b910-3b4fb079327b",
        "from": "wavynoteadmin"
    }
    ```
- 응답
  - 본문
    - `result`: 랜덤 매칭을 통한 임의의 대상에게 보내기에 대한 성공/실패 여부(true/false)
    - `msg`: 실패 시 반환하는 에러 메시지
  - 응답 본문 예시
    ```json
    {
        "result": "true",
        "msg": ""
    }
    ```

### 내가 쓴 노트 조회
*응답값에 `conversation_id`와 `to` 정보가 필요한가?*

- 메서드: GET
- URL: https://{ip}:{port}/wavynote/v1.0/write/show?id={note_id}
- 요청
  - 헤더
    - Authorization: Basic {ID:PW Base64 인코딩 데이터}
    - Content-Type: application/x-www-form-urlencoded
  - 본문
    - ~~`type`: 글의 종류(my/send/recv/open)~~
    - `id`: 조회 대상이 되는 노트의 고유 id 값
- 응답
  - 본문
    - `note_id`: 내가 조회한 노트의 고유 id 값
    - `folder_id`: 내가 조회한 노트가 포함되어 있는 폴더의 고유 id 값
    - `from`: 내가 조회한 노트 작성자의 id
    - `save_at`: 내가 조회한 노트를 저장한 마지막 날짜 및 시간 정보
    - `title`: 내가 조회한 노트의 제목 
    - `content`: 내가 조회한 노트의 본문 내용
    - `keyword`: 내가 조회한 노트의 키워드
  - 응답 본문 예시
    ```json
    {
        "note_id": "09d05df1-2958-4a3d-b910-3b4fb079327b",
        "folder_id": "980e71ba-0395-49aa-833e-3ebc76b3ec88",
        "from": "wavynoteadmin",
        "save_at": "2023-11-01 21:00:00",
        "title": "나의첫번째노트",
        "content": "나의 첫 웨이비노트 본문 내용입니다.",
        "keyword": "일상"
    }
    ```

## Box 페이지
### 노트를 주고 받는 대화방 목록 조회(최대 3개)
- 메서드: GET
- URL: https://{ip}:{port}/wavynote/v1.0/box/conversations
- 요청
  - 헤더
    - Authorization: Basic {ID:PW Base64 인코딩 데이터}
    - Content-Type: application/x-www-form-urlencoded
- 응답
  - 본문
    - `data`: 노트를 주고 받는 대화방 정보를 배열 형태로 전달
      - `conversation_id`: 대화방 고유의 id 값
      - `opp_nickname`: 대화 상대의 별명
      - `note_count`: 대화방에 존재하는 노트의 총 개수
  - 응답 본문 예시
    ```json
    {
        "data": [
            {
                "conversation_id": "1afc571d-61bf-4cef-95ce-ab791f999297",
                "opp_nickname": "누군가",
                "note_count": "20"
            },
            {
                "conversation_id": "e1ce587c-c6c0-46fc-b59d-fd8316a4502a",
                "opp_nickname": "또다른누군가",
                "note_count": "2"
            }
        ]
    }
    ```

### 특정 친구와 주고받은 노트 목록
*받은 노트와 보낸 노트 구분은 from, to 필드를 통해서 + timestamp를 통해 노트의 순서 배치*<br>
*노트의 개수가 많은 경우 페이징 처리가 필요함*

- 메서드: GET
- URL: https://{ip}:{port}/wavynote/v1.0/box/notelist/{conversation_id}
- 요청
  - 헤더
    - Authorization: Basic {ID:PW Base64 인코딩 데이터}
    - Content-Type: application/x-www-form-urlencoded
  - 본문
    - `conversation_id`: 특정 친구와 주고받은 노트 목록이 저장되는 대화방의 고유 id 값
- 응답
  - 본문
    - `data`: 특정 친구와 주고받은 노트 정보가 배열 형태로 전달(send_at을 기준으로 정렬됨)
      - `note_id`: 노트의 고유 id 값
      - `from`: 노트의 송신자 id
      - `to`: 노트의 수신자 id
      - `title`: 노트의 제목
      - `preview`: 노트의 본문 내용 미리보기
      - `send_at`: 노트를 보낸 날짜 및 시간 정보
    - 응답 본문 예시
    ```json
    {
        "data": [
            {
                "note_id": "1e5b9d7b-dfaa-4ad0-bcdb-d82cd3819650",
                "from": "somebody",
                "to": "wavynoteadmin",
                "title": "일상공유해요",
                "preview": "일상공유에 대한 본문내용입니다.",
                "send_at": "2023-10-29 21:00:48"
            },
            {
                "note_id": "09d05df1-2958-4a3d-b910-3b4fb079327b",
                "from": "wavynoteadmin",
                "to": "somebody",
                "title": "나의첫번째노트",
                "preview": "나의 첫 웨이비노트 본문 내용입니다.",
                "send_at": "2023-11-01 23:20:12",
            }
        ]
    }
    ```

### 특정 노트 조회
- 메서드: GET
- URL: https://{ip}:{port}/wavynote/v1.0/box/show?id={note_id}
- 요청
  - 헤더
    - Authorization: Basic {ID:PW Base64 인코딩 데이터}
    - Content-Type: application/x-www-form-urlencoded
  - 본문
    - `id`: 조회할 노트의 고유 id 값
- 응답
  - 본문
    - `note_id`: 내가 조회한 노트의 고유 id 값
    - `folder_id`: 내가 조회한 노트가 포함되어 있는 폴더의 고유 id 값
    - `from`: 내가 조회한 노트의 송신자 id
    - `to`: 내가 조회한 노트의 수신자 id
    - `save_at`: 내가 조회한 노트를 저장한 마지막 날짜 및 시간 정보 
    - `send_at`: 내가 조회한 노트를 송신한 날짜 및 시간 정보
    - `title`: 내가 조회한 노트의 제목 
    - `content`: 내가 조회한 노트의 본문 내용
    - `keyword`: 내가 조회한 노트의 키워드
  - 응답 본문 예시
    ```json
    {
        "note_id": "09d05df1-2958-4a3d-b910-3b4fb079327b",
        "folder_id": "980e71ba-0395-49aa-833e-3ebc76b3ec88",
        "from": "wavynoteadmin",
        "to": "somebody",
        "save_at": "2023-11-01 21:00:00",
        "send_at": "2023-11-01 23:20:12",
        "title": "나의첫번째노트",
        "content": "나의 첫 웨이비노트 본문 내용입니다.",
        "keyword": "일상"
    }
    ```

### 특정 대화방 삭제
- 메서드: DELETE
- URL: https://{ip}:{port}/wavynote/v1.0/box/{conversation_id}
- 요청
  - 헤더
    - Authorization: Basic {ID:PW Base64 인코딩 데이터}
    - Content-Type: application/x-www-form-urlencoded
  - 본문
    - `conversation_id`: 삭제할 대화방의 고유 id 값
- 응답
  - 본문
    - `result`: 특정 대화방 삭제에 대한 성공/실패 여부(true/false)
    - `msg`: 실패 시 반환하는 에러 메시지
  - 응답 본문 예시
    ```json
    {
        "result": "true",
        "msg": ""
    }
    ```

## Profile 페이지