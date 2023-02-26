package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/suite"
	"gokeeper/internal/core"
	"os"
	"reflect"
	"testing"
	"time"
)

const defaultTestPostgresAddr = "postgres://homestead:homestead@localhost:54321/homestead"

type userRepoSuite struct {
	suite.Suite
	pool *pgxpool.Pool
}

func Test_userRepoSuite(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	suite.Run(t, new(userRepoSuite))
}

func (ur *userRepoSuite) SetupSuite() {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = defaultTestPostgresAddr
	}
	conf, err := pgxpool.ParseConfig(addr)
	if err != nil {
		log.Error().Err(err).Msg("Unable to parse Database config")
		return
	}

	ur.pool, err = pgxpool.ConnectConfig(context.Background(), conf)
	if err != nil {
		log.Error().Err(err).Msg("Unable to connect to db")
		return
	}
}

func (ur *userRepoSuite) SetupTest() {
	query := `CREATE OR REPLACE FUNCTION truncate_tables(username IN VARCHAR) RETURNS void AS $$
DECLARE
    statements CURSOR FOR
        SELECT tablename FROM pg_tables
        WHERE tableowner = username AND schemaname = 'public';
BEGIN
    FOR stmt IN statements LOOP
        EXECUTE 'TRUNCATE TABLE ' || quote_ident(stmt.tablename) || ' CASCADE;';
    END LOOP;
END;
$$ LANGUAGE plpgsql;`
	_, err := ur.pool.Exec(context.Background(), query)
	ur.Require().NoError(err)
}

func (ur *userRepoSuite) TestUserRepository_Create() {
	type args struct {
		ctx  context.Context
		user *core.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.Background(),
				user: &core.User{
					ID:        1,
					Login:     "user_test",
					Password:  "hashed_password",
					AesSecret: "fsdfds",
					RsaSecret: "grthtyjty",
					CreatedAt: time.Now(),
					UpdatedAt: nil,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		ur.Run(tt.name, func() {
			users := &UserRepository{
				db: ur.pool,
			}
			if err := users.Create(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				ur.T().Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			gotUser, err := users.GetByID(tt.args.ctx, 1)
			ur.Require().NoError(err)
			ur.Require().NotNil(gotUser)
			ur.Require().Equal(gotUser.ID, 1)
		})
	}
}

func (ur *userRepoSuite) TestUserRepository_GetByID() {
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		args    args
		want    *core.User
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.Background(),
				id:  45,
			},
			want: &core.User{
				ID:        45,
				Login:     "test_user",
				Password:  "password",
				AesSecret: "aes_secret",
				RsaSecret: "rsa_secret",
				CreatedAt: time.Now(),
				UpdatedAt: nil,
			},
		},
	}
	for _, tt := range tests {
		ur.Run(tt.name, func() {
			users := &UserRepository{
				db: ur.pool,
			}
			err := users.Create(context.Background(), &core.User{
				ID:        45,
				Login:     "test_user",
				Password:  "password",
				AesSecret: "aes_secret",
				RsaSecret: "rsa_secret",
				CreatedAt: time.Now(),
				UpdatedAt: nil,
			})
			ur.Require().NoError(err)

			gotUser, err := users.GetByLogin(tt.args.ctx, "test_user")
			if (err != nil) != tt.wantErr {
				ur.T().Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, err := users.GetByID(tt.args.ctx, gotUser.ID)
			if (err != nil) != tt.wantErr {
				ur.T().Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				ur.T().Errorf("GetByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func (ur *userRepoSuite) TestUserRepository_GetByLogin() {
	type args struct {
		ctx   context.Context
		login string
	}
	tests := []struct {
		name    string
		args    args
		want    *core.User
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:   context.Background(),
				login: "test_user",
			},
			want: &core.User{
				ID:        45,
				Login:     "test_user_2",
				Password:  "password",
				AesSecret: "aes_secret",
				RsaSecret: "rsa_secret",
				CreatedAt: time.Now(),
				UpdatedAt: nil,
			},
		},
	}
	for _, tt := range tests {
		ur.Run(tt.name, func() {
			users := &UserRepository{
				db: ur.pool,
			}
			err := users.Create(context.Background(), &core.User{
				ID:        45,
				Login:     "test_user_2",
				Password:  "password",
				AesSecret: "aes_secret",
				RsaSecret: "rsa_secret",
				CreatedAt: time.Now(),
				UpdatedAt: nil,
			})
			ur.Require().NoError(err)

			got, err := users.GetByLogin(tt.args.ctx, tt.args.login)
			if (err != nil) != tt.wantErr {
				ur.T().Errorf("GetByLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				ur.T().Errorf("GetByLogin() got = %v, want %v", got, tt.want)
			}
		})
	}
}
