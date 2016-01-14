package redmine

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

//type projectRequest struct {
	//Project Project `json:"project"`
//}

//type projectResult struct {
	//Project Project `json:"project"`
//}

type versionsResult struct {
	Versions []Version `json:"versions"`
}

type Version struct {
	Id          int    `json:"id"`
	Project     IdName `json:"project"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status		string `json:"status"`
	Sharing		string `json:"sharing"`
	DueDate		string `json:"due_date"`
	CreatedOn   string `json:"created_on"`
	UpdatedOn   string `json:"updated_on"`
}

func (c *client) Versions(projectId int) ([]Version, error) {
	res, err := c.Get(c.endpoint + "/projects/" + strconv.Itoa(projectId) + "/versions.json?key=" + c.apikey)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var r versionsResult
	if res.StatusCode == 404 {
		return nil, errors.New("Not Found")
	}
	if res.StatusCode != 200 {
		var er errorsResult
		err = decoder.Decode(&er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	} else {
		err = decoder.Decode(&r)
	}
	if err != nil {
		return nil, err
	}

	return r.Versions, nil
}