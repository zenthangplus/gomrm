package gomrm

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)

type DB struct {
    db *sqlx.DB
}

func Connect(driverName string, dataSource string) (*DB, error) {
    db, err := sqlx.Open(driverName, dataSource)
    if err != nil {
        return nil, err
    }
    return &DB{db: db}, nil
}

func (c *DB) Close() error {
    return c.db.Close()
}

func (c *DB) QueryRaw(query string) (Dataset, error) {
    rows, err := c.db.Queryx(query)
    if err != nil {
        return nil, err
    }
    results := make(Dataset, 0)
    for rows.Next() {
        result := make(map[string]interface{})
        if err := rows.MapScan(result); err != nil {
            return nil, err
        }
        results = append(results, result)
    }
    return results, nil
}

func (c *DB) Query(query string) (*DataCollection, error) {
    results, err := c.QueryRaw(query)
    if err != nil {
        return nil, err
    }
    return NewDataCollection(results), nil
}

type Dataset []map[string]interface{}

type DataCollection struct {
    dataset Dataset
    size    int
    cursor  *int
    current *map[string]interface{}
}

func NewDataCollection(dataset Dataset) *DataCollection {
    return &DataCollection{
        dataset: dataset,
        size:    len(dataset),
        cursor:  nil,
    }
}

func (d *DataCollection) Dataset() *Dataset {
    return &d.dataset
}

func (d *DataCollection) First() map[string]interface{} {
    if d.size == 0 {
        return nil
    }
    return d.dataset[0]
}

func (d *DataCollection) Last() map[string]interface{} {
    if d.size == 0 {
        return nil
    }
    return d.dataset[d.size-1]
}

func (d *DataCollection) Get() map[string]interface{} {
    return *d.current
}

func (d *DataCollection) Next() bool {
    if d.cursor == nil {
        zero := 0
        d.cursor = &zero
    } else {
        *d.cursor++
    }
    if *d.cursor >= d.size {
        return false
    }
    d.current = &d.dataset[*d.cursor]
    return true
}

func (d *DataCollection) Rewind() {
    d.cursor = nil
    d.current = nil
}
