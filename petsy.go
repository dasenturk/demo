package petsy

import "github.com/google/uuid"

type Thread struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
}

type Post struct {
	ID            uuid.UUID `db:"id"`
	ThreadID      uuid.UUID `db:"thread_id"`
	Title         string    `db:"title"`
	Content       string    `db:"content"`
	Votes         int       `db:"votes"`
	CommentsCount int       `db:"comments_count"`
	ThreadTitle   string    `db:"thread_title"`
}

type Comment struct {
	ID      uuid.UUID `db:"id"`
	PostID  uuid.UUID `db:"post_id"`
	Content string    `db:"content"`
	Votes   int       `db:"votes"`
}

type PetOwner struct{
	ID        uuid.UUID `db:"id"`
	Name 	  string    `db:"ownername"`
	Surname   string    `db:"surname"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	Email	  string    `db:"email"`
	CityName  string 	 `db:"cityname"`
}

type Pet struct {
	ID        uuid.UUID `db:"id"`
	Name 	  string    `db:"petname"`
	PetType   string    `db:"pettype"`
	Breed     string    `db:"breed_type"`
	OwnerID   uuid.UUID	`db:"owner_id"`
	Gender	  string	`db:"gender"`
	Weight	  int		`db:"weight"`
	Age		  int		`db:"age"` 	  
}

type Breed struct {
	Name 		   string 	`db:"breed_name"`
	BType 		   string   `db:"btype"`
	AvgWeight      int 	    `db:"avgweight"`
	ActivityLevel  string 	`db:"activity"`
	CoatType 	   string 	`db:"coat"`
	Shedding       string   `db:"shedding"`
}

type PetOwnerStore interface{
	PetOwner(id uuid.UUID) (PetOwner, error)
	PetOwnerByUsername(username string) (PetOwner, error)
	PetOwners()([]PetOwner, error)
	PetOwnerByName(name, surname string) (PetOwner, error)
	PetOwnerByEmail(email string)(PetOwner, error)
	PetOwnerByCity(cityname string)([]PetOwner, error)
	CreatePetOwner(u *PetOwner) error
	UpdatePetOwner(u *PetOwner) error
	DeletePetOwner(id uuid.UUID) error
}

type PetStore interface{
	Pet(id uuid.UUID) (Pet, error)
	PetByName(name string)([]Pet, error)
	PetByOwner(ownerID uuid.UUID)([]Pet, error)
	PetByType(pettype string)([]Pet,error)
	PetByBreed(breed string)([]Pet, error)
	PetByGender(gender string)([]Pet,error)
	CreatePet(p *Pet) error
	UpdatePet(p *Pet) error
	DeletePet(id uuid.UUID) error
}

type BreedStore interface{
	Breed(name string)(Breed, error)
	BreedByType(btype string)([]Breed, error)
	CreateBreed(b *Breed) error
	UpdateBreed(b *Breed) error
	DeleteBreed(name string) error
}

type ThreadStore interface {
	Thread(id uuid.UUID) (Thread, error)
	Threads() ([]Thread, error)
	CreateThread(t *Thread) error
	UpdateThread(t *Thread) error
	DeleteThread(id uuid.UUID) error
}

type PostStore interface {
	Post(id uuid.UUID) (Post, error)
	Posts() ([]Post, error)
	PostsByThread(threadID uuid.UUID) ([]Post, error)
	CreatePost(t *Post) error
	UpdatePost(t *Post) error
	DeletePost(id uuid.UUID) error
}

type CommentStore interface {
	Comment(id uuid.UUID) (Comment, error)
	CommentsByPost(postID uuid.UUID) ([]Comment, error)
	CreateComment(t *Comment) error
	UpdateComment(t *Comment) error
	DeleteComment(id uuid.UUID) error
}

type Store interface{
	ThreadStore
	PostStore
	CommentStore
	PetOwnerStore
	PetStore
	BreedStore
}

