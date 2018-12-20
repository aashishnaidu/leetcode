package leetcode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
	"unicode"
)

func QuestionData(titleSlug string) (qd QuestionDataType) {
	jsonStr := `{
  		"operationName": "questionData",
  		"variables": {
    		"titleSlug": "` + titleSlug + `"
  		},
  		"query": "query questionData($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    boundTopicId\n    title\n    titleSlug\n    content\n    translatedTitle\n    translatedContent\n    isPaidOnly\n    difficulty\n    likes\n    dislikes\n    isLiked\n    similarQuestions\n    contributors {\n      username\n      profileUrl\n      avatarUrl\n      __typename\n    }\n    langToValidPlayground\n    topicTags {\n      name\n      slug\n      translatedName\n      __typename\n    }\n    companyTagStats\n    codeSnippets {\n      lang\n      langSlug\n      code\n      __typename\n    }\n    stats\n    hints\n    solution {\n      id\n      canSeeDetail\n      __typename\n    }\n    status\n    sampleTestCase\n    metaData\n    judgerAvailable\n    judgeType\n    mysqlSchemas\n    enableRunCode\n    enableTestMode\n    envInfo\n    __typename\n  }\n}\n"
	}`
	resp, err := http.Post(GraphqlUrl, "application/json", strings.NewReader(jsonStr))
	checkErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	err = json.Unmarshal(body, &qd)
	checkErr(err)
	return
}

type QuestionDataType struct {
	Errors []errorType `json:"errors"`
	Data   dataType    `json:"data"`
}

type errorType struct {
	Message string `json:"message"`
}

type dataType struct {
	Question questionType `json:"question"`
}

type questionType struct {
	QuestionId         string             `json:"questionId"`
	QuestionFrontendId string             `json:"questionFrontendId"`
	BoundTopicId       string             `json:"boundTopicId"`
	Title              string             `json:"title"`
	TitleSlug          string             `json:"titleSlug"`
	Content            string             `json:"content"`
	TranslatedTitle    string             `json:"translatedTitle"`
	TranslatedContent  string             `json:"translatedContent"`
	IsPaidOnly         bool               `json:"isPaidOnly"`
	Difficulty         string             `json:"difficulty"`
	Likes              int                `json:"likes"`
	Dislikes           int                `json:"dislikes"`
	IsLiked            int                `json:"isLiked"`
	SimilarQuestions   string             `json:"similarQuestions"`
	CodeSnippets       []codeSnippetsType `json:"codeSnippets"`
}

type codeSnippetsType struct {
	Lang     string `json:"lang"`
	LangSlug string `json:"langSlug"`
	Code     string `json:"code"`
}

type similarQuestionType struct {
	Title           string `json:"title"`
	TitleSlug       string `json:"titleSlug"`
	Difficulty      string `json:"difficulty"`
	TranslatedTitle string `json:"translatedTitle"`
}

func (question questionType) SaveContent() {
	fmt.Println(question.QuestionFrontendId, "\t", question.Title, " saving...")
	if question.Content != "" {
		filePutContents(question.getFilePath("README.md"), question.getDescContent())
	}
}

func (question questionType) getDescContent() []byte {
	wb := bytes.Buffer{}
	if question.Difficulty != "" {
		question.Difficulty = fmt.Sprintf(" (%s)", question.Difficulty)
	}
	wb.WriteString(fmt.Sprintf("## %s. %s%s\n\n", question.QuestionFrontendId, question.Title, question.Difficulty))
	wb.WriteString(question.Content)
	wb.Write(question.getSimilarQuestion())
	return wb.Bytes()
}

func (question questionType) getSimilarQuestion() []byte {
	var sq []similarQuestionType
	if question.SimilarQuestions != "" {
		json.Unmarshal([]byte(question.SimilarQuestions), &sq)
	}
	var bf bytes.Buffer
	if len(sq) > 0 {
		bf.WriteString("\n\n### Similar Questions\n")
	}
	format := "  1. [%s](https://github.com/openset/leetcode/tree/master/solution/%s)%s\n"
	for _, q := range sq {
		if q.Difficulty != "" {
			q.Difficulty = fmt.Sprintf(" (%s)", q.Difficulty)
		}
		bf.WriteString(fmt.Sprintf(format, q.Title, q.TitleSlug, q.Difficulty))
	}
	return bf.Bytes()
}

func (question questionType) getFilePath(filename string) string {
	return path.Join("solution", question.TitleSlug, filename)
}

func (question questionType) TitleSnake() string {
	return strings.Replace(question.TitleSlug, "-", "_", -1)
}

func (question questionType) PackageName() string {
	snake := question.TitleSnake()
	if snake != "" && unicode.IsNumber(rune(snake[0])) {
		snake = "p_" + snake
	}
	return snake
}

func (question questionType) SaveCodeSnippet() {
	for _, code := range question.CodeSnippets {
		if code.LangSlug == "golang" {
			file := question.getFilePath(question.TitleSnake() + ".go")
			bf := bytes.Buffer{}
			bf.WriteString(fmt.Sprintf("package %s\n\n", question.PackageName()))
			bf.WriteString(code.Code)
			filePutContents(file, bf.Bytes())
		}
	}
}