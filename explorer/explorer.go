package explorer

import (
	"_nadochain_com/blockchain"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type cls_HomeData struct {
	PageTitle string
	Blocks    []*blockchain.Cls_Block
}

const (
	port        string = ":4000"      //	기본 PORT , 앞에 : 안 찍으면 PORT 입력할때 오류남
	templateDir string = "templates/" //	텝플릿 기본 DIR, go파일은 src 밑으로 옮기는데 템플릿은 작업폴더에 있어야 하는듯
)

var templates *template.Template

func Fn_Start() {
	// template.ParseFiles("index.html")  - 템플릿 파싱 여려개의 인자를 받아 중첩 템플릿 형태를 만들 수 있다.
	// template.ParseGlob("*.html")  - 패턴 매칭을 이용한 템플릿 파싱
	// template.Parse(data) - string의 템플릿을 파싱
	// .Must : 에러가 nil이 아닐 때 발생할경우를 처리할 때 사용

	// 한번에 여러개 폴더를 로딩하지는 못함, 2개를 & 로 이어줌
	templates = template.Must(template.ParseGlob(templateDir + "_inc/*.gohtml"))   //	_inc에 있는 모든 파일
	templates = template.Must(templates.ParseGlob(templateDir + "_page/*.gohtml")) //	_page에 있는 모든 파일

	// 라우팅
	http.HandleFunc("/", fn_home)
	http.HandleFunc("/add", fn_Add)

	// port 4000 서비스 대기
	fmt.Printf("listening : %s\n", port)

	// 오류시 메시지 표시
	log.Fatal(http.ListenAndServe(port, nil))
}

// "/" 에 접속할때
func fn_home(rw http.ResponseWriter, r *http.Request) {
	data := cls_HomeData{"Home", blockchain.Fn_getBlockChain().Fn_allBlocks()}
	err := templates.ExecuteTemplate(rw, "home", data)
	// 오류 메시지 출력
	if err != nil {
		fmt.Println("error:", err.Error())
	}
}

// "/add" 에 접속할때
func fn_Add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData") //	form 에서 blockData 값을 읽어냄
		blockchain.Fn_getBlockChain().Fn_addBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}
