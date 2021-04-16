package main

import (
	"context"
	"time"
)

// Common fields to all tables
type Base struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// Before insert hooks into insert opeations, settings CreatedAt and UpdatedAt to current time
func (b *Base) BeforeInsert(ctx context.Context) (context.Context, error) {
	now := time.Now()
	b.CreatedAt = now
	b.UpdatedAt = now
	return ctx, nil
}

// BeforeUpdate hooks into update operations, setting UpdatedAt to current time
func (b *Base) BeforeUpdate(ctx context.Context) (context.Context, error) {
	now := time.Now()
	b.UpdatedAt = now
	return ctx, nil
}
