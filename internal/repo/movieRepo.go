package repo

import (
	"context"
	"database/sql"
	"errors"
	"github.com/fouched/go-movies-htmx/internal/models"
)

func AllMovies() ([]*models.Movie, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select 
			id, title, release_date, runtime, mpaa_rating, description,
			coalesce(image, ''), created_at, updated_at
		from
		    movies
		order by
		    title		
	`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*models.Movie
	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.ReleaseDate,
			&movie.RunTime,
			&movie.MPAARating,
			&movie.Description,
			&movie.Image,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		movies = append(movies, &movie)
	}

	return movies, nil
}

func GetMovieByID(id int) (*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select 
		    id, title, release_date, runtime, mpaa_rating, description,
			coalesce(image, ''), created_at, updated_at
		from
		    movies
		where 
		    id = $1
	`
	row := db.QueryRowContext(ctx, query, id)

	var movie models.Movie
	err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.ReleaseDate,
		&movie.RunTime,
		&movie.MPAARating,
		&movie.Description,
		&movie.Image,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	// get genres, if any
	query = `
		select 
		    g.id, g.genre
		from 
		    movies_genres mg
		left join genres g on mg.genre_id = g.id
		where 
		    mg.movie_id = $1
		order by
		    g.genre
	`
	rows, err := db.QueryContext(ctx, query, movie.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	defer rows.Close()

	var genres []*models.Genre
	for rows.Next() {
		var g models.Genre
		err := rows.Scan(
			&g.ID,
			&g.Genre,
		)
		if err != nil {
			return nil, err
		}
		genres = append(genres, &g)
	}

	movie.Genres = genres

	return &movie, nil
}

func GetMovieByIDForEdit(id int) (*models.Movie, []*models.Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select 
		    id, title, release_date, runtime, mpaa_rating, description,
			coalesce(image, ''), created_at, updated_at
		from
		    movies
		where 
		    id = $1
	`
	row := db.QueryRowContext(ctx, query, id)

	var movie models.Movie
	err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.ReleaseDate,
		&movie.RunTime,
		&movie.MPAARating,
		&movie.Description,
		&movie.Image,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)
	if err != nil {
		return nil, nil, err
	}

	// get genres, if any
	query = `
		select 
		    g.id, g.genre
		from 
		    movies_genres mg
		left join genres g on mg.genre_id = g.id
		where 
		    mg.movie_id = $1
		order by
		    g.genre
	`
	rows, err := db.QueryContext(ctx, query, movie.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, nil, err
	}
	defer rows.Close()

	var genres []*models.Genre
	var genresArray []int
	for rows.Next() {
		var g models.Genre
		err := rows.Scan(
			&g.ID,
			&g.Genre,
		)
		if err != nil {
			return nil, nil, err
		}
		genres = append(genres, &g)
		genresArray = append(genresArray, g.ID)
	}

	movie.Genres = genres
	movie.GenresArray = genresArray

	var allGenres []*models.Genre
	query = "select id, genre from genres order by genre"
	gRows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, nil, err
	}
	defer gRows.Close()

	for gRows.Next() {
		var g models.Genre
		err := gRows.Scan(
			&g.ID,
			&g.Genre,
		)
		if err != nil {
			return nil, nil, err
		}
		allGenres = append(allGenres, &g)
	}

	return &movie, allGenres, nil
}

func GetAllGenres() ([]*models.Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, genre, created_at, updated_at from genres order by genre`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []*models.Genre
	for rows.Next() {
		var g models.Genre
		err := rows.Scan(
			&g.ID,
			&g.Genre,
			&g.CreatedAt,
			&g.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		genres = append(genres, &g)
	}

	return genres, nil
}

func InsertMovie(movie *models.Movie) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `insert into movies (title, description, release_date, runtime, mpaa_rating, image, created_at, updated_at) 
			values ($1, $2, $3, $4, $5, $6, $7, $8) returning id`

	var id int

	err := db.QueryRowContext(ctx, stmt,
		movie.Title,
		movie.Description,
		movie.ReleaseDate,
		movie.RunTime,
		movie.MPAARating,
		movie.Image,
		movie.CreatedAt,
		movie.UpdatedAt,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func UpdateMovieGenres(id int, genreIDs []int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from movies_genres where movie_id = $1`
	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	for _, n := range genreIDs {
		stmt = `insert into movies_genres (movie_id, genre_id) values ($1, $2)`
		_, err = db.ExecContext(ctx, stmt, id, n)
		if err != nil {
			return err
		}
	}

	return nil
}
