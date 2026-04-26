package generator

import (
	"fmt"
	"log/slog"
	"math/rand"

	"github.com/gocql/gocql"
)

func Generate(
	id int,
	ks string,
	session *gocql.Session,
	jobs <-chan struct{},
) {
	for range jobs {
		err := session.
			Query(
				`INSERT INTO `+ks+`.sensor_data (sensor_id, created_at, value) VALUES (?, toTimestamp(now()), ?)`,
				fmt.Sprintf(
					"sensor_%02d",
					rand.Intn(5)+1,
				),
				rand.Float32()*40+10,
			).
			Exec()
		if err != nil {
			slog.Error(
				fmt.Sprintf(
					"%d: failed to insert: %v", id,
					err,
				),
			)
		}
	}
}
