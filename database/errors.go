// Copyright 2022 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package database

import (
	"strings"

	"github.com/juju/errors"
	"github.com/mattn/go-sqlite3"

	"github.com/juju/juju/database/driver"
)

// IsErrConstraintUnique returns true if the input error was
// returned by SQLite due to violation of a unique constraint.
func IsErrConstraintUnique(err error) bool {
	if err == nil {
		return false
	}

	var sqliteErr sqlite3.Error
	if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
		return true
	}

	// TODO (manadart 2022-12-16): The logic above works in unit tests using an
	// in-memory SQLite DB, but appears to fail when running with Dqlite.
	// Extended error codes can be enabled via PRAGMA, but we need to
	// investigate further.
	if strings.Contains(strings.ToLower(err.Error()), "unique constraint failed") {
		return true
	}

	return false
}

// IsErrRetryable returns true if the given error might be
// transient and the interaction can be safely retried.
func IsErrRetryable(err error) bool {
	var dErr *driver.Error

	if errors.As(err, &dErr) && dErr.Code == driver.ErrBusy {
		return true
	}

	if errors.Is(err, sqlite3.ErrLocked) || errors.Is(err, sqlite3.ErrBusy) {
		return true
	}

	// Unwrap errors one at a time.
	for ; err != nil; err = errors.Unwrap(err) {
		if strings.Contains(err.Error(), "database is locked") {
			return true
		}

		if strings.Contains(err.Error(), "cannot start a transaction within a transaction") {
			return true
		}

		if strings.Contains(err.Error(), "bad connection") {
			return true
		}

		if strings.Contains(err.Error(), "checkpoint in progress") {
			return true
		}
	}

	return false
}
