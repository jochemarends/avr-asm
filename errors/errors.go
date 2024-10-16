package errors

type What string

type Error struct {
    What What
    Line uint
}

