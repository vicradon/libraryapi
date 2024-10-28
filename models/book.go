package models

// Book represents the book model
// @Description Book model
type Book struct {
	Id          int     `json:"id" example:"1"`                                                                                         // @Description The unique ID of a book
	Title       string  `json:"title" binding:"required" example:"The Book Thief"`                                                      // @Description The title of the book
	Author      string  `json:"author" binding:"required" example:"Markus Zusak"`                                                       // @Description The author of this book
	Genre       *string `json:"genre,omitempty" example:"Young Adult Literature"`                                                       // @Description The genre of the book
	IsAvailable *bool   `json:"is_available,omitempty"`                                                                                 // @Description Whether the book is available for borrowing or already borrowed
	Edition     *string `json:"edition,omitempty" example:"1st"`                                                                        // @Description the edition of the book
	Summary     *string `json:"summary,omitempty" example:"Liesel's life during WW2 with her stealing books that defy Nazi principles"` // @Description the summary of this book
	CreatedAt   *string `json:"created_at,omitempty" example:"2024-10-27T17:41:52Z"`                                                    // @Description When the book was created
	UpdatedAt   *string `json:"updated_at,omitempty" example:"2024-10-27T17:41:52Z"`                                                    // @Description When the book was updated                                                                               // @Description when the book was updated
}
