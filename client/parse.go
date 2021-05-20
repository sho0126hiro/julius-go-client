package client

import (
	"errors"
	"strings"
)

// isRecogout Determine if it is an analysis target
func isRecogout(msg string) bool {
	return strings.Contains(msg, "<RECOGOUT>")
}

// getParseTarget extraction of parse target
func getParseTarget(msg string) (string, error) {
	tmp := strings.Split(msg, "\n")
	result := ""
	flag := false
	for _, v := range tmp {
		if v == "<RECOGOUT>" {
			flag = true
		} else if v == "</RECOGOUT>" {
			return result, nil
		}
		if flag {
			result += v + "\n"
		}
	}
	return "", errors.New("</RECOGOUT> tag is not found")
}

// parseMessage parse logic
func parseMessage(msg string) (*Result, error) {
	msg, err := getParseTarget(msg)
	if err != nil {
		return nil, err
	}
	msg = deleteSubstring(msg, "<RECOGOUT>")
	msg = deleteSubstring(msg, "<SHYPO ")
	msg = deleteSubstring(msg, "<WHYPO ")
	msg = deleteSubstring(msg, "</SHYPO>")
	msg = strings.TrimRight(msg, ".")
	lines := strings.Split(msg, "\n")

	var parseTarget []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		line = deleteSubstring(line, "/")
		line = deleteSubstring(line, ">")
		parseTarget = append(parseTarget, line)
	}

	var result Result
	var details []Detail
	for i := 0; i < len(parseTarget); i++ {
		if i == 0 {
			// SHYPO Tag
			shypo := strings.Split(parseTarget[i], " ")
			for _, str := range shypo {
				tmp := strings.Split(str, "=")
				if len(tmp) != 2 {
					continue
				}
				switch tmp[0] {
				case "RANK":
					result.Rank = strings.Trim(tmp[1], "\"")
				case "SCORE":
					result.Score = strings.Trim(tmp[1], "\"")
				case "GRAM":
					result.Gram = strings.Trim(tmp[1], "\"")
				}
			}
			continue
		}
		// WHYPO Tag
		var detail Detail
		whypo := strings.Split(parseTarget[i], "\" ")
		for _, str := range whypo {
			tmp := strings.Split(str, "=")
			if len(tmp) != 2 {
				continue
			}
			switch tmp[0] {
			case "WORD":
				detail.Word = strings.Trim(tmp[1], "\"")
			case "CLASSID":
				detail.ClassID = strings.Trim(tmp[1], "\"")
			case "PHONE":
				detail.Phone = strings.Trim(tmp[1], "\"")
			case "CM":
				detail.CM = strings.Trim(tmp[1], "\"")
			}
		}
		details = append(details, detail)
	}
	result.Details = details
	return &result, nil
}

// deleteSubstring delete substring from string
func deleteSubstring(s, deleteStr string) string {
	s = strings.ReplaceAll(s, deleteStr, "")
	return s
}
