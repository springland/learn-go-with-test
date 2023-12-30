package test

import(

	"testing"
	"errors"
	"testing/fstest"
	"fundamentals/blogposts"  
	"reflect"
	"io/fs"
	
)

type StubFailingFS struct {

}

func (s *StubFailingFS) Open( name string)(fs.File , error){
	return nil , errors.New("oh no, i always fail")
}

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
			Description: Description 1
			Tags: tdd, go`
		secondBody = `Title: Post 2
			Description: Description 2
			Tags: rust, borrow-checker`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Errorf( "%v" , err)
	}
	// rest of test code cut for brevity
	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
	})

	
}

func assertPost( t *testing.T , got blogposts.Post , want blogposts.Post){

	if !reflect.DeepEqual(got , want){
		t.Errorf("got %+v, want %+v", got, want)
	}



}