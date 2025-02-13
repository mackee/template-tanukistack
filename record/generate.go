package record

//go:generate go tool genddl -driver pg -outpath ../sql/pg.sql
//go:generate go tool sqlla --dir-all --dialect postgresql
