package repository

import (
	"database/sql"
	"errors"
	"github.com/JulisTafita/go-savannahTech/internal/api"
	"github.com/JulisTafita/go-savannahTech/internal/config"
	"github.com/JulisTafita/go-savannahTech/internal/db"
	"strconv"
	"time"
)

// IRepository represents the repository information.
type IRepository struct {
	Id                  int       `json:"id"`
	CreatedAt           time.Time `json:"created_at"`
	RepositoryCreatedAt time.Time `json:"repository_created_at"`
	RepositoryUpdatedAt time.Time `json:"repository_updated_at"`
	Name                string    `json:"name"`
	Owner               string    `json:"owner"`
	Description         string    `json:"description"`
	Url                 string    `json:"url"`
	Language            string    `json:"language"`
	ForksCount          int       `json:"forks_count"`
	StarsCount          int       `json:"stars_count"`
	OpenIssuesCount     int       `json:"open_issues_count"`
	WatchersCount       int       `json:"watchers_count"`
	Commits             []ICommit `json:"commits" q:"_"`
}

// ICommit represents the commit information.
type ICommit struct {
	Id            int       `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	IRepositoryId int       `json:"i_repository_id"`
	AuthorName    string    `json:"author"`
	AuthorEmail   string    `json:"author_email"`
	AuthorLogin   string    `json:"author_login"`
	Date          time.Time `json:"date"`
	Message       string    `json:"message"`
	Url           string    `json:"url"`
}

// CommitAuthor represents the author of a commit.
type CommitAuthor struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Login       string `json:"login"`
	CommitCount int    `json:"commit_count"`
}

// GetRepositoryData retrieves and updates repository data.
func GetRepositoryData(repositoryName string) (err error) {
	var repository = &IRepository{}

	// Retrieve repository data from the API.
	repository, err = repository.GetRepositoryFromApi(repositoryName)
	if err != nil {
		return
	}

	// Add or update the repository in the database.
	err = repository.AddOrUpdate()
	if err != nil {
		return
	}

	// Sync repository commits.
	err = repository.syncRepositoryCommits()
	if err != nil {
		return err
	}

	return err
}

// GetRepositoryFromApi retrieves repository data from the API.
func (repository *IRepository) GetRepositoryFromApi(name string) (newRepository *IRepository, err error) {
	var repositoryService api.RepositoryService
	var repositories []api.Repository

	// Check if a private repository should be used.
	if config.UsePrivateRepository() {
		repos, _ := repositoryService.GetRepository(name)
		if repos.ID > 0 {
			return parseRepository(&repos), err
		}

		repositories, _ = repositoryService.GetUserRepositoryList()

		newRepository = look(name, repositories)
		if newRepository != nil {
			return newRepository, err
		}
	}

	// Search for the repository by name.
	repositories, _ = repositoryService.SearchRepositoryByName(name)

	newRepository = look(name, repositories)
	if newRepository != nil {
		return newRepository, err
	}

	return nil, errors.New("repository not found")
}

// AddOrUpdate inserts or updates the repository in the database.
func (repository *IRepository) AddOrUpdate() (err error) {
	var _repository IRepository
	query := `
		INSERT INTO i_repository(
		                         repository_created_at, 
		                         repository_updated_at,
		                         name, 
		                         owner, 
		                         description,
		                         url,
		                         language,
		                         forks_count,
		                         stars_count,
		                         open_issues_count,
		                         watchers_count
		                        )
			VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE 
			                                                    repository_created_at = VALUES(repository_created_at),
			                                                    repository_updated_at = VALUES(repository_updated_at),
			                                                    owner = VALUES(owner),
			                                                    description = VALUES(description),
			                                                    url = VALUES(url),
			                                                    language = VALUES(language),
			                                                    forks_count = VALUES(forks_count),
			                                                    stars_count = VALUES(stars_count),
			                                                    open_issues_count = VALUES(open_issues_count),
			                                                    watchers_count = VALUES(watchers_count)
	`

	var args = []any{
		repository.RepositoryCreatedAt, repository.RepositoryUpdatedAt,
		repository.Name, repository.Owner, repository.Description,
		repository.Url, repository.Language,
		repository.ForksCount, repository.StarsCount,
		repository.OpenIssuesCount, repository.WatchersCount,
	}

	// Execute the insert or update query.
	err = db.Exec(query, args...)
	if err != nil {
		return err
	}

	// Retrieve the inserted or updated repository.
	err = db.Get(&_repository, `SELECT * FROM i_repository WHERE name = ?`, repository.Name)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	// Update the repository fields with the retrieved data.
	repository.Id = _repository.Id
	repository.CreatedAt = _repository.CreatedAt
	repository.RepositoryCreatedAt = _repository.RepositoryCreatedAt
	repository.RepositoryUpdatedAt = _repository.RepositoryUpdatedAt
	repository.Name = _repository.Name
	repository.Owner = _repository.Owner
	repository.Description = _repository.Description
	repository.Url = _repository.Url
	repository.Language = _repository.Language
	repository.ForksCount = _repository.ForksCount
	repository.StarsCount = _repository.StarsCount
	repository.OpenIssuesCount = _repository.OpenIssuesCount
	repository.WatchersCount = _repository.WatchersCount

	return err
}

// syncRepositoryCommits synchronizes the repository commits.
func (repository *IRepository) syncRepositoryCommits() (err error) {
	var repositoryService api.RepositoryService
	var subQueryPart string
	var subQueryArgs []any

	// Retrieve the commits from the API.
	repositoryCommits, err := repositoryService.GetRepositoryCommits(repository.Owner, repository.Name, config.Default.Option.PullingStartDate, config.Default.Option.PullingEndDate)
	if err != nil {
		return err
	}

	// Process each commit.
	for i := 0; i < len(repositoryCommits); i++ {
		commit := *parseCommit(&repositoryCommits[i])
		subQueryPart += "(?,?,?,?,?,?,?),"
		subQueryArgs = append(
			subQueryArgs,
			repository.Id,
			commit.AuthorName,
			commit.AuthorEmail,
			commit.AuthorLogin,
			commit.Date.Format(time.DateTime),
			commit.Message, commit.Url)
	}

	if len(subQueryPart) > 0 {
		// Delete existing commits from the database based on the pulling options.
		query := `DELETE FROM i_commit WHERE i_repository_id = ?`

		if !config.Default.Option.ResetCollection {
			if !config.Default.Option.PullingEndDate.IsZero() {
				query += ` AND date < ` + "'" + config.Default.Option.PullingEndDate.Format(time.DateTime) + "' "
			}

			if !config.Default.Option.PullingStartDate.IsZero() {
				query += ` AND date > ` + "'" + config.Default.Option.PullingStartDate.Format(time.DateTime) + "' "
			}
		}

		err = db.Exec(query, repository.Id)
		if err != nil {
			return err
		}

		// Insert new commits into the database.
		subQueryPart = "INSERT INTO i_commit(i_repository_id, author_name, author_email,author_login, date, message, url) VALUES " + subQueryPart[:len(subQueryPart)-1]
		err = db.Exec(subQueryPart, subQueryArgs...)
		if err != nil {
			return err
		}
	}

	return err
}

// GetRepository retrieves a repository from the database.
func GetRepository() (repository IRepository, err error) {
	err = db.Get(&repository, `SELECT * FROM i_repository WHERE name = ?`, config.Default.Github.RepositoryName)
	if err != nil {
		return repository, err
	}

	err = db.Select(&repository.Commits, `SELECT * FROM i_commit WHERE i_repository_id = ?`, repository.Id)
	if err != nil {
		return repository, err
	}

	return repository, err
}

// GetTopNCommitAuthors retrieves the top N commit authors.
func GetTopNCommitAuthors(N int) (author []CommitAuthor, err error) {
	query := `
		SELECT
		    ic.author_name as 'name',
		    ic.author_email as 'email',
		    ic.author_login as 'login',
		    count(ic.id) as 'commit_count' 
		FROM i_repository ir
			JOIN i_commit ic On ir.id = ic.i_repository_id
		WHERE ir.name = ? group by ic.author_login ORDER BY count(ic.id) DESC
	`

	if N > 0 {
		query += ` LIMIT ` + strconv.Itoa(N)
	}

	err = db.Select(&author, query, config.Default.Github.RepositoryName)
	if err != nil {
		return nil, err
	}

	return author, err
}
