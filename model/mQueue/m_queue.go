package mQueue

import "cloud.google.com/go/firestore"

const Collection = `queue_entries`

type QueueEntry struct {
	Id string 
	Name string // customer name
	Priority int // lower = higher priority
}

func (q *QueueEntry) Insert(c *firestore.Client) error {
	// TODO: insert to db
	return nil
}

func (q *QueueEntry) Update(c *firestore.Client) error {
	// TODO: update based on id
	return nil
}

func (q *QueueEntry) Delete(c *firestore.Client, limit int) ([]QueueEntry,error) {
	res, err := q.List(c, limit)
	if err != nil {
		return nil, err
	}
	// TODO: remove each returned items
	return res, err
}

func (q *QueueEntry) List(c *firestore.Client, limit int) ([]QueueEntry, error) {
	// TODO: get top
	return nil, nil
}
