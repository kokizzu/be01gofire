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

func (q *QueueEntry) Delete(c *firestore.Client) error {
	entries := c.Collection(Collection)
	qe := entries.Doc(q.Id)
	_, err := qe.Get(context.Background())
	if err != nil {
		return err // does not exists or other error
	}
	_, err = qe.Delete(context.Background())
	return err
}

func (q *QueueEntry) List(c *firestore.Client, limit int) ([]QueueEntry, error) {
	if limit <= 0 {
		limit = 10
	}
	entries := c.Collection(Collection)
	que := entries.
		Where(`Priority`,`>`,0).
		OrderBy(`Priority`,firestore.Asc).
		Limit(limit)
	rows := que.Documents(context.Background())
	defer rows.Stop()
	all, err := rows.GetAll()
	if err != nil {
		return nil, err
	}
	res := []QueueEntry{}
	for _, v := range all {
		row := QueueEntry{}
		err = v.DataTo(&row)
		if err != nil {
			return res, err
		}
		res = append(res, row)
	}
	return res, nil
}
