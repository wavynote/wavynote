package restapi

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	METHOD_POST   = 0
	METHOD_GET    = 1
	METHOD_PUT    = 2
	METHOD_DELETE = 3

	RESTAPI_NAME     = "/wavynote"
	RESTAPI_VERSION  = "v1.0"
	RESTAPI_BASEPATH = RESTAPI_NAME + "/" + RESTAPI_VERSION

	RESTAPI_SERVICENAME_MAIN     = "/main"
	RESTAPI_SERVICENAME_SEARCH   = "/search"
	RESTAPI_SERVICENAME_WRITE    = "/write"
	RESTAPI_SERVICENAME_BOX      = "/box"
	RESTAPI_SERVICENAME_PROFILE  = "/profile"
	RESTAPI_SERVICENAME_OPENNOTE = "/opennote"

	BASIC_AUTH_USER     = "wavynote"
	BASIC_AUTH_PASSWORD = "wavy20230914"

	LOCATION_FOR_MAIN_FOLDERLIST = "/folderlist"
	LOCATION_FOR_MAIN_NOTELIST   = "/notelist"
	LOCATION_FOR_MAIN_NOTEFOLDER = "/notefolder"
	LOCATION_FOR_MAIN_FOLDER     = "/folder"
	LOCATION_FOR_MAIN_NOTE       = "/note"

	LOCATION_FOR_SEARCH_FROM_TOP      = "/top"
	LOCATION_FOR_SEARCH_FROM_FOLDER   = "/folder"
	LOCATION_FOR_SEARCH_FROM_OPENNOTE = "/opennote"

	LOCATION_FOR_WRITE_SAVE     = "/save"
	LOCATION_FOR_WRITE_SEND     = "/send"
	LOCATION_FOR_WRITE_OPENNOTE = "/opennote"
	LOCATION_FOR_WRITE_RANDOM   = "/random"
	LOCATION_FOR_WRITE_SHOW     = "/show"

	LOCATION_FOR_BOX_CONVERSATION_LIST = "/conversationlist"
	LOCATION_FOR_BOX_CONVERSATION      = "/conversation"
	LOCATION_FOR_BOX_NOTELIST          = "/notelist"
	LOCATION_FOR_BOX_SHOW              = "/show"

	LOCATION_FOR_PROFILE_SIGNIN            = "/signin"
	LOCATION_FOR_PROFILE_CHECKDULPLICATEID = "/duplicate"
	LOCATION_FOR_PROFILE_SIGNUP            = "/signup"

	LOCATION_FOR_OPENNOTE_LIST = "/list"
	LOCATION_FOR_OPENNOTE_SHOW = "/show"
)

type Response400 struct {
	Code int    `json:"code" example:"400"`
	Msg  string `json:"msg" example:"Bad request"`
}

type Response401 struct {
	Code int    `json:"code" example:"401"`
	Msg  string `json:"msg" example:"Fail to parse Authorization header"`
}

type Response403 struct {
	Code int    `json:"code" example:"403"`
	Msg  string `json:"msg" example:"you can't access this target resource when using swagger API test"`
}

type Response404 struct {
	Code int    `json:"code" example:"404"`
	Msg  string `json:"msg" example:"Not exist the target resource in our server"`
}

type Response500 struct {
	Code int    `json:"code" example:"500"`
	Msg  string `json:"msg" example:"Internal server error"`
}

type DefaultResponse struct {
	Result string `json:"result" example:"true"` // 요청에 대한 처리 성공/실패 여부
	Msg    string `json:"msg"`                   // 실패 시 반환하는 에러 메시지
}

type FolderInfo struct {
	FolderId   string `json:"folder_id" example:"a3106a0c-5ce7-40f6-81f4-ff9b8ebb240b"` // 폴더의 고유 id 값
	UserId     string `json:"user_id" example:"wavynoteadmin@gmail.com"`                // 폴더를 소유하고 있는 사용자 ID
	FolderName string `json:"folder_name" example:"whatever"`                           // 폴더 이름
}

