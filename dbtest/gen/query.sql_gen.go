// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: query.sql

package dbtest

import (
	"context"
	"database/sql"
	"time"
)

const createExercise = `-- name: CreateExercise :execresult
INSERT INTO
        gowebapp.exercises (Exercise_Name)
values
        (1)
`

// insert a new exercise
func (q *Queries) CreateExercise(ctx context.Context) (sql.Result, error) {
	return q.db.ExecContext(ctx, createExercise)
}

const createSet = `-- name: CreateSet :execresult
INSERT INTO
        gowebapp.sets (Exercise_Id, Weight)
values
        (1, 2)
`

// insert new exercise sets
func (q *Queries) CreateSet(ctx context.Context) (sql.Result, error) {
	return q.db.ExecContext(ctx, createSet)
}

const createUserImage = `-- name: CreateUserImage :execresult
INSERT INTO
        gowebapp.images (User_ID, Content_Type, Image_Data)
values
        (1, 2, 3)
`

// insert a new image
func (q *Queries) CreateUserImage(ctx context.Context) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUserImage)
}

const createUsers = `-- name: CreateUsers :execresult
INSERT INTO
        gowebapp.users (User_Name, Pass_Word_Hash, name)
VALUES
        (1, 2, 3)
`

// insert new user
func (q *Queries) CreateUsers(ctx context.Context) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUsers)
}

const createWorkout = `-- name: CreateWorkout :execresult
INSERT INTO
        gowebapp.workouts (User_ID, Set_ID, Start_Date)
values
        (1, 2, 3)
`

// insert new workouts
func (q *Queries) CreateWorkout(ctx context.Context) (sql.Result, error) {
	return q.db.ExecContext(ctx, createWorkout)
}

const deleteExercise = `-- name: DeleteExercise :exec
DELETE FROM
        gowebapp.exercises e
WHERE
        e.exercise_id = 1
`

// delete a particular exercise
func (q *Queries) DeleteExercise(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteExercise)
	return err
}

const deleteSets = `-- name: DeleteSets :exec
DELETE FROM
        gowebapp.sets s
WHERE
        s.set_id = 1
`

// delete a particular exercise sets
func (q *Queries) DeleteSets(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteSets)
	return err
}

const deleteUserImage = `-- name: DeleteUserImage :exec
DELETE FROM
        gowebapp.images i
WHERE
        i.user_id = 1
`

// delete a particular user's image
func (q *Queries) DeleteUserImage(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteUserImage)
	return err
}

const deleteUserWorkouts = `-- name: DeleteUserWorkouts :exec
DELETE FROM
        gowebapp.workouts w
WHERE
        w.user_id = 1
`

// delete a particular user's workouts
func (q *Queries) DeleteUserWorkouts(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteUserWorkouts)
	return err
}

const deleteUsers = `-- name: DeleteUsers :exec
DELETE FROM
        gowebapp.users
WHERE
        user_id = 1
`

// delete a particular user
func (q *Queries) DeleteUsers(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteUsers)
	return err
}

const getUser = `-- name: GetUser :one
SELECT
        user_id, user_name, pass_word_hash, name, config, created_at, is_enabled
FROM
        gowebapp.users
WHERE
        user_id = 1
`

// get users of a particular user_id
func (q *Queries) GetUser(ctx context.Context) (GowebappUser, error) {
	row := q.db.QueryRowContext(ctx, getUser)
	var i GowebappUser
	err := row.Scan(
		&i.UserID,
		&i.UserName,
		&i.PassWordHash,
		&i.Name,
		&i.Config,
		&i.CreatedAt,
		&i.IsEnabled,
	)
	return i, err
}

const getUserImage = `-- name: GetUserImage :one
SELECT
        u.name,
        u.user_id,
        i.image_data
FROM
        gowebapp.users u,
        gowebapp.images i
WHERE
        u.user_id = i.user_id
        AND u.user_id = 1
`

type GetUserImageRow struct {
	Name      string `db:"name" json:"name"`
	UserID    int64  `db:"user_id" json:"userId"`
	ImageData []byte `db:"image_data" json:"imageData"`
}

// get a particular user image
func (q *Queries) GetUserImage(ctx context.Context) (GetUserImageRow, error) {
	row := q.db.QueryRowContext(ctx, getUserImage)
	var i GetUserImageRow
	err := row.Scan(&i.Name, &i.UserID, &i.ImageData)
	return i, err
}

const getUserSets = `-- name: GetUserSets :many
SELECT
        u.user_id,
        w.workout_id,
        w.start_date,
        s.set_id,
        s.weight
FROM
        gowebapp.users u,
        gowebapp.workouts w,
        gowebapp.sets s
WHERE
        u.user_id = w.user_id
        AND w.set_id = s.set_id
        AND u.user_id = 1
`

type GetUserSetsRow struct {
	UserID    int64     `db:"user_id" json:"userId"`
	WorkoutID int64     `db:"workout_id" json:"workoutId"`
	StartDate time.Time `db:"start_date" json:"startDate"`
	SetID     int64     `db:"set_id" json:"setId"`
	Weight    int32     `db:"weight" json:"weight"`
}

