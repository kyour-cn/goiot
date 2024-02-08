package gimp

import (
	"bufio"
	"fmt"
	"strings"
)

type Gimp struct {
	cmd     string
	headers map[string]string
	body    string
}

func (g *Gimp) SetCmd(cmd string) {
	g.cmd = cmd
}

func (g *Gimp) GetCmd() string {
	return g.cmd
}

func (g *Gimp) SetHeader(key, value string) {
	g.headers[key] = value
}

func (g *Gimp) GetHeader(key string) string {
	return g.headers[key]
}

func (g *Gimp) SetBody(body string) {
	g.body = body
}

func (g *Gimp) GetBody() string {
	return g.body
}

func (g *Gimp) Encode() string {
	msg := g.cmd

	if g.body == "" && len(g.headers) == 0 {
		return msg
	}

	// 指令结束符
	msg += "\n"

	for k, v := range g.headers {
		msg += k + ":" + v + "\n"
	}

	if g.body == "" {
		return msg
	}

	// 消息头结束符（形成空行），后面是消息体
	return msg + "\n" + g.body
}

func (g *Gimp) Decode(data string) error {

	reader := strings.NewReader(data)

	scanner := bufio.NewScanner(reader)

	// 读取并处理第一行（指令）
	if scanner.Scan() {
		g.cmd = scanner.Text()
	}

	// 初始化map
	g.headers = make(map[string]string)

	// 处理消息头部分
	for scanner.Scan() {
		currentLine := scanner.Text()
		if currentLine == "" {
			break // 遇到空行，结束消息头部分的读取
		}
		headerParts := strings.SplitN(currentLine, ":", 2)
		if len(headerParts) == 2 {
			g.headers[headerParts[0]] = headerParts[1]
		} else {
			return fmt.Errorf("invalid header format: %s", currentLine)
		}
	}

	// 读取消息体部分
	bodyBuilder := &strings.Builder{}
	isFirstLine := true
	for scanner.Scan() {
		if isFirstLine {
			isFirstLine = false
		} else {
			bodyBuilder.WriteRune('\n') // 保持与原始数据相同的换行格式
		}
		bodyBuilder.WriteString(scanner.Text())
	}
	g.body = bodyBuilder.String()

	return nil
}