type FolderSimpleInfo struct {
	FolderId   string `json:"folder_id" example:"a3106a0c-5ce7-40f6-81f4-ff9b8ebb240b"` // 폴더의 고유 id 값
	FolderName string `json:"folder_name" example:"my diary"`                           // 폴더 이름
	NoteCount  int    `json:"note_count" example:"5"`                                   // 폴더에 존재하는 노트의 총 개수
}

type FolderListResponse struct {
	Folders []FolderSimpleInfo `json:"data"`
}

type RemoveFolderInfo struct {
	FolderId   string `json:"folder_id" example:"a3106a0c-5ce7-40f6-81f4-ff9b8ebb240b"` // 삭제할 폴더의 고유 id 값
	FolderName string `json:"folder_name" example:"my diary"`                           // 삭제할 폴더의 이름
}

type RemoveFolderRequest struct {
	UserId        string             `json:"user_id" example:"wavynoteadmin@gmail.com"`
	RemoveFolders []RemoveFolderInfo `json:"data"`
}

type AddFolderRequest struct {
	UserId     string `json:"user_id" example:"wavynoteadmin@gmail.com"`
	FolderName string `json:"folder_name" example:"my wavywavy"`
}

type ChangeFolderNameRequest struct {
	FolderId   string `json:"folder_id" example:"a3106a0c-5ce7-40f6-81f4-ff9b8ebb240b"` // 이름을 변경할 폴더의 고유 id 값
	UserId     string `json:"user_id" example:"wavynoteadmin@gmail.com"`                // 이름을 변경할 폴더를 소유하고 있는 사용자 ID
	FolderName string `json:"folder_name" example:"whatever"`                           // 변경할 폴더 이름
}

type NoteInfo struct {
	NoteId         string   `json:"note_id" example:"09d05df1-2958-4a3d-b910-3b4fb079327b"`         // 노트의 고유 id 값
	FolderId       string   `json:"folder_id" example:"980e71ba-0395-49aa-833e-3ebc76b3ec88"`       // 노트가 포함되어 있는 폴더의 고유 id 값
	ConversationId string   `json:"conversation_id" example:"1afc571d-61bf-4cef-95ce-ab791f999297"` // 대화방의 고유 id 값
	FromId         string   `json:"from_id" example:"wavynoteadmin@gmail.com"`                      // 노트 작성자(또는 송신자)의 id
	ToId           string   `json:"to_id" example:"somebody@naver.com"`                             // 노트 수신자의 id
	SaveAt         string   `json:"save_at" example:"2023-11-01 21:00:00"`                          // 노트를 저장한 마지막 날짜 및 시간 정보
	SendAt         string   `json:"send_at" example:"2023-11-01 23:20:12"`                          // 노트를 송신한 날짜 및 시간 정보
	Title          string   `json:"title" example:"my first note"`                                  // 노트의 제목
	Content        string   `json:"content" example:"This is the main text of my first wavy note."` // 노트의 본문 내용
	Keywords       []string `json:"keywords" example:"b0d88d67-01fd-47f8-b426-6ca0657d0f6e"`        // 노트의 키워드
}

type NoteSimpleInfo struct {
	NoteId  string `json:"note_id" example:"09d05df1-2958-4a3d-b910-3b4fb079327b"`         // 해당 폴더에 존재하는 노트의 고유 id 값
	Title   string `json:"title" example:"my first note"`                                  // 해당 폴더에 존재하는 노트의 제목
	Preview string `json:"preview" example:"This is the main text of my first wavy note."` // 해당 폴더에 존재하는 노트의 본문 미리보기(글자수 제한)
}

type NoteListResponse struct {
	Notes []NoteSimpleInfo `json:"data"`
}

type RemoveNoteInfo struct {
	FolderId string `json:"folder_id" example:"980e71ba-0395-49aa-833e-3ebc76b3ec88"` // 삭제할 노트가 포함되어 있는 폴더의 고유 id 값
	NoteId   string `json:"note_id" example:"09d05df1-2958-4a3d-b910-3b4fb079327b"`   // 삭제할 노트의 고유 id 값
}