// get a particular user information, exercise sets and workouts
func (q *Queries) GetUserSets(ctx context.Context) ([]GetUserSetsRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserSets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserSetsRow
	for rows.Next() {
		var i GetUserSetsRow
		if err := rows.Scan(
			&i.UserID,
			&i.WorkoutID,
			&i.StartDate,
			&i.SetID,
			&i.Weight,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserWorkout = `-- name: GetUserWorkout :many
SELECT
        u.user_id,
        w.workout_id,
        w.start_date,
        w.set_id
FROM
        gowebapp.users u,
        gowebapp.workouts w
WHERE
        u.user_id = w.user_id
        AND u.user_id = 1
`

type GetUserWorkoutRow struct {
	UserID    int64     `db:"user_id" json:"userId"`
	WorkoutID int64     `db:"workout_id" json:"workoutId"`
	StartDate time.Time `db:"start_date" json:"startDate"`
	SetID     int64     `db:"set_id" json:"setId"`
}

// get a particular user information and workouts
func (q *Queries) GetUserWorkout(ctx context.Context) ([]GetUserWorkoutRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserWorkout)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserWorkoutRow
	for rows.Next() {
		var i GetUserWorkoutRow
		if err := rows.Scan(
			&i.UserID,
			&i.WorkoutID,
			&i.StartDate,
			&i.SetID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listExercises = `-- name: ListExercises :many
SELECT
        exercise_id, exercise_name
FROM
        gowebapp.exercises
ORDER BY
        exercise_name
`

// get all exercises ordered by the exercise name
func (q *Queries) ListExercises(ctx context.Context) ([]GowebappExercise, error) {
	rows, err := q.db.QueryContext(ctx, listExercises)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GowebappExercise
	for rows.Next() {
		var i GowebappExercise
		if err := rows.Scan(&i.ExerciseID, &i.ExerciseName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listImages = `-- name: ListImages :many
SELECT
        image_id, user_id, content_type, image_data
FROM
        gowebapp.images
ORDER BY
        image_id
`

// get all images ordered by the id
func (q *Queries) ListImages(ctx context.Context) ([]GowebappImage, error) {
	rows, err := q.db.QueryContext(ctx, listImages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GowebappImage
	for rows.Next() {
		var i GowebappImage
		if err := rows.Scan(
			&i.ImageID,
			&i.UserID,
			&i.ContentType,
			&i.ImageData,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSets = `-- name: ListSets :many
SELECT
        set_id, exercise_id, weight
FROM
        gowebapp.sets
ORDER BY
        weight
`

// get all exercise sets ordered by weight
func (q *Queries) ListSets(ctx context.Context) ([]GowebappSet, error) {
	rows, err := q.db.QueryContext(ctx, listSets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GowebappSet
	for rows.Next() {
		var i GowebappSet
		if err := rows.Scan(&i.SetID, &i.ExerciseID, &i.Weight); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsers = `-- name: ListUsers :many
SELECT
        user_id, user_name, pass_word_hash, name, config, created_at, is_enabled
FROM
        gowebapp.users
ORDER BY
        user_name
`

// get all users ordered by the username
func (q *Queries) ListUsers(ctx context.Context) ([]GowebappUser, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GowebappUser
	for rows.Next() {
		var i GowebappUser
		if err := rows.Scan(
			&i.UserID,
			&i.UserName,
			&i.PassWordHash,
			&i.Name,
			&i.Config,
			&i.CreatedAt,
			&i.IsEnabled,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listWorkouts = `-- name: ListWorkouts :many
SELECT
        workout_id, set_id, user_id, exercise_id, start_date
FROM
        gowebapp.workouts
ORDER BY
        workout_id
`

// get all workouts ordered by id
func (q *Queries) ListWorkouts(ctx context.Context) ([]GowebappWorkout, error) {
	rows, err := q.db.QueryContext(ctx, listWorkouts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GowebappWorkout
	for rows.Next() {
		var i GowebappWorkout
		if err := rows.Scan(
			&i.WorkoutID,
			&i.SetID,
			&i.UserID,
			&i.ExerciseID,
			&i.StartDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSet = `-- name: UpdateSet :execresult
UPDATE
        gowebapp.sets
SET
        Exercise_Id = 1,
        Weight = 2
WHERE
        set_id = 3
`

// insert a sets id
func (q *Queries) UpdateSet(ctx context.Context) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateSet)
}

const upsertExercise = `-- name: UpsertExercise :execresult
INSERT INTO
        gowebapp.exercises (Exercise_Name)
VALUES
        (1) ON DUPLICATE KEY
UPDATE
        Exercise_Name = EXCLUDED.Exercise_Name
`

// insert or update exercise of a particular id
func (q *Queries) UpsertExercise(ctx context.Context) (sql.Result, error) {
	return q.db.ExecContext(ctx, upsertExercise)
}

const upsertUserImage = `-- name: UpsertUserImage :execresult
INSERT INTO
        gowebapp.images (Image_Data)
VALUES
        (1) ON DUPLICATE KEY
UPDATE
        Image_Data = EXCLUDED.Image_Data
`

// insert or update image of a particular id
func (q *Queries) UpsertUserImage(ctx context.Context) (sql.Result, error) {
	return q.db.ExecContext(ctx, upsertUserImage)
}

const upsertWorkout = `-- name: UpsertWorkout :execresult
INSERT INTO
        gowebapp.workouts (User_ID, Set_ID, Start_Date)
values
        (1, 2, 3) ON DUPLICATE KEY
UPDATE
        User_ID = EXCLUDED.User_ID,
        Set_ID = EXCLUDED.Set_ID,
        Start_Date = EXCLUDED.Start_Date
`

// insert or update workouts based of a particular ID
func (q *Queries) UpsertWorkout(ctx context.Context) (sql.Result, error) {
	return q.db.ExecContext(ctx, upsertWorkout)
}
