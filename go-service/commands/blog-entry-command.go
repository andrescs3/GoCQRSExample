package blogEntryCommand

import (
	"database/sql"
	"fmt"
	"go-service/database"
	"go-service/entity"
)

func GetAll() ([]entity.BlogEntry, error) {
	ctx := database.Context

	var blogentries []entity.BlogEntry
	// Check if database is alive.
	err := database.Db.PingContext(ctx)
	if err != nil {
		return blogentries, err
	}

	tsql := "SELECT [Id],[Title],[Content],[CreatedDate] FROM [devdb].[TestSchema].[BlogEntry]"

	// Execute query
	rows, err := database.Db.QueryContext(ctx, tsql)
	if err != nil {
		return blogentries, err
	}

	defer rows.Close()

	// Iterate through the result set.
	for rows.Next() {
		var content, title string
		var id, createddate int

		// Get values from row.
		err := rows.Scan(&id, &title, &content, &createddate)
		if err != nil {
			return blogentries, err
		}

		fmt.Printf("ID: %d, Name: %s, Location: %s\n", id, title, content)
		var blogEntry entity.BlogEntry
		blogEntry.Content = content
		blogEntry.Title = title
		blogEntry.CreatedDate = createddate
		blogEntry.ID = id
		blogentries = append(blogentries, blogEntry)
	}

	return blogentries, nil
}

func GetByID(ID int) (entity.BlogEntry, error) {
	ctx := database.Context

	var blogEntry entity.BlogEntry
	// Check if database is alive.
	err := database.Db.PingContext(ctx)
	if err != nil {
		return blogEntry, err
	}

	tsql := "SELECT [Id],[Title],[Content],[CreatedDate] FROM [devdb].[TestSchema].[BlogEntry] WHERE Id = @id "

	// Execute query
	rows, err := database.Db.QueryContext(ctx, tsql, sql.Named("id", ID))
	if err != nil {
		return blogEntry, err
	}

	defer rows.Close()

	// Iterate through the result set.
	for rows.Next() {
		var content, title string
		var id, createddate int

		// Get values from row.
		err := rows.Scan(&id, &title, &content, &createddate)
		if err != nil {
			return blogEntry, err
		}

		fmt.Printf("ID: %d, Name: %s, Location: %s\n", id, title, content)
		var blogEntry entity.BlogEntry
		blogEntry.Content = content
		blogEntry.Title = title
		blogEntry.CreatedDate = createddate
		blogEntry.ID = id
		return blogEntry, nil
	}

	return blogEntry, nil
}

func UpdateBlogEntry(blogEntry entity.BlogEntry) (int64, error) {
	ctx := database.Context

	// Check if database is alive.
	err := database.Db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("UPDATE [BlogEntry] SET [Title] = @title ,[Content] = @content ,[CreatedDate] = @createdDate WHERE ID = @id ")

	// Execute non-query with named parameters
	result, err := database.Db.ExecContext(
		ctx,
		tsql,
		sql.Named("title", blogEntry.Title),
		sql.Named("content", blogEntry.Content),
		sql.Named("createdDate", blogEntry.CreatedDate),
		sql.Named("id", blogEntry.ID))
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}

func DeleteBlogEntry(ID int) (int64, error) {
	ctx := database.Context

	// Check if database is alive.
	err := database.Db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("DELETE FROM TestSchema.Employees WHERE ID = @id;")

	// Execute non-query with named parameters
	result, err := database.Db.ExecContext(ctx, tsql, sql.Named("id", ID))
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}

func CreateBlogEntry(blogEntry entity.BlogEntry) (int, error) {
	ctx := database.Context

	// Check if database is alive.
	err := database.Db.PingContext(ctx)

	if err != nil {
		return -1, err
	}

	tsql := `
      INSERT INTO [TestSchema].[BlogEntry]
           ([Title]
           ,[Content]
           ,[CreatedDate])
     VALUES
           (@title
           ,@content
           ,@createdDate)
    `

	stmt, err := database.Db.Prepare(tsql)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(
		ctx,
		sql.Named("title", blogEntry.Title),
		sql.Named("content", blogEntry.Content),
		sql.Named("createdDate", blogEntry.CreatedDate))
	var newID int
	err = row.Scan(&newID)
	if err != nil {
		return -1, err
	}

	return newID, nil
}

func main() {
	var config database.Config
	config.Server = "PCANDRES\\MSSQL"
	config.Port = 1433
	config.User = "sa"
	config.Password = "test123"
	config.DataBaseName = "devdb"

	var connectionString = database.GetConnectionString(config)
	var err = database.Connect(connectionString)

	if err != nil {

		fmt.Printf(fmt.Sprintf("server=%s", err.Error()))
	}
	if err == nil {
		var err2 error
		blogentries, err2 := GetAll()
		fmt.Printf(fmt.Sprintf("server=%d", len(blogentries)))
		if err != nil {

			fmt.Printf(fmt.Sprintf("server=%s", err2.Error()))
		}
	}

}
