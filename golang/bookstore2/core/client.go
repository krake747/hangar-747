package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"bookstore2/core/books"
)

type Client struct {
	addr string
}

func NewClient(addr string) *Client {
	return &Client{
		addr: addr,
	}
}

func (client *Client) MakeGetRequest(URI string, result any) error {
	url := client.addr + "/v1/" + URI
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close() //nolint:errcheck
	if resp.StatusCode == http.StatusNotFound {
		return errors.New("not found")
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status %d", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if result != nil {
		err = json.Unmarshal(data, result)
		if err != nil {
			return fmt.Errorf("%v in %q", err, data)
		}
	}
	return nil
}

func (client *Client) MakePostRequest(URI string, result any) error {
	url := client.addr + "/v1/" + URI
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close() //nolint:errcheck
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status %d", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if result != nil {
		err = json.Unmarshal(data, result)
		if err != nil {
			return fmt.Errorf("%v in %q", err, data)
		}
	}
	return nil
}

func (client *Client) GetAllBooks() ([]books.Book, error) {
	bookList := []books.Book{}
	err := client.MakeGetRequest("list", &bookList)
	if err != nil {
		return nil, err
	}
	return bookList, nil
}

func (client *Client) GetBook(ID string) (books.Book, error) {
	book := books.Book{}
	err := client.MakeGetRequest("find/"+ID, &book)
	if err != nil {
		return books.Book{}, err
	}
	return book, nil
}

func (client *Client) GetCopies(ID string) (int, error) {
	copies := 0
	err := client.MakeGetRequest("getcopies/"+ID, &copies)
	if err != nil {
		return 0, err
	}
	return copies, nil
}

func (client *Client) AddCopies(ID string, copies int) (int, error) {
	url := client.addr + "/v1/addcopies/" + ID + "/" + fmt.Sprintf("%d", copies)
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close() //nolint:errcheck
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status %d", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var stock int
	err = json.Unmarshal(data, &stock)
	if err != nil {
		return 0, fmt.Errorf("%v in %q", err, data)
	}
	return stock, nil
}

func (client *Client) SubCopies(ID string, copies int) (int, error) {
	url := client.addr + "/v1/subcopies/" + ID + "/" + fmt.Sprintf("%d", copies)
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close() //nolint:errcheck
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status %d", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var stock int
	err = json.Unmarshal(data, &stock)
	if err != nil {
		return 0, fmt.Errorf("%v in %q", err, data)
	}
	return stock, nil
}
