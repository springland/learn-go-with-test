package blogposts

import(

	"io/fs"
	"io"
	"bufio"
	"strings"
	"bytes"
	"fmt"
)

type Post struct {
	Title string 
	Description string 
	//Body string
	Tags []string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator = "Tags: "
)


func NewPostsFromFS( filesystem fs.FS) ([]Post  , error) {

	dir , err := fs.ReadDir(filesystem , ".")

	if err != nil {
		return nil , err
	}
	var posts []Post
	
	for _ , f := range dir {


		post , err := getPost(filesystem , f)
		if err != nil {
			return nil , err
		}
		posts  = append(posts , post)


		
	}
	return posts , nil 
}

func getPost(fileSystem fs.FS , f fs.DirEntry) (Post , error){

	postFile , err := fileSystem.Open(f.Name())
	if err != nil {
		return Post{} , err
	}
	defer postFile.Close()

	return newPost(postFile)
	// postData , err := io.ReadAll(postFile)
	// if err != nil {
	// 	return Post{} , err
	// }

	// post := Post{Title : string(postData)[7:]}
	// return post , nil
}

func newPost(postBody io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postBody)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		text := strings.TrimSpace(scanner.Text())
		//fmt.Printf(" text %q \n" , text)
		return strings.TrimPrefix(text, tagName)
	}

	title := readMetaLine(titleSeparator)
	description := readMetaLine(descriptionSeparator)
	tags := strings.Split(readMetaLine(tagsSeparator), ", ")

	fmt.Printf(" title %q , description %q , tags %v \n" , title , description , tags)

	scanner.Scan() // ignore a line

	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	//body := strings.TrimSuffix(buf.String(), "\n")

	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
		//Body:        body,
	}, nil
}