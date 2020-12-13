package entities

import "time"

type Article struct {
    Id int64
    Title string
    Content string
    CreateAt time.Time
}
