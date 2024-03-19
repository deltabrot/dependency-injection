package psqlstore

type Migration struct {
	Up   string
	Down string
}

var migrations = []Migration{
	{
		Up: `CREATE TABLE IF NOT EXISTS songs (
      id SERIAL PRIMARY KEY,
      title TEXT,
      artist TEXT
    )`,
		Down: `DROP TABLE IF EXISTS songs`,
	},
}
