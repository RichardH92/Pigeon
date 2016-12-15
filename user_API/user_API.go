package user_API

import (
	"net/http"
	//"log"
	"Pigeon/utilities"
)

type Email_Type struct {
	Type_Name      string
	HTML_File_Name string
	Query_Vals     []string
	Subject        string
}

func HandleGetSendEmail(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		w.WriteHeader(401)
		return
	}

	dests, cont := ValidateDestQuery(req.Form.Get("dest"))
	if !cont {
		sendErrResponse(w, GetInvalidQueryParameterMessage("dest"))
		return
	}

	//TODO: Validate email addresses here

	emailType, cont := ValidateEmailType(req.Form.Get("type"))
	if !cont {
		sendErrResponse(w, GetInvalidQueryParameterMessage("type"))
		return
	}

	m := getFindReplaceQueryMap(emailType.Query_Vals, req)
	cont, errs := ValidateHTMLReplacementValues(m)
	if !cont {
		sendErrResponse(w, GetInvalidQueryParametersMessage(errs))
		return
	}

	cont, html := utilities.GetFormattedHTMLFromFile(m, emailType.HTML_File_Name)
	if !cont {
		sendErrResponse(w, GetFileOpeningErrorMessage(emailType.HTML_File_Name))
	}

	utilities.SendEmail(dests, html, emailType.Subject)

}

func sendErrResponse(w http.ResponseWriter, resp ErrorMessage) {
	w.WriteHeader(resp.Code)

	for e := resp.Messages.Front(); e != nil; e = e.Next() {
		w.Write([]byte(e.Value.(string)))
	}
}

func getFindReplaceQueryMap(queries []string, req *http.Request) map[string]string {
	m := make(map[string]string)

	for _, element := range queries {
		m[element] = req.Form.Get(element)
	}

	return m
}
