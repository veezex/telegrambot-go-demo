package postgress

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	applePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/apple"
	colorPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/color"
	storagePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
)

type appleRes struct {
	Id        uint64  `db:"id"`
	Price     float64 `db:"price"`
	ColorId   uint64  `db:"color_id"`
	ColorName string  `db:"name"`
}

type bucket struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) storagePkg.AppleStorage {
	return &bucket{
		pool: pool,
	}
}

func (b *bucket) Add(ctx context.Context, a *applePkg.Apple) error {
	tx, err := b.pool.Begin(ctx)
	defer tx.Rollback(ctx)

	c, err := getColor(ctx, tx, &a.Color)
	a.Color = *c

	if err != nil {
		return err
	}

	// insert apple
	query, args, err := squirrel.Insert("apples").
		Columns("color_id, price").
		Values(c.Id, a.Price).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("Apples.Inc: to sql: %w", err)
	}

	row := tx.QueryRow(ctx, query, args...)
	if err := row.Scan(&a.Id); err != nil {
		return fmt.Errorf("Apples.Inc: insert: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (b *bucket) Delete(ctx context.Context, id uint64) error {
	query, args, err := squirrel.Delete("apples").
		Where(squirrel.Eq{
			"id": id,
		},
		).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("Apples.Delete: to sql: %w", err)
	}

	_, err = b.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("Apples.Delete: exec: %w", err)
	}

	return nil
}

func (b *bucket) Get(ctx context.Context, id uint64) (*applePkg.Apple, error) {
	query, args, err := squirrel.Select("apples.id, apples.price, apples.color_id, colors.name").
		From("apples").
		Where(squirrel.Eq{
			"apples.id": id,
		},
		).
		LeftJoin("colors on colors.id = apples.color_id").
		Limit(1).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("Apples.Get: to sql: %w", err)
	}

	var apples []appleRes
	if err := pgxscan.Select(ctx, b.pool, &apples, query, args...); err != nil {
		return nil, fmt.Errorf("Apples.Get: select: %w", err)
	}

	if len(apples) == 0 {
		return nil, storagePkg.ErrNotExists
	}

	return &mapApples(apples)[0], nil
}

func (b *bucket) List(ctx context.Context, opts *storagePkg.PaginationOpts) ([]applePkg.Apple, error) {
	statement := squirrel.Select("apples.id, apples.price, apples.color_id, colors.name").
		From("apples").
		LeftJoin("colors on colors.id = apples.color_id").
		PlaceholderFormat(squirrel.Dollar)

	if opts != nil {
		statement.
			OrderBy("apples.id").
			Limit(opts.Limit).
			Offset(opts.Offset)
	}

	query, args, err := statement.ToSql()
	if err != nil {
		return nil, fmt.Errorf("Apples.List: to sql: %w", err)
	}

	var apples []appleRes
	if err := pgxscan.Select(ctx, b.pool, &apples, query, args...); err != nil {
		return nil, fmt.Errorf("Apples.List: select: %w", err)
	}

	return mapApples(apples), nil
}

func (b *bucket) Update(ctx context.Context, a *applePkg.Apple) error {
	tx, err := b.pool.Begin(ctx)
	defer tx.Rollback(ctx)

	c, err := getColor(ctx, tx, &a.Color)
	a.Color = *c
	if err != nil {
		return err
	}

	query, args, err := squirrel.Update("apples").
		Where(squirrel.Eq{
			"id": a.Id,
		},
		).
		Set("price", a.Price).
		Set("color_id", c.Id).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("Apples.Update: to sql: %w", err)
	}

	row := tx.QueryRow(ctx, query, args...)
	if err := row.Scan(&a.Id); err != nil {
		return fmt.Errorf("Apples.(update color): update: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func getColor(ctx context.Context, tx pgx.Tx, color *colorPkg.Color) (*colorPkg.Color, error) {
	result, err := checkColor(ctx, tx, color)
	if err != nil {
		col, err := addColor(ctx, tx, color)

		if err != nil {
			return nil, err
		}

		result = col
	}

	return result, nil
}

func checkColor(ctx context.Context, tx pgx.Tx, color *colorPkg.Color) (*colorPkg.Color, error) {
	query, args, err := squirrel.Select("id, name").
		From("colors").
		Where(squirrel.Eq{
			"name": color.Name,
		},
		).
		Limit(1).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("Apples.(check color): to sql: %w", err)
	}

	var colors []colorPkg.Color
	if err := pgxscan.Select(ctx, tx, &colors, query, args...); err != nil {
		return nil, fmt.Errorf("Apples.(check color): select: %w", err)
	}

	if len(colors) == 0 {
		return nil, errors.New("Not found")
	}

	return &colors[0], nil
}

func addColor(ctx context.Context, tx pgx.Tx, color *colorPkg.Color) (*colorPkg.Color, error) {
	query, args, err := squirrel.Insert("colors").
		Columns("name").
		Values(color.Name).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("Apples.(add color): to sql: %w", err)
	}
	row := tx.QueryRow(ctx, query, args...)
	if err := row.Scan(&color.Id); err != nil {
		return nil, fmt.Errorf("Apples.(add color): insert: %w", err)
	}
	return color, nil
}

func mapApples(input []appleRes) []applePkg.Apple {
	result := make([]applePkg.Apple, len(input))
	for index, val := range input {
		result[index] = applePkg.Apple{
			Id:    val.Id,
			Price: val.Price,
			Color: colorPkg.Color{
				Id:   val.ColorId,
				Name: val.ColorName,
			},
		}
	}
	return result
}
