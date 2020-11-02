package mQueue

import (
	"cloud.google.com/go/firestore"
	"context"
	"time"
)

const Collection = `queue_entries`

type QueueEntry struct {
	Id string 
	Name string // customer name
	Priority int // lower = higher priority
	CreatedAt int64
}

func (q *QueueEntry) Insert(c *firestore.Client) error {
	entries := c.Collection(Collection)
	qe := entries.NewDoc()
	q.CreatedAt = time.Now().UnixNano()
	q.Id = qe.ID
	_, err := qe.Set(context.Background(), q)
	return err
}

func (q *QueueEntry) Update(c *firestore.Client) error {
	entries := c.Collection(Collection)
	qe := entries.Doc(q.Id)
	_, err := qe.Get(context.Background())
	if err != nil {
		return err // does not exists or other error
	}
	_, err = qe.Set(context.Background(), q)
	return err
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