type RemoveNoteRequest struct {
	UserId      string           `json:"user_id" example:"wavynoteadmin@gmail.com"`
	RemoveNotes []RemoveNoteInfo `json:"data"`
}

type ChangeNoteFolderRequest struct {
	UserId   string   `json:"user_id" example:"wavynoteadmin@gmail.com"`                // 사용자 id
	FolderId string   `json:"folder_id" example:"980e71ba-0395-49aa-833e-3ebc76b3ec88"` // 노트가 저장되는 폴더의 고유 id 값
	Notes    []string `json:"notes" example:"09d05df1-2958-4a3d-b910-3b4fb079327b"`     // 폴더를 변경할 노트의 고유 id 값
}

type SaveNoteRequest struct {
	FolderId string   `json:"folder_id" example:"980e71ba-0395-49aa-833e-3ebc76b3ec88"`       // 내가 쓴 노트가 포함되어 있는 폴더의 고유 id 값
	FromId   string   `json:"from_id" example:"wavynoteadmin@gmail.com"`                      // 작성자의 id
	Title    string   `json:"title" example:"my first note"`                                  // 내가 쓴 노트의 제목
	Content  string   `json:"content" example:"This is the main text of my first wavy note."` // 내가 쓴 노트의 본문 내용
	Keywords []string `json:"keywords" example:"b0d88d67-01fd-47f8-b426-6ca0657d0f6e"`        // 내가 쓴 노트의 키워드
}

type UpdateNoteRequest struct {
	NoteId   string   `json:"note_id" example:"09d05df1-2958-4a3d-b910-3b4fb079327b"`         // 내가 쓴 노트의 고유 id 값
	FolderId string   `json:"folder_id" example:"980e71ba-0395-49aa-833e-3ebc76b3ec88"`       // 내가 쓴 노트가 포함되어 있는 폴더의 고유 id 값
	FromId   string   `json:"from_id" example:"wavynoteadmin@gmail.com"`                      // 작성자의 id
	Title    string   `json:"title" example:"my first note"`                                  // 내가 쓴 노트의 제목
	Content  string   `json:"content" example:"This is the main text of my first wavy note."` // 내가 쓴 노트의 본문 내용
	Keywords []string `json:"keywords" example:"b0d88d67-01fd-47f8-b426-6ca0657d0f6e"`        // 내가 쓴 노트의 키워드
}

type SendNoteRequest struct {
	NoteId         string `json:"note_id" example:"09d05df1-2958-4a3d-b910-3b4fb079327b"`         // 내가 쓴 노트의 고유 id 값
	FromId         string `json:"from_id" example:"wavynoteadmin@gmail.com"`                      // 작성자의 id
	ToId           string `json:"to_id" example:"somebody@naver.com"`                             // 내가 쓴 노트를 보내는 대상의 id
	ConversationId string `json:"conversation_id" example:"1afc571d-61bf-4cef-95ce-ab791f999297"` // 대화방의 고유 id 값
}

type ShareNoteRequest struct {
	NoteId string `json:"note_id" example:"09d05df1-2958-4a3d-b910-3b4fb079327b"` // 오픈 노트에 공유할 노트의 고유 id 값
	HostId string `json:"host_id" example:"wavynoteadmin@gmail.com"`              // 오픈 노트에 공유한 사용자 id
}

type RandomMatchRequest struct {
	NoteId string `json:"note_id" example:"09d05df1-2958-4a3d-b910-3b4fb079327b"` // 랜덤 매칭을 통해 임의의 대상에게 보낼 노트의 고유 id 값
	FromId string `json:"from_id" example:"wavynoteadmin@gmail.com"`              // 노트를 보내는 사용자의 id
}

type ConversationInfo struct {
	ConverstaionId string `json:"conversation_id" example:"1afc571d-61bf-4cef-95ce-ab791f999297"` // 대화방 고유의 id 값
	OppNickName    string `json:"opp_nickname" example:"somebody"`                                // 대화 상대의 별명
	NoteCount      int    `json:"note_count" example:"20"`                                        // 대화방에 존재하는 노트의 총 개수
	NewNote        bool   `json:"new_note" example:"true"`                                        // 대화방에 읽지않은 노트가 존재하는지 여부
}

