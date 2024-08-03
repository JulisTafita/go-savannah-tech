package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/JulisTafita/go-savannahTech/internal/config"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func fetchUserRepository(owner, repositoryName string) (repository Repository, err error) {
	repositoryUrl := fmt.Sprintf("/repos/%s/%s", url.QueryEscape(owner), url.QueryEscape(repositoryName))
	response, err := fetch(repositoryUrl)
	if err != nil {
		return repository, err
	}

	err = json.Unmarshal(response, &repository)
	if err != nil {
		return repository, err
	}

	return repository, nil
}

func fetchUserRepositoryList() (repositories []Repository, err error) {
	response, err := fetch("/user/repos")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &repositories)
	if err != nil {
		return nil, err
	}

	return repositories, err
}

func fetchSearchResultRepositoryList(keyWord string) (searchResult RepositorySearchResult, err error) {
	searchUrl := fmt.Sprintf("/search/repositories?q=%s&order=desc&per_page=100", url.QueryEscape(keyWord))
	response, err := fetch(searchUrl)
	if err != nil {
		return searchResult, err
	}

	err = json.Unmarshal(response, &searchResult)
	if err != nil {
		return searchResult, err
	}

	return searchResult, err
}

func fetchRepositoryCommit(owner, repositoryName string, since, until time.Time, page int) (commits []RepositoryCommit, err error) {
	commitUrl := fmt.Sprintf("/repos/%s/%s/commits?per_page=100", owner, repositoryName)

	if !since.IsZero() {
		commitUrl += "&since=" + since.Format(time.RFC3339)
	}

	if !until.IsZero() {
		commitUrl += "&until=" + until.Format(time.RFC3339)
	}

	commitUrl += "&page=" + strconv.Itoa(page)

	response, err := fetch(commitUrl)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &commits)
	if err != nil {
		return nil, err
	}

	return commits, err
}

func fetchProfile() (profile Profile, err error) {
	response, err := fetch("/user")
	if err != nil {
		return profile, err
	}

	err = json.Unmarshal(response, &profile)
	if err != nil {
		return profile, err
	}

	return profile, err
}

func fetch(url string) (response []byte, err error) {
	var (
		req    *http.Request
		res    *http.Response
		client *http.Client
	)

	apiFullUrl := config.Default.Github.ApiEndpoint + url

	client = &http.Client{}
	req, err = http.NewRequest(http.MethodGet, apiFullUrl, nil)
	if err != nil {
		return nil, err
	}

	if config.UsePrivateRepository() {
		req.Header.Add("Authorization", config.Default.Github.UserToken)
	}

	res, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.Status != http.StatusText(http.StatusOK) {
		fmt.Println(res.Status)
		fmt.Println(string(body))
		return nil, errors.New(string(body))
	}

	return body, err
}
