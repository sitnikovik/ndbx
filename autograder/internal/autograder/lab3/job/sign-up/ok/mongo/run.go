package mongo

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/app/mongo/user/collection"
	userdocs "github.com/sitnikovik/ndbx/autograder/internal/app/mongo/user/doc"
	"github.com/sitnikovik/ndbx/autograder/internal/app/mongo/user/doc/key"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/variable"
	mongodoc "github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/strings"
	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run checks that the user with the given variables
// exists in the MongoDB collection and that the password is not equal
// to the provided password (i.e. must be hashed).
func (s *Step) Run(
	ctx context.Context,
	vars step.Variables,
) error {
	vals := variable.NewValues(vars)
	usr := vals.MustUser()
	all, err := s.mongo.AllBy(
		ctx,
		collection.Name,
		mongodoc.NewKVs(
			mongodoc.NewKV(
				key.FullName,
				usr.FullName(),
			),
			mongodoc.NewKV(
				key.Username,
				usr.Username(),
			),
		),
	)
	if err != nil {
		return errs.WrapJoin(
			"failed to get users from MongoDB",
			errs.ErrExternalDependencyFailed,
			err,
		)
	}
	err = numbers.AssertEquals(1, all.Len())
	if err != nil {
		return err
	}
	pwd, ok := all.
		First().
		KVs().
		MustGet(key.Password).(string)
	if !ok {
		panic("password must be a string")
	}
	err = strings.AssertNotEquals(
		vals.MustUserPassword(),
		pwd,
	)
	if err != nil {
		return err
	}
	idxx, err := s.mongo.Indexes(ctx, collection.Name)
	if err != nil {
		return errs.WrapJoin(
			"failed to get indexes from MongoDB",
			errs.ErrExternalDependencyFailed,
			err,
		)
	}
	idx := idxx.For(key.Username)
	if idx.Empty() {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"expected an index on the "+
				log.String(key.Username)+
				" field, but it does not exist",
		)
	}
	if !idx.Unique() {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"expected an index on the "+
				log.String(key.Username)+
				" field to be unique, but it is not",
		)
	}
	vars.Set(
		variable.User,
		userdocs.
			NewUserDocument(all.First()).
			ToUser(),
	)
	return nil
}
