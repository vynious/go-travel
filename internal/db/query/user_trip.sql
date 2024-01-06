
-- CreateUserTrip :one
INSERT INTO user_trip (
                       trip_id,
                       user_id
) VALUES (
          $1, $2
         ) RETURNING *;


-- DeleteUserTrip :one
DELETE FROM USER_TRIP
WHERE user_id = $1
  AND trip_id = $2
RETURNING *;