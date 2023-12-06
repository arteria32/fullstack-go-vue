-- name: ListUsers :many
-- get all users ordered by the username
SELECT
        *
FROM
        gowebapp.users
ORDER BY
        user_name;

-- name: ListImages :many
-- get all images ordered by the id
SELECT
        *
FROM
        gowebapp.images
ORDER BY
        image_id;

-- name: ListExercises :many
-- get all exercises ordered by the exercise name
SELECT
        *
FROM
        gowebapp.exercises
ORDER BY
        exercise_name;

-- name: ListSets :many
-- get all exercise sets ordered by weight
SELECT
        *
FROM
        gowebapp.sets
ORDER BY
        weight;

-- name: ListWorkouts :many
-- get all workouts ordered by id
SELECT
        *
FROM
        gowebapp.workouts
ORDER BY
        workout_id;

-- name: GetUser :one
-- get users of a particular user_id
SELECT
        *
FROM
        gowebapp.users
WHERE
        user_id = 1;

-- name: GetUserWorkout :many
-- get a particular user information and workouts
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
        AND u.user_id = 1;

-- name: GetUserSets :many
-- get a particular user information, exercise sets and workouts
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
        AND u.user_id = 1;

-- name: GetUserImage :one
-- get a particular user image
SELECT
        u.name,
        u.user_id,
        i.image_data
FROM
        gowebapp.users u,
        gowebapp.images i
WHERE
        u.user_id = i.user_id
        AND u.user_id = 1;

-- name: DeleteUsers :exec
-- delete a particular user
DELETE FROM
        gowebapp.users
WHERE
        user_id = 1;

-- name: DeleteUserImage :exec
-- delete a particular user's image
DELETE FROM
        gowebapp.images i
WHERE
        i.user_id = 1;

-- name: DeleteUserWorkouts :exec
-- delete a particular user's workouts
DELETE FROM
        gowebapp.workouts w
WHERE
        w.user_id = 1;

-- name: DeleteExercise :exec
-- delete a particular exercise
DELETE FROM
        gowebapp.exercises e
WHERE
        e.exercise_id = 1;

-- name: DeleteSets :exec
-- delete a particular exercise sets
DELETE FROM
        gowebapp.sets s
WHERE
        s.set_id = 1;

-- name: CreateExercise :execresult
-- insert a new exercise
INSERT INTO
        gowebapp.exercises (Exercise_Name)
values
        (1);
-- Retrieve the last inserted ID using the LAST_INSERT_ID() function
SELECT
        LAST_INSERT_ID() as last_inserted_id;

-- name: UpsertExercise :execresult
-- insert or update exercise of a particular id
INSERT INTO
        gowebapp.exercises (Exercise_Name)
VALUES
        (1) ON DUPLICATE KEY
UPDATE
        Exercise_Name = EXCLUDED.Exercise_Name;

SELECT
        LAST_INSERT_ID() as Exercise_ID;

-- name: CreateUserImage :execresult
-- insert a new image
INSERT INTO
        gowebapp.images (User_ID, Content_Type, Image_Data)
values
        (1, 2, 3);

SELECT
        LAST_INSERT_ID() as LAST_INSERT_ID;

-- name: UpsertUserImage :execresult
-- insert or update image of a particular id
INSERT INTO
        gowebapp.images (Image_Data)
VALUES
        (1) ON DUPLICATE KEY
UPDATE
        Image_Data = EXCLUDED.Image_Data;

SELECT
        LAST_INSERT_ID() as Image_ID;

-- name: CreateSet :execresult
-- insert new exercise sets
INSERT INTO
        gowebapp.sets (Exercise_Id, Weight)
values
        (1, 2);

SELECT
        LAST_INSERT_ID() as LAST_INSERT_ID;

-- name: UpdateSet :execresult
-- insert a sets id
UPDATE
        gowebapp.sets
SET
        Exercise_Id = 1,
        Weight = 2
WHERE
        set_id = 3;

SELECT
        LAST_INSERT_ID() as LAST_INSERT_ID;

-- name: CreateWorkout :execresult
-- insert new workouts
INSERT INTO
        gowebapp.workouts (User_ID, Set_ID, Start_Date)
values
        (1, 2, 3);

SELECT
        LAST_INSERT_ID() as last_workout_id;

-- name: UpsertWorkout :execresult
-- insert or update workouts based of a particular ID
INSERT INTO
        gowebapp.workouts (User_ID, Set_ID, Start_Date)
values
        (1, 2, 3) ON DUPLICATE KEY
UPDATE
        User_ID = EXCLUDED.User_ID,
        Set_ID = EXCLUDED.Set_ID,
        Start_Date = EXCLUDED.Start_Date;

SELECT
        LAST_INSERT_ID() as Workout_ID;

-- name: CreateUsers :execresult
-- insert new user
INSERT INTO
        gowebapp.users (User_Name, Pass_Word_Hash, name)
VALUES
        (1, 2, 3);

SELECT
        LAST_INSERT_ID();