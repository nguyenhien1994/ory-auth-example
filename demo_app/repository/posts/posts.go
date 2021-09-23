package posts

import (
	"simple-go-auth/services/mysql"
	"strconv"

	"github.com/pkg/errors"
)

type Post struct {
	ID          int
	Title       string
	PostDetails string
	OwnerId     int
}

func parsePostFromDBData(rows [][]interface{}, columns []string) (Post, error) {
	var post Post
	for _, row := range rows {
		for columnIndex, columnValue := range row {
			b, ok := columnValue.([]byte)
			if !ok {
				return post, errors.New("Invalid data type retrieving from database")
			}
			switch columns[columnIndex] {
			case "id":
				id, err := strconv.Atoi(string(b))
				if err != nil {
					return post, errors.Wrap(err, "Can't parse id")
				}
				post.ID = id
			case "title":
				post.Title = string(b)
			case "post_details":
				post.PostDetails = string(b)
			case "owner_id":
				id, err := strconv.Atoi(string(b))
				if err != nil {
					return post, errors.Wrap(err, "Can't parse id")
				}
				post.OwnerId = id
			default:
				return post, errors.New("Invalid data type retrieving from database")
			}
		}
	}
	return post, nil
}

func FindPostByID(id int) (Post, error) {
	query := "SELECT * from posts where id = '" + strconv.Itoa(id) + "';"
	rows, columns, err := mysql.GetClientClient().ExecuteQuery(query)
	if err != nil {
		return Post{}, err
	}

	return parsePostFromDBData(rows, columns)
}

func UpdatePost(id int, detail string) error {
	return nil
}
