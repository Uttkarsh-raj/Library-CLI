package lib

import (
	"encoding/json"
	"errors"

	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
)

type Book struct {
	NAME         string
	Genre        string
	Issued       bool
	IssueDate    time.Time
	ReturnedDate time.Time
}

type Library []Book

func (b *Library) Add(newBook, genre string) error {
	B := Book{
		NAME:         newBook,
		Genre:        genre,
		Issued:       false,
		IssueDate:    time.Now(),
		ReturnedDate: time.Time{},
	}
	ls := *b
	for idx := range ls {
		if ls[idx].NAME == B.NAME && ls[idx].Genre == B.Genre {
			return errors.New("Book already exists.\n")
		}
	}
	*b = append(*b, B)
	return nil
}

func (b *Library) Load(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, b)
	if err != nil {
		return err
	}
	return nil
}

func (b *Library) Store(fileLocation string) error {
	data, err := json.Marshal(b)
	if err != nil {
		return nil
	}
	return ioutil.WriteFile(fileLocation, data, 0644)
}

func (b *Library) Delete(name string) error {
	ls := *b
	for i := range ls {
		if ls[i].NAME == name {
			*b = append(ls[:i], ls[i+1:]...)
			return nil
		}
	}
	return errors.New("No such book exists.\n")
}

func (b *Library) List() {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Title"},
			{Align: simpletable.AlignCenter, Text: "Genre"},
			{Align: simpletable.AlignCenter, Text: "Issued"},
			{Align: simpletable.AlignCenter, Text: "Issued on"},
			{Align: simpletable.AlignCenter, Text: "Returned on"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, item := range *b {
		idx++
		task := blue(item.NAME)
		done := blue("No")
		issOn := blue(item.IssueDate.Format(time.RFC822))
		retOn := blue(item.ReturnedDate.Format(time.RFC822))
		if item.Issued {
			task = red(fmt.Sprintf("%s", item.NAME))
			done = red("Yes")
			issOn = red(fmt.Sprintf("%s", item.IssueDate.Format(time.RFC822)))
			retOn = red(fmt.Sprintf("%s", item.ReturnedDate.Format(time.RFC822)))
		}
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: item.Genre},
			{Text: done},
			{Text: issOn},
			{Text: retOn},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
	fmt.Println(red(fmt.Sprintf("%d books yet to be returned.\n", b.CountPending())))
}

func (b *Library) CountPending() int {
	total := 0
	for _, item := range *b {
		if item.Issued {
			total++
		}
	}
	return total
}

func (b *Library) Issue(name string) error {
	ls := *b
	for idx, item := range ls {
		if item.NAME == name && item.Issued == false {
			ls[idx].IssueDate = time.Now()
			ls[idx].ReturnedDate = time.Now()
			ls[idx].Issued = true
			return nil
		} else if item.NAME == name && item.Issued == true {
			return errors.New("The book is already issued.\n")
		}
	}
	return errors.New("No such book is present.\n")
}

func (b *Library) Return(name string) error {
	ls := *b
	for idx, item := range ls {
		if item.NAME == name && item.Issued == true {
			ls[idx].IssueDate = time.Now()
			ls[idx].ReturnedDate = time.Now()
			ls[idx].Issued = false
			return nil
		} else if item.NAME == name && item.Issued == false {
			return errors.New("The book was not issued.\n")
		}
	}
	return errors.New("No such book is present.\n")
}