type ConversationListResponse struct {
	Conversations []ConversationInfo `json:"data"`
}

type ConversationNoteInfo struct {
	NoteId  string `json:"note_id" example:"09d05df1-2958-4a3d-b910-3b4fb079327b"`
	FromId  string `json:"from_id" example:"wavynoteadmin@gmail.com"`
	ToId    string `json:"to_id" example:"somebody@naver.com"`
	Title   string `json:"title" example:"my first note"`
	Preview string `json:"preview" example:"This is the main text of my first wavy note."`
	SendAt  string `json:"send_at" example:"2023-10-29 21:00:48"`
}

type ConverstaionNoteListResponse struct {
	ConversationNotes []ConversationNoteInfo `json:"data"`
}

type LandingPageResonse struct {
	FolderId string `json:"folder_id" example:"283744fe-7074-4ffa-abc6-dae8cb786058"` // 랜딩 페이지에서 보여줄 폴더의 고유 id 값
	UserId   string `json:"user_id" example:"somebody@naver.com"`                     // 사용자 id
}

type SignInRequest struct {
	Id       string `json:"id" example:"somebody@naver.com"`                                                     // 로그인시 입력한 ID
	Password string `json:"password" example:"04F0BF9AB5963706415A0EC4C0E8167EBC8640B0B9A40870CDC39BE8A8DA813B"` // 로그인시 입력한 패스워드
}

type SignUpRequest struct {
	Id       string   `json:"id" example:"somebody@naver.com"`                                                     // 회원 가입시 입력한 ID
	Password string   `json:"password" example:"04F0BF9AB5963706415A0EC4C0E8167EBC8640B0B9A40870CDC39BE8A8DA813B"` // 회원 가입시 입력한 패스워드
	NickName string   `json:"nickname" example:"somebody"`                                                         // 회원 가입시 입력한 닉네임
	Keywords []string `json:"keywords" example:"b0d88d67-01fd-47f8-b426-6ca0657d0f6e"`                             // 회원 가입시 입력한 키워드
	Emoji    string   `json:"emoji" example:"😎"`                                                                   // 회원 가입시 입력한 이모지
}

type OpenNoteSimpleInfo struct {
	NoteId         string `json:"note_id" example:"09d05df1-2958-4a3d-b910-3b4fb079327b"`
	ConversationId string `json:"conversation_id" example:"1afc571d-61bf-4cef-95ce-ab791f999297"`
	Title          string `json:"title" example:"my first note"`
	Preview        string `json:"preview" example:"This is the main text of my first wavy note."`
	PostAt         string `json:"post_at" example:"2023-12-19 15:55:48"`
}

type OpenNoteListResponse struct {
	Notes []OpenNoteSimpleInfo `json:"data"`
}

func BasicAuth(c *gin.Context) {
	user, pwd, ok := c.Request.BasicAuth()
	if !ok {
		c.Abort()
		c.Header("WWW-Authenticate", `Basic realm="webkeeper"`)
		c.IndentedJSON(http.StatusUnauthorized, Response401{
			Code: http.StatusUnauthorized,
			Msg:  fmt.Sprintf("%d %s fail to parse Authorization header", http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)),
		})
		return
	}

	if user != BASIC_AUTH_USER {
		c.Abort()
		c.Header("WWW-Authenticate", `Basic realm="webkeeper"`)
		c.IndentedJSON(http.StatusUnauthorized, Response401{
			Code: http.StatusUnauthorized,
			Msg:  fmt.Sprintf("%d %s user incorrect", http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)),
		})
		return
	}

	if pwd != BASIC_AUTH_PASSWORD {
		c.Abort()
		c.Header("WWW-Authenticate", `Basic realm="webkeeper"`)
		c.IndentedJSON(http.StatusUnauthorized, Response401{
			Code: http.StatusUnauthorized,
			Msg:  fmt.Sprintf("%d %s password incorrect", http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)),
		})
		return
	}
}
